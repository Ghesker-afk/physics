package hydrostatic

import "errors"

func CalculateDensity(mass, volume float64) (float64, error) {
	if mass <= 0 || volume <= 0 {
		return 0.0, errors.New("negative? no")
	}

	density := mass / volume
	return density, nil
}

func CalculateBuoyancy(density, volume, gravity float64) (float64, error) {
	if density <= 0 || volume <= 0 || gravity <= 0 {
		return 0.0, errors.New("negative? no")
	}

	buoyancy := (density * volume * gravity)
	return buoyancy, nil
}
