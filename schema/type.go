package schema

type IdxType string

const (
	IdxTypePrimary IdxType = "primary"
	IdxTypeUnique  IdxType = "unique"
	IdxTypeIndex   IdxType = "index"
)
