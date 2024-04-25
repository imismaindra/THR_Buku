package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"thr/controller"
	"thr/handler"
	"thr/model"
	"thr/view"
)

func main_program() {
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
			pilih = scanner.Text()
		}
		switch pilih {
		case "1":
			view.BukuInsert()
			break
		case "2":
			view.BukuUpdate()
			break
		case "3":
			view.BukuDelete()
			break

		case "4":
			view.BukuSearch()
			break
		case "5":
			view.BukuView()
			break
		case "6":
			return
		default:
			fmt.Println("Pilihan tidak ada")

		}

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
	//test search member
	fmt.Println(controller.InsertMember("indra", "Casanova", "12345", "A", 3))
	// main_program()
	// webProgram()
}
