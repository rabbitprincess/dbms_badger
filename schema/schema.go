package schema

import "encoding/json"

const (
	TBL_NAME_INFORMATION_SCHEMA = "information_schema"
)

type Schema struct {
	Tables []*Table
}

func (s *Schema) Load(bt []byte) error {
	// 임시 - 현재는 json decoding 으로 복구
	return json.Unmarshal(bt, s)
}

func (s *Schema) Save() (bt []byte, err error) {
	// 임시 - 현재는 json encoding 으로 저장
	return json.Marshal(s)
}

func (s *Schema) AddTable(tblName string) {
	if s.Tables == nil {
		s.Tables = make([]*Table, 0, 10)
	}

	s.Tables = append(s.Tables, &Table{
		Seq:  len(s.Tables),
		Name: tblName,
	})
}

func (s *Schema) GetTable(tblName string) *Table {
	if s.Tables == nil {
		return nil
	}
	for _, tbl := range s.Tables {
		if tbl.Name == tblName {
			return tbl
		}
	}
	return nil
}

type Table struct {
	Seq     int
	Name    string
	Indexes []*Index
	Fields  []*Field
}

func (t *Table) AddIndex() {
	if t.Indexes == nil {
		t.Indexes = make([]*Index, 0, 10)
	}

	t.Indexes = append(t.Indexes, &Index{
		Seq: len(t.Indexes),
	})
}

func (t *Table) GetIndex(idxName string) *Index {
	if t.Indexes == nil {
		return nil
	}
	for _, idx := range t.Indexes {
		if idx.Name == idxName {
			return idx
		}
	}
	return nil
}

type Index struct {
	Seq    int
	Type   IdxType
	Name   string
	Fields []*Field
}

type Record interface {
	Get(fieldName string) *Field
	Encode(fields ...*Field) ([]byte, error)
	Decode(bt []byte, fields ...*Field) error
}

type Field struct {
	Seq   int
	Name  string
	value interface{}
}
