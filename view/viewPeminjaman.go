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

func InsertPeminjaman() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("== Insert Peminjaman ==")

	var member node.Member
	fmt.Println("Enter your name:")
	if scanner.Scan() {
		member.Nama = scanner.Text()
	}
	fmt.Println("Enter your address:")
	if scanner.Scan() {
		member.Alamat = scanner.Text()
	}
	fmt.Println("Enter your phone number:")
	if scanner.Scan() {
		member.NoTelp = scanner.Text()
	}

	var bukuIDs []int
	bookIDMap := make(map[int]bool) // Map to track book IDs

	for {
		fmt.Println("Enter book ID to borrow (press Enter to finish):")
		var idBukuStr string
		if scanner.Scan() {
			idBukuStr = scanner.Text()
		}
		if idBukuStr == "" {
			break
		}
		idBuku, err := strconv.Atoi(idBukuStr)
		if err != nil {
			fmt.Println("Error: Enter a valid integer for book ID.")
			continue
		}
		if bookIDMap[idBuku] {
			fmt.Println("You have already entered this book ID.")
			continue
		}
		bukuIDs = append(bukuIDs, idBuku)
		bookIDMap[idBuku] = true // Mark book ID as entered
	}

	controller.InsertPeminjaman(member, bukuIDs)

	fmt.Println("== Loan Successfully Saved ==")
}
func DisplayAllPeminjaman() {
	fmt.Println("== List of All Loans ==")
	peminjamanList := controller.GetAllPeminjaman()

	for _, peminjaman := range peminjamanList {
		fmt.Printf("ID Peminjaman: %d\n", peminjaman.IdPeminjaman)
		fmt.Println("Member Information:")
		fmt.Printf("- Nama: %s\n", peminjaman.Member.Nama)
		fmt.Printf("- Alamat: %s\n", peminjaman.Member.Alamat)
		fmt.Printf("- Nomor Telepon: %s\n", peminjaman.Member.NoTelp)
		fmt.Println("Detail Peminjaman:")
		for _, detail := range peminjaman.DetailPeminjaman {
			fmt.Printf("- ID Buku: %d\n", detail.IdBuku)
			fmt.Printf("- Judul Buku: %s\n", detail.Jdl)
		}
		fmt.Println("----------------------------")
	}
}
