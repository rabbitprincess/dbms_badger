package dbms

import (
	"encoding/binary"

	"github.com/gokch/dbms_badger/engine"
	"github.com/gokch/dbms_badger/schema"
)

type DBMS struct {
	engine *engine.Engine
	schema *schema.Schema
}

func (t *DBMS) Init() error {
	// TODO: input path or badger option
	err := t.InitEngine("")
	if err != nil {
		return err
	}

	err = t.InitSchema()
	if err != nil {
		return err
	}
	return nil
}

func (t *DBMS) InitEngine(path string) error {
	t.engine = engine.NewEngine()

	// open engine
	err := t.engine.Open(path)
	if err != nil {
		return err
	}
	return nil
}

func (t *DBMS) InitSchema() error {
	t.schema = schema.NewSchema()

	// get information schema from db
	return t.engine.TxView(func(view *engine.TxView) error {
		keyInformationSchema := make([]byte, 0, 3)
		keyInformationSchema = append(keyInformationSchema, "T"...)
		keyInformationSchema = binary.LittleEndian.AppendUint16(keyInformationSchema, uint16(schema.SeqTblInformationSchema))
		value, err := view.Get(keyInformationSchema)
		if err != nil {
			return err
		}
		if value == nil {
			return nil
		}
		// load schema
		err = t.schema.Load(value)
		if err != nil {
			return err
		}
		return nil
	})
}

func (t *DBMS) UpdateSchema() error {
	return t.engine.TxUpdate(func(update *engine.TxUpdate) error {
		keyInformationSchema := make([]byte, 0, 3)
		keyInformationSchema = append(keyInformationSchema, "T"...)
		keyInformationSchema = binary.LittleEndian.AppendUint16(keyInformationSchema, uint16(schema.SeqTblInformationSchema))
		value, err := t.schema.Save()
		if err != nil {
			return err
		}
		err = update.Set(keyInformationSchema, value)
		if err != nil {
			return err
		}
		return nil
	})
}

// 초기화 시 engine 의 kv 저장소에 저장되어있는 information_schema 에 따라 테이블 스키마 복구
// 시나리오 정리 - 테이블과 인덱스 추가
// 테이블에 레코드 추가
// 테이블에 레코드 조회
// 테이블에 레코드 수정
// 테이블에 레코드 삭제
// 인덱스 추가 ( 이미 존재하는 kv 를 )
