package controller

import (
	"thr/model"
	"thr/node"
)

func InsertBuku(jdl string, pengarang string, penerbit string, thn string, stok int, img string) bool {
	if jdl != "" && pengarang != "" && penerbit != "" && thn != "" {
		model.BukuInsert(jdl, pengarang, penerbit, thn, stok, img)
		return true
	}
	return false
}
func ViewBuku() []node.Buku {
	buku := model.BukuReadAll()
	if buku == nil {
		return nil
	}
	return buku

}
func SearchBuku(id int) []node.Buku {
	buku := model.BukuSearch(id)
	if buku == nil {
		return nil
	}
	var vBuku []node.Buku
	vBuku = append(vBuku, buku.Buku)

	return vBuku
}
func UpdateBuku(id int, jdl string, pengarang string, penerbit string, thn string, stok int, img string) bool {

	if jdl != "" && pengarang != "" && penerbit != "" && thn != "" {
		model.BukuUpdate(id, jdl, pengarang, penerbit, thn, stok, img)
		return true
	}
	return false
}
func CheckBukuID(id int) bool {
	_, IsIdBukuTrue := model.IsIdBukuAda(id)
	if IsIdBukuTrue != nil {
		return true
	}
	return false
}
func DeleteBuku(id int) bool {
	if id != 0 {
		model.BukuDelete(id)
		return true
	}
	return false
}
