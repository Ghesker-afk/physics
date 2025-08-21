package eletromagnetism

import "math"

// Para um corpo interagir com um campo
// magnético, ele precisa ter carga e estar
// em movimento (v # 0 e q # 0).

// calculateMagneticForceInternal verifica se a força magnética pode existir, ou seja,
// se o corpo tem carga e está em movimento.
// Esta é uma função de uso interno.

func calculateMagneticForceInternal(charge, velocity float64) bool {
	return charge != 0 && velocity != 0
}

// CalculateMagneticForce calcula a magnitude da força magnética sobre uma partícula em movimento.

// Parâmetros:
// charge: a carga elétrica da partícula (q)
// velocity: o valor da velocidade da partícula (v)
// magneticField: o módulo do campo magnético (B)
// angle: o ângulo em radianos entre a velocidade e campo magnético

func CalculateMagneticForce(magneticField, charge, velocity, angle float64) float64 {
	if !calculateMagneticForceInternal(charge, velocity) {
		return 0
	}

	return magneticField * charge * velocity * math.Sin(angle)
}
