package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templ.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{Number: "M", Name: "Maths", Units: "three"},
				course{Number: "E", Name: "English", Units: "four"},
				course{Number: "S", Name: "Science", Units: "five"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{Number: "G", Name: "Geography", Units: "one"},
				course{Number: "M", Name: "Morals", Units: "four"},
				course{Number: "H", Name: "History", Units: "five"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				course{Number: "P", Name: "Politics", Units: "three"},
				course{Number: "C", Name: "Cooking", Units: "four"},
				course{Number: "K", Name: "Kannada", Units: "five"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
