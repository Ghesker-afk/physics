package espelhos

/* Podemos classificar um espelho esférico em côncavo ou em
* convexo. Sabendo que ambos surgem a partir do corte de uma
* esfera, podemos admitir que aqueles em que a parte refleto-
* ra é para a direção interna são côncavos, e aqueles em que
* a parte refletora é para fora são convexos.
 */

/* O centro de curvatura (C) do espelho é o ponto que é o
 * centro da esfera que originou o espelho. O raio de curvatu-
 * ra (R) é o raio do esfera refletora que originou o espelho.
 * O eixo principal do espelho é um diâmetro da esfera que
 * parte o espelho no meio, e o ponto médio do espelho é cha-
 * mado de vértice do espelho (V). A distância focal (f) é a
 * distância do foco até o vértice do espelho, e nas nossas
 * condições (espelho gaussiano), ele será igual à metade do
 * raio de curvatura (f = R/2). O foco (F) é o ponto médio
 * entre o vértice do espelho e o centro de curvatura.
 */

/* No espelho côncavo, o centro de curvatura, o vértice, o
 * raio de curvatura e todos os demais elementos ficam todos
 * na região em que a luz incide - a superfície refletora.
 * Quando esses elementos ficam nas "costas" do espelho, o
 * espelho é classificado como convexo.
 */

/* Trabalhamos com raios de luz classificados como
 * paraxiais, que são praticamente paralelos em relação ao
 * eixo principal, e também bem próximos do eixo principal.
 * São esses espelhos, ditos como gaussianos, que estudaremos.
 * O foco principal é o ponto onde convergem todos os raios
 * refletidos que incidem paralelos ao eixo principal do
 * espelho.
 *
 */

type EspelhoEsferico struct {
	MirrorType      string
	FocalDistance   float64
	CurvatureRadius float64
}

func (e *EspelhoEsferico) CalculateFocalDistance() float64 {
	return e.CurvatureRadius / 2
}
