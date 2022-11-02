package schema

import (
	"fmt"
)

const (
	TBL_NAME_INFORMATION_SCHEMA = "information_schema"
)

//go:generate msgp

// 임시 - todo - 동시 접근을 막기 위한 mutex lock 필요
type Schema struct {
	Tables   []*Table
	tblNames map[string]*Table
}

func (s *Schema) Load(bt []byte) error {
	// 임시 - 현재는 msgpack decoding 으로 복구
	_, err := s.UnmarshalMsg(bt)
	return err
}

func (s *Schema) Save() (bt []byte, err error) {
	// 임시 - 현재는 msgpack encoding 으로 저장
	return s.MarshalMsg(nil)
}

func (s *Schema) AddTable(tblName string) error {
	if s.Tables == nil {
		s.Tables = make([]*Table, 0, 10)
		s.tblNames = make(map[string]*Table)
	}
	if _, ok := s.tblNames[tblName]; ok {
		return fmt.Errorf("table already exists: %s", tblName)
	}

	tbl := &Table{
		Seq:  len(s.Tables),
		Name: tblName,
	}
	s.Tables = append(s.Tables, tbl)
	s.tblNames[tblName] = tbl
	return nil
}

func (s *Schema) GetTable(tblName string) *Table {
	if s.Tables == nil {
		return nil
	}
	return s.tblNames[tblName]
}

type Table struct {
	Seq      int
	Name     string
	Primary  *Index
	Indexes  []*Index
	idxNames map[string]*Index
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
		t.idxNames = make(map[string]*Index)
	}

	if _, ok := t.idxNames[idxName]; ok {
		return fmt.Errorf("index already exists: %s", idxName)
	}

	idx := &Index{
		Seq:  len(t.Indexes) + 1,
		Name: idxName,
	}

	t.Indexes = append(t.Indexes, idx)
	t.idxNames[idxName] = idx
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
	Value interface{}
}
