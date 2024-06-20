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
func CalculateLateFee(returnDate, dueDate time.Time) (int, int) {
	lateDays := int(returnDate.Sub(dueDate).Hours() / 24)
	if lateDays <= 0 {
		return 0, 0
	}
	fee := lateDays * 2000
	return lateDays, fee
}

func InsertPeminjaman(member node.Member, bukuIDs []int) (node.PeminjamanBuku, bool) {
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
	return newPeminjaman, true

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
func UpdatePeminjamanStatus(Id int, nwStatus int) bool {
	_, temp := IsIdPeminjamanExist(Id)
	if nwStatus == 1 {
		temp.Peminjaman.Status = nwStatus
		backAt := time.Now().AddDate(0, 0, 3)
		// formattedTime := backAt.Format("Monday, 02 January 2006 15:04 MST")
		temp.Peminjaman.BackAt = backAt
		return true

	} else if nwStatus == 2 {
		temp.Peminjaman.Status = nwStatus
		// Mendapatkan ID buku yang dipinjam dalam peminjaman ini
		for _, detail := range temp.Peminjaman.DetailPeminjaman {
			fmt.Println(detail.IdBuku)
			UpdateStokBuku(detail.IdBuku, 3)
		}

		return true
	} else if nwStatus == 3 {
		temp.Peminjaman.Status = nwStatus
		return true
	} else {
		return false
	}

}

func ReturnBook(peminjamanID, userID int, bukuIDs []int) {
	_, peminjaman := IsIdPeminjamanExist(peminjamanID)
	if peminjaman == nil {
		fmt.Printf("ID peminjaman %d tidak ditemukan\n", peminjamanID)
		return
	}

	for _, bukuID := range bukuIDs {
		found := false
		for _, detail := range peminjaman.Peminjaman.DetailPeminjaman {
			if detail.IdBuku == bukuID {
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("ID buku %d tidak sesuai dengan peminjaman\n", bukuID)
			return
		}

		// Perbarui stok buku
		UpdateStokBuku(bukuID, 3)
	}

	now := time.Now()
	dueDate := peminjaman.Peminjaman.BackAt

	lateDays, fee := CalculateLateFee(now, dueDate)
	if lateDays > 0 {
		fmt.Printf("Anda terlambat mengembalikan buku selama %d hari. Denda: %d\n", lateDays, fee)
		if lateDays > 2 {
			UpdateUserStatus(userID, 1) // 1 untuk banned
			fmt.Println("Status pengguna telah diperbarui menjadi 'dilarang' karena keterlambatan lebih dari 2 hari.")
		}
	} else {
		fmt.Println("Buku dikembalikan tepat waktu. Tidak ada denda.")
	}

	// Perbarui status peminjaman menjadi selesai/dikembalikan
	UpdatePeminjamanStatus(peminjamanID, 3)
}

func UpdateStokBuku(id int, status int) {
	var temp *node.LinkedList
	temp = &database.DbBuku
	for temp != nil {
		if temp.Buku.Id == id {
			if status == 1 && temp.Buku.Stok > 0 {
				temp.Buku.Stok--
				fmt.Println("Stok buku berkurang")
				return
			} else if status == 3 {
				temp.Buku.Stok++
				return
			}
		}
		temp = temp.Next
	}
}
func PeminjamanCount() int {
	var count int
	var temp *node.PeminjamanLL
	temp = &database.DbPeminjaman

	for temp.Next != nil {
		temp = temp.Next
		count++
	}
	return count
}
