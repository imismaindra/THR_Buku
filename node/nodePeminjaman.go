package node

import "time"

type PeminjamanBuku struct {
	IdPeminjaman int
	IdMember     int

	CreateAt         time.Time
	UpdateAt         time.Time
	DetailPeminjaman []DetailPeminjaman
	ReturnAt         time.Time
}
type DetailPeminjaman struct {
	IdBuku int
	Jdl    string
}
type PeminjamanLL struct {
	Peminjaman PeminjamanBuku
	Next       *PeminjamanLL
}
