package pointer

import "fmt"

func Pointer() {
	var k string = "pandu"

	fmt.Println("K (value): ", k)
	fmt.Println("K (address): ", &k)

}
