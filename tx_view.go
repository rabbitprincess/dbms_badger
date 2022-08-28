package db_badger

import (
	badger "github.com/dgraph-io/badger/v2"
)

//------------------------------------------------------------------------------------//
// base

type TxView struct {
	db  *Badger
	txn *badger.Txn
}

func (t *TxView) Init(_db *Badger, _txn *badger.Txn) {
	t.db = _db
	t.txn = _txn
}

//------------------------------------------------------------------------------------//
// get set del

func (t *TxView) IsExist(_key []byte) (bool, error) {
	_, err := t.Get(_key)
	if err != nil {
		return false, err
	} else if err == badger.ErrKeyNotFound {
		return false, nil
	}
	return true, nil
}

func (t *TxView) Get(_key []byte) (val []byte, err error) {
	item, err := t.txn.Get(_key)
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, nil
		}
		return nil, err
	}

	fn := func(_val []byte) error {
		val = _val
		return nil
	}
	item.Value(fn)
	return
}

func (t *TxView) GetStr(_key string) (val []byte, err error) {
	return t.Get([]byte(_key))
}
