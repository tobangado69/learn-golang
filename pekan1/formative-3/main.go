package main

import (
	"fmt"
	"strings"
)

func main() {

	// 	Yang paling sering dipakai (dan wajib kamu hafal)
	// Dari dokumentasi resmi:
	// %s → string
	// %d → integer (decimal)
	// %f → float
	// %v → value (default, paling aman)
	// %T → tipe data
	// %% → cetak tanda %

	// Soal 1
	kalimat := "halo halo bandung"
	angka := 2021

	newKaliamt := strings.Replace(kalimat, "halo", "hi", -1)
	fmt.Printf("\"%s\" - %d\n", newKaliamt, angka)

	// Soal 2
	nilaiJohn := 80
	nilaiDoe := 50

	fmt.Println("Indeks John:", indeksNilai(nilaiJohn))
	fmt.Println("Indeks Doe:", indeksNilai(nilaiDoe))

	// Soal 3
	tanggal := 17
	bulan := 8
	tahun := 1945

	var namaBulan string

	switch bulan {
	case 1:
		namaBulan = "Januari"
		break
	case 2:
		namaBulan = "February"
		break
	case 3:
		namaBulan = "Maret"
		break
	case 4:
		namaBulan = "April"
		break
	case 5:
		namaBulan = "Mei"
		break
	case 6:
		namaBulan = "Juni"
		break
	case 7:
		namaBulan = "Juli"
		break
	case 8:
		namaBulan = "Agustus"
		break
	case 9:
		namaBulan = "September"
		break
	case 10:
		namaBulan = "Oktober"
		break
	case 11:
		namaBulan = "November"
		break
	case 12:
		namaBulan = "Desember"
		break
	default:
		namaBulan = "Nama Bulan Tidak valid"
	}

	fmt.Println(tanggal, namaBulan, tahun)

	// Soal 4
	year := 1945
	generasi := " "
	name := "Iqbal"

	if year >= 1944 && year <= 1964 {
		generasi = "Baby Boomer"
	} else if year >= 1965 && year <= 1979 {
		generasi = "X"
	} else if year >= 1980 && year <= 1994 {
		generasi = "Y (Milenials)"
	} else if year >= 1995 && year <= 2015 {
		generasi = "Z"
	} else if year > 2015 {
		generasi = "Alpha"
	} else {
		generasi = "Maaf masukkan tahun kelahiran yang valid"
	}
	fmt.Printf("%s lahir pada tahun %d dan merupakan generasi %s\n", name, year, generasi)

	// Soal 5
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Println(i, "I love coding")
		} else if i%2 == 2 {
			fmt.Println(i, "Berkualitas")
		} else {
			fmt.Println(i, "Santai")
		}
	}

}

// Soal 2
func indeksNilai(nilai int) string {
	if nilai >= 80 && nilai <= 100 {
		return "A"
	} else if nilai >= 70 && nilai < 80 {
		return "B"
	} else if nilai >= 60 && nilai < 70 {
		return "C"
	} else if nilai >= 50 && nilai < 60 {
		return "D"
	} else if nilai < 50 {
		return "E"
	} else {
		return "Harap masukkan nilai yang valid"
	}
}
