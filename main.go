package db_badger

import (
	badger "github.com/dgraph-io/badger/v2"
)

type Badger struct {
	badger *badger.DB

	chanGCEnd chan bool
}

func (t *Badger) Open(_dbpath string) error {
	var err error
	t.badger, err = badger.Open(badger.DefaultOptions(_dbpath))
	if err != nil {
		return err
	}

	return nil
}

func (t *Badger) RunGC(_ratio float64) {
	t.badger.RunValueLogGC(_ratio)
}

func (t *Badger) Info() []badger.TableInfo {
	TableInfo := t.badger.Tables(true)
	return TableInfo
}

func (t *Badger) DropAll() error {
	return t.badger.DropAll()
}

func (t *Badger) SequenceGet(_key []byte, _bandwidth uint64) (*badger.Sequence, error) {
	return t.badger.GetSequence(_key, _bandwidth)
}

func (t *Badger) TX_view(_fn_cb func(_view *TxView) error) error {
	fn := func(_tx *badger.Txn) error {
		pt_tx__view := &TxView{}
		pt_tx__view.Init(t, _tx)
		return _fn_cb(pt_tx__view)
	}
	return t.badger.View(fn)
}

func (t *Badger) TX_update(_fn_cb func(_update *TxUpdate) error) error {
	fn := func(_tx *badger.Txn) error {
		pt_tx__update := &TxUpdate{}
		pt_tx__update.Init(t, _tx)
		return _fn_cb(pt_tx__update)
	}
	return t.badger.Update(fn)
}
