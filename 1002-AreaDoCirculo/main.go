package main

import (
	"fmt"
	"math"
)

func main() {
	var raio, area float64
	n := 3.14159

	fmt.Scan(&raio)
	area = n * math.Pow(raio, 2)
	//area = math.Round(area) esse arredonda tudo, deixa so o primeiro, por isso precisa multiplicar com a quantidade de zero igual casas que quer e depois dividir
	area = (math.Round(area*10000) / 10000)
	fmt.Print("A=", area, "\n")

}
