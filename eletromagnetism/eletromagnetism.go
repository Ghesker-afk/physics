package eletromagnetism

import "math"

// Para um corpo interagir com um campo magnético, ele precisa ter carga e estar em movimento (v # 0 e q # 0).

func HasMagneticForce(hasMoviment, hasCarga bool) bool {
	if hasMoviment && hasCarga {
		return true
	}

	return false
}

func CalculateMagneticForce(magneticCamp, electricCharge, velocity, magneticAndVelocityAngle float64) float64 {
	return magneticCamp * electricCharge * velocity * math.Sin(magneticAndVelocityAngle)
}
