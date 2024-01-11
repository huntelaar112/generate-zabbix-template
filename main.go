package main

import (
	"os"
	"strings"
	"text/template"

	"github.com/huntelaar112/goutils/utils"
)

type Template struct {
	GroupTemplateUuid string
	TemplateName      string
	Endpoints         []Enpoint
}

type Enpoint struct {
	TemplateName string
	EndpointName string
}

func main() {
	zb_template := Template{
		GroupTemplateUuid: "7df96b18c230490a9a0a9e2307226338",
		TemplateName:      "PRODVN2-2GPU",
		Endpoints: []Enpoint{
			{
				TemplateName: "PRODVN2-2GPU",
				EndpointName: "id",
			},
			{
				TemplateName: "PRODVN2-2GPU",
				EndpointName: "vnhff",
			},
			{
				TemplateName: "PRODVN2-2GPU",
				EndpointName: "vnhff1",
			},
			{
				TemplateName: "PRODVN2-2GPU",
				EndpointName: "ghvn",
			},
			{
				TemplateName: "PRODVN2-2GPU",
				EndpointName: "vnpff",
			},
		},
	}

	funcMap := template.FuncMap{
		"idv4gen": idv4gen,
	}

	utils.FileCreate("./template-gen.yaml")
	f, _ := os.OpenFile("./template-gen.yaml", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	var tmplFile = "zabbix_templates.tmpl"
	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, zb_template)
	if err != nil {
		panic(err)
	}
}

func idv4gen() string {
	idv4, _ := utils.GenerateIdv4("")
	return strings.ReplaceAll(idv4, "-", "")
}
