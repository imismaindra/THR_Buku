package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"thr/controller"
	"thr/model"
)

// Handler untuk membaca semua anggota dan menampilkannya
func MemberReadAllHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/viewMember.html"))
	members := controller.ReadAllMember()
	if err := tmpl.Execute(w, members); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
