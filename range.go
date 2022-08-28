package db_badger

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
)

func (t *TxView) NewIterator(_reverse bool, _keyOnly bool, _prefix, _start []byte) *Iterator {
	iterator := &Iterator{}
	iterator.Init(t, _reverse, _keyOnly, _prefix, _start)
	return iterator
}

//---------------------------------------------------------------------------------------//
type Iterator struct {
	badgerIt *badger.Iterator

	reverse bool
	keyOnly bool
	prefix  []byte
	start   []byte
}

func (t *Iterator) Init(_tx *TxView, _reverse, _keyOnly bool, _prefix []byte, _start []byte) {
	t.reverse = _reverse
	t.keyOnly = _keyOnly
	t.prefix = _prefix
	t.start = _start

	t.badgerIt = _tx.txn.NewIterator(badger.IteratorOptions{
		Reverse: t.reverse,
		Prefix:  t.prefix,
	})
	t.badgerIt.Seek(t.start)
}

func (t *Iterator) Config() (_reverse bool, _prefix, _start []byte) {
	return t.reverse, t.prefix, t.start
}

func (t *Iterator) Valid() bool {
	return t.badgerIt.Valid()
}

func (t *Iterator) Next() {
	t.badgerIt.Next()
}

func (t *Iterator) Key() ([]byte, error) {
	if t.Valid() != true {
		return nil, fmt.Errorf("invalid key")
	}
	return t.badgerIt.Item().Key(), nil
}

func (t *Iterator) Value() ([]byte, error) {
	var value []byte
	err := t.badgerIt.Item().Value(func(val []byte) error {
		value = val
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (t *Iterator) ValueCopy() ([]byte, error) {
	var val []byte
	return t.badgerIt.Item().ValueCopy(val)
}

//-----------------------------------------------------------------//
// 임시 - range type
/*
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
*/
