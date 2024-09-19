package main

import (
	f "fmt"
	"math"
)

func calculaArea(raio float64) (float64){
	const pi = 3.14159
	return pi * raio * raio
}
func soma2(a int, b int) (int){
	return a + b
}
func produto(a int, b int) (int){
	return a * b
}
func media(a float32, b float32) (float32){
	return (a*3.5 + b*7.5)/11
}
func media3(a float32, b float32,c float32) (float32){
	return (a*2 + b*3 + c*5)/10
}
func diferenca(a int, b int,c int,d int) (int){
	return (a*b - c*d)
}
func salario(a int, b float64) (float64){
	return float64(a) * b
}

func salarioBonus(a float64, b float64) (float64){
	var comissao float64
	comissao = (15*b)/100
	return comissao + a
}
func esfera(raio float64) (float64){
	var pi float64
	pi = 3.14159
	return (4.0 / 3.0) * pi * math.Pow(raio, 3)
}
func forms() (float64,float64,float64,float64,float64){
	var A, B, C float64
    const pi = 3.14159

    f.Scan(&A, &B, &C)

    areaTriangulo := (A * C) / 2.0
    areaCirculo := pi * math.Pow(C, 2)
    areaTrapezio := ((A + B) * C) / 2.0
    areaQuadrado := B * B
    areaRetangulo := A * B
	return areaCirculo, areaTriangulo, areaTrapezio, areaQuadrado, areaRetangulo
}