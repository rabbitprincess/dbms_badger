package schema

//go:generate msgp
type IdxType string

const (
	IdxTypeUnique IdxType = "unique"
	IdxTypeIndex  IdxType = "index"
)
