package view

import (
	"bufio"
	"fmt"
	"os"
	"thr/controller"
)

func MemberInsert() {
	var name, uname, pass, role string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("== Insert Member ==")
	fmt.Print("== Nama: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Print("== Username: ")
	if scanner.Scan() {
		uname = scanner.Text()
	}
	fmt.Print("== Password: ")
	if scanner.Scan() {
		pass = scanner.Text()
	}
	fmt.Print("== Role: ")
	if scanner.Scan() {
		role = scanner.Text()
	}
	cek := controller.InsertMember(name, uname, pass, role, 1)
	if cek {
		fmt.Println("== Data Berhasil Ditambahkan ==")
	} else {
		fmt.Println("== Data Gagal Ditambahkan ==")
	}
}

func MemberView() {
	members := controller.ReadAllMember()
	if members != nil {
		fmt.Println("-------------------Data Member-----------------")
		fmt.Println("| ID | Nama | Username | Password | Role | Status |")
		for _, Member := range members {
			fmt.Printf("| %d | %s | %s | %s | %s | %d |\n",
				Member.Id, Member.Nama, Member.Username, Member.Password, Member.Role, Member.Status)
		}
		fmt.Println("----------------------------------------------")
	} else {
		fmt.Println("== Data Tidak Ada ==")
	}
}
func MemberUpdate() {
	var id, status int
	var role, nama, username string
	fmt.Println("--- Id Member yang ingin di Update ---")
	fmt.Print("-- ID : ")
	fmt.Scan(&id)
	if controller.CheckMemberID(id) {
		fmt.Println("--- Data dengan Id", id, " Ditemukan ---")
		fmt.Println("--- Nama: ")
		fmt.Scanln(&nama)
		fmt.Println("--- Username: ")
		fmt.Scanln(&username)
		fmt.Print("--- Status[0/1] : ")
		fmt.Scan(&status)
		fmt.Print("--- Role[A/M] : ")
		fmt.Scan(&role)
		controller.UpdateMember(id, nama, username, role, status)
		fmt.Println("Data Buku Berhasil di Update!!")

	} else {
		fmt.Println("Buku dengan Id", id, "Tidak ditemukan")
	}

}
func MemberSearch() {
	var id int
	fmt.Println("== Search Member ==")
	fmt.Print("== ID : ")
	fmt.Scan(&id)
	members := controller.SearchMember(id)
	if members != nil {
		fmt.Println("-------------------Data Member-----------------")
		fmt.Println("| ID | Nama | Username | Passsword | Role | Status |  ")
		for _, Member := range members {
			fmt.Printf("| %d | %s | %s | %s | %s | %d |\n",
				Member.Id, Member.Nama, Member.Username, Member.Password, Member.Role, Member.Status)
		}
		fmt.Println("----------------------------------------------")
	} else {
		fmt.Println("Id member", id, "Tidak ditemukan")
	}
}
func MemberDelete() {
	var id int
	fmt.Println("== Delete Member ==")
	fmt.Print("== ID : ")
	fmt.Scan(&id)
	if controller.CheckMemberID(id) {
		controller.DeleteMember(id)
		fmt.Println("Data Member Berhasil di Hapus!!")
	} else {
		fmt.Println("Id member", id, "Tidak ditemukan")
	}
}
