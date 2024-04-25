package controller

import (
	"errors"
	"thr/model"
)

func InsertMember(nama string, uname string, pass string, role string, status int) (error, bool) {
	// Validasi input
	if nama == "" || uname == "" || pass == "" || role == "" {
		return errors.New("semua parameter harus diisi"), false
	}

	// Validasi status
	if status <= 0 || status >= 1 {
		return errors.New("status harus berada dalam rentang 0-1"), false
	}
	model.InsertMember(nama, uname, pass, role, status)
	return nil, true
}
func UpdateMember(id int, pass string, role string, status int) bool {
	if pass == "" && role == "" {
		return false
	}
	if status <= 0 || status >= 1 {
		return false
	}
	model.UpdateMember(id, pass, role, status)
	return true
}
