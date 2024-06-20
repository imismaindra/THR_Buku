// view.go

package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"thr/controller"
	"thr/node"
)

func InsertPeminjaman(nama string, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	var member node.Member
	member.User.Nama = nama
	member.User.Id = id

	fmt.Println("== Insert Peminjaman ==")
	fmt.Println("== ID Member:", id)
	fmt.Println("== Nama Member:", nama)
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
			fmt.Print("Kesalahan: Masukkan angka yang valid untuk ID buku.")

			continue
		}
		if bookIDMap[idBuku] {
			fmt.Println("Anda telah memasukkan ID buku ini.")
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
}

func DisplayAllPeminjaman() {
	fmt.Println("== History Peminjaman ==")
	peminjamanList := controller.GetAllPeminjaman()
	if peminjamanList == nil {
		fmt.Println("Belum Ada Peminjaman")
	} else {

		for _, peminjaman := range peminjamanList {

			fmt.Printf("== ID Peminjaman: %d\n", peminjaman.IdPeminjaman)
			fmt.Println("== Informasi Peminjam:")
			fmt.Printf("- IdMember: %d\n", peminjaman.Member.User.Id)
			fmt.Printf("- Nama: %s\n", peminjaman.Member.User.Nama)
			fmt.Printf("- Alamat: %s\n", peminjaman.Member.Alamat)
			fmt.Printf("- Nomor Telepon: %s\n", peminjaman.Member.NoTelp)
			fmt.Println("Detail Peminjaman:")
			for _, detail := range peminjaman.DetailPeminjaman {
				fmt.Printf("- ID Buku: %d\n", detail.IdBuku)
				fmt.Printf("- Judul Buku: %s\n", detail.Jdl)
			}
			if peminjaman.Status == 0 {
				fmt.Printf("== Status Peminjaman: %s\n", "Diajukan")
			} else if peminjaman.Status == 1 {
				fmt.Printf("== Status Peminjaman: %s\n", "Disetujui")
			} else if peminjaman.Status == 2 {
				fmt.Printf("== Status Peminjaman: %s\n", "Ditolak")
			} else if peminjaman.Status == 3 {
				fmt.Printf("== Status Peminjaman: %s\n", "Dikembalikan")
			}
			if peminjaman.Status != 0 && peminjaman.Status != 2 {
				fmt.Println("== Batas pengembalian: ", peminjaman.BackAt)
			}
			fmt.Println("----------------------------")
			fmt.Println()

		}
	}
}
func DisplayAllPeminjamanByUser(idUser int) {
	fmt.Println("==== History Peminjaman ====")
	peminjamanList := controller.GetAllPeminjaman()
	if peminjamanList == nil {
		fmt.Println("Belum Ada Peminjaman")
	} else {

		for _, peminjaman := range peminjamanList {
			if peminjaman.User.Id == idUser {
				fmt.Printf("== ID Peminjaman: %d\n", peminjaman.IdPeminjaman)
				fmt.Println("== Informasi Peminjam:")
				fmt.Printf("- IdMember: %d\n", peminjaman.Member.User.Id)
				fmt.Printf("- Nama: %s\n", peminjaman.Member.User.Nama)
				fmt.Printf("- Alamat: %s\n", peminjaman.Member.Alamat)
				fmt.Printf("- Nomor Telepon: %s\n", peminjaman.Member.NoTelp)
				fmt.Println("Detail Peminjaman:")
				for _, detail := range peminjaman.DetailPeminjaman {
					fmt.Printf("- ID Buku: %d\n", detail.IdBuku)
					fmt.Printf("- Judul Buku: %s\n", detail.Jdl)
				}
				if peminjaman.Status == 0 {
					fmt.Printf("== Status Peminjaman: %s\n", "Diajukan")
				} else if peminjaman.Status == 1 {
					fmt.Printf("== Status Peminjaman: %s\n", "Disetujui")
				} else if peminjaman.Status == 2 {
					fmt.Printf("== Status Peminjaman: %s\n", "Ditolak")
				} else if peminjaman.Status == 3 {
					fmt.Printf("== Status Peminjaman: %s\n", "Dikembalikan")
				}
				if peminjaman.Status != 0 || peminjaman.Status != 2 {
					fmt.Println("== Batas pengembalian: ", peminjaman.BackAt)
				}
				fmt.Println("----------------------------")
				fmt.Println()
			}

		}
	}

}
func UpdateStsPeminjaman() {
	fmt.Println("== Update Peminjaman ==")
	var idPeminjaman int
	fmt.Print("Masukkan ID Peminjaman: ")
	fmt.Scanln(&idPeminjaman)
	peminjaman, _ := controller.CheckPeminjamanID(idPeminjaman)
	if peminjaman {
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
func VreturnBook(id int) {
	scanner := bufio.NewScanner(os.Stdin)
	var peminjamanID int
	var bukuIDs []int

	fmt.Print("Masukkan ID Peminjaman: ")
	fmt.Scanln(&peminjamanID)
	checkid, _ := controller.CheckPeminjamanID(peminjamanID)
	if !checkid {
		fmt.Println("Id Peminjaman Tidak valid")
		return
	} else {

		fmt.Print("Masukkan ID Buku yang dikembalikan (pisahkan dengan koma): ")
		if scanner.Scan() {
			bukuIDStr := strings.Split(scanner.Text(), ",")
			for _, idStr := range bukuIDStr {
				id, err := strconv.Atoi(strings.TrimSpace(idStr))
				if err == nil {
					bukuIDs = append(bukuIDs, id)
				} else {
					fmt.Printf("ID buku tidak valid: %s\n", idStr)
				}
			}
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}

		success := controller.CreturnBook(peminjamanID, id, bukuIDs)
		if !success {
			fmt.Println("Pengembalian buku gagal.")
		} else {
			fmt.Println("Pengembalian buku berhasil.")
		}
	}
}
