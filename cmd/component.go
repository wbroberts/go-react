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
	"github.com/wbroberts/go-react/internal/config"
	"github.com/wbroberts/go-react/internal/templates"
)

type Component struct {
	Name  string
	Props bool
	Dir   string
}

type ComponentConfig struct {
	Dir string
}

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Creates a React component and test",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		withProps, _ := cmd.Flags().GetBool("props")
		path, _ := cmd.Flags().GetString("dir")
		skipTests, _ := cmd.Flags().GetBool("skip-tests")
		wd, _ := os.Getwd()
		dir := strings.Join([]string{wd, path}, "/")

		checkDir(dir)

		data := Component{
			Name:  name,
			Props: withProps,
			Dir:   dir,
		}

		wg := new(sync.WaitGroup)

		routines := 1

		if !skipTests {
			routines = 2
		}

		wg.Add(routines)

		go createComponent(&data, wg)

		if !skipTests {
			go createTest(&data, wg)
		}

		wg.Wait()
	},
}

func checkDir(d string) {
	if _, err := os.Stat(d); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(d, os.ModePerm)
		} else {
			log.Fatal(err)
		}
	}
}

func createComponent(data *Component, wg *sync.WaitGroup) {
	filename := data.Name + ".component.tsx"
	filepath := strings.Join([]string{data.Dir, filename}, "/")
	t, _ := template.New("component").Parse(string(templates.Component()))
	createFile(data, filepath, t)

	fmt.Println(green("created"), filename)

	wg.Done()
}

func createTest(data *Component, wg *sync.WaitGroup) {
	filename := data.Name + ".component.test.tsx"
	filepath := strings.Join([]string{data.Dir, filename}, "/")
	t, _ := template.New("test").Parse(string(templates.ComponentTest()))

	createFile(data, filepath, t)
	fmt.Println(green("created"), filename)

	wg.Done()
}

func createFile(data *Component, p string, t *template.Template) {
	f, err := os.Create(p)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	t.Execute(f, data)
}

func green(t string) string {
	return color.New(color.FgGreen).SprintFunc()(t)
}

var (
	defaultDir   = "components"
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
