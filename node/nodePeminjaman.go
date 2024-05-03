package node

import "time"

type Member struct {
	Nama   string
	Alamat string
	NoTelp string
}

type PeminjamanBuku struct {
	IdPeminjaman int
	Member       Member

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
