package dbms

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/sroar"
	"github.com/gokch/dbms_badger/engine"
	"github.com/gokch/dbms_badger/schema"
	"github.com/valyala/bytebufferpool"
)

// TODO: 임시 - key 제작 시 구분자 기준 대신 길이 처리 필요 ( 동적 길이 또는 고정 길이 )
func (t *DBMS) Insert(txn *engine.TxUpdate, tblName string, record schema.Record) error {
	// get schema
	tbl := t.schema.GetTable(tblName)
	if tbl == nil {
		return fmt.Errorf("table not found: %s", tblName)
	}

	// TODO: 임시 - 중복 값 검사 필요 ( primary key, unique key )
	var pk uint64 = 1
	if len(tbl.Indexes) == 0 {
		return fmt.Errorf("empty idx in table: %s", tblName)
	}
	// set pk record in index
	pkIdx := tbl.Indexes[0]
	key, err := schema.AppendIndexKey(nil, tbl, pkIdx, record)
	if err != nil {
		return err
	}
	value, err := record.EncodeAll(nil)
	if err != nil {
		return err
	}
	err = txn.Set(key, value)
	if err != nil {
		return err
	}

	// set record by index
	for i := 1; i < len(tbl.Indexes); i++ {
		idx := tbl.Indexes[i]
		if idx == nil {
			return fmt.Errorf("idx not found: %d", i)
		}

		key, err := schema.AppendIndexKey(nil, tbl, idx, record)
		if err != nil {
			return err
		}
		switch tbl.Indexes[i].Type {
		case schema.IdxTypeBitset:
			// IF BITSET INDEX
			v, err := txn.GetItem(key)
			if err != nil && err != badger.ErrKeyNotFound {
				return err
			}

			// TODO: 비트맵 샤딩 처리 필요
			b := bytebufferpool.Get()
			defer bytebufferpool.Put(b)
			b.B, err = v.ValueCopy(b.B)
			if err != nil {
				return err
			}

			bitmap := sroar.FromBuffer(b.B)
			bitmap.Set(pk)
			err = txn.Set(key, bitmap.ToBuffer())
			if err != nil {
				return err
			}
		case schema.IdxTypeRange, schema.IdxTypeUnique:
			// Value = <nil>
			err = txn.Set(key, nil)
			if err != nil {
				return err
			}
		case schema.IdxTypePrimary:
			// Table has only one primary key, and it is always at index 0
			// But, this for loop starts from 1
			panic("unreachable")
			/*
				// Value = < serialized record >
				v, err := record.Encode(nil)
				if err != nil {
					return err
				}
				err = txn.Set(key, v)
				if err != nil {
					return err
				}
			*/
		default:
			return schema.ErrInvalidIndexType
		}
	}

	return nil
}

func (t *DBMS) Exist(txn *engine.TxView, tblName string, idxName string, record schema.Record) (exist bool, err error) {
	value, err := t.Get(txn, tblName, idxName, record)
	if err != nil {
		return false, err
	}
	if value == nil {
		return false, nil
	}
	return true, nil
}

func (t *DBMS) Get(txn *engine.TxView, tblName string, idxName string, record schema.Record) (value []byte, err error) {
	// get schema
	tbl := t.schema.GetTable(tblName)
	if tbl == nil {
		return nil, fmt.Errorf("table not found: %s", tblName)
	}
	idx := tbl.GetIndex(idxName)
	if idx == nil {
		return nil, fmt.Errorf("index not found: %s", idxName)
	}

	// record 를 이용해 가져와서 key 제작
	key, err := schema.AppendIndexKey(nil, tbl, idx, record)
	if err != nil {
		return nil, err
	}

	// pk 가 아닐 경우 pk get

	// key 로 value get
	return txn.Get(key)
}

func (t *DBMS) Delete(txn *engine.TxUpdate, tblName string, idxName string, record schema.Record) (err error) {
	// get schema
	tbl := t.schema.GetTable(tblName)
	if tbl == nil {
		return fmt.Errorf("table not found: %s", tblName)
	}
	idx := tbl.GetIndex(idxName)
	if idx == nil {
		return fmt.Errorf("index not found: %s", idxName)
	}

	// record 를 이용해 가져와서 key 제작
	key, err := schema.AppendIndexKey(nil, tbl, idx, record)
	if err != nil {
		return err
	}

	// pk 가 아닐 경우 pk get

	// key 로 value delete
	return txn.Delete(key)
}

// TODO: Implement Range
func (t *DBMS) Range(txn *engine.TxView, tblName string, idxName string, record schema.Record) error {

	return nil
}
