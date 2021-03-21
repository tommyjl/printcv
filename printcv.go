package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

type manifest struct {
	Author     author
	Summary    string
	Experience []experience
	Education  []education
	Skills     []skill
}

type author struct {
	Name     string
	Email    string
	Address  string
	GitHub   string
	LinkedIn string
}

type experience struct {
	From     string
	To       string
	Employer string
	Title    string
	Items    []string
}

type education struct {
	From        string
	To          string
	Institution string
	Degree      string
	Items       []string
}

type skill struct {
	Category string
	Items    []string
}

var cvtemplate string = `{{ printf "%-70v" .Author.Name }}{{ now | date "2006-01-02" }}

{{ printf "%-10v" "Email:" }}{{ .Author.Email }}
{{ printf "%-10v" "Address:" }}{{ .Author.Address }}
{{ printf "%-10v" "GitHub:" }}{{ .Author.GitHub }}
{{ printf "%-10v" "LinkedIn:" }}{{ .Author.LinkedIn }}


{{ .Summary | wrap 64 | indent 8 }}


Experience
{{ range .Experience }}
{{ printf "%-38v" .Employer | indent 2 }}{{printf "%40v" (printf "%s--%s" .From .To) }}
{{ .Title | indent 2 }}{{ range .Items }}
{{ printf "- %s" . | wrap 78 | indent 2 }}{{ end }}
{{ end }}
Education
{{ range .Education }}
{{ printf "%-38v" .Institution | indent 2 }}{{printf "%40v" (printf "%s--%s" .From .To) }}
{{ .Degree | indent 2 }}{{ range .Items }}
{{ printf "- %s" . | wrap 78 | indent 2 }}{{ end }}
{{ end }}

Skills
{{ range .Skills }}{{ $items := join ", " .Items }}
	{{- printf "%s: " .Category | nindent 2 }}
	{{- if gt 80 (add 4 (len .Category) (len $items)) }}
		{{- $items }}
	{{- else }}
		{{- $items | wrap 76 | nindent 4 }}
	{{- end }}
{{- end }}
`

func main() {
	t := template.Must(template.New("cv").Funcs(sprig.TxtFuncMap()).Parse(cvtemplate))

	jsonBytes, err := ioutil.ReadFile("example.json")
	if err != nil {
		log.Fatal("Unable to read file")
	}

	var cvData manifest
	err = json.Unmarshal(jsonBytes, &cvData)
	if err != nil {
		log.Fatal("Unable to unmarshal json", err)
	}

	err = t.Execute(os.Stdout, cvData)
	if err != nil {
		log.Fatal("Invalid data")
	}
}
