package controller

import (
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
