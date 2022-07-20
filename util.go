package db_badger

const (
	DEF_b1_max = byte(255)
)

func Util__conv_to_next_value(_key []byte) (keyNext []byte) {
	keyNext = Util__fill_0xff_max(_key)

	// 성능향상을 위한 대체 코드 (봉인)
	//	성능을 위해 다음을 완성하여 사용 할 수 있다.
	// 	255 는 fill 0xff 로 처리하고
	//  255 != 는 (마지막 값 + 1) 형식으로 처리
	//		단 (마지막 값 + 1) 과 같은 값은 탐색에서 제외 해야함 ( prefix 가 다름으로 탐색 대상이 아님 )

	// if len(_bt_key) == 0 {
	// 	return nil
	// }
	// b1_key__last := _bt_key[len(_bt_key)-1]
	// if b1_key__last == DEF_b1_max { // last byte == 255
	// 	bt_key__next = Util__fill_0xff_max(_bt_key)
	// } else { // last byte < 255
	// 	bt_key__next = make([]byte, len(_bt_key))
	// 	copy(bt_key__next, _bt_key)
	// 	bt_key__next[len(_bt_key)-1]++
	// }

	return keyNext
}

func Util__fill_0xff_max(_key []byte) []byte {
	bt_key__seek := make([]byte, DEF_maxKeyLen)
	var i int
	for i = 0; i < len(_key); i++ {
		bt_key__seek[i] = _key[i]
	}
	for ; i < len(bt_key__seek); i++ {
		bt_key__seek[i] = DEF_b1_max
	}
	return bt_key__seek
}
