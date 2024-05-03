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
