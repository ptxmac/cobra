package cobra

import (
	"html/template"
	"io"
)

var (
	fishCompletionText = `
# fish completion

{{range .Commands}}{{ if not .Hidden }}
{{template "selectCmdTemplate" .}}
{{- end }}{{end}}
{{define "selectCmdTemplate" -}}
complete -f -c {{.Root.Name}} -n '__fish_use_subcommand' -a {{.Name}} -d '{{.Short}}'
{{end}}
`
)

func (c *Command) GenFishCompletion(w io.Writer) error {
	tmpl, err := template.New("Main").Parse(fishCompletionText)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, c.Root())
}
