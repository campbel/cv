package main

import (
	"html/template"
	"os"
	"strings"

	"github.com/campbel/yoshi"
	"gopkg.in/yaml.v3"
)

type App struct {
	CV     func(CVOptions)
	Letter func(LetterOptions)
}

type CVOptions struct {
	Config string `yoshi:"CONFIG;The yaml data to be used for generation;"`
}

type LetterOptions struct {
	Company string `yoshi:"COMPANY;The company to address the letter to;"`
	Role    string `yoshi:"ROLE;The role to address the letter to;"`
}

func main() {
	yoshi.New("cvgen").Run(App{
		CV: func(options CVOptions) {
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
			cvMarkdownTemplate.Execute(os.Stdout, cv)
		},
		Letter: func(options LetterOptions) {
			if options.Company == "" {
				panic("no company provided")
			}
			if options.Role == "" {
				panic("no role provided")
			}
			coverLetterMakrdownTemplate.Execute(os.Stdout, options)
		},
	})
}

var cvMarkdownTemplate = template.Must(template.New("md").Funcs(template.FuncMap{
	"join": func(s []string) string {
		return strings.Join(s, ", ")
	},
}).Parse(cvMarkdownTemplateText))

const cvMarkdownTemplateText = `
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
{{ if .Description }}{{ .Description }}{{end}}
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
	Description     string   `yaml:"description"`
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

var coverLetterMakrdownTemplate = template.Must(template.New("md").Parse(coverLetterTemplate))

const coverLetterTemplate = `
Dear Hiring Manager,


I am excited to submit my application for the {{.Role}} role at {{.Company}}. With fourteens years of experience in the industry, I am confident that I possess the skills and knowledge necessary to excel in this position.

Throughout my career, I have honed my skills in designing, developing, and deploying cloud-based software applications. I have worked on a wide range of projects, from simple web applications to complex distributed systems, and have gained valuable experience working with various cloud computing technologies such as Amazon Web Services (AWS) and Google Cloud Platform (GCP).

My expertise in cloud architecture and infrastructure design, as well as my strong coding skills in programming languages such as Go, Python, and JavaScript, make me an ideal candidate for this role. I am also experienced in working with containerization technologies such as Docker and Kubernetes.

In addition to my technical skills, I am a strong team player with excellent communication skills. I enjoy collaborating with colleagues to tackle complex challenges and find innovative solutions. I am also committed to staying up-to-date with the latest industry trends and best practices to ensure that my skills are always current.

Thank you for considering my application. I am excited about the opportunity to bring my skills and experience to {{.Company}} and contribute to the company's success.


Sincerely,

Chris Campbell
`
