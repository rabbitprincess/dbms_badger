package schema

import "math"

//go:generate go run github.com/tinylib/msgp@latest

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=TblSeq
type TblSeq uint16

const (
	SeqTblInformationSchema TblSeq = iota
	SeqTblTable             TblSeq = math.MaxUint16
)

type TblName string

const (
	NameTblInformationSchema TblName = "information_schema"
)

type IdxSeq uint16

const (
	SeqIdxPK  IdxSeq = 0
	SeqIdxMax IdxSeq = math.MaxUint16
)

type IdxName string

const (
	NameIdxPk IdxName = "0"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=IdxType
type IdxType uint8

const (
	IdxTypePrimary IdxType = iota
	IdxTypeUnique
	IdxTypeRange
	IdxTypeBitset
)

type FLdSeq uint16

const (
	FldSeqMax FLdSeq = math.MaxUint16
)
