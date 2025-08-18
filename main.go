package main

import (
	"fmt"
	"physics/calorimetry"
)

func main() {
	latent1, err := calorimetry.CalculateLatentHeat(2, 80)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(latent1)
}
