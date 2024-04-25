package node

type MemberNode struct {
	Id       int
	Nama     string
	Username string
	Password string
	Role     string
	Status   int
}

type MemberLinkedList struct {
	Member MemberNode
	Next   *MemberLinkedList
}
