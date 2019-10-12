package main

import "fmt"

func main() {
	var a string = "nama"
	b := 1
	cetak(a, b)

	d := indra("dana cukup", "calon ada")
	fmt.Printf("%s \n", d)

}

func cetak(x string, y int) {
	fmt.Printf("hallo %s %d \n", x, y)
}

func indra(x string, y string) string {
	var a string
	if x == "dana cukup" && y == "calon ada" {
		a = "nikah"
	} else {
		a = "ditinggal nikah"
	}
	return a
}
