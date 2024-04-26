package controller

import (
	"errors"
	"thr/model"
	"thr/node"
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
func ReadAllMember(id int) []node.MemberNode {
	isMember := model.SearchMember(id)
	if isMember == nil {
		return nil
	}
	var Tbl_member []node.MemberNode
	Tbl_member = append(Tbl_member)
	return Tbl_member
}

func DeleteMember(id int) bool {
	if id != 0 {
		model.BukuDelete(id)
		return true
	}
	return false
}
func SearchMember(id int) []node.MemberNode {
	IsMember := model.SearchMember(id)
	if IsMember == nil {
		return nil
	}
	var Tbl_member []node.MemberNode
	Tbl_member = append(Tbl_member, IsMember.Member)

	return Tbl_member
}
