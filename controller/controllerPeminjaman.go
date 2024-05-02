package controller

import (
	"thr/model"
	"thr/node"
)

func InsertPenjualan(idmember int, details []node.DetailPeminjaman) bool {
	if idmember != 0 {
		model.InsertPeminjaman(idmember, details)
		return true
	}
	return false
}
