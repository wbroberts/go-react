package templates

func Component() []byte {
	return []byte(
`import React from 'react';
{{ if .Props }}
type {{ .Name }}Props = {}

export const {{ .Name }}: React.ComponentType<{{ .Name }}Props> = ({}) => {
	return <div>{{ .Name }} renders</div>;
}
{{ else }}
export const {{ .Name }} = () => {
	return <div>{{ .Name }} renders</div>;
}

{{ end }}
`)
}

func ComponentTest() []byte {
	return []byte(
`import { screen, render } from '@testing-library/react';

import { {{ .Name }} } from './{{ .Name }}.component'

describe('{{ .Name }}', () => {
	it('renders', () => {
		render(<{{ .Name }} />)

		expect(screen.getByText(/{{ .Name }} renders/)).toBeDefined();
	})
})
`)
}