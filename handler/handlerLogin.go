package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"thr/controller"
	"thr/model"
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
func StoreHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"view/store.html"))
	users := controller.ViewBuku()
	// Menampilkan data ke template HTML
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
