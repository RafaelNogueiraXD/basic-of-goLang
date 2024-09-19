package main

import (
	f "fmt"
	// "math"
)

func notas(N float64)(float64){

    notas := []int{100, 50, 20, 10, 5, 2}
    f.Printf("NOTAS:\n")
    for _, nota := range notas {
        quantidade := int(N)/ nota
        f.Printf("%d nota(s) de R$ %d.00\n", quantidade, nota)
        N -= float64(quantidade * nota)
    }
    return N
}

func moedas(N float64) {
    f.Printf("MOEDAS:\n")
    moedas := []float64{1.00, 0.50, 0.25, 0.10, 0.05, 0.01}
    for _, moeda := range moedas {
        quantidade := int(N * 100) / int(moeda*100)
        f.Printf("%d moeda(s) de R$ %.2f\n", quantidade, moeda)
        N -= float64(quantidade) * moeda
    }
}

func tempo(N int){
	horas := N / 3600

    minutos := (N % 3600) / 60

    segundos := N % 60

    f.Printf("%d:%d:%d\n", horas, minutos, segundos)
} 
func CalculaIdade(diasTotais int){
    anos := diasTotais / 365

    meses := (diasTotais % 365) / 30
    dias := (diasTotais % 365) % 30
    f.Printf("%d ano(s)\n", anos)
    f.Printf("%d mes(es)\n", meses)
    f.Printf("%d dia(s)\n", dias)
}

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func operacaoSoma(n1 int, n2 int, d1 int, d2 int) (int, int) {
	// Soma: (N1*D2 + N2*D1) / (D1*D2)
	numerador := (n1*d2 + n2*d1)
	denominador := (d1 * d2)
	return numerador, denominador
}

func operacaoSub(n1 int, n2 int, d1 int, d2 int) (int, int) {
	// Subtração: (N1*D2 - N2*D1) / (D1*D2)
	numerador := (n1*d2 - n2*d1)
	denominador := (d1 * d2)
	return numerador, denominador
}

func operacaoMul(n1 int, n2 int, d1 int, d2 int) (int, int) {
	// Multiplicação: (N1*N2) / (D1*D2)
	numerador := n1 * n2
	denominador := d1 * d2
	return numerador, denominador
}

func operacaoDiv(n1 int, n2 int, d1 int, d2 int) (int, int) {
	// Divisão: (N1*D2) / (N2*D1)
	numerador := n1 * d2
	denominador := n2 * d1
	return numerador, denominador
}

func simplicador(numerado int, denominador int)(int, int){
    var cond int
    for cond == 0 {
        if numerado % 2 == 0 && denominador % 2 == 0{
            numerado /= 2
            denominador /= 2
        }else if numerado % 3 == 0 && denominador % 3 == 0{
            numerado /= 3
            denominador /= 3
        }else if numerado % 5 == 0 && denominador % 5 == 0{
            numerado /= 5
            denominador /= 5
        }else{
            cond = 1
        }
    }

    return numerado,denominador
}

// Função para calcular o máximo divisor comum (MDC)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Função para simplificar uma fração
func simplifyFraction(numerator, denominator int) (int, int) {
	commonDivisor := gcd(numerator, denominator)
	return numerator / commonDivisor, denominator / commonDivisor
}

// Função para processar a expressão e realizar a operação
func processExpression(n1, d1 int, operator string, n2, d2 int) (int, int, int, int) {
	var numerator, denominator int

	switch operator {
	case "+":
		numerator = n1*d2 + n2*d1
		denominator = d1 * d2
	case "-":
		numerator = n1*d2 - n2*d1
		denominator = d1 * d2
	case "*":
		numerator = n1 * n2
		denominator = d1 * d2
	case "/":
		numerator = n1 * d2
		denominator = n2 * d1
	}

	simplifiedNumerator, simplifiedDenominator := simplifyFraction(numerator, denominator)
	return numerator, denominator, simplifiedNumerator, simplifiedDenominator
}

func main() {
	var nCases int
	f.Scan(&nCases)

	for i := 0; i < nCases; i++ {
		var n1, d1, n2, d2 int
		var operator string

		// Lê a entrada na forma: "N1 / D1 operador N2 / D2"
		f.Scanf("%d / %d %s %d / %d", &n1, &d1, &operator, &n2, &d2)

		numerator, denominator, simplifiedNumerator, simplifiedDenominator := processExpression(n1, d1, operator, n2, d2)

		// Saída do resultado
		f.Printf("%d/%d = %d/%d\n", numerator, denominator, simplifiedNumerator, simplifiedDenominator)
	}
}