package waveTheory

import (
	"fmt"
	"strings"
)

type Wave struct {
	Wavelength       float64
	PropagationSpeed float64
	Period           float64
	Frequency        float64
}

func HasPhaseInversion(end string) {
	end = strings.ToLower(end)

	words := strings.Fields(end)
	finalWords := strings.Join(words, " ")
	if finalWords == "fixed end" {
		fmt.Println("This wave has phase shift.")
	} else if finalWords == "free end" {
		fmt.Println("This wave doesn't have phase inversion.")
	} else {
		fmt.Println("Invalid input.")
	}
}
