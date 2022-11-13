package engine

import (
	"github.com/dgraph-io/badger/v3"
)

func NewEngine() *Engine {
	engine := &Engine{}
	return engine
}

type Engine struct {
	badger *badger.DB
}

func (t *Engine) Open(path string) error {
	return t.OpenAdv(badger.DefaultOptions(path))
}

func (t *Engine) OpenAdv(opt badger.Options) error {
	var err error
	t.badger, err = badger.Open(opt)
	if err != nil {
		return err
	}
	return nil
}

func (t *Engine) RunGC(ratio float64) {
	t.badger.RunValueLogGC(ratio)
}

func (t *Engine) Info() []badger.TableInfo {
	return t.badger.Tables()
}

func (t *Engine) DropAll() error {
	return t.badger.DropAll()
}

func (t *Engine) SequenceGet(key []byte, bandwidth uint64) (*badger.Sequence, error) {
	return t.badger.GetSequence(key, bandwidth)
}

func (t *Engine) TxView(fnCb func(view *TxView) error) error {
	return t.badger.View(func(_tx *badger.Txn) error {
		txView := &TxView{}
		txView.Init(t, _tx)
		return fnCb(txView)
	})
}

func (t *Engine) TxUpdate(fnCb func(update *TxUpdate) error) error {
	return t.badger.Update(func(_tx *badger.Txn) error {
		txUpdate := &TxUpdate{}
		txUpdate.Init(t, _tx)
		return fnCb(txUpdate)
	})
}
