package main

import "fmt"

func main() {
	var numb int
	var horas int
	var valorHora float64
	var salario float64

	fmt.Scan(&numb, &horas, &valorHora)
	fhoras := float64(horas)
	salario = fhoras * valorHora
	fsalario := fmt.Sprintf("%.2f", salario)
	fmt.Println("NUMBER =", numb)
	fmt.Println("SALARY = U$", fsalario)

}
