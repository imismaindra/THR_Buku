package node

type MemberNode struct {
	Nama     string
	Username string
	Password string
	Role     string
}

type MemberLinkedList struct {
	Member *MemberNode
	Next   *MemberLinkedList
}
