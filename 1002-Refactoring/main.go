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
	areaS := fmt.Sprintf("%.4f", area)
	fmt.Print("A=", areaS, "\n")

}
