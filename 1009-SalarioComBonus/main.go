package main

import "fmt"

func main() {

	var nome string
	var salario float64
	var numVendas float64
	var comissao float64

	fmt.Scan(&nome, &salario, &numVendas)
	comissao = (numVendas * 15) / 100
	salario = comissao + salario
	fSalario := fmt.Sprintf("%.2f", salario)

	fmt.Println("TOTAL = R$", fSalario)
}
