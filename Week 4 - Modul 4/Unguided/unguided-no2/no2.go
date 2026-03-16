package main

import "fmt"

const Phi = 3.14

func hitungPersegi(sisi int) {
	var luas int
	var keliling int
	luas = sisi * sisi
	keliling = 4 * sisi
	fmt.Println("Luas persegi : ", luas)
	fmt.Println("Keliling persegi : ", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	var luas int
	var keliling int
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	fmt.Println("Luas persegi panjang : ", luas)
	fmt.Println("Keliling persegi panjang : ", keliling)
}

func hitungLingkaran(jarijari float64) {
	var luas float64
	var keliling float64
	luas = Phi * jarijari * jarijari
	keliling = 2 * Phi * jarijari
	fmt.Println("Luas lingkaran : ", luas)
	fmt.Println("Keliling lingkaran : ", keliling)
}

func main() {
	var pilihan int
	var sisi int
	var panjang int
	var lebar int
	var jarijari float64

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("\nMasukkan sisi : ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		fmt.Print("\nMasukkan panjang : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar : ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)
	case 3:
		fmt.Print("\nMasukkan jari-jari : ")
		fmt.Scan(&jarijari)
		hitungLingkaran(jarijari)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}