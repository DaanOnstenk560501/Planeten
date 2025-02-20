package main

import (
	"flag"
	"fmt"
	"math"
)

func berekenSnelheid(afstandKm float64, tijdUren float64) float64 {
	/*
	   Berekent de snelheid in km/u.
	*/
	omtrekKm := 2 * math.Pi * afstandKm // Omtrek van de cirkelbaan
	snelheidKmU := omtrekKm / tijdUren
	return snelheidKmU
}

func main() {
	afstandPtr := flag.Float64("afstand", 150.0, "Afstand tot de ster in miljoenen kilometers")
	tijdPtr := flag.Float64("tijd", 8766.0, "Tijd voor een volledige omwenteling in uren")
	flag.Parse()

	afstand := *afstandPtr * 1000000 // Zet miljoenen km om naar km
	tijd := *tijdPtr

	snelheid := berekenSnelheid(afstand, tijd)
	fmt.Printf("De planeet beweegt zich met %.2f km/u door de ruimte.\n", snelheid)
}
