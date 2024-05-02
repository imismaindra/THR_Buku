package node

type Buku struct {
	Id        int
	Judul     string
	Pengarang string
	Penerbit  string
	Tahun     string
	Stok      int
}

type LinkedList struct {
	Buku Buku
	Next *LinkedList
}
