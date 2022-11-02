package dbms

import (
	"fmt"

	"github.com/gokch/db_badger/engine"
	"github.com/gokch/db_badger/schema"
)

func (t *DBMS) Insert(txn *engine.TxUpdate, tblName string, record schema.Record) error {
	// 스키마 가져오기
	tbl := t.schema.GetTable(tblName)
	if tbl == nil {
		return fmt.Errorf("table not found: %s", tblName)
	}

	// 중복 값 검사 ( primary key, unique key )

	arrKey := make([][]byte, 0, len(tbl.Indexes)+1)
	arrVal := make([][]byte, 0, len(tbl.Indexes)+1)

	// set record to primary key
	{
		key := make([]byte, 0, 1024)
		key = append(key, []byte(tblName)...)
		val, err := record.Encode()
		if err != nil {
			return err
		}
		arrKey = append(arrKey, key)
		arrVal = append(arrVal, val)
	}

	// set record to index
	for _, idx := range tbl.Indexes {
		var key []byte = make([]byte, 0, 1024)
		key = append(key, []byte(tblName+":")...)
		key = append(key, []byte(idx.Name+":")...)

		fields, err := record.Encode(idx.Fields...)
		if err != nil {
			return err
		}
		key = append(key, fields...)

		arrKey = append(arrKey, key)
		arrVal = append(arrVal, arrKey[0])
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
