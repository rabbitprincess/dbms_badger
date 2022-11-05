package schema

import (
	"encoding/binary"
)

/*

# BitSet Index Format

key = _t/<table_id>/<index_id>/<index_key>
value = <bitset>

# Unique Index Format

key = _t/<table_id>/<index_id>/<index_key>/<primary_key>
value = <null>

# Range Index Format

key = _t/<table_id>/<index_id>/<index_key>/<primary_key>
value = <null>

# Primary Key Format

key = _t/<table_id>/0/<primary_key>
value = <data>

*/

func MakeKV(tbl *Table, idx *Index, record Record) (key, val []byte, err error) {
	key = make([]byte, 5, 1024)

	copy(key[:2], []byte("_t"))
	binary.BigEndian.PutUint16(key[2:4], uint16(tbl.Seq))
	key[4] = byte(idx.Seq)

	switch idx.Type {
	case IdxTypePrimary:
		val, err = record.Encode(nil) // set record to value
		if err != nil {
			return nil, nil, err
		}
	case IdxTypeUnique, IdxTypeIndex:
		key, err = record.Encode(key, idx.Fields...) // append index fields
		if err != nil {
			return nil, nil, err
		}
		// key = append(key, arrKey[0]...) // append primary key
	case IdxTypeBitset:
		// TODO: 임시 - idxType == bitset 시 key,value 세팅 방법 필요
	}

	return key, val, nil
}
