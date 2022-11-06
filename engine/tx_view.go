package engine

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

//------------------------------------------------------------------------------------//
// base

type TxView struct {
	db  *Engine
	Txn *badger.Txn
}

func (t *TxView) Init(db *Engine, txn *badger.Txn) {
	t.db = db
	t.Txn = txn
}

//------------------------------------------------------------------------------------//
// get set del

func (t *TxView) IsExist(key []byte) (bool, error) {
	_, err := t.Get(key)
	if err != nil {
		return false, err
	} else if err == badger.ErrKeyNotFound {
		return false, nil
	}
	return true, nil
}

func (t *TxView) Get(key []byte) (value []byte, err error) {
	item, err := t.GetItem(key)
	if err != nil {
		return nil, err
	}

	return item.ValueCopy(nil)
}

func (t *TxView) GetItem(key []byte) (*badger.Item, error) {
	item, err := t.Txn.Get(key)
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}

func (t *TxView) GetStr(key string) (val []byte, err error) {
	return t.Get([]byte(key))
}

//------------------------------------------------------------------------------------//
// iterate

func (t *TxView) NewIterator(reverse bool, keyOnly bool, prefix, start []byte) *Iterator {
	iterator := &Iterator{}
	iterator.Init(t, reverse, keyOnly, prefix, start)
	return iterator
}

type Iterator struct {
	badgerIt *badger.Iterator

	reverse bool
	keyOnly bool
	prefix  []byte
	start   []byte
}

func (t *Iterator) Init(tx *TxView, reverse, keyOnly bool, prefix []byte, start []byte) {
	t.reverse = reverse
	t.keyOnly = keyOnly
	t.prefix = prefix
	t.start = start

	t.badgerIt = tx.Txn.NewIterator(badger.IteratorOptions{
		Reverse: t.reverse,
		Prefix:  t.prefix,
	})
	t.badgerIt.Seek(t.start)
}

func (t *Iterator) Config() (reverse bool, prefix, start []byte) {
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
