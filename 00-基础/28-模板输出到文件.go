package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t := template.Must(template.New("").Parse(`
{{- range .}}{{.}}:
	echo "from {{.}}"
{{end}}
`))
	//t.Execute(os.Stdout, []string{"app1", "app2", "app3"})

	f, err := os.Create("./myfile.txt")
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = t.Execute(f, []string{"app1", "app2", "app3"})
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}