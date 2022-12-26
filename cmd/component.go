package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"text/template"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wbroberts/mrf/internal/config"
	"github.com/wbroberts/mrf/internal/templates"
)

type Component struct {
  Name string
	Props bool
	Dir string
}

type ComponentConfig struct {
	Dir string
}

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Creates a React component and test",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		withProps, _ := cmd.Flags().GetBool("props")
		path, _ := cmd.Flags().GetString("dir")
		skipTests, _ := cmd.Flags().GetBool("skip-tests")
		wd, _ := os.Getwd()
		checkComponentsDir(wd, path)
		
		data := Component{
			Name: name,
			Props: withProps,
			Dir: 	strings.Join([]string{wd, path}, "/"),
		}

		wg := new(sync.WaitGroup);
		
		routines := 2

		if skipTests {
			routines = 1
		}

		wg.Add(routines)

		go createComponent(&data, wg)

		if !skipTests {
			go createTest(&data, wg)
		}

		wg.Wait()
	},
}

func green(t string) string {
	return color.New(color.FgGreen).SprintFunc()(t)
}

func createComponent(data *Component, wg *sync.WaitGroup) {
	filename := data.Name + ".component.tsx"
	f := createFile(data.Dir, filename)
	defer f.Close()

	t, _ := template.New("component").Parse(string(templates.Component()))
	t.Execute(f, &data)
	wg.Done()

	fmt.Println(green("created"), filename)
}

func createTest(data *Component, wg *sync.WaitGroup) {
	filename := data.Name + ".component.test.tsx"
	f := createFile(data.Dir, filename)
	defer f.Close()

	t, _ := template.New("test").Parse(string(templates.ComponentTest()))
	t.Execute(f, &data)
	wg.Done()

	fmt.Println(green("created"), filename)
}

func createFile(dir, filename string) *os.File {
	filepath := strings.Join([]string{dir, filename}, "/")
	file, err := os.Create(filepath)
		
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func checkComponentsDir(dir string, path string) {
	dirPath := strings.Join([]string{dir, path}, "/")

	if _, err := os.Stat(dirPath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dirPath, os.ModePerm)
		} else {
			log.Fatal(err)
		}
	}
}

var ( 
	defaultDir = "components"
	defaultProps = false
)

func init() {
	config := config.GetConfig()

	if config.Component.Dir != "" {
		defaultDir = config.Component.Dir
	}

	if config.Component.Props {
		defaultProps = config.Component.Props
	}

	rootCmd.AddCommand(componentCmd)
	componentCmd.Flags().StringP("dir", "d", defaultDir, "Customizes the path from the root")
	componentCmd.Flags().BoolP("props", "p", defaultProps, "Adds props type for the component")
	componentCmd.Flags().Bool("skip-tests", false, "Skips adding a test file")
}
