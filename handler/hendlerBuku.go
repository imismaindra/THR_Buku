package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"thr/controller"
	"thr/model"
	"thr/node"
)

func sub(x, y int) int {
	return x - y
}

func add(x, y int) int {
	return x + y
}
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	//memanggil memberView.html dengan tamplate

	tmpl := template.Must(template.ParseFiles(
		"view/viewLogin.html"))
	users := controller.ViewBuku()
	// Menampilkan data ke template HTML
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BukuInsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Menampilkan form inputan
		tmpl := template.Must(template.ParseFiles("view/insertBuku.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Handle form submission
		r.ParseForm()
		judul := r.FormValue("judul")
		pengarang := r.FormValue("pengarang")
		penerbit := r.FormValue("penerbit")
		tahun := r.FormValue("tahun")
		stokStr := r.Form.Get("stok")
		stok, err := strconv.Atoi(stokStr)
		if err != nil {
			http.Error(w, "Stok harus berupa angka", http.StatusBadRequest)
			return
		}

		// Check for other potential issues with submitted data
		if judul == "" || pengarang == "" || penerbit == "" || tahun == "" {
			http.Error(w, "Semua field harus diisi", http.StatusBadRequest)
			return
		}

		// Memanggil controller untuk insert data
		model.BukuInsert(judul, pengarang, penerbit, tahun, stok)

		// Redirect kembali ke halaman utama setelah proses insert
		http.Redirect(w, r, "/buku", http.StatusSeeOther)
		return
	}
}

func BukuReadAllHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"sub": sub,
		"add": add,
	}

	tmpl := template.Must(template.New("viewBuku.html").Funcs(funcMap).ParseFiles("view/viewBuku.html"))

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10 // Jumlah buku per halaman
	totalBooks := model.BukuCount()
	totalPages := (totalBooks + pageSize - 1) / pageSize

	books := model.GetBooksByPage(page, pageSize)

	// Filter out books with invalid ids (e.g., id == 0)
	var validBooks []node.Buku
	for _, book := range books {
		if book.Id != 0 {
			validBooks = append(validBooks, book)
		}
	}

	data := struct {
		Books      []node.Buku
		Page       int
		TotalPages int
	}{
		Books:      validBooks,
		Page:       page,
		TotalPages: totalPages,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func BukuUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	judul := r.FormValue("judul")
	pengarang := r.FormValue("pengarang")
	penerbit := r.FormValue("penerbit")
	tahun := r.FormValue("tahun")
	stokStr := r.Form.Get("stok")
	stok, err := strconv.Atoi(stokStr)
	if err != nil {
		http.Error(w, "Stok harus berupa angka", http.StatusBadRequest)
		return
	}

	success := model.BukuUpdate(id, judul, pengarang, penerbit, tahun, stok)
	if !success {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK) // Baris ini tidak diperlukan
}

func EditBukuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/updatebuku/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	buku := model.BukuSearch(id)
	if buku == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("view/UpdateBuku.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Buku *node.Buku
	}{
		Buku: &buku.Buku,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func BukuDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	deleted := model.BukuDelete(id)
	if deleted == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func BukuSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	buku := model.BukuSearch(id)
	if buku == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buku.Buku)
}
