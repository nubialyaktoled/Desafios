package main

import "fmt"

func main() {

	var codigo1, codigo2 string
	var numPecas1, numPecas2 int
	var valor1, valor2 float64

	fmt.Scan(&codigo1, &numPecas1, &valor1, &codigo2, &numPecas2, &valor2)
	fnum1 := float64(numPecas1)
	fnum2 := float64(numPecas2)
	total1 := fnum1 * valor1
	total2 := fnum2 * valor2
	totalGeral := total1 + total2
	formatFloat := fmt.Sprintf("%.2f", totalGeral)

	fmt.Println("VALOR A PAGAR: R$", formatFloat)

}
