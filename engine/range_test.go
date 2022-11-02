package engine

import (
	"encoding/binary"
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

	e.TxUpdate(func(update *TxUpdate) error {
		var b [8]byte
		for i := 0; i < 1000000; i++ {
			binary.BigEndian.PutUint64(b[:], uint64(i))
			update.Set(b[:], b[:])
		}
		return nil
	})

	err = e.TxView(func(view *TxView) error {
		var start, end [8]byte
		binary.BigEndian.PutUint64(start[:], 100)
		binary.BigEndian.PutUint64(end[:], 200)
		var count int = 100
		view.RangeBETWEEN(
			nil,
			start[:],
			end[:],
			false,
			func(key, value []byte) error {
				if string(key) != string(value) {
					t.Fatal("key != value")
				}

				v := binary.BigEndian.Uint64(value)
				if v != uint64(count) {
					t.Fatalf("expected %d, got %d", count, v)
				}
				count++
				return nil
			},
		)

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
