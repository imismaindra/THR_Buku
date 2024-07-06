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
	"thr/node"
	"thr/view"
)

func MenuBuku(nama string, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	var pilih string
	for {
		fmt.Println("======= Menu Buku =======")
		fmt.Println("== Menu:")
		fmt.Println("== 1. Insert")
		fmt.Println("== 2. Update")
		fmt.Println("== 3. Delete")
		fmt.Println("== 4. Search")
		fmt.Println("== 5. ReadAll")
		fmt.Println("== 6. Kembali")
		fmt.Println("========================")
		fmt.Println()
		fmt.Print("== Pilih Menu: ")
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
			main_program(nama, id)
		default:
			fmt.Println("Pilihan tidak ada")
		}
		scanner.Scan()

	}
}
func MenuMember(nama string, id int) {
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
			main_program(nama, id)
		default:
			fmt.Println("Pilihan tidak ada")
		}
		scanner.Scan()

	}
}
func MenuPeminjaman(nama string, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("==== Menu Peminjaman =====")
		fmt.Println("== Menu:")
		fmt.Println("== 1. Insert Peminjaman")
		fmt.Println("== 2. Update Peminjaman")
		fmt.Println("== 3. Return Peminjaman")
		fmt.Println("== 4. Search Peminjaman")
		fmt.Println("== 5. History Peminjaman")
		fmt.Println("== 6. Kembali")
		fmt.Print("Pilih: ")

		if !scanner.Scan() {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}

		pilih := strings.TrimSpace(scanner.Text())

		switch pilih {
		case "1":
			view.InsertPeminjaman(nama, id)
		case "2":
			view.UpdateStsPeminjaman()
		case "3":
			view.VreturnBook(id)
		case "4":
			view.MemberSearch()
		case "5":
			view.DisplayAllPeminjaman()
		case "6":
			main_program(nama, id)
			return
		default:
			fmt.Println("Pilihan tidak ada")
		}
	}
}

func main_program(nama string, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	var pilih string
	for {
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
			MenuBuku(nama, id)
		case "2":
			MenuMember(nama, id)
		case "3":
			MenuPeminjaman(nama, id)
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak ada")
		}
		scanner.Scan()
	}

}
func VLogin() (string, string, int) {
	var uname, password string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("=== LOGIN ===")
	fmt.Print("=== Username : ")
	if scanner.Scan() {
		uname = strings.TrimSpace(scanner.Text())
	} else {
		fmt.Println("Error reading input:", scanner.Err())
		return "", "", 0
	}
	fmt.Print("=== Password: ")
	if scanner.Scan() {
		password = strings.TrimSpace(scanner.Text())
	} else {
		fmt.Println("Error reading input:", scanner.Err())
		return "", "", 0
	}
	role, name, id := controller.Login(uname, password)
	return role, name, id
}
func methodOverrideMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if overrideMethod := r.FormValue("_method"); overrideMethod != "" {
				r.Method = strings.ToUpper(overrideMethod)
			}
		}
		next.ServeHTTP(w, r)
	}
}
func WebBukuHendler() {
	http.HandleFunc("/buku", handler.BukuReadAllHandler)
	http.HandleFunc("/insertbuku", handler.BukuInsertHandler)
	http.HandleFunc("/updatebuku/", handler.EditBukuHandler)   // Menampilkan form edit
	http.HandleFunc("/updatebuku", handler.BukuUpdateHandler)  // Menangani permintaan PUT untuk update buku
	http.HandleFunc("/buku/delete", handler.BukuDeleteHandler) // Menangani permintaan DELETE

}
func WebMemberHendler() {
	http.HandleFunc("/memberupdate/", methodOverrideMiddleware(handler.EditMemberHandler))
	http.HandleFunc("/member", handler.MemberReadAllHandler)
	http.HandleFunc("/member/delete", methodOverrideMiddleware(handler.MemberDeleteHandler))
	http.HandleFunc("/member/insert", methodOverrideMiddleware(handler.MemberInsertHandler))
}
func WebPeminjamanHendler() {
	http.HandleFunc("/peminjaman", handler.PeminjamanReadAllHandler)
	http.HandleFunc("/peminjaman/update/", handler.PeminjamanUpdateStatusHandler)
	http.HandleFunc("/peminjaman/detail/", handler.PeminjamanDetailHandler) // Tambahkan ini
	http.HandleFunc("/addToCart", handler.AddToCartHandler)
	http.HandleFunc("/checkout", handler.CheckoutHandler)
	http.HandleFunc("/history", handler.PeminjamanAndReturnHandler)
	// http.HandleFunc("/return", handler.ReturnBookHandler)

}

func webProgram() {
	http.HandleFunc("/", handler.ViewHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/dashboard", handler.DashboardHandler)
	http.HandleFunc("/store", handler.StoreHandler)
	WebBukuHendler()
	WebMemberHendler()
	WebPeminjamanHendler()
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func AdminMenu(nama string, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	var pilih string
	for {
		fmt.Println("Menu Admin:")
		fmt.Println("1. Buku")
		fmt.Println("2. Member")
		fmt.Println("3. Peminjaman")
		fmt.Println("4. Exit")
		fmt.Println()
		fmt.Print("Pilih: ")
		if scanner.Scan() {
			pilih = strings.TrimSpace(scanner.Text())
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}
		switch pilih {
		case "1":
			MenuBuku(nama, id)
		case "2":
			MenuMember(nama, id)
		case "3":
			MenuPeminjaman(nama, id)
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak ada")
		}
		scanner.Scan()
	}
}
func tester() {
	//test search member
	// fmt.Println(controller.InsertMember("Mulira", "Vaco", "12345", "A", 1))
	// fmt.Println(controller.UpdateMember(1, "M", 0))
	// fmt.Println(controller.ReadAllMember())
	model.BukuInsert("Sangkuriang", "Andi Harahap", "Gramedia", "2002", 10)
	model.BukuInsert("Timun Emas", "Mustakim", "JKutBook", "2004", 12)
	model.BukuInsert("Merah Putih", "Rudolf", "Kompas", "1989", 2)
	model.BukuInsert("Laskar Pelangi", "Andrea Hirata", "Bentang Pustaka", "2005", 8)
	model.BukuInsert("Bumi Manusia", "Pramoedya Ananta Toer", "Hasta Mitra", "1980", 5)
	model.BukuInsert("Hujan", "Tere Liye", "Gramedia Pustaka Utama", "2016", 20)
	model.BukuInsert("Perahu Kertas", "Dewi Lestari", "Bentang Pustaka", "2009", 15)
	model.BukuInsert("5 CM", "Donny Dhirgantoro", "Grasindo", "2007", 25)
	model.BukuInsert("Dilan: Dia Adalah Dilanku Tahun 1990", "Pidi Baiq", "Pastel Books", "2014", 30)
	model.BukuInsert("Ayat-Ayat Cinta", "Habiburrahman El Shirazy", "Republika", "2004", 10)
	model.BukuInsert("Negeri 5 Menara", "A. Fuadi", "Gramedia Pustaka Utama", "2009", 18)
	model.BukuInsert("Supernova: Ksatria, Puteri, dan Bintang Jatuh", "Dewi Lestari", "Truedee Books", "2001", 7)
	model.BukuInsert("Ronggeng Dukuh Paruk", "Ahmad Tohari", "Gramedia Pustaka Utama", "1982", 4)
	model.BukuInsert("Bulan", "Tere Liye", "Gramedia Pustaka Utama", "2015", 15)
	model.BukuInsert("Orang-Orang Biasa", "Andrea Hirata", "Bentang Pustaka", "2019", 6)
	model.BukuInsert("Gajah Mada", "Langit Kresna Hariadi", "Tiga Serangkai", "2003", 9)
	model.BukuInsert("Cantik Itu Luka", "Eka Kurniawan", "Gramedia Pustaka Utama", "2002", 12)
	model.BukuInsert("Di Bawah Lindungan Ka'bah", "Hamka", "Bulantiga Press", "1938", 3)
	model.BukuInsert("Lelaki Harimau", "Eka Kurniawan", "Gramedia Pustaka Utama", "2004", 8)
	model.BukuInsert("Garis Waktu", "Fiersa Besari", "Mediakita", "2016", 20)
	model.BukuInsert("Kambing Jantan", "Raditya Dika", "Gagas Media", "2005", 25)
	model.BukuInsert("Sabtu Bersama Bapak", "Adhitya Mulya", "Gagas Media", "2014", 18)
	model.BukuInsert("Tenggelamnya Kapal Van der Wijck", "Hamka", "Balai Pustaka", "1938", 5)
	model.BukuInsert("Rindu", "Tere Liye", "Republika", "2014", 13)
	model.BukuInsert("Pulang", "Leila S. Chudori", "Kepustakaan Populer Gramedia", "2012", 6)

	//test insert member
	model.InsertMember("indra", "Casanova", "12345", "A", 1)
	model.InsertMember("Zayn", "Zayn", "1111", "M", 1)
	model.InsertMember("Malik", "Malik", "2222", "M", 1)

	model.InsertMember("Rohman Ayai", "Rhm", "12345", "M", 0)
	model.InsertPeminjaman(node.Member{node.MemberNode{5, "Hanabi", "hanabi", "1212", "M", 1}, "Malang", "085849584"}, []int{1, 3})
	model.InsertPeminjaman(node.Member{node.MemberNode{2, "Zayn", "Zayn", "1111", "M", 1}, "Pasuruan", "085849584"}, []int{2, 3})
	fmt.Println("Dah Jalan")
	// fmt.Println(model.ReadAllMember())
	// fmt.Println(controller.Login("Casanova", "12345"))
	// main_program()
	//view.BukuUpdate()
	// webProgram()

}
func UserMenu(nama string, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	var pilih string
	for {
		fmt.Println("Menu Member:")
		fmt.Println("1. Peminjaman")
		fmt.Println("2. History Peminjaman")
		fmt.Println("3. Exit")
		fmt.Println()
		fmt.Print("Pilih: ")
		if scanner.Scan() {
			pilih = strings.TrimSpace(scanner.Text())
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}
		switch pilih {
		case "1":
			view.InsertPeminjaman(nama, id)
		case "2":
			view.DisplayAllPeminjamanByUser(id)
		case "3":
			os.Exit(0)
		case "4":
			view.DisplayAllPeminjaman()
		default:
			fmt.Println("Pilihan tidak ada")
		}
		scanner.Scan()
	}
}

func main() {
	tester()
	// role, name, id := VLogin()
	// if role == "A" {
	// 	fmt.Println("Selamat Datang ", name, ":)")
	// 	fmt.Println("Login Berhasil")
	// 	fmt.Println()
	// 	AdminMenu(name, id)
	// } else if role == "M" {
	// 	fmt.Println("Halo", name)
	// 	fmt.Println("Login Berhasil")
	// 	fmt.Println()
	// 	UserMenu(name, id)
	// } else {
	// 	fmt.Println("Login Gagal")
	// }
	webProgram()
}
