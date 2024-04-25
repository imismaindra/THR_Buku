package model

import (
	"thr/database"
	"thr/node"
)

func MemberId() int {
	var temp *node.MemberLinkedList
	temp = &database.DbMember
	if temp.Next == nil {
		return 1

	} else {
		for temp.Next != nil {
			temp = temp.Next

		}
		return temp.Member.Id + 1
	}
}
func IsMemberExist(id int) (*node.MemberLinkedList, *node.MemberLinkedList) {
	var prev, temp *node.MemberLinkedList
	temp = &database.DbMember
	for temp != nil {
		if temp.Member.Id == id {
			return prev, temp
		}
		prev = temp
		temp = temp.Next
	}
	return nil, nil
}

func InsertMember(nama string, Uname string, pass string, role string, status int) {
	var temp *node.MemberLinkedList
	temp = &database.DbMember
	member := node.MemberNode{
		Id:       MemberId(),
		Nama:     nama,
		Username: Uname,
		Password: pass,
		Role:     role,
		Status:   status,
	}
	newLL := node.MemberLinkedList{
		Member: member,
	}
	if temp.Next == nil {
		temp.Next = &newLL
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newLL
	}

}
func ReadAllMember() []node.MemberNode {
	var temp *node.MemberLinkedList
	temp = &database.DbMember
	var TableMember []node.MemberNode
	for temp.Next != nil {
		temp = temp.Next
		TableMember = append(TableMember, temp.Member)

	}
	return TableMember
}
func UpdateMember(id int, pass string, role string, status int) bool {
	_, IsMember := IsMemberExist(id)
	IsMember.Member.Password = pass
	IsMember.Member.Role = role
	IsMember.Member.Status = status
	return true

}
func SearchMember(id int) *node.MemberLinkedList {
	_, IsMember := IsMemberExist(id)
	return IsMember
}

func DeleteMember(id int) *node.MemberLinkedList {
	_, IsMemeber := IsMemberExist(id)
	return IsMemeber

}