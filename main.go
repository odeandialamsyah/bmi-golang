package main

import (
	"fmt"
)

func main() {
	fmt.Println("Selamat Datang di Sistem BMI")
	fmt.Println("-----------------------------")

	// Memulai perhitungan BMI
	menghitungBMI()

	// Loop perhitungan BMI jika pengguna ingin mengulang
	for perulangan() {
		menghitungBMI()
	}

	fmt.Println("Terima Kasih Telah Menggunakan Sistem BMI!")
}

func menghitungBMI() {
	var beratBadan, tinggiBadan float64

	// Meminta input berat badan
	for {
		fmt.Print("Masukkan Berat Badan Anda (kg): ")
		fmt.Scanln(&beratBadan)
		if beratBadan > 0 {
			break
		}
		fmt.Println("Berat badan harus lebih dari 0. Coba lagi.")
	}

	// Meminta input tinggi badan
	for {
		fmt.Print("Masukkan Tinggi Badan Anda (m): ")
		fmt.Scanln(&tinggiBadan)
		if tinggiBadan > 0 {
			break
		}
		fmt.Println("Tinggi badan harus lebih dari 0. Coba lagi.")
	}

	// Menghitung BMI
	bmi := beratBadan / (tinggiBadan * tinggiBadan)

	// Menentukan kategori BMI
	fmt.Printf("Hasil BMI Anda: %.2f\n", bmi)
	switch {
	case bmi < 18.5:
		fmt.Println("Kategori: Berat Badan Kurang")
	case bmi >= 18.5 && bmi < 24.9:
		fmt.Println("Kategori: Berat Badan Normal")
	case bmi >= 25 && bmi < 29.9:
		fmt.Println("Kategori: Berat Badan Berlebih")
	default:
		fmt.Println("Kategori: Obesitas")
	}
}

func perulangan() bool {
	for {
		var pilihan string
		fmt.Print("Apakah Anda Ingin Menghitung Ulang? (y/n): ")
		fmt.Scanln(&pilihan)

		// Validasi input
		if pilihan == "y" || pilihan == "Y" {
			return true
		} else if pilihan == "n" || pilihan == "N" {
			return false
		}

		// Jika input tidak valid
		fmt.Println("Input tidak valid. Masukkan 'y' untuk ya atau 'n' untuk tidak.")
	}
}
