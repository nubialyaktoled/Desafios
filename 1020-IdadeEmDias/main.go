package main

import "fmt"

func main() {

	var idadeDias int
	fmt.Scan(&idadeDias)
	anos := idadeDias / 365
	sobraAno := idadeDias - (365 * anos)
	mes := sobraAno / 30
	sobraMes := sobraAno - (30 * mes)
	dia := sobraMes

	fmt.Println(anos, "ano(s)")
	fmt.Println(mes, "mes(es)")
	fmt.Println(dia, "dia(s)")
}
