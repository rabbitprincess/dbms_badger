package dbms

import (
	"github.com/gokch/dbms_badger/engine"
	"github.com/gokch/dbms_badger/schema"
)

type DBMS struct {
	schema *schema.Schema
	engine *engine.Engine
}

func (t *DBMS) Init() {

}

// 초기화 시 engine 의 kv 저장소에 저장되어있는 information_schema 에 따라 테이블 스키마 복구
// 시나리오 정리 - 테이블과 인덱스 추가
// 테이블에 레코드 추가
// 테이블에 레코드 조회
// 테이블에 레코드 수정
// 테이블에 레코드 삭제
// 인덱스 추가 ( 이미 존재하는 kv 를 )
