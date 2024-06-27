package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"thr/model"
	"thr/node"
)

type LoginResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Role string `json:"role"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		member := model.CheckLogin(username, password)
		if member != nil {
			response := LoginResponse{
				ID:   member.Member.Id,
				Nama: member.Member.Nama,
				Role: member.Member.Role,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

			// Redirect based on role
			if member.Member.Role == "A" {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/store", http.StatusSeeOther)
			}
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
		return
	}

	tmpl, err := template.ParseFiles("ViewLogin.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/dashboard.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Dapatkan nama pengguna dari sesi atau dari data lainnya
	// Misalnya dari cookie atau token

	namaPengguna := "Neil Sims" // Ganti dengan cara Anda mendapatkan nama pengguna

	data := struct {
		BukuReport       int
		PeminjamanReport int
		MemberReport     int
		Nama             string
	}{
		BukuReport:       model.BukuCount(),
		PeminjamanReport: model.PeminjamanCount(),
		MemberReport:     model.MemberCount(),
		Nama:             namaPengguna,
	}
	log.Print(model.MemberCount())

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bookIdStr := r.FormValue("bookId")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Pada bagian ini, Anda bisa menambahkan logika untuk menambah buku ke dalam cart.
	// Misalnya, menyimpan ID buku ke dalam session atau database sesuai dengan kebutuhan aplikasi Anda.

	// Setelah berhasil menambahkan, Anda bisa memberikan respons ke klien.
	fmt.Fprintf(w, "Book with ID %d added to cart successfully", bookId)
}

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	books := model.GetRandomBooks(10)
	allBooks := model.BukuReadAll()
	booksJSON, err := json.Marshal(allBooks)
	if err != nil {
		http.Error(w, "Failed to encode books to JSON", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("view/store.html"))
	data := struct {
		AllBooks  []node.Buku
		Books     []node.Buku
		BooksJSON string
	}{
		AllBooks:  allBooks,
		Books:     books,
		BooksJSON: string(booksJSON),
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
