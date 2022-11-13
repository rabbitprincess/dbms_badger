package schema

import (
	"fmt"
)

const (
	TBL_NAME_INFORMATION_SCHEMA = "information_schema"
	FLD_IDX_PK                  = 0
	FID_NAME_PK                 = "0"
)

//go:generate go run github.com/tinylib/msgp@latest

func NewSchema() *Schema {
	schema := &Schema{}
	return schema
}

// TODO: 임시 - 동시 접근을 막기 위한 mutex lock 필요
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

	tbl := &Table{}
	tbl.init(TblSeq(len(s.Tables)+1), tblName)
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
	Seq      TblSeq
	Name     string
	Indexes  []*Index
	idxNames map[string]*Index
}

func (t *Table) init(seq TblSeq, name string) {
	t.Seq = seq
	t.Name = name
	t.Indexes = make([]*Index, 0, 10)
	t.idxNames = make(map[string]*Index)

	// init pk
	pk := &Index{
		Seq:  FLD_IDX_PK,
		Name: FID_NAME_PK,
		Type: IdxTypePrimary,
	}
	t.Indexes = append(t.Indexes, pk)
	t.idxNames[FID_NAME_PK] = pk
}

func (t *Table) AddIndex(idxName string, idxType IdxType) error {
	if _, ok := t.idxNames[idxName]; ok {
		return fmt.Errorf("index already exists: %s", idxName)
	}

	idx := &Index{
		Seq:  IdxSeq(len(t.Indexes) + 1),
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
	Seq    IdxSeq
	Type   IdxType
	Name   string
	Fields []*Field
}

type Record interface {
	Get(fieldName string) *Field

	// EncodeAll appends the binary representation of the record to the byte slice,
	// and returns the updated slice.
	EncodeAll(b []byte) ([]byte, error)

	// DecodeAll decodes the binary representation of the record from the byte slice.
	DecodeAll(b []byte) error

	// EncodeField appends the binary representation of the field to the byte slice,
	// and returns the updated slice.
	EncodeField(b []byte, id FLdSeq) ([]byte, error)

	// DecodeField decodes the binary representation of the field from the byte slice.
	DecodeField(b []byte, id FLdSeq) error
}

type Field struct {
	Seq   FLdSeq
	Name  string
	Value interface{}
}
