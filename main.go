package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"thr/controller"
	"thr/handler"
	"thr/model"
	"thr/view"
)

func MenuBuku() {
	scanner := bufio.NewScanner(os.Stdin)
	var pilih string
	for {

		fmt.Println("Menu:")
		fmt.Println("1. Insert")
		fmt.Println("2. Update")
		fmt.Println("3. Delete")
		fmt.Println("4. Search")
		fmt.Println("5. Read All")
		fmt.Println("6. EXIT")
		fmt.Println()
		fmt.Print("Pilih: ")
		if scanner.Scan() {
			pilih = strings.TrimSpace(scanner.Text()) // Trim any leading or trailing whitespace
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}

		switch pilih {
		case "1":
			view.BukuInsert()
			//view.MemberInsert()
		case "2":
			view.BukuUpdate()
			//view.MemberUpdate()
			break
		case "3":
			view.BukuDelete()
			break
		case "4":
			view.BukuSearch()
			//view.MemberSearch()
		case "5":
			view.BukuView()
			//view.MemberView()
		case "6":
			main_program()
		default:
			fmt.Println("Pilihan tidak ada")
		}
		scanner.Scan()

	}
}
func MenuMember() {
	scanner := bufio.NewScanner(os.Stdin)
	var pilih string
	for {

		fmt.Println("Menu Member:")
		fmt.Println("1. Insert Member")
		fmt.Println("2. Update Member")
		fmt.Println("3. Delete Member")
		fmt.Println("4. Search Member")
		fmt.Println("5. Read All Member")
		fmt.Println("6. Kembali")
		fmt.Println()
		fmt.Print("Pilih: ")
		if scanner.Scan() {
			pilih = strings.TrimSpace(scanner.Text()) // Trim any leading or trailing whitespace
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}

		switch pilih {
		case "1":

			view.MemberInsert()
		case "2":
			view.MemberUpdate()
			break
		case "3":
			view.MemberDelete()
			break
		case "4":
			view.MemberSearch()
		case "5":
			view.MemberView()
		case "6":
			main_program()
		default:
			fmt.Println("Pilihan tidak ada")
		}
		scanner.Scan()

	}
}
func main_program() {
	scanner := bufio.NewScanner(os.Stdin)
	var pilih string
	for {
		fmt.Println("Perpustakaan GWE")
		fmt.Println("Menu:")
		fmt.Println("1. Buku")
		fmt.Println("2. Member")
		fmt.Println("3. Peminjaman")
		fmt.Println("4. Exit")
		fmt.Println()
		fmt.Print("Pilih: ")
		if scanner.Scan() {
			pilih = strings.TrimSpace(scanner.Text()) // Trim any leading or trailing whitespace
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}
		switch pilih {
		case "1":
			MenuMember()
		case "2":
			MenuBuku()
		case "3":
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak ada")
		}
		scanner.Scan()
	}

}

func webProgram() {
	http.HandleFunc("/", handler.ViewHandler)
	http.HandleFunc("/insert", handler.BukuInsertHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("'localhost:8080'")
}
func main() {
	model.BukuInsert("Sangkuriang", " Andi Harahap", "Gramedia", "2002")
	model.BukuInsert("Timun Emas", " Mustakim", "JKutBook", "2004")
	model.BukuInsert("Merah Putih", " Rudolf", "Kompas", "1989")
	//test insert member
	model.InsertMember("indra", "Casanova", "12345", "A", 1)
	model.InsertMember("Firda", "PPP", "jagonyaAyam", "M", 1)
	model.InsertMember("Rohman Ayai", "Rhm", "12345", "M", 0)
	fmt.Println(model.ReadAllMember())
	fmt.Println(controller.Login("Casanova", "12345"))
	//test search member
	// fmt.Println(controller.InsertMember("Mulira", "Vaco", "12345", "A", 1))
	// fmt.Println(controller.UpdateMember(1, "M", 0))
	// fmt.Println(controller.ReadAllMember())
	view.VLogin()
	// main_program()
	//view.BukuUpdate()
	// webProgram()
}
