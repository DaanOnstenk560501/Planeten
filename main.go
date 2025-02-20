package main

import (
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
	var afstand, tijd float64

	fmt.Print("Voer de afstand tot de ster in miljoenen kilometers in: ")
	fmt.Scanln(&afstand)

	fmt.Print("Voer de tijd voor een volledige omwenteling in uren in: ")
	fmt.Scanln(&tijd)

	afstand *= 1000000 // Zet miljoenen km om naar km

	snelheid := berekenSnelheid(afstand, tijd)
	fmt.Printf("De planeet beweegt zich met %.2f km/u door de ruimte.\n", snelheid)

	fmt.Print("Druk op Enter om af te sluiten...")
	fmt.Scanln() // Wacht op Enter voordat het programma wordt afgesloten
}
