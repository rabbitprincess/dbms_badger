package db_badger

import "errors"

//-----------------------------------------------------------------//
type TD_range int

const (
	TD_range_empty TD_range = iota
	TD_range_beginWith
	TD_range_between
	TD_range_lt
	TD_range_le
	TD_range_ge
	TD_range_gt
	TD_range_eq
	TD_range_ne
)

//-----------------------------------------------------------------//
type TD_dir int

const (
	TD_dir_asc TD_dir = iota
	TD_dir_desc
)

//-----------------------------------------------------------------//
type TD_keyOnly int

const (
	TD_keyOnly_false TD_keyOnly = iota
	TD_keyOnly_true
)

const (
	DEF_maxKeyLen = 65000
)

//-----------------------------------------------------------------//
// error

var (
	ErrKeyAlreadyExist = errors.New("already exist key")
)
