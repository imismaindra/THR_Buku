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
func IsIdBukuAda(id int) (*node.LinkedList, *node.LinkedList) {
	var prev, temp *node.LinkedList
	temp = &database.DbBuku
	for temp != nil {
		if temp.Buku.Id == id {
			return prev, temp
		}
		prev = temp
		temp = temp.Next
	}
	return nil, nil
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

	_, alBuku := IsIdBukuAda(id)
	alBuku.Buku.Judul = jdl
	alBuku.Buku.Penerbit = penerbit
	alBuku.Buku.Pengarang = pengarang
	alBuku.Buku.Tahun = tahun
	return true

}

func BukuDelete(id int) *node.LinkedList {
	prev, current := IsIdBukuAda(id)
	if current == nil {
		return nil
	}

	// Menghapus node yang ditemukan
	if prev == nil {
		// Jika yang dihapus adalah head
		database.DbBuku = *database.DbBuku.Next
	} else {
		prev.Next = current.Next
	}
	// Bebaskan memori yang digunakan oleh node yang dihapus
	current.Next = nil

	return &database.DbBuku
}

func BukuSearch(id int) *node.LinkedList {
	_, alBuku := IsIdBukuAda(id)
	return alBuku
}
