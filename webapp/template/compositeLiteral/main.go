package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

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

func init() {
	tpl = template.Must(template.ParseFiles("compositeLiteral.gohtml"))
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

	err := tpl.ExecuteTemplate(os.Stdout, "compositeLiteral.gohtml", cityCars)
	if err != nil {
		log.Fatal(err)
	}
}
