package eletromagnetism

import "math"

const k = 8.99e9

// Parâmetros:
// float64 charge1: a carga do primeiro corpo em coulombs (C)
// float64 charge2: a carga do segundo corpo em coulombs (C)
// float64 distance: o quadrado da distância entre os centros dos corpos em metros quadrados (m²)
func CoulombForce(charge1, charge2, distance float64) float64 {
	return k * math.Abs(charge1*charge2) / math.Pow(distance, 2)
}
