package model

import (
	"thr/database"
	"thr/node"
)

func BukuId() int {
	var temp *node.LinkedList
	temp = &database.DbBuku
	if temp.Next == nil {
		return 1

	} else {
		for temp.Next != nil {
			temp = temp.Next

		}
		return temp.Buku.Id + 1
	}
}
func IsIdBukuAda(id int) *node.LinkedList {
	var temp *node.LinkedList
	temp = &database.DbBuku
	for temp != nil {
		if temp.Buku.Id == id {
			return temp
		}
		temp = temp.Next
	}
	return nil
}

func BukuInsert(judul string, pengarang string, penerbit string, tahun string) {
	var temp *node.LinkedList
	temp = &database.DbBuku
	buku := node.Buku{
		Id:        BukuId(),
		Judul:     judul,
		Pengarang: pengarang,
		Penerbit:  penerbit,
		Tahun:     tahun,
	}
	newLL := node.LinkedList{
		Buku: buku,
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
func BukuReadAll() []node.Buku {
	var temp *node.LinkedList
	temp = &database.DbBuku
	var TableBuku []node.Buku
	for temp.Next != nil {
		temp = temp.Next
		TableBuku = append(TableBuku, temp.Buku)
	}
	return TableBuku

}
func BukuUpdate(id int, jdl string, pengarang string, penerbit string, tahun string) bool {

	alBuku := IsIdBukuAda(id)
	alBuku.Buku.Judul = jdl
	alBuku.Buku.Penerbit = penerbit
	alBuku.Buku.Pengarang = pengarang
	alBuku.Buku.Tahun = tahun
	return true

}

func BukuDelete(id int) *node.LinkedList {
	var temp *node.LinkedList
	temp = &database.DbBuku
	if temp.Next != nil {
		for temp.Next != nil {
			if temp.Next.Buku.Id == id {
				temp.Next = temp.Next.Next
				return &database.DbBuku
			}
			temp = temp.Next
		}
	}
	return nil
}
func BukuSearch(id int) *node.LinkedList {
	// var temp *node.LinkedList
	// temp = &database.DbBuku
	alBuku := IsIdBukuAda(id)
	if alBuku != nil {
		return alBuku
	}

	// if temp.Next != nil {
	// 	for temp.Next != nil {
	// 		if temp.Next.Buku.Id == id {
	// 			return temp.Next
	// 		}
	// 		temp = temp.Next

	// 	}
	// } else {
	// 	return nil
	// }
	return nil
}
