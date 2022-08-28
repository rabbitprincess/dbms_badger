package db_badger

import (
	"log"
	"time"

	badger "github.com/dgraph-io/badger/v2"
)

type Badger struct {
	badger *badger.DB

	chanGCEnd chan bool
}

func (t *Badger) Open(_dbpath string) error {
	var err error
	t.badger, err = badger.Open(badger.DefaultOptions(_dbpath))
	if err != nil {
		return err
	}

	return nil
}

func (t *Badger) RunGC(_ratio float64) {
	t.badger.RunValueLogGC(_ratio)
}

func (t *Badger) runGCLoop(_ratio float64, _intervalSec int) {
	// 사용안함 설정 시 바로 종료
	{
		if _ratio == 0.0 || _intervalSec == 0 {
			return
		}
	}

	// DB 가 준비되어 있지 않으면 종료
	{
		if t.badger == nil {
			log.Fatal("db is not ready")
		}
	}

	t.chanGCEnd = make(chan bool)

	go func() {
		isEnd := false
		for isEnd == false {
			select {
			case <-t.chanGCEnd:
				{
					// 종료 요청이 있으면 루프 탈출
					isEnd = true
				}
			case <-time.After(time.Duration(_intervalSec) * time.Second):
				{
					// 지정 시간이 경과 할 때 마다 GC 수행
					t.RunGC(_ratio)
				}
			}
		}
		t.chanGCEnd <- true // 채널에 종료 사실을 알림
	}()
}

func (t *Badger) Close() error {
	// gc 옵션이 켜져 있는경우만
	if t.chanGCEnd != nil {
		t.chanGCEnd <- true // 종료 시작 알림
		<-t.chanGCEnd       // 종료 완료 수신 까지 대기
	}
	t.badger.Close()
	return nil
}

func (t *Badger) Info() []badger.TableInfo {
	TableInfo := t.badger.Tables(true)
	return TableInfo
}

func (t *Badger) DropAll() error {
	return t.badger.DropAll()
}

func (t *Badger) SequenceGet(_key []byte, _bandwidth uint64) (*badger.Sequence, error) {
	return t.badger.GetSequence(_key, _bandwidth)
}

func (t *Badger) TX_view(_fn_cb func(_view *TxView) error) error {
	fn := func(_tx *badger.Txn) error {
		pt_tx__view := &TxView{}
		pt_tx__view.Init(t, _tx)
		return _fn_cb(pt_tx__view)
	}
	return t.badger.View(fn)
}

func (t *Badger) TX_update(_fn_cb func(_update *TxUpdate) error) error {
	fn := func(_tx *badger.Txn) error {
		pt_tx__update := &TxUpdate{}
		pt_tx__update.Init(t, _tx)
		return _fn_cb(pt_tx__update)
	}
	return t.badger.Update(fn)
}
