package engine

import (
	"bytes"
	"sync"

	"github.com/dgraph-io/badger/v3"
)

// 임시 - !!!테스트 필요!!!

func (t *TxView) RangeBETWEEN(prefix, start, end []byte, reverse bool, limit int, read func(key, value []byte) error) (nextStart []byte, err error) {
	rng := &Range{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		start:        start,
		end:          end,
		reverse:      reverse,
		limit:        limit,
		read:         read,
		includeStart: true,
		includeEnd:   true,
	}
	return rng.Range()
}

func (t *TxView) RangeGT(prefix, start []byte, reverse bool, limit int, read func(key, value []byte) error) (nextStart []byte, err error) {
	rng := &Range{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		start:        start,
		reverse:      reverse,
		limit:        limit,
		read:         read,
		includeStart: false,
		includeEnd:   true,
	}
	return rng.Range()
}

func (t *TxView) RangeGE(prefix, start []byte, reverse bool, limit int, read func(key, value []byte) error) (nextStart []byte, err error) {
	rng := &Range{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		start:        start,
		reverse:      reverse,
		limit:        limit,
		read:         read,
		includeStart: true,
		includeEnd:   true,
	}
	return rng.Range()
}

func (t *TxView) RangeLT(prefix, end []byte, reverse bool, limit int, read func(key, value []byte) error) (nextStart []byte, err error) {
	rng := &Range{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		end:          end,
		reverse:      reverse,
		limit:        limit,
		read:         read,
		includeStart: true,
		includeEnd:   false,
	}
	return rng.Range()
}

func (t *TxView) RangeLE(prefix, end []byte, reverse bool, limit int, read func(key, value []byte) error) (nextStart []byte, err error) {
	rng := &Range{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		end:          end,
		reverse:      reverse,
		limit:        limit,
		read:         read,
		includeStart: true,
		includeEnd:   true,
	}
	return rng.Range()
}

// 임시 - todo
// if reverse == false -> lt + gt
// if reverse == true -> gt + lt
func (t *TxView) RangeNE(prefix, target []byte, reverse bool, limit int, read func(key, value []byte)) (nextStart []byte, err error) {
	return nil, nil
}

//------------------------------------------------------------------------------//
// range

type Range struct {
	mtx *sync.Mutex

	txn          *badger.Txn
	prefix       []byte
	start        []byte
	end          []byte
	includeStart bool
	includeEnd   bool

	reverse bool
	keyOnly bool
	limit   int

	read func(key []byte, value []byte) (err error)
}

func (t *Range) Range() (next []byte, err error) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	// init iterator
	var iter *badger.Iterator
	defer iter.Close()
	{
		opt := badger.DefaultIteratorOptions
		copy(opt.Prefix, t.prefix) // set prefix
		if t.keyOnly == true {     // set key only
			opt.PrefetchValues = false
		}
		if t.reverse == true { // set reverse
			opt.Reverse = true
			t.includeStart, t.includeEnd = t.includeEnd, t.includeStart
			t.start, t.end = t.end, t.start
		}
		iter = t.txn.NewIterator(opt)
	}

	// set start, end
	var iterStart, iterEnd []byte
	{
		iterStart = make([]byte, 0, len(t.prefix)+len(t.start))
		iterStart = append(iterStart, t.prefix...)
		iterStart = append(iterStart, t.start...)

		iterEnd = make([]byte, 0, len(t.prefix)+len(t.end))
		iterEnd = append(iterEnd, t.prefix...)
		iterEnd = append(iterEnd, t.end...)
	}

	// 시작값이 prefix 로만 지정되었는데, reverse == true 일 경우 처리 필요
	if t.reverse == true && len(iterStart) == len(t.prefix) {
		iterStart = append(iterStart, 0xFF)
	}

	// iterate
	for iter.Seek(iterStart); iter.Valid(); iter.Next() {
		// set key
		var key []byte = iter.Item().Key()
		if t.includeStart != true && bytes.Equal(iterStart, key) == true { // include start 조건
			continue
		}
		if t.includeEnd != true && bytes.Equal(iterEnd, key) == true { // include end 조건
			break
		}
		if (t.reverse == true && bytes.Compare(iterEnd, key) > 0) ||
			(t.reverse == false && bytes.Compare(iterEnd, key) < 0) { // end 조건
			break
		}

		// set value
		var value []byte
		if t.keyOnly != true {
			value, err = iter.Item().ValueCopy(nil)
			if err != nil {
				return nil, err
			}
		}

		// read key, value
		err = t.read(key, value)
		if err != nil {
			return nil, err
		}

		// decrease limit
		t.limit--
		if t.limit == 0 {
			break
		}
	}

	// return next key
	if iter.Next(); iter.Valid() == true {
		return iter.Item().Key(), nil
	}
	return nil, nil
}
