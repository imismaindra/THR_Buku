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

func InsertPeminjaman(idmember int, details []node.DetailPeminjaman) {
	now := time.Now()

	for i, detail := range details {
		bukuTemp := BukuSearch(detail.IdBuku)
		details[i].Jdl = bukuTemp.Buku.Judul
	}
	var temp *node.PeminjamanLL
	temp = &database.DbPeminjaman
	newPeminjaman := node.PeminjamanBuku{
		IdPeminjaman:     PeminjamanId(),
		IdMember:         idmember,
		CreateAt:         now,
		UpdateAt:         now,
		DetailPeminjaman: details,
	}
	if temp.Next == nil {
		temp.Next = &node.PeminjamanLL{Peminjaman: newPeminjaman, Next: nil}
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &node.PeminjamanLL{Peminjaman: newPeminjaman, Next: nil}
	}

}
