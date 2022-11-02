package engine

import (
	"bytes"
	"sync"

	"github.com/dgraph-io/badger/v3"
)

// 임시 - !!!테스트 필요!!!

func (t *TxView) RangeBETWEEN(prefix, start, end []byte, reverse bool, read func(key, value []byte) error) *Scroll {
	return &Scroll{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		start:        start,
		end:          end,
		reverse:      reverse,
		read:         read,
		includeStart: true,
		includeEnd:   true,
	}
}

func (t *TxView) RangeGT(prefix, start []byte, reverse bool, read func(key, value []byte) error) *Scroll {
	return &Scroll{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		start:        start,
		reverse:      reverse,
		read:         read,
		includeStart: false,
		includeEnd:   true,
	}
}

func (t *TxView) RangeGE(prefix, start []byte, reverse bool, read func(key, value []byte) error) *Scroll {
	return &Scroll{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		start:        start,
		reverse:      reverse,
		read:         read,
		includeStart: true,
		includeEnd:   true,
	}
}

func (t *TxView) RangeLT(prefix, end []byte, reverse bool, read func(key, value []byte) error) *Scroll {
	return &Scroll{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		end:          end,
		reverse:      reverse,
		read:         read,
		includeStart: true,
		includeEnd:   false,
	}
}

func (t *TxView) RangeLE(prefix, end []byte, reverse bool, read func(key, value []byte) error) *Scroll {
	return &Scroll{
		mtx:          &sync.Mutex{},
		txn:          t.txn,
		prefix:       prefix,
		end:          end,
		reverse:      reverse,
		read:         read,
		includeStart: true,
		includeEnd:   true,
	}
}

// 임시 - todo
// continue if target == key
func (t *TxView) RangeNE(prefix, target []byte, reverse bool, read func(key, value []byte)) *Scroll {
	return nil
}

//------------------------------------------------------------------------------//
// range

type Scroll struct {
	mtx *sync.Mutex

	txn          *badger.Txn
	prefix       []byte
	start        []byte
	end          []byte
	includeStart bool
	includeEnd   bool

	reverse bool
	keyOnly bool

	read func(key []byte, value []byte) (err error)
}

// 임시 - iter 재사용 필요
func (t *Scroll) Next(next *[]byte, limit *int) (err error) {
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
		if next != nil {
			iterStart = make([]byte, 0, len(t.prefix)+len(*next))
			iterStart = append(iterStart, *next...)
			iterStart = append(iterStart, t.start...)
			t.includeStart = true
		} else {
			iterStart = make([]byte, 0, len(t.prefix)+len(t.start))
			iterStart = append(iterStart, t.prefix...)
			iterStart = append(iterStart, t.start...)
		}

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
			break
		}
	}

	// return next key

	if iter.Next(); iter.Valid() == true {
		*next = iter.Item().Key()
	}
	return nil
}
