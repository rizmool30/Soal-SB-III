package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"manajemen-buku/models"
	"manajemen-buku/services"
)

// Scanner untuk membaca input pengguna
var scanner = bufio.NewScanner(os.Stdin)

// Fungsi membaca input
func input(teks string) string {
	fmt.Print(teks)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// Fungsi input data buku
func inputBuku() models.Buku {
	var buku models.Buku

	buku.Judul = input("Masukkan judul       : ")
	buku.Penulis = input("Masukkan penulis     : ")
	buku.Kategori = input("Masukkan kategori    : ")
	buku.Penerbit = input("Masukkan penerbit    : ")

	for {
		tahun := input("Masukkan tahun terbit: ")

		angka, err := strconv.Atoi(tahun)

		if err == nil {
			buku.TahunTerbit = angka
			break
		}

		fmt.Println("Tahun harus berupa angka!")
	}

	return buku
}

// Fungsi menampilkan menu
func tampilkanMenu() {
	fmt.Println("\n================================")
	fmt.Println("     PROGRAM MANAJEMEN BUKU")
	fmt.Println("================================")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Tampilkan Semua Buku")
	fmt.Println("3. Cari Buku Berdasarkan ID")
	fmt.Println("4. Ubah Data Buku")
	fmt.Println("5. Hapus Buku")
	fmt.Println("6. Keluar")
	fmt.Println("================================")
}

func main() {
	var daftarBuku []models.Buku

	for {
		tampilkanMenu()

		pilihan := input("Pilih menu (1-6): ")

		switch pilihan {

		case "1":
			fmt.Println("\n=== TAMBAH BUKU ===")

			buku := inputBuku()

			daftarBuku = services.TambahBuku(daftarBuku, buku)

			fmt.Println("\nBuku berhasil ditambahkan!")
			fmt.Println("ID Buku:", daftarBuku[len(daftarBuku)-1].ID)

		case "2":
			fmt.Println("\n=== DAFTAR SEMUA BUKU ===")

			services.TampilkanBuku(daftarBuku)

		case "3":
			fmt.Println("\n=== CARI BUKU ===")

			id := input("Masukkan ID buku: ")

			buku, ditemukan := services.CariBuku(daftarBuku, id)

			if ditemukan {
				fmt.Println("\nBuku ditemukan!")
				fmt.Println("ID          :", buku.ID)
				fmt.Println("Judul       :", buku.Judul)
				fmt.Println("Penulis     :", buku.Penulis)
				fmt.Println("Kategori    :", buku.Kategori)
				fmt.Println("Penerbit    :", buku.Penerbit)
				fmt.Println("Tahun Terbit:", buku.TahunTerbit)
			} else {
				fmt.Println("Buku tidak ditemukan!")
			}

		case "4":
			fmt.Println("\n=== UBAH DATA BUKU ===")

			id := input("Masukkan ID buku: ")

			_, ditemukan := services.CariBuku(daftarBuku, id)

			if !ditemukan {
				fmt.Println("Buku tidak ditemukan!")
				break
			}

			fmt.Println("\nMasukkan data buku yang baru:")

			dataBaru := inputBuku()

			daftarBuku, ditemukan = services.UbahBuku(
				daftarBuku,
				id,
				dataBaru,
			)

			if ditemukan {
				fmt.Println("Data buku berhasil diubah!")
			}

		case "5":
			fmt.Println("\n=== HAPUS BUKU ===")

			id := input("Masukkan ID buku: ")

			var ditemukan bool

			daftarBuku, ditemukan = services.HapusBuku(daftarBuku, id)

			if ditemukan {
				fmt.Println("Buku berhasil dihapus!")
			} else {
				fmt.Println("Buku tidak ditemukan!")
			}

		case "6":
			fmt.Println("\nTerima kasih telah menggunakan program!")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
