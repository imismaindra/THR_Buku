// view.go

package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"thr/controller"
	"thr/node"
)

func InsertPeminjaman(nama string, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("== Insert Peminjaman ==")
	var member node.Member
	fmt.Println("ID Member:", id)
	fmt.Println("Nama Member:", nama)
	member.Nama = nama
	fmt.Print("Alamat:")
	if scanner.Scan() {
		member.Alamat = scanner.Text()
	}
	fmt.Print("NoTelp:")
	if scanner.Scan() {
		member.NoTelp = scanner.Text()
	}
	var bukuIDs []int
	bookIDMap := make(map[int]bool) // Map to track book IDs

	for {
		fmt.Print("Masukkan ID Buku ( Enter untuk menyelesaikan):")
		var idBukuStr string
		if scanner.Scan() {
			idBukuStr = scanner.Text()
		}
		if idBukuStr == "" {
			break
		}
		idBuku, err := strconv.Atoi(idBukuStr)
		if err != nil {
			fmt.Print("Error: Enter a valid integer for book ID.")

			continue
		}
		if bookIDMap[idBuku] {
			fmt.Println("You have already entered this book ID.")
			continue
		}
		if !controller.CheckBukuID(idBuku) {
			fmt.Println("Buku dengan Id", idBuku, "Tidak Ada. Tolong Masukkan Id Buku yang Valid.")
			continue
		}
		bukuIDs = append(bukuIDs, idBuku)
		bookIDMap[idBuku] = true // Mark book ID as entered\
		controller.CheckStokBuku(idBuku)

	}

	controller.InsertPeminjaman(member, bukuIDs)

	fmt.Println("== Pengajuan Peminjaman Buku Sedang di Tinjau Admin ==")
}

func DisplayAllPeminjaman() {
	fmt.Println("== List Peminjaman ==")
	peminjamanList := controller.GetAllPeminjaman()

	for _, peminjaman := range peminjamanList {
		fmt.Printf("ID Peminjaman: %d\n", peminjaman.IdPeminjaman)
		fmt.Println("Informasi Peminjam:")
		fmt.Printf("- IdMember: %s\n", peminjaman.Member.IdMember)
		fmt.Printf("- Nama: %s\n", peminjaman.Member.Nama)
		fmt.Printf("- Alamat: %s\n", peminjaman.Member.Alamat)
		fmt.Printf("- Nomor Telepon: %s\n", peminjaman.Member.NoTelp)
		fmt.Println("Detail Peminjaman:")
		for _, detail := range peminjaman.DetailPeminjaman {
			fmt.Printf("- ID Buku: %d\n", detail.IdBuku)
			fmt.Printf("- Judul Buku: %s\n", detail.Jdl)
		}
		if peminjaman.Status == 0 {
			fmt.Printf("Status Peminjaman: %s\n", "Ditinjau Admin")
		} else if peminjaman.Status == 1 {
			fmt.Printf("Status Peminjaman: %s\n", "Disetujui Admin")
		} else if peminjaman.Status == 2 {
			fmt.Printf("Status Peminjaman: %s\n", "Ditolak Admin")
		} else if peminjaman.Status == 3 {
			fmt.Printf("Status Peminjaman: %s\n", "Dikembalikan")
		}
		fmt.Println("Batas pengembalian: ", peminjaman.BackAt)
		fmt.Println("----------------------------")
	}
}
func UpdateStsPeminjaman() {
	fmt.Println("== Update Peminjaman ==")
	var idPeminjaman int
	fmt.Print("Masukkan ID Peminjaman: ")
	fmt.Scanln(&idPeminjaman)
	if controller.CheckPeminjamanID(idPeminjaman) {
		fmt.Println("Status: \n0 -> Ditinjau\n1 -> Disetujui/Dipinjam\n2 -> Ditolak\n3 -> Selesai/Dikembalikan")
		fmt.Print("Status Peminjaman:")
		var status int
		fmt.Scanln(&status)
		controller.UpStatusPeminjaman(idPeminjaman, status)

		fmt.Println("Status Peminjaman Berhasil Diupdate")
	} else {
		fmt.Println("Id peminjaman tidak ditemukan")
	}
}
