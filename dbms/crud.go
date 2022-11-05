package dbms

import (
	"fmt"

	"github.com/gokch/dbms_badger/engine"
	"github.com/gokch/dbms_badger/schema"
)

// TODO: 임시 - key 제작 시 구분자 기준 대신 길이 처리 필요 ( 동적 길이 또는 고정 길이 )
func (t *DBMS) Insert(txn *engine.TxUpdate, tblName string, record schema.Record) error {
	// get schema
	tbl := t.schema.GetTable(tblName)
	if tbl == nil {
		return fmt.Errorf("table not found: %s", tblName)
	}

	// TODO: 임시 - 중복 값 검사 필요 ( primary key, unique key )

	// set record by index
	for _, idx := range tbl.Indexes {
		key, val, err := schema.MakeKV(tbl, idx, record)
		if err != nil {
			return err
		}
		err = txn.Set(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: Implement Exist
func (t *DBMS) Exist(txn *engine.TxView, tblName string, idxName string, record schema.Record) (exist bool, err error) {

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
	key, _, err := schema.MakeKV(tbl, idx, record)
	if err != nil {
		return nil, err
	}

	// pk 가 아닐 경우 pk get

	// key 로 value get

	txn.Get(key)
	return
}

// TODO: Implement Range
func (t *DBMS) Range(txn *engine.TxView, tblName string, idxName string, record schema.Record) error {

	return nil
}
