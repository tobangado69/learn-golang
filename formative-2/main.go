package main

import (
	"fmt"
)

func main() {

	// Soal 1
	a := "Bootcamp"
	b := "Digital"
	c := "Skill"
	d := "Sanbercode"
	e := "Golang"

	Kalimat := a + " " + b + " " + c + " " + d + " " + e
	fmt.Println(Kalimat)

	// Soal 2
	halo := "Halo Golang"
	fmt.Println(halo)

	// Soal 3
	kataPertama := "saya"
	kataKedua := "senang"
	kataKetiga := "belajar"
	kataKeempat := "golang"

	Kalimat2 := kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat
	fmt.Println(Kalimat2)

	// Soal 4
	angkaPertama := "8"
	angkaKedua := "5"
	angkaKetiga := "6"
	angkaKeempat := "7"

	kalimat3 := angkaPertama + " " + angkaKedua + " " + angkaKetiga + " " + angkaKeempat
	fmt.Println(kalimat3)

	// Soal 5
	var panjangPersegiPanjang int = 8
	var lebarPersegiPanjang int = 5

	var alasSegitiga int = 6
	var tinggiSegitiga int = 7

	kalimat4 := panjangPersegiPanjang + lebarPersegiPanjang + alasSegitiga + tinggiSegitiga

	fmt.Println(kalimat4)
}
