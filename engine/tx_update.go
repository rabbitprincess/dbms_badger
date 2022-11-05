package engine

//------------------------------------------------------------------------------------//
// base

type TxUpdate struct {
	TxView
}

func (t *TxUpdate) Set(key, val []byte) error {
	return t.Txn.Set(key, val)
}

func (t *TxUpdate) SetStr(key, val string) error {
	return t.Set([]byte(key), []byte(val))
}

func (t *TxUpdate) SetIfNotExist(key, val []byte) error {
	isExist, err := t.IsExist(key)
	if err != nil {
		return err
	}

	// 키 없음이 아니면 오류
	if isExist == true {
		return ErrKeyAlreadyExist
	}

	// 키가 존재하지 않으면 set
	return t.Set(key, val)
}

func (t *TxUpdate) Delete(key []byte) error {
	return t.Txn.Delete(key)
}

func (t *TxUpdate) DeleteIfExist(key []byte) error {
	_, err := t.Get(key)
	if err != nil {
		return err
	}
	return t.Delete(key)
}
