package schema

import "math"

//go:generate msgp
type TblSeq uint16

const (
	SeqTblInformationSchema TblSeq = iota
	SeqTblTable             TblSeq = math.MaxUint16
)

type TblName string

const (
	NameTblInformationSchema TblName = "information_schema"
)

type IdxSeq uint8

const (
	SeqIdxPK  IdxSeq = 0
	SeqIdxMax IdxSeq = math.MaxUint8
)

type IdxName string

const (
	NameIdxPk IdxName = "0"
)

type IdxType uint8

const (
	IdxTypePrimary IdxType = iota
	IdxTypeUnique
	IdxTypeIndex
	IdxTypeBitset
)

func (t IdxType) String() string {
	switch t {
	case IdxTypePrimary:
		return "primary"
	case IdxTypeUnique:
		return "unique"
	case IdxTypeIndex:
		return "index"
	case IdxTypeBitset:
		return "bitset"
	}
	return ""
}

type FLdSeq uint16

const (
	FldSeqMax FLdSeq = math.MaxUint16
)
