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
	scroll.init(t.txn, prefix, start, end, true, true, reverse, keyOnly, read)
	return scroll
}

func (t *TxView) RangeGT(prefix, start []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}
	scroll.init(t.txn, prefix, start, nil, false, true, reverse, keyOnly, read)
	return scroll
}

func (t *TxView) RangeGE(prefix, start []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}
	scroll.init(t.txn, prefix, start, nil, true, true, reverse, keyOnly, read)
	return scroll
}

func (t *TxView) RangeLT(prefix, end []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}
	scroll.init(t.txn, prefix, nil, end, true, false, reverse, keyOnly, read)
	return scroll
}

func (t *TxView) RangeLE(prefix, end []byte, reverse, keyOnly bool, read func(key, value []byte) error) *Scroll {
	scroll := &Scroll{}
	scroll.init(t.txn, prefix, nil, end, true, true, reverse, keyOnly, read)
	return scroll
}

// 임시 - todo
// continue if target == key
func (t *TxView) RangeNE(prefix, target []byte, reverse, keyOnly bool, read func(key, value []byte)) *Scroll {
	return nil
}

//------------------------------------------------------------------------------//
// range

type Scroll struct {
	mtx *sync.Mutex

	iterStart    []byte
	iterEnd      []byte
	includeStart bool
	includeEnd   bool

	reverse bool
	keyOnly bool

	read func(key []byte, value []byte) (err error)

	iter   *badger.Iterator
	finish bool
}

func (t *Scroll) init(txn *badger.Txn, prefix, start, end []byte, includeStart, includeEnd, reverse, keyOnly bool, read func(key, value []byte) error) {
	t.mtx = &sync.Mutex{}

	// init iterator
	{
		opt := badger.DefaultIteratorOptions
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

	t.includeStart = includeStart
	t.includeEnd = includeEnd
	t.reverse = reverse
	t.keyOnly = keyOnly
	t.read = read

	// 시작값이 prefix 로만 지정되었는데, reverse == true 일 경우 처리 필요
	if t.reverse == true && len(t.iterStart) == len(prefix) {
		t.iterStart = append(t.iterStart, 0xFF)
	}

	// seek iterator
	t.iter.Seek(t.iterStart)
	if t.includeStart != true {
		t.iter.Next()
	}
}

func (t *Scroll) Next(limit *int) (err error) {
	if t.finish == true {
		return errors.New("EOF")
	}

	t.mtx.Lock()
	defer t.mtx.Unlock()

	// iterate
	for ; t.iter.Valid(); t.iter.Next() {
		// set key
		var key []byte = t.iter.Item().Key()

		// break end
		if (t.includeEnd != true && bytes.Equal(t.iterEnd, key) == true) ||
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
			return nil
		}
	}
	t.finish = true
	t.iter.Close()

	return nil
}
