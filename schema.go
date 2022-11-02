package db_badger

type Schema struct {
	Tables []*Table
}

type Table struct {
	seq     uint64
	Indexes []*Index
	Fields  []*Field
}

type Index struct {
	seq    uint64
	Fields []*Field
}

type Field struct {
	seq   uint64
	value interface{}
}
