package db_badger

//------------------------------------------------------------------------------------//
// base

type Tx_update struct {
	Tx_view
}

func (t *Tx_update) Set(_key, _val []byte) error {
	return t.txn.Set(_key, _val)
}

func (t *Tx_update) SetStr(_key, _val string) error {
	return t.Set([]byte(_key), []byte(_val))
}

func (t *Tx_update) SetIfNotExist(_key, _val []byte) error {
	isExist, err := t.IsExist(_key)
	if err != nil {
		return err
	}

	// 키 없음이 아니면 오류
	if isExist == true {
		return ErrKeyAlreadyExist
	}

	// 키가 존재하지 않으면 set
	return t.Set(_key, _val)
}

func (t *Tx_update) Delete(_key []byte) error {
	return t.txn.Delete(_key)
}

func (t *Tx_update) Delete__if_exist(_key []byte) error {
	_, err := t.Get(_key)
	if err != nil {
		return err
	}
	return t.Delete(_key)
}
