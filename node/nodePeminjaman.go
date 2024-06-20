package node

import "time"

type Member struct {
	User   MemberNode
	Alamat string
	NoTelp string
}

type PeminjamanBuku struct {
	IdPeminjaman int
	Member
	CreateAt         time.Time
	UpdateAt         time.Time
	DetailPeminjaman []DetailPeminjaman
	ReturnAt         time.Time
	Status           int
	BackAt           time.Time
}

type DetailPeminjaman struct {
	IdBuku int
	Jdl    string
}

type PeminjamanLL struct {
	Peminjaman PeminjamanBuku
	Next       *PeminjamanLL
}
