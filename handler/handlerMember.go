package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"thr/controller"
	"thr/model"
	"thr/node"
)

// Handler untuk membaca semua anggota dan menampilkannya
func MemberReadAllHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/viewMember.html"))
	members := controller.ReadAllMember()
	if err := tmpl.Execute(w, members); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Print(members)
}

// Handler untuk menampilkan halaman update member
func MemberUpdatePageHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/updatemember/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	memberNode := model.SearchMember(id)
	if memberNode == nil {
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	member := memberNode.Member
	tmpl := template.Must(template.ParseFiles("view/updateMember.html"))
	if err := tmpl.Execute(w, struct {
		Member node.MemberNode
	}{Member: member}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handler untuk proses update member
func MemberUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	nama := r.FormValue("nama")
	username := r.FormValue("username")
	role := r.FormValue("role")
	status, err := strconv.Atoi(r.FormValue("status"))
	if err != nil {
		http.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}

	success := controller.UpdateMember(id, nama, username, role, status)
	if !success {
		http.Error(w, "Failed to update member", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/member", http.StatusSeeOther)
}

func MemberDeleteHandler(w http.ResponseWriter, r *http.Request) {
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

	deleted := model.DeleteMember(id)
	if deleted == nil {
		http.Error(w, "Member not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handler untuk menampilkan halaman edit member
func EditMemberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := r.URL.Path[len("/memberupdate/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		member := model.SearchMember(id)
		if member == nil {
			http.Error(w, "Member not found", http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles("view/updateMember.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Member node.MemberNode
		}{
			Member: member.Member,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPut {
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		nama := r.FormValue("nama")
		username := r.FormValue("username")
		role := r.FormValue("role")
		statusStr := r.FormValue("status")
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			http.Error(w, "Status harus berupa angka", http.StatusBadRequest)
			return
		}

		success := model.UpdateMember(id, nama, username, role, status)
		if !success {
			http.Error(w, "Failed to update member", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/member", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func MemberInsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Menampilkan form inputan
		tmpl := template.Must(template.ParseFiles("view/insertMember.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Handle form submission
		r.ParseForm()
		nama := r.FormValue("nama")
		username := r.FormValue("username")
		password := r.FormValue("password")
		role := r.FormValue("role")
		statusStr := r.FormValue("status")
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			http.Error(w, "Status harus berupa angka", http.StatusBadRequest)
			return
		}

		// Check for other potential issues with submitted data
		if nama == "" || username == "" || password == "" {
			http.Error(w, "Semua field harus diisi", http.StatusBadRequest)
			return
		}

		// Memanggil controller untuk insert data
		model.InsertMember(nama, username, password, role, status)

		// Redirect kembali ke halaman utama setelah proses insert
		http.Redirect(w, r, "/member", http.StatusSeeOther)
		return
	}
}
