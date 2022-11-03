package engine

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/badger/v3"
)

func TestTxView_RangeBETWEEN(t *testing.T) {
	opt := badger.DefaultOptions("").WithInMemory(true)
	db, err := badger.Open(opt)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer db.Close()

	e := &Engine{db}
	e.DropAll()

	err = e.TxUpdate(func(update *TxUpdate) error {
		update.Set([]byte{0, 0}, []byte{0, 0})
		update.Set([]byte{0, 1}, []byte{0, 1})
		update.Set([]byte{0, 2}, []byte{0, 2})
		update.Set([]byte{0, 3}, []byte{0, 3})
		update.Set([]byte{0, 4}, []byte{0, 4})

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	err = e.TxView(func(view *TxView) error {
		prefix := []byte{0}
		start := []byte{0}
		end := []byte{3}
		var count int = 100
		scroll := view.RangeBETWEEN(
			prefix,
			start[:],
			end[:],
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
		err = scroll.Next(&limit)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
