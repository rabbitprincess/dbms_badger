package dbms

import (
	"fmt"
	"strconv"

	"github.com/gokch/dbms_badger/engine"
	"github.com/gokch/dbms_badger/schema"
)

// TODO: 임시 - key 제작 시 구분자 기준 대신 길이 처리 필요 ( 동적 길이 또는 고정 길이 )
func (t *DBMS) Insert(txn *engine.TxUpdate, tblName string, record schema.Record) error {
	// 스키마 가져오기
	var err error

	tbl := t.schema.GetTable(tblName)
	if tbl == nil {
		return fmt.Errorf("table not found: %s", tblName)
	}

	// TODO: 임시 - 중복 값 검사 필요 ( primary key, unique key )

	arrKey := make([][]byte, 0, len(tbl.Indexes)+1)
	arrVal := make([][]byte, 0, len(tbl.Indexes)+1)

	// set record to index
	for _, idx := range tbl.Indexes {
		var key []byte = make([]byte, 0, 1024)
		key = append(key, []byte(strconv.FormatInt(int64(tbl.Seq), 10)+":")...)
		key = append(key, []byte(strconv.FormatInt(int64(idx.Seq), 10)+":")...)

		var val []byte
		if idx.Type == schema.IdxTypePrimary { // pk 일 경우
			val, err = record.Encode(nil) // set record to value
			if err != nil {
				return err
			}
		} else { // non pk 일 경우
			key, err = record.Encode(key, idx.Fields...) // append index fields
			if err != nil {
				return err
			}
			key = append(key, arrKey[0]...) // append primary key
		}

		arrKey = append(arrKey, key)
		arrVal = append(arrVal, val) // pk key
	}

	// set kv
	for i := 0; i < len(arrKey); i++ {
		err := txn.Set(arrKey[i], arrVal[i])
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

// TODO: Implement Get
func (t *DBMS) Get(txn *engine.TxView, tblName string, idxName string, record schema.Record) (value []byte, err error) {
	// 스키마 가져오기
	tbl := t.schema.GetTable(tblName)
	if tbl == nil {
		return nil, fmt.Errorf("table not found: %s", tblName)
	}
	idx := tbl.GetIndex(idxName)
	if idx == nil {
		return nil, fmt.Errorf("index not found: %s", idxName)
	}

	// record 를 이용해 가져와서 key 제작
	var key []byte

	// key 로 value get
	return txn.Get(key)
}

// TODO: Implement Range
func (t *DBMS) Range(txn *engine.TxView, tblName string, idxName string, record schema.Record) error {

	return nil
}
