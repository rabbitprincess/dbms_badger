package engine

import (
	"bytes"
	"errors"
	"sync"

	"github.com/dgraph-io/badger/v3"
)

// 임시 - !!!테스트 필요!!!

func (t *TxView) RangeBETWEEN(prefix, start, end []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}
	scroll.init(t.Txn, prefix, start, end, true, true, reverse, keyOnly, read)
	return scroll
}

func (t *TxView) RangeGT(prefix, start []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}
	scroll.init(t.Txn, prefix, start, nil, false, true, reverse, keyOnly, read)
	return scroll
}

func (t *TxView) RangeGE(prefix, start []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}
	scroll.init(t.Txn, prefix, start, nil, true, true, reverse, keyOnly, read)
	return scroll
}

func (t *TxView) RangeLT(prefix, end []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}
	scroll.init(t.Txn, prefix, nil, end, true, false, reverse, keyOnly, read)
	return scroll
}

func (t *TxView) RangeLE(prefix, end []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}
	scroll.init(t.Txn, prefix, nil, end, true, true, reverse, keyOnly, read)
	return scroll
}

// TODO: 임시 continue if target == key
func (t *TxView) RangeNE(prefix, target []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	return nil
}

func (t *TxView) RangeBeginsWith(prefix []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}

	end := make([]byte, len(prefix))
	copy(end, prefix) // copy end
	if end[len(end)-1] < 0xFF {
		end[len(end)-1]++
	} else {
		end = append(end, bytes.Repeat([]byte{0xFF}, 65535)...)
	}
	scroll.init(t.Txn, nil, prefix, end, true, false, reverse, keyOnly, read)
	return scroll
}

//------------------------------------------------------------------------------//
// range

type Scroll struct {
	mtx *sync.Mutex

	includeStart bool
	includeEnd   bool
	reverse      bool
	keyOnly      bool
	read         func(key []byte, value []byte) (err error)

	iterStart []byte
	iterEnd   []byte

	iter   *badger.Iterator
	finish bool
}

func (t *Scroll) init(txn *badger.Txn, prefix, start, end []byte, includeStart, includeEnd, reverse, keyOnly bool, read func(key, value []byte) error) {
	t.mtx = &sync.Mutex{}
	t.mtx.Lock()
	defer t.mtx.Unlock()

	t.includeStart = includeStart
	t.includeEnd = includeEnd
	t.reverse = reverse
	t.keyOnly = keyOnly
	t.read = read

	// init iterator
	{
		opt := badger.DefaultIteratorOptions
		opt.Prefix = make([]byte, len(prefix))
		copy(opt.Prefix, prefix) // set prefix
		if t.reverse == true {   // set reverse
			opt.Reverse = true
			includeStart, includeEnd = includeEnd, includeStart
			start, end = end, start
		}
		if t.keyOnly == true { // set key only
			opt.PrefetchValues = false
		}
		t.iter = txn.NewIterator(opt)
	}

	// init start, end
	{
		t.iterStart = make([]byte, 0, len(prefix)+len(start))
		t.iterStart = append(t.iterStart, prefix...)
		t.iterStart = append(t.iterStart, start...)

		t.iterEnd = make([]byte, 0, len(prefix)+len(end))
		t.iterEnd = append(t.iterEnd, prefix...)
		t.iterEnd = append(t.iterEnd, end...)
	}

	// 시작값이 prefix 로만 지정되었는데, reverse == true 일 경우 처리 필요
	if t.reverse == true && len(t.iterStart) == len(prefix) {
		t.iterStart = append(t.iterStart, 0xFF)
	}

	// seek iterator
	t.iter.Seek(t.iterStart)
}

func (t *Scroll) Next(limit *int) (err error) {
	if t.finish == true {
		return errors.New("EOF")
	}

	t.mtx.Lock()
	defer t.mtx.Unlock()

	if t.includeStart != true {
		t.iter.Next()
	} else {
		t.includeStart = false
	}

	// iterate
	for ; t.iter.Valid(); t.iter.Next() {
		// set key
		var key []byte = t.iter.Item().Key()

		// break end
		if (t.includeEnd != true && bytes.Compare(t.iterEnd, key) == 0) ||
			(t.reverse == false && bytes.Compare(t.iterEnd, key) < 0) ||
			(t.reverse == true && bytes.Compare(t.iterEnd, key) > 0) { // end 조건
			break
		}

		// set value
		var value []byte
		if t.keyOnly != true {
			value, err = t.iter.Item().ValueCopy(nil)
			if err != nil {
				return err
			}
		}

		// read key, value
		err = t.read(key, value)
		if err != nil {
			return err
		}

		// decrease limit
		*limit--
		if *limit == 0 {
			return nil // return ( without close )
		}
	}
	t.Close()
	return nil
}

func (t *Scroll) Close() {
	t.finish = true
	t.iter.Close()
}
