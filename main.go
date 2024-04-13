package main

import (
	"fmt"
	"thr/model"
	"thr/view"
)

func main_program() {
	var pilih int
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
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			view.BukuInsert()
			break
		case 2:
			view.BukuUpdate()
			break
		case 3:
			view.BukuDelete()
			break

		case 4:
			view.BukuSearch()
			break
		case 5:
			view.BukuView()
			break
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak ada")

		}

	}
}
func main() {
	model.BukuInsert("Sangkuriang", " Andi Harahap", "Gramedia", "2002")
	model.BukuInsert("Timun Emas", " Mustakim", "JKutBook", "2004")
	model.BukuInsert("Merah Putih", " Rudolf", "Kompas", "1989")
	main_program()
}
