package db_badger

import (
	"bytes"

	badger "github.com/dgraph-io/badger/v2"
)

func (t *Tx_view) Range_all(
	_limit int,
	_fn_cb_read func(_prefix []byte, _key []byte, _val []byte) error,
) (
	keyNext []byte,
	err error,
) {
	return t.Range(nil, nil, nil, nil, false, false, TD_dir_asc, TD_keyOnly_true, _limit, _fn_cb_read)
}

func (t *Tx_view) Range_beginWith(
	_keyPrefix []byte,
	_keyTarget []byte,
	_keyNext []byte,
	_dir TD_dir,
	_keyOnly TD_keyOnly,
	_limit int,
	_fn_cb_read func(_prefix []byte, _key []byte, _value []byte) error,
) (
	keyNext []byte,
	err error,
) {
	bt_key__target__max := Util__fill_0xff_max(_keyTarget)
	return t.Range(_keyPrefix, _keyTarget, bt_key__target__max, _keyNext, false, false, _dir, _keyOnly, _limit, _fn_cb_read)
}

func (t *Tx_view) Range_between(
	_keyPrefix []byte,
	_keyStart []byte,
	_keyEnd []byte,
	_keyNext []byte,
	_dir TD_dir,
	_keyOnly TD_keyOnly,
	_limit int,
	_fn_cb_read func(_prefix []byte, _key []byte, _value []byte) error,
) (
	keyNext []byte,
	err error,
) {
	return t.Range(_keyPrefix, _keyStart, _keyEnd, _keyNext, false, false, _dir, _keyOnly, _limit, _fn_cb_read)
}

func (t *Tx_view) Range_ne(
	_keyPrefix []byte,
	_keyStart []byte,
	_keyEnd []byte,
	_keyTarget []byte,
	_keyNext []byte,
	_dir TD_dir,
	_keyOnly TD_keyOnly,
	_limit int,
	_fn_cb_read func(_prefix []byte, _key []byte, _value []byte) error,
) (
	keyNext []byte,
	err error,
) {

	// 전처리 - limit 처리
	if _limit >= 0 {
		var bt_key__prefix_with_key []byte
		if _keyNext == nil {
			bt_key__prefix_with_key = append(_keyPrefix, _keyTarget...)
		} else {
			bt_key__prefix_with_key = append(_keyPrefix, _keyNext...)
		}
		bt_value, err := t.Get(bt_key__prefix_with_key)
		if err != nil {
			return nil, err
		}
		// get 값이 존재하면 limit 을 하나 늘린다 ( +1 )
		// 순회 중 get 값을 지나간다는 뜻이기 때문에, 하나를 더 늘림
		if bt_value != nil {
			_limit++
		}
	}

	fn_cb_read := func(_prefix []byte, _key []byte, _val []byte) error {
		if bytes.Equal(_key, _keyTarget) == true {
			return nil
		}
		return _fn_cb_read(_prefix, _key, _val)
	}

	return t.Range(_keyPrefix, _keyStart, _keyEnd, _keyNext, false, false, _dir, _keyOnly, _limit, fn_cb_read)

}

func (t *Tx_view) Range_eq(
	_keyPrefix []byte,
	_keyTarget []byte,
	_keyOnly TD_keyOnly,
	_fn_cb_read func(_prefix []byte, _key []byte, _val []byte) error,
) (
	err error,
) {
	keyWithPrefix := append(_keyPrefix, _keyTarget...)
	val, err := t.Get(keyWithPrefix)
	if err != nil {
		return err
	}
	return _fn_cb_read(_keyPrefix, _keyTarget, val)
}

func (t *Tx_view) Range__lt(
	_keyPrefix []byte,
	_keyTarget []byte,
	_keyNext []byte,
	_dir TD_dir,
	_keyOnly TD_keyOnly,
	_limit int,
	_fn_cb_read func(_prefix []byte, _key []byte, _val []byte) error,
) (
	bt_key__next []byte,
	err error,
) {
	return t.Range(_keyPrefix, nil, _keyTarget, _keyNext, false, true, _dir, _keyOnly, _limit, _fn_cb_read)
}

func (t *Tx_view) Range__le(
	_keyPrefix []byte,
	_keyTarget []byte,
	_keyNext []byte,
	_dir TD_dir,
	_keyOnly TD_keyOnly,
	_limit int,
	_fn_cb_read func(_prefix []byte, _key []byte, _val []byte) error,
) (
	bt_key__next []byte,
	err error,
) {
	return t.Range(_keyPrefix, nil, _keyTarget, _keyNext, false, false, _dir, _keyOnly, _limit, _fn_cb_read)
}

func (t *Tx_view) Range_gt(
	_keyPrefix []byte,
	_keyTarget []byte,
	_keyNext []byte,
	_dir TD_dir,
	_keyOnly TD_keyOnly,
	_limit int,
	_fn_cb_read func(_prefix []byte, _key []byte, _val []byte) error,
) (
	keyNext []byte,
	err error,
) {
	return t.Range(_keyPrefix, _keyTarget, nil, _keyNext, true, false, _dir, _keyOnly, _limit, _fn_cb_read)
}

func (t *Tx_view) Range_ge(
	_keyPrefix []byte,
	_keyTarget []byte,
	_keyNext []byte,
	_dir TD_dir,
	_keyOnly TD_keyOnly,
	_limit int,
	_fn_cb_read func(_prefix []byte, _key []byte, _val []byte) error,
) (
	keyNext []byte,
	err error,
) {
	return t.Range(_keyPrefix, _keyTarget, nil, _keyNext, false, false, _dir, _keyOnly, _limit, _fn_cb_read)
}

//---------------------------------------------------------------------------------------//

func (t *Tx_view) Range(
	_keyPrefix []byte,
	_keyStart []byte,
	_keyEnd []byte,
	_keyNext []byte,
	_isWithoutStart bool,
	_isWithoutEnd bool,
	_dir TD_dir,
	_keyOnly TD_keyOnly,
	_limit int,
	_fn_cb_read func(_prefix []byte, _key []byte, _val []byte) error,
) (
	keyNext []byte,
	err error,
) {
	var it *badger.Iterator
	var n_cmp_end int
	{
		opt_it := badger.DefaultIteratorOptions

		// get type
		{
			switch _keyOnly {
			case TD_keyOnly_true:
				opt_it.PrefetchValues = false
			case TD_keyOnly_false:
				opt_it.PrefetchValues = true
			}
		}

		// dir
		{
			switch _dir {
			case TD_dir_asc:
				opt_it.Reverse = false
				n_cmp_end = 1
			case TD_dir_desc:
				opt_it.Reverse = true
				n_cmp_end = -1

				bt_key__tmp := _keyStart
				_keyStart = _keyEnd
				_keyEnd = bt_key__tmp

				is_without__tmp := _isWithoutStart
				_isWithoutStart = _isWithoutEnd
				_isWithoutEnd = is_without__tmp
			}
		}

		it = t.txn.NewIterator(opt_it)
		defer it.Close()
	}

	// 시작 위치 설정
	{
		if len(_keyNext) == 0 {
			if len(_keyStart) == 0 {
				if len(_keyPrefix) == 0 {
					it.Rewind() // ""
				} else {
					// 역방향인경우 = 역정렬 시작위치를 위해 prefix 내부 마지막 값을 찾는다
					if _dir == TD_dir_desc {
						var bt_key__seek []byte
						bt_key__seek = Util__conv_to_next_value(_keyPrefix)
						it.Seek(bt_key__seek)
					} else {
						it.Seek(_keyPrefix) // "[ prefix ]"
					}
				}
			} else {
				if len(_keyPrefix) == 0 {
					it.Seek(_keyStart) // "[ start ]"
				} else {
					_bt_prefix__with__start := append(_keyPrefix, _keyStart...)
					it.Seek(_bt_prefix__with__start) // "[ prefix ][ start ]"
				}
			}
		} else {
			// next key 는 이전 조회의 마지막 key 이기 때문에
			// next key 로 seek() 후 valid 인 경우 next() 를 실행 해야지만
			// 이전 자료의 다음 자료 위치에서 시작하게 된다.
			_bt_prefix__with__next := append(_keyPrefix, _keyNext...)
			it.Seek(_bt_prefix__with__next)

			if it.Valid() == false {
				return nil, nil // next key 가 없는 위치면 자료 없음 반환
			}
			it.Next()
		}
	}

	// 반복문

	var n_cnt int
	n_len_prefix := len(_keyPrefix)
	for ; it.Valid(); it.Next() {
		pt_item := it.Item()
		bt_key_with_prefix := pt_item.Key()

		// check prefix
		if bytes.HasPrefix(bt_key_with_prefix, _keyPrefix) == false {
			break
		}

		bt_key := bt_key_with_prefix[n_len_prefix:] // prefix 제거

		// 스킵 조건 검사
		{
			// start
			{
				// start key 를 제외하는 옵션
				if _isWithoutStart == true && len(_keyStart) != 0 {
					_isWithoutStart = false

					n_cmp_start_now := bytes.Compare(bt_key, _keyStart)
					if n_cmp_start_now == 0 {
						continue
					}
				}
			}
		}

		{
			// end 조건이 있으면서 현재 키가 end 를 넘어서면 종료
			if len(_keyEnd) != 0 {
				n_cmp_end__now := bytes.Compare(bt_key, _keyEnd)

				// end key 제외 옵션 처리
				if _isWithoutEnd == true && n_cmp_end__now == 0 {
					return nil, nil
				}

				// 현재키가 end key 를 넘어서면 종료
				if n_cmp_end__now == n_cmp_end {
					return nil, nil
				}
			}
		}

		var value []byte

		switch _keyOnly {
		case TD_keyOnly_true:
			value = nil
		case TD_keyOnly_false:
			err = pt_item.Value(func(_value []byte) error {
				value = _value
				return nil
			})
			if err != nil {
				return nil, err
			}
		}

		err = _fn_cb_read(_keyPrefix, bt_key, value)
		if err != nil {
			return nil, err
		}
		n_cnt++
		if _limit == n_cnt {
			keyNext = bt_key
			break
		}
	}

	return bt_key__next, nil
}
