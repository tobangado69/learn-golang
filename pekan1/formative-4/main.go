package main

import "fmt"

func main() {

	// Soal 1
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// Soal 2
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Print("[")

	for i := 2; i < len(kalimat); i++ {
		fmt.Print(kalimat[i])

		if i != len(kalimat)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")

	// Soal 3
	sayuran := []string{}
	sayuran = append(
		sayuran,
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	)
	for i, item := range sayuran {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// Soal 4
	satuan := map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	volumeSatuan := 1

	for _, nilai := range satuan {
		volumeSatuan *= nilai
	}

	fmt.Println("volume balok", volumeSatuan)

	// Soal 5
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
}

func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}
