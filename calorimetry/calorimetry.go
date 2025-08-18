package calorimetry

import (
	"errors"
)

func CalculateHeatCapacity(heat, temperatureVariation float64) float64 {
	return (heat / temperatureVariation)
}

func TemperatureDiffer(initial, final float64) float64 {
	return final - initial
}

func JoulesToCalories(joules float64) float64 {
	return joules / 4.2
}

func CalculateSensibleHeat(mass, specificHeat, temperatureVariation float64) (float64, error) {
	if mass <= 0 {
		return 0.0, errors.New("mass doesn't allowed to be negative")
	}

	if specificHeat <= 0 {
		return 0.0, errors.New("specific heat doesn't allowed to be negative")
	}

	sensibleHeat := (mass * specificHeat * temperatureVariation)
	return sensibleHeat, nil
}

func CalculateLatentHeat(mass, latentHeat float64) (float64, error) {
	if latentHeat <= 0 {
		return 0.0, errors.New("latent heat doesn't allowed to be less or equal to zero")
	}

	heat := (mass * latentHeat)
	return heat, nil
}
