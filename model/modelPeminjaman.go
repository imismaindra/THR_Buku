package model

import (
	"thr/database"
	"thr/node"
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

func InsertPeminjaman(idmember int, total int, details []node.DetailPeminjaman) {

	for i, detail := range details {
		bukuTemp := BukuSearch(detail.IdBuku)
		details[i].Jdl = bukuTemp.Buku.Judul
	}
	var temp *node.PeminjamanLL
	temp = &database.DbPeminjaman
	newPeminjaman := node.PeminjamanBuku{
		IdPeminjaman: PeminjamanId(),
		IdMember:     idmember,
		CreateAt:     now.Format("2006-01-02 15:04:05"),
		UpdateAt:     now.Format("2006-01-02 15:04:05"),
		Total:        total,
		Details:      details,
	}

}
