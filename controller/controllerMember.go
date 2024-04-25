package controller

import "errors"

func InsertMember(nama string, uname string, pass string, role string, status int) error {
	// Validasi input
	if nama == "" || uname == "" || pass == "" || role == "" {
		return errors.New("semua parameter harus diisi")
	}

	// Validasi status
	if status <= 0 || status >= 1 {
		return errors.New("status harus berada dalam rentang 0-1")
	}
	return nil
}
