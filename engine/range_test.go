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
		update.Set([]byte{0, 5}, []byte{0, 5})
		update.Set([]byte{0, 6}, []byte{0, 6})
		update.Set([]byte{0, 7}, []byte{0, 7})
		update.Set([]byte{0, 8}, []byte{0, 8})
		update.Set([]byte{0, 9}, []byte{0, 9})
		update.Set([]byte{0, 10}, []byte{0, 10})
		update.Set([]byte{0, 11}, []byte{0, 11})
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if err := e.TxView(func(view *TxView) error {
		var count int
		scroll := view.RangeBETWEEN(
			nil,
			[]byte{0, 1},
			[]byte{0, 10},
			true,
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
		defer scroll.Close()

		limit := 3
		if err := scroll.Next(&limit); err != nil {
			return err
		}
		fmt.Printf("count : %v | left limit : %v\n", count, limit)

		limit = 3
		if err := scroll.Next(&limit); err != nil {
			return err
		}
		fmt.Printf("count : %v | left limit : %v\n", count, limit)
		limit = 10
		if err := scroll.Next(&limit); err != nil {
			return err
		}
		fmt.Printf("count : %v | left limit : %v\n", count, limit)

		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
