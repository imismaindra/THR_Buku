package model

import (
	"fmt"
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

func IsIdPeminjamanExist(id int) (*node.PeminjamanLL, *node.PeminjamanLL) {
	var prev, temp *node.PeminjamanLL
	temp = &database.DbPeminjaman
	for temp != nil {
		if temp.Peminjaman.IdPeminjaman == id {
			return prev, temp
		}
		prev = temp
		temp = temp.Next
	}
	return nil, nil
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
		Status:           0,
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
func UpdateStsPeminjaman(Id int, nwStatus int) bool {
	_, temp := IsIdPeminjamanExist(Id)
	if nwStatus == 1 {
		temp.Peminjaman.Status = nwStatus
		backAt := time.Now().AddDate(0, 0, 3)
		// formattedTime := backAt.Format("Monday, 02 January 2006 15:04 MST")
		temp.Peminjaman.BackAt = backAt
		return true

	} else if nwStatus == 3 {
		temp.Peminjaman.Status = nwStatus

		return true
	}
	return true

}
func UpdateStokBuku(id int, Sts int) {
	var temp *node.LinkedList
	temp = &database.DbBuku
	for temp != nil {
		if temp.Buku.Id == id {

			if Sts == 1 && temp.Buku.Stok > 0 {
				temp.Buku.Stok--
				fmt.Println("Stok buku berkurang")
				return
			} else if Sts == 3 {
				temp.Buku.Stok++
				return

			}
		}
		temp = temp.Next
	}
}
