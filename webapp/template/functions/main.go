package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type City struct {
	Name       string
	population int
}

type Car struct {
	Name         string
	Manufacturer string
}

type CityCar struct {
	City City
	Car  Car
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	return s[:3]
}
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("functions.gohtml"))
}

func main() {
	vancouver := City{
		Name:       "Vancouver",
		population: 234234,
	}
	istanbul := City{
		Name:       "Istanbul",
		population: 42333,
	}
	milan := City{
		Name:       "Milan",
		population: 23434,
	}
	nano := Car{
		Name:         "Nano",
		Manufacturer: "TATA",
	}
	eon := Car{
		Name:         "Eon",
		Manufacturer: "Hyundai",
	}
	etios := Car{
		Name:         "Etios",
		Manufacturer: "Toyota",
	}

	var cityCars = []CityCar{
		CityCar{vancouver, nano},
		CityCar{istanbul, eon},
		CityCar{milan, etios},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "functions.gohtml", cityCars)
	if err != nil {
		log.Fatal(err)
	}
}
