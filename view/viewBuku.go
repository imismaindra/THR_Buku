package view

import (
	"fmt"
	"thr/controller"
)

func BukuInsert() {
	var judul, pengarang, penerbit, tahun string
	fmt.Println("== Insert Buku ==")
	fmt.Print("== Judul : ")
	fmt.Scan(&judul)
	fmt.Print("== Pengarang : ")
	fmt.Scan(&pengarang)
	fmt.Print("== Penerbit : ")
	fmt.Scan(&penerbit)
	fmt.Print("== Tahun : ")
	fmt.Scan(&tahun)
	cek := controller.InsertBuku(judul, pengarang, penerbit, tahun)
	if cek == true {
		fmt.Println("== Data Berhasil Ditambahkan ==")
	} else {
		fmt.Println("== Data Gagal Ditambahkan ==")
	}

}
func BukuView() {
	books := controller.ViewBuku()
	if books != nil {
		fmt.Println("-------------------Data Buku-----------------")
		fmt.Println("| ID | Judul | Pengarang | Penerbit | Tahun |")

		for _, book := range books {
			fmt.Printf("| %d | %s | %s | %s | %s |\n",
				book.Id, book.Judul, book.Pengarang, book.Penerbit, book.Tahun)
		}
		fmt.Println("_____________________________________________")

	} else {
		fmt.Println("== Data Buku Kosong ==")
	}
}
func BukuSearch() {
	var id int
	fmt.Println("== Search Buku ==")
	fmt.Print("== ID : ")
	fmt.Scan(&id)
	books := controller.SearchBuku(id)
	if books != nil {
		fmt.Println("-------------------Data Buku-----------------")
		fmt.Println("| ID | Judul | Pengarang | Penerbit | Tahun |")
		for _, book := range books {

			fmt.Printf("| %d | %s | %s | %s | %s |\n",
				book.Id, book.Judul, book.Pengarang, book.Penerbit, book.Tahun)
		}
	} else {
		fmt.Println("Id buku", id, "Tidak ditemukan")
	}
}
func BukuUpdate() {
	var id int
	var judul, pengarang, penerbit, tahun string
	fmt.Println("--- Id Buku yang ingin di Update ---")
	fmt.Print("-- ID : ")
	fmt.Scan(&id)
	if controller.CheckBukuID(id) {
		fmt.Println("--- Data dengan Id", id, " Ditemukan ---")
		fmt.Print("-- Judul : ")
		fmt.Scan(&judul)
		fmt.Print("-- Pengarang : ")
		fmt.Scan(&pengarang)
		fmt.Print("-- Penerbit : ")
		fmt.Scan(&penerbit)
		fmt.Print("-- Tahun : ")
		fmt.Scan(&tahun)
		controller.UpdateBuku(id, judul, pengarang, penerbit, tahun)
		fmt.Println("Data Buku Berhasil di Update!!")

	} else {
		fmt.Println("Buku dengan Id", id, "Tidak ditemukan")
	}

}
func BukuDelete() {
	var id int
	fmt.Println("--- Delete Buku ---")
	fmt.Print("-- ID : ")
	fmt.Scan(&id)
	if controller.CheckBukuID(id) {
		controller.DeleteBuku(id)
		fmt.Println("Data Buku Berhasil di Hapus")

	} else {
		fmt.Println("buku dengan id =", id, ", Tidak ditemukan")
	}
}
