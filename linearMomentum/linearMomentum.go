package linearMomentum

import (
	"errors"
	"fmt"
)

func CalculateLinearMomentum(mass, velocity float64) (float64, error) {
	if mass <= 0 {
		return 0.0, errors.New("the mass must be greater than zero")
	}

	fmt.Println("The linear momentum is a vectorial physical quantity, and has same direction and of the velocity vector.")

	linearMomentum := (mass * velocity)
	return linearMomentum, nil
}

func CalculateImpulse(strength, timeInterval float64) (float64, error) {
	if timeInterval < 0 {
		return 0.0, errors.New("negative time? must be joking")
	}

	fmt.Println("The impulse is a vectorial physical quantity, and has same direction and of the velocity vector.")

	impulse := (strength * timeInterval)
	return impulse, nil
}
