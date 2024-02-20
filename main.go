package main

import (
	"fmt"
	"os"
)

type peserta struct {
	ID        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func displayData(nama string, listPeserta []peserta) {
	available := false

	for _, peserta := range listPeserta {
		if peserta.nama == nama {
			available = true
			fmt.Printf("Data peserta:\n")
			fmt.Printf("ID peserta:%d\n", peserta.ID)
			fmt.Printf("Nama: %s\n", peserta.nama)
			fmt.Printf("Alamat: %s\n", peserta.alamat)
			fmt.Printf("Pekerjaan: %s\n", peserta.pekerjaan)
			fmt.Printf("Alasan memilih kelas Golang: %s\n", peserta.alasan)
			break
		}
	}

	if !available {
		fmt.Printf("Data dengan ID %s tidak tersedia.\n", nama)
	}
}

func main() {

	listPeserta := []peserta{
		{1, "Fitri", "Jakarta", "Programmer", "Memperbanyak skill"},
		{2, "Dion", "Jakarta", "Mahasiswa", "Memperbanyak skill"},
	}
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Gunakan perintah: go run main.go [nama peserta]")
		return
	}

	idPeserta := args[1]

	displayData(idPeserta, listPeserta)

}
