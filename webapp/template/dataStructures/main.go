package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("data.gohtml"))

}

func main() {
	cities := map[string]string{"1": "Vancouver", "2": "Toronto",
		"3": "Istanbul", "4": "Mecca", "5": "Oslo"}
	err := tpl.ExecuteTemplate(os.Stdout, "data.gohtml", cities)
	if err != nil {
		log.Fatal(err)
	}
}
