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

// Parâmetros:
//   mass: a massa da partícula (m) em kg
//   velocity: a velocidade da partícula     (v) em m/s
//   charge: a carga da partícula (q) em C
//   magneticField: o módulo do campo magnético (B) em T

func CalculateCircularRadius(mass, velocity, charge, magneticField float64) float64 {
	if charge == 0 || magneticField == 0 {
		return 0
	}

	// o uso de math.Abs(charge) é para garantir que o raio sempre seja positivo.

	return (mass * velocity) / (math.Abs(charge) * magneticField)
}

// Parâmetros:
// mass: a massa da partícula (m) em kg
// charge: a carga da partícula (q) em C
// magneticField: o módulo do campo magnético (B) em T

func CalculatePeriod(mass, charge, magneticField float64) float64 {
	if charge == 0 || magneticField == 0 {
		return 0
	}

	return (2 * math.Pi * mass) / (math.Abs(charge) * magneticField)
}

// Parâmetros:
// magneticField: o módulo do campo magnético (B) em T
// current: a corrente elétrica no fio (i) em A
// length: o comprimento do fio imerso no campo (L) em m
// angle: o ângulo em radianos entre o fio e o campo magnético

func CalculateWireForce(magneticField, current, length, angle float64) float64 {
	if magneticField == 0 || current == 0 || length == 0 {
		return 0
	}

	sinTheta := math.Sin(angle)
	if sinTheta == 0 {
		return 0
	}

	return magneticField * current * length * sinTheta
}
