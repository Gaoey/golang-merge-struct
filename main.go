package main

import (
	"encoding/json"
	"fmt"
)

type Planet struct {
	Name       string  `json:"name"`
	Aphelion   float64 `json:"aphelion"`   // in million km
	Perihelion float64 `json:"perihelion"` // in million km
	Axis       int64   `json:"Axis"`       // in km
	Radius     float64 `json:"radius"`
}

type Relation struct {
	Name string
	With *Planet
}

type PlanetWithMass struct {
	Planet
	Relation
	Mass float64
}

func main() {
	var mars = new(Planet)
	mars.Name = "Mars"
	mars.Aphelion = 249.2
	mars.Perihelion = 206.7
	mars.Axis = 227939100
	mars.Radius = 3389.5

	relation := Relation{
		Name: "single",
		With: mars,
	}

	marsWithMass := PlanetWithMass{
		Planet:   *mars,
		Mass:     639e21,
		Relation: relation,
	}

	b, _ := json.Marshal(marsWithMass)

	fmt.Printf("%s\n", b)
}
