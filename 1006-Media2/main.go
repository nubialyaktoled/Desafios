package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	media := (a*2 + b*3 + c*5) / 10
	fMedia := fmt.Sprintf("%.1f", media)
	fmt.Println("MEDIA =", fMedia)
}
