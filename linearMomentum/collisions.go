package linearMomentum

import (
	"fmt"
	"math"
)

// CalculateCoefficientOfRestitution calcula o coeficiente
// de restituição com base nos valores das velocidades.
// O primeiro parâmetro `separationVelocity` é a veloci
// dade de afastamento entre os corpos depois da colisão.
// O segundo parâmetro `approachVelocity` é a velocidade
// de aproximação entre os corpos antes da colisão.
// O retorno é o coeficiente de restituição.
func CalculateCoefficientOfRestitution(separationVelocity, approachVelocity float64) float64 {
	if approachVelocity == 0 {
		fmt.Println("Approach velocity must be different from zero because if it's equal to zero, how would occur a collision?")
		return 0.0
	}

	return math.Abs(separationVelocity) / math.Abs(approachVelocity)
}
