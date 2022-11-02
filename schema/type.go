package schema

type IdxType string

const (
	IdxTypeUnique IdxType = "unique"
	IdxTypeIndex  IdxType = "index"
)
