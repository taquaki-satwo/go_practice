package main

import (
	"os"
	"text/template"
)

type Wheather struct {
	Place string
}

func main() {
	tpl := template.Must(template.ParseFiles("template.tpl"))
	wheather := Wheather{"東京"}

	tpl.Execute(os.Stdout, wheather)
}
