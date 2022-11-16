package engine

type QueryType uint16

const (
	QueryTypeSelect QueryType = iota
	QueryTypeInsert
	QueryTypeUpdate
	QueryTypeDelete

	QueryTypeExplain
)

type Query struct {
	Type QueryType
}

type Field interface {
	// Encode appends the binary representation of the field to the byte slice,
	// and returns the updated slice.
	Encode(b []byte) ([]byte, error)

	// Decode decodes the binary representation of the field from the byte slice.
	Decode(b []byte) error
}

type Operator uint16

const (
	OperatorEqual Operator = 1 << iota
	OperatorNotEqual
	OperatorLess
	OperatorLessEqual
	OperatorGreater
	OperatorGreaterEqual
)

type Where struct {
	Table  string
	Column string
	Field  Field
}
