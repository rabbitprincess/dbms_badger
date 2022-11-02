package db_badger

//------------------------------------------------------------------------------------//
// base

type TxUpdate struct {
	TxView
}

func (t *TxUpdate) Set(_key, _val []byte) error {
	return t.txn.Set(_key, _val)
}

func (t *TxUpdate) SetStr(_key, _val string) error {
	return t.Set([]byte(_key), []byte(_val))
}

func (t *TxUpdate) SetIfNotExist(_key, _val []byte) error {
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

func (t *TxUpdate) Delete(_key []byte) error {
	return t.txn.Delete(_key)
}

func (t *TxUpdate) DeleteIfExist(_key []byte) error {
	_, err := t.Get(_key)
	if err != nil {
		return err
	}
	return t.Delete(_key)
}
