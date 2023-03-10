package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	media := (a*3.5 + b*7.5) / 11
	fmedia := fmt.Sprintf("%.5f", media)
	fmt.Println("MEDIA =", fmedia)
}
