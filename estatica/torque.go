package estatica

import "math"

/* O torque ou momento de uma força é a capacidade dela
* de produzir rotação. Vake lembrar que a rotação de um corpo
* é sempre com relação a algum determinado ponto.
* O torque é diretamente proporcional à força que é aplicada
* perpendicularmente à superfície tomada como referência, e * também é diretamente proporcional à distância de onde a força * é aplicada até o ponto de referência. O torque é medido em
* newtons por metros, N/m, no S.I.
 */

/*
* Perceba que o torque em relação a um determinado ponto A é
* igual à componente perpendicular da força aplicada ao corpo
* multiplicada pela distância do ponto de aplicação da força
* até o ponto tomada como referência - nesse caso, A.
* Lembre-se que somente a componente de força que é
* perpendicular ao objeto é que produz rotação (torque), e
* podemos generalizar que essa força perpendicular é igual ao
* produto da força pelo seno do ângulo formado com
* a superfície.
 */

/*
 * O torque é uma grandeza vetorial, apesar de não nos
 * preocuparmos muito com essa informação no momento.
 * Novamente, sua unidade de medida no S.I é o N/m, e não J.
 */

func DegreeToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

type Object struct {
	TotalStrength    float64
	AngleInDegrees   float64
	DistanceInMeters float64
}

func (o *Object) TorqueStrength() float64 {
	AngleInRadians := DegreeToRadians(o.AngleInDegrees)
	return o.TotalStrength * math.Sin(AngleInRadians) * o.DistanceInMeters
}
