package main

import (
	"fmt"
	"os"
	"strconv"
)

// Struct untuk merepresentasikan teman sekelas
type Classmate struct {
	Name       string
	Address    string
	Occupation string
	Motivation string
}

func main() {
	// Slice yang berisi data teman-teman sekelas
	classmates := []Classmate{
		{"Wahyu", "Jakarta, Indonesia", "Menganggur", "Ingin uang"},
		{"Zuck", "NYC, USA", "Tidur", "Coba-coba aja"},
		{"Rasmus L", "Sabang, Indonesia", "IT Support", "Ingin ganti job"},
		{"Torvald", "Batam, Indonesia", "Petani Batam", "Ingin bikin jasa ekspedisi bebas cukai"},
		{"Larry Page", "Merauke, Indonesia", "Top Global Gusion", "Ingin membuat cheat"},
	}

	// Mengambil input dari command line arguments (os.Args)
	query := os.Args

	// Cek apakah jumlah argument kurang dari dua, karena argument pertama adalah nama program itu sendiri
	if len(query) < 2 {
		fmt.Println(`
      Index belum dimasukkan.
      Cara penggunaan yang benar seperti ini:

      go run main.go 0
      `)
		// Menghentikan program agar tidak melanjutkan eksekusi ke bawah
		return
	}

	// Mengambil data berdasarkan index yang diinputkan (query[1])
	result, err := getClassmate(classmates, query[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Menampilkan informasi yang berhasil diambil
	fmt.Printf("Nama: %v\nAlamat: %v\nPekerjaan: %v\nMotivasi: %v\n",
		result.Name, result.Address, result.Occupation, result.Motivation)
}

// Function untuk mendapatkan teman sekelas berdasarkan index
func getClassmate(classmates []Classmate, indexArg string) (Classmate, error) {
	// Mengubah string dari os.Args menjadi integer
	index, err := strconv.Atoi(indexArg)
	if err != nil {
		return Classmate{}, fmt.Errorf("index hanya boleh angka")
	}

	// Menangani kasus jika index terlalu kecil atau terlalu besar
	if index < 0 || index >= len(classmates) {
		return Classmate{}, fmt.Errorf("index diluar jangkauan")
	}

	// Return data yang sesuai dengan index
	return classmates[index], nil
}
