package schema

import (
	"encoding/json"
	"fmt"
)

const (
	TBL_NAME_INFORMATION_SCHEMA = "information_schema"
)

// 임시 - todo - 동시 접근을 막기 위한 mutex lock 필요
type Schema struct {
	Tables  []*Table
	tblName map[string]*Table
}

func (s *Schema) Load(bt []byte) error {
	// 임시 - 현재는 json decoding 으로 복구
	return json.Unmarshal(bt, s)
}

func (s *Schema) Save() (bt []byte, err error) {
	// 임시 - 현재는 json encoding 으로 저장
	return json.Marshal(s)
}

func (s *Schema) AddTable(tblName string) error {
	if s.Tables == nil {
		s.Tables = make([]*Table, 0, 10)
		s.tblName = make(map[string]*Table)
	}
	if _, ok := s.tblName[tblName]; ok {
		return fmt.Errorf("table already exists: %s", tblName)
	}

	tbl := &Table{
		Seq:  len(s.Tables),
		Name: tblName,
	}
	s.Tables = append(s.Tables, tbl)
	s.tblName[tblName] = tbl
	return nil
}

func (s *Schema) GetTable(tblName string) *Table {
	if s.Tables == nil {
		return nil
	}
	return s.tblName[tblName]
}

type Table struct {
	Seq       int
	Name      string
	Primary   *Index
	Indexes   []*Index
	indexName map[string]*Index
	Fields    []*Field
}

func (t *Table) AddPrimaryIndex(idxName string) error {
	if t.Primary != nil {
		return fmt.Errorf("primary index already exists: %s", idxName)
	}

	idx := &Index{
		Seq:  0,
		Name: idxName,
	}
	t.Primary = idx
	return nil
}

func (t *Table) GetPrimaryIndex() *Index {
	return t.Primary
}

func (t *Table) AddIndex(idxName string, idxType IdxType) error {
	if t.Indexes == nil {
		t.Indexes = make([]*Index, 0, 10)
		t.indexName = make(map[string]*Index)
	}

	if _, ok := t.indexName[idxName]; ok {
		return fmt.Errorf("index already exists: %s", idxName)
	}

	idx := &Index{
		Seq:  len(t.Indexes) + 1,
		Name: idxName,
	}

	t.Indexes = append(t.Indexes, idx)
	t.indexName[idxName] = idx
	return nil
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
