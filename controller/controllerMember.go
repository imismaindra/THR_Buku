package controller

import (
	"thr/model"
	"thr/node"
)

func InsertMember(nama string, uname string, pass string, role string, status int) bool {
	// Validasi input
	if nama == "" || uname == "" || pass == "" || role == "" {
		return false
	}

	// Validasi status
	if status < 0 || status > 1 {
		return false
	}
	model.InsertMember(nama, uname, pass, role, status)
	return true
}
func UpdateMember(id int, role string, status int) bool {
	if role == "" {
		return false
	}
	if status < 0 || status > 1 {
		return false
	}
	model.UpdateMember(id, role, status)
	return true
}
func ReadAllMember() []node.MemberNode {
	member := model.ReadAllMember()
	if member == nil {
		return nil
	}
	return member
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
func CheckMemberID(id int) bool {
	_, IsIdMemberTrue := model.IsMemberExist(id)
	// if IsIdMemberTrue != nil {
	// 	return true
	// }
	return IsIdMemberTrue != nil
}
func Login(username, password string) (string, string, int) {
	member := model.CheckLogin(username, password)
	if member == nil {

		return "", "", 1
	}
	if member.Member.Username != username && member.Member.Password != password {
		return "Password atau Username salah", "", 1
	}
	return member.Member.Role, member.Member.Nama, member.Member.Id
}
