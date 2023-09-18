package main

import (
	"os"
	"testing"
	"text/template"
)

func Test_text_template(t *testing.T) {
	tempText1 := "value is {{.}}\n"
	t1 := template.New("t1")
	t1, err := t1.Parse(tempText1)
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse(tempText1))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name string, txt string) *template.Template {
		return template.Must(template.New(name).Parse(txt))
	}

	t2 := Create("t2", "name is {{.Name}}\n")
	t2.Execute(os.Stdout,
		struct {
			Name string
		}{
			Name: "Alice",
		})
	t2.Execute(os.Stdout,
		map[string]bool{
			"Name": true,
		})

	t3 := Create("t3",
		`
		value is '{{.}}'
		// zero value will return false
		{{if .}} it's not zero value
		{{else}} it's zero value
		{{end}}`)
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, " ")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", `
	Range: 
	{{range .}}
		{{.}}
	{{end}}`)
	t4.Execute(os.Stdout, []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
	})
}
