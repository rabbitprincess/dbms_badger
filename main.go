package db_badger

import (
	"github.com/dgraph-io/badger/v3"
)

type Badger struct {
	badger *badger.DB
	schema *Schema
}

func (t *Badger) Open(path string) error {
	var err error
	t.badger, err = badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return err
	}

	return nil
}

func (t *Badger) RunGC(ratio float64) {
	t.badger.RunValueLogGC(ratio)
}

func (t *Badger) Info() []badger.TableInfo {
	return t.badger.Tables()
}

func (t *Badger) DropAll() error {
	return t.badger.DropAll()
}

func (t *Badger) SequenceGet(key []byte, bandwidth uint64) (*badger.Sequence, error) {
	return t.badger.GetSequence(key, bandwidth)
}

func (t *Badger) TxView(fnCb func(view *TxView) error) error {
	return t.badger.View(func(_tx *badger.Txn) error {
		txView := &TxView{}
		txView.Init(t, _tx)
		return fnCb(txView)
	})
}

func (t *Badger) TxUpdate(fnCb func(update *TxUpdate) error) error {
	return t.badger.Update(func(_tx *badger.Txn) error {
		txUpdate := &TxUpdate{}
		txUpdate.Init(t, _tx)
		return fnCb(txUpdate)
	})
}
