package main

import (
	"fmt"

	"github.com/refqihussein/golangforlife/golangforlife/Introduction/cetak"
	"github.com/refqihussein/golangforlife/golangforlife/Introduction/handler"
	"github.com/refqihussein/golangforlife/golangforlife/Introduction/model"
	"github.com/refqihussein/golangforlife/golangforlife/Introduction/pointer"
)

func main() {
	var a string = "nama"
	b := 1
	cetak.Cetak(a, b)

	d := cetak.Indra("dana cukup", "calon ada")
	fmt.Printf("%s \n", d)

	_, x := cetak.Adi()
	fmt.Printf("Istri adi :\n " + x)

	pointer.Pointer()

	var createdPerson *model.Person
	createdPerson = handler.Create()

	handler.Edit(createdPerson)
}
