package controller

import (
	"fmt"
	"thr/model"
	"thr/node"
)

func InsertPeminjaman(member node.Member, bukuIDs []int) {
	model.InsertPeminjaman(member, bukuIDs)

}
func GetAllPeminjaman() []node.PeminjamanBuku {
	return model.GetAllPeminjaman()
}
func UpStatusPeminjaman(idPeminjaman int, newStatus int) bool {
	success := model.UpdatePeminjamanStatus(idPeminjaman, newStatus)
	if success {
		return true
	} else {
		return false
	}
}
func CheckPeminjamanID(id int) (bool, int) {
	prev, temp := model.IsIdPeminjamanExist(id)
	if temp != nil {
		return true, temp.Peminjaman.Status
	} else if prev != nil {
		return true, prev.Peminjaman.Status
	} else {
		return false, -1
	}
}

func CheckStokBuku(id int) bool {
	_, buku := model.IsIdBukuAda(id)
	if buku != nil {
		if buku.Buku.Stok > 0 {
			model.UpdateStokBuku(id, 1)
			return true
		} else {
			return false
		}
	}
	return false
}
func CreturnBook(peminjamanid int, userid int, bukuid []int) bool {
	_, Peminjaman := model.IsIdPeminjamanExist(peminjamanid)
	// Jika ID peminjaman tidak ada
	// if Peminjaman == nil {
	// 	fmt.Printf("ID peminjaman %d tidak ditemukan\n", peminjamanid)
	// 	return false
	// }
	// Jika peminjaman statusnya 0 = ditinjau atau 3 = sudah dikembalikan
	if Peminjaman.Peminjaman.Status == 0 || Peminjaman.Peminjaman.Status == 3 {
		fmt.Printf("Peminjaman dengan ID %d tidak valid untuk pengembalian\n", peminjamanid)
		return false
	}
	model.ReturnBook(peminjamanid, userid, bukuid)
	return true
}
