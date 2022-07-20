package db_badger

import (
	"fmt"
	"log"
	"time"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
)

type Badger struct {
	badger *badger.DB

	dbPath    string
	chanGCEnd chan bool
}

func (t *Badger) Open(_s_path__db string, _is_debug bool) error {
	return t.Open_Adv(_s_path__db, _is_debug, 0.0, 0, false, 0, 0, false)
}

func (t *Badger) Open_Adv(
	_path string,
	_debug bool,
	_gcRatio float64,
	_gcIntervalSec int,
	_isCompression bool,
	_logSize int,
	_version int,
	_loadingMode bool,
) error {

	var err error

	t.dbPath = _path

	// option
	t_option := badger.DefaultOptions(_path)
	{
		t_option.Truncate = true
		t_option.SyncWrites = true
		t_option.CompactL0OnClose = true
		t_option.NumVersionsToKeep = _version

		if _debug == false {
			t_option.Logger = nil
		}

		if _isCompression == true {
			t_option.Compression = options.Snappy
		}

		if _logSize != 0 {
			t_option.ValueLogFileSize = int64(_logSize) * 1024 * 1024
		}

		if _loadingMode == true {
			t_option.TableLoadingMode = options.FileIO
			t_option.ValueLogLoadingMode = options.FileIO
		}

		t.badger, err = badger.Open(t_option)
		if err != nil {
			return err
		}
	}

	// GC
	if _gcRatio > 0 && _gcIntervalSec > 0 {
		t.runGCLoop(_gcRatio, _gcIntervalSec)
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
		is_end := false
		for is_end == false {
			select {
			case <-t.chanGCEnd:
				{
					// 종료 요청이 있으면 루프 탈출
					is_end = true
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
	// gc thread 종료대기 후 close
	{
		// gc 옵션이 켜져 있는경우만
		if t.chanGCEnd != nil {
			t.chanGCEnd <- true // 종료 시작 알림
			<-t.chanGCEnd       // 종료 완료 수신 까지 대기
		}
	}
	t.badger.Close()
	return nil
}

func (t *Badger) Info() []badger.TableInfo {
	TableInfo := t.badger.Tables(true)
	SizeLSM, SizeVlog := t.badger.Size()
	fmt.Printf("table_info : %d - %d\n", SizeLSM, SizeVlog)
	return TableInfo
}

func (t *Badger) DropAll() error {
	return t.badger.DropAll()
}

func (t *Badger) SequenceGet(_key []byte, _bandwidth uint64) (*badger.Sequence, error) {
	return t.badger.GetSequence(_key, _bandwidth)
}

func (t *Badger) TX_view(_fn_cb func(_view *Tx_view) error) error {
	fn := func(_tx *badger.Txn) error {
		pt_tx__view := &Tx_view{}
		pt_tx__view.Init(t, _tx)
		return _fn_cb(pt_tx__view)
	}
	return t.badger.View(fn)
}

func (t *Badger) TX_update(_fn_cb func(_update *Tx_update) error) error {
	fn := func(_tx *badger.Txn) error {
		pt_tx__update := &Tx_update{}
		pt_tx__update.Init(t, _tx)
		return _fn_cb(pt_tx__update)
	}
	return t.badger.Update(fn)
}
