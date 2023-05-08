package main

import (
	"html/template"
	"os"
	"strings"

	"github.com/campbel/yoshi"
	"gopkg.in/yaml.v3"
)

type Options struct {
	Config string `yoshi:"CONFIG;The yaml data to be used for generation;"`
}

func main() {
	yoshi.New("cvgen").Run(func(options Options) {
		// check if options.config exists
		if options.Config == "" {
			panic("no config provided")
		}
		data, err := os.ReadFile(options.Config)
		if err != nil {
			panic(err)
		}

		var cv CV
		if err = yaml.Unmarshal(data, &cv); err != nil {
			panic(err)
		}

		mdTemplate.Execute(os.Stdout, cv)
	})
}

var mdTemplate = template.Must(template.New("md").Funcs(template.FuncMap{
	"join": func(s []string) string {
		return strings.Join(s, ", ")
	},
}).Parse(markdownTemplate))

const markdownTemplate = `
# {{ .Info.Name }}
{{ .Info.Email }} | {{ .Info.Phone }} | [{{ .Info.GitHub }}](https://{{.Info.GitHub}})

{{ .Info.Summary }}

## Skills

{{ range .Skills }}
**{{.Title}}**: {{ join .Values }}
{{ end }}

## History
{{ range .History }}
### {{ .Company }}, {{ .Title }} ({{ .Start }} - {{ .End }})
{{ range .Accomplishments }}
- {{ . }}
{{ end }}
{{ end }}

## Publications
{{ range .Publications }}
- {{ if .Link }} [{{ .Title }}]({{ .Link }}) {{ else }} {{ .Title }} {{ end }} ({{ .Date}})
{{ .Description }}
{{ end }}

## Talks
{{ range .Talks }}
- {{ if .Link }} [{{ .Title }}]({{ .Link }}) {{ else }} {{ .Title }} {{ end }} ({{ .Date }})
{{ .Description }}
{{ end }}

## Community
{{ range .Community }}
- {{ if .Link }} [{{ .Title }}]({{ .Link }}) {{ else }} {{ .Title }} {{ end }} ({{ .Start }} - {{ .End }})
{{ .Description }}
{{ end }}

## Education
{{ range .Education }}
- {{ .School }} - {{ .Degree }}
{{ end }}
`

type CV struct {
	Info struct {
		Name    string `yaml:"name"`
		Email   string `yaml:"email"`
		GitHub  string `yaml:"github"`
		Phone   string `yaml:"phone"`
		Summary string `yaml:"summary"`
	} `yaml:"info"`
	Skills []struct {
		Title  string   `yaml:"title"`
		Values []string `yaml:"values"`
	} `yaml:"skills"`
	History      []History      `yaml:"history"`
	Publications []Publications `yaml:"publications"`
	Talks        []Talks        `yaml:"talks"`
	Community    []Community    `yaml:"community"`
	Education    []Education    `yaml:"education"`
}

type History struct {
	Company         string   `yaml:"company"`
	Title           string   `yaml:"title"`
	Start           string   `yaml:"start"`
	End             string   `yaml:"end"`
	Accomplishments []string `yaml:"accomplishments"`
}

type Publications struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
	Date        string `yaml:"date"`
}

type Talks struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
	Date        string `yaml:"date"`
}

type Community struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
	Start       string `yaml:"start"`
	End         string `yaml:"end"`
}

type Education struct {
	School string `yaml:"school"`
	Degree string `yaml:"degree"`
}
