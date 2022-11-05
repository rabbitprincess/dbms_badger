package schema

import (
	"encoding/binary"
	"errors"
)

/*

# BitSet Index Format

key = T/<table_id>/<index_id>/<index_key>
value = <bitset>

# Unique Index Format

key = T/<table_id>/<index_id>/<index_key>/<primary_key>
value = <null>

# Range Index Format

key = T/<table_id>/<index_id>/<index_key>/<primary_key>
value = <null>

# Primary Key Format

key = T/<table_id>/0/<primary_key>
value = <data>

*/

func appendIndexPrefix(b []byte, tblID TblSeq, idxID IdxSeq) []byte {
	b = append(b, "T"...)
	b = binary.LittleEndian.AppendUint16(b, uint16(tblID))
	b = binary.LittleEndian.AppendUint16(b, uint16(idxID))
	return b
}

var ErrInvalidIndexType = errors.New("invalid index type")

func AppendIndexKey(b []byte, tbl *Table, idx *Index, record Record) (key []byte, err error) {
	b = appendIndexPrefix(b, tbl.Seq, idx.Seq)
	switch idx.Type {
	case IdxTypePrimary, IdxTypeUnique, IdxTypeRange:
		b, err = record.EncodeField(b, idx.Fields[0].Seq)
		if err != nil {
			return nil, err
		}
		pk := tbl.Indexes[0]
		b, err = record.EncodeField(b, pk.Fields[0].Seq)
		if err != nil {
			return nil, err
		}
		return b, nil
	case IdxTypeBitset:
		b, err = record.EncodeField(b, idx.Fields[0].Seq)
		if err != nil {
			return nil, err
		}
		// No primary key
		return b, nil
	default:
		err = ErrInvalidIndexType
		return
	}
}
