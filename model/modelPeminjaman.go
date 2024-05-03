package model

import (
	"thr/database"
	"thr/node"
	"time"
)

func PeminjamanId() int {
	var temp *node.PeminjamanLL
	temp = &database.DbPeminjaman
	if temp.Next == nil {
		return 1

	} else {
		for temp.Next != nil {
			temp = temp.Next

		}
		return temp.Peminjaman.IdPeminjaman + 1
	}
}

func InsertPeminjaman(member node.Member, bukuIDs []int) {
	now := time.Now()
	var details []node.DetailPeminjaman

	for _, id := range bukuIDs {
		bukuTemp := BukuSearch(id)
		details = append(details, node.DetailPeminjaman{IdBuku: id, Jdl: bukuTemp.Buku.Judul})
	}

	newPeminjaman := node.PeminjamanBuku{
		IdPeminjaman:     PeminjamanId(),
		Member:           member,
		CreateAt:         now,
		UpdateAt:         now,
		DetailPeminjaman: details,
	}

	temp := &database.DbPeminjaman
	if temp.Next == nil {
		temp.Next = &node.PeminjamanLL{Peminjaman: newPeminjaman, Next: nil}
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &node.PeminjamanLL{Peminjaman: newPeminjaman, Next: nil}
	}
}
func GetAllPeminjaman() []node.PeminjamanBuku {
	var peminjamanList []node.PeminjamanBuku

	temp := database.DbPeminjaman.Next
	for temp != nil {
		peminjamanList = append(peminjamanList, temp.Peminjaman)
		temp = temp.Next
	}

	return peminjamanList
}
