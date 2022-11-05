package schema

//go:generate msgp
type IdxType string

const (
	IdxTypePrimary IdxType = "PRIMARY"
	IdxTypeUnique  IdxType = "unique"
	IdxTypeIndex   IdxType = "index"
)
