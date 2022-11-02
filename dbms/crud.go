package dbms

import (
	"fmt"
	"strconv"

	"github.com/gokch/db_badger/engine"
	"github.com/gokch/db_badger/schema"
)

func (t *DBMS) Insert(txn *engine.TxUpdate, tblName string, record schema.Record) error {
	// 스키마 가져오기
	tbl := t.schema.GetTable(tblName)
	if tbl == nil {
		return fmt.Errorf("table not found: %s", tblName)
	}

	// 임시 - todo - 중복 값 검사 필요 ( primary key, unique key )

	arrKey := make([][]byte, 0, len(tbl.Indexes)+1)
	arrVal := make([][]byte, 0, len(tbl.Indexes)+1)

	// set record to primary key
	{
		key := make([]byte, 0, 1024)
		key = append(key, []byte(strconv.FormatInt(int64(tbl.Seq), 10)+":")...)
		key = append(key, []byte(strconv.FormatInt(int64(tbl.Primary.Seq), 10))...)
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
		key = append(key, []byte(strconv.FormatInt(int64(tbl.Seq), 10)+":")...)
		key = append(key, []byte(strconv.FormatInt(int64(idx.Seq), 10)+":")...)
		fields, err := record.Encode(idx.Fields...)
		if err != nil {
			return err
		}
		key = append(key, fields...)

		arrKey = append(arrKey, key)
		arrVal = append(arrVal, arrKey[0]) // pk key
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
