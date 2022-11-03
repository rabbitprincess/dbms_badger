package engine

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/badger/v3"
)

func TestTxView_RangeBETWEEN(t *testing.T) {
	opt := badger.DefaultOptions("").WithInMemory(true).WithLogger(nil)
	e := &Engine{}
	if err := e.OpenAdv(opt); err != nil {
		t.Fatal(err)
	}
	if err := e.DropAll(); err != nil {
		t.Fatal(err)
	}

	if err := e.TxUpdate(func(update *TxUpdate) error {
		update.Set([]byte{0, 0}, []byte{0, 0})
		update.Set([]byte{0, 1}, []byte{0, 1})
		update.Set([]byte{0, 2}, []byte{0, 2})
		update.Set([]byte{0, 3}, []byte{0, 3})
		update.Set([]byte{0, 4}, []byte{0, 4})

		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if err := e.TxView(func(view *TxView) error {
		var count int = 100
		scroll := view.RangeBETWEEN(
			[]byte{0},
			[]byte{1},
			[]byte{3},
			false,
			false,
			func(key, value []byte) error {
				if string(key) != string(value) {
					return fmt.Errorf("key != value")
				}
				count++
				fmt.Printf("key : %v, val : %v\n", key, value)
				return nil
			},
		)

		var limit int = 10
		if err := scroll.Next(&limit); err != nil {
			return err
		}
		fmt.Println("left limit :", limit)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
