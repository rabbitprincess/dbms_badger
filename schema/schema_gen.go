package schema

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Field) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Value":
			z.Value, err = dc.ReadIntf()
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Field) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Seq"
	err = en.Append(0x83, 0xa3, 0x53, 0x65, 0x71)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Seq)
	if err != nil {
		err = msgp.WrapError(err, "Seq")
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "Value"
	err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteIntf(z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Field) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Seq"
	o = append(o, 0x83, 0xa3, 0x53, 0x65, 0x71)
	o = msgp.AppendInt(o, z.Seq)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o, err = msgp.AppendIntf(o, z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Field) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Value":
			z.Value, bts, err = msgp.ReadIntfBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Field) Msgsize() (s int) {
	s = 1 + 4 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 6 + msgp.GuessSize(z.Value)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Index) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "Type":
			err = z.Type.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Fields":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Fields")
				return
			}
			if cap(z.Fields) >= int(zb0002) {
				z.Fields = (z.Fields)[:zb0002]
			} else {
				z.Fields = make([]*Field, zb0002)
			}
			for za0001 := range z.Fields {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Fields", za0001)
						return
					}
					z.Fields[za0001] = nil
				} else {
					if z.Fields[za0001] == nil {
						z.Fields[za0001] = new(Field)
					}
					var zb0003 uint32
					zb0003, err = dc.ReadMapHeader()
					if err != nil {
						err = msgp.WrapError(err, "Fields", za0001)
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							err = msgp.WrapError(err, "Fields", za0001)
							return
						}
						switch msgp.UnsafeString(field) {
						case "Seq":
							z.Fields[za0001].Seq, err = dc.ReadInt()
							if err != nil {
								err = msgp.WrapError(err, "Fields", za0001, "Seq")
								return
							}
						case "Name":
							z.Fields[za0001].Name, err = dc.ReadString()
							if err != nil {
								err = msgp.WrapError(err, "Fields", za0001, "Name")
								return
							}
						case "Value":
							z.Fields[za0001].Value, err = dc.ReadIntf()
							if err != nil {
								err = msgp.WrapError(err, "Fields", za0001, "Value")
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								err = msgp.WrapError(err, "Fields", za0001)
								return
							}
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Index) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Seq"
	err = en.Append(0x84, 0xa3, 0x53, 0x65, 0x71)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Seq)
	if err != nil {
		err = msgp.WrapError(err, "Seq")
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = z.Type.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "Fields"
	err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Fields)))
	if err != nil {
		err = msgp.WrapError(err, "Fields")
		return
	}
	for za0001 := range z.Fields {
		if z.Fields[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 3
			// write "Seq"
			err = en.Append(0x83, 0xa3, 0x53, 0x65, 0x71)
			if err != nil {
				return
			}
			err = en.WriteInt(z.Fields[za0001].Seq)
			if err != nil {
				err = msgp.WrapError(err, "Fields", za0001, "Seq")
				return
			}
			// write "Name"
			err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
			if err != nil {
				return
			}
			err = en.WriteString(z.Fields[za0001].Name)
			if err != nil {
				err = msgp.WrapError(err, "Fields", za0001, "Name")
				return
			}
			// write "Value"
			err = en.Append(0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			if err != nil {
				return
			}
			err = en.WriteIntf(z.Fields[za0001].Value)
			if err != nil {
				err = msgp.WrapError(err, "Fields", za0001, "Value")
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Index) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Seq"
	o = append(o, 0x84, 0xa3, 0x53, 0x65, 0x71)
	o = msgp.AppendInt(o, z.Seq)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o, err = z.Type.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Fields"
	o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fields)))
	for za0001 := range z.Fields {
		if z.Fields[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 3
			// string "Seq"
			o = append(o, 0x83, 0xa3, 0x53, 0x65, 0x71)
			o = msgp.AppendInt(o, z.Fields[za0001].Seq)
			// string "Name"
			o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.Fields[za0001].Name)
			// string "Value"
			o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
			o, err = msgp.AppendIntf(o, z.Fields[za0001].Value)
			if err != nil {
				err = msgp.WrapError(err, "Fields", za0001, "Value")
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Index) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "Type":
			bts, err = z.Type.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Fields":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Fields")
				return
			}
			if cap(z.Fields) >= int(zb0002) {
				z.Fields = (z.Fields)[:zb0002]
			} else {
				z.Fields = make([]*Field, zb0002)
			}
			for za0001 := range z.Fields {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Fields[za0001] = nil
				} else {
					if z.Fields[za0001] == nil {
						z.Fields[za0001] = new(Field)
					}
					var zb0003 uint32
					zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Fields", za0001)
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "Fields", za0001)
							return
						}
						switch msgp.UnsafeString(field) {
						case "Seq":
							z.Fields[za0001].Seq, bts, err = msgp.ReadIntBytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "Fields", za0001, "Seq")
								return
							}
						case "Name":
							z.Fields[za0001].Name, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "Fields", za0001, "Name")
								return
							}
						case "Value":
							z.Fields[za0001].Value, bts, err = msgp.ReadIntfBytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "Fields", za0001, "Value")
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								err = msgp.WrapError(err, "Fields", za0001)
								return
							}
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Index) Msgsize() (s int) {
	s = 1 + 4 + msgp.IntSize + 5 + z.Type.Msgsize() + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Fields {
		if z.Fields[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Fields[za0001].Name) + 6 + msgp.GuessSize(z.Fields[za0001].Value)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Schema) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tables":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Tables")
				return
			}
			if cap(z.Tables) >= int(zb0002) {
				z.Tables = (z.Tables)[:zb0002]
			} else {
				z.Tables = make([]*Table, zb0002)
			}
			for za0001 := range z.Tables {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Tables", za0001)
						return
					}
					z.Tables[za0001] = nil
				} else {
					if z.Tables[za0001] == nil {
						z.Tables[za0001] = new(Table)
					}
					err = z.Tables[za0001].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Tables", za0001)
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Schema) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Tables"
	err = en.Append(0x81, 0xa6, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Tables)))
	if err != nil {
		err = msgp.WrapError(err, "Tables")
		return
	}
	for za0001 := range z.Tables {
		if z.Tables[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Tables[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Tables", za0001)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Schema) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Tables"
	o = append(o, 0x81, 0xa6, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Tables)))
	for za0001 := range z.Tables {
		if z.Tables[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Tables[za0001].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Tables", za0001)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Schema) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tables":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Tables")
				return
			}
			if cap(z.Tables) >= int(zb0002) {
				z.Tables = (z.Tables)[:zb0002]
			} else {
				z.Tables = make([]*Table, zb0002)
			}
			for za0001 := range z.Tables {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Tables[za0001] = nil
				} else {
					if z.Tables[za0001] == nil {
						z.Tables[za0001] = new(Table)
					}
					bts, err = z.Tables[za0001].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Tables", za0001)
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Tables {
		if z.Tables[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Tables[za0001].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Table) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Primary":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Primary")
					return
				}
				z.Primary = nil
			} else {
				if z.Primary == nil {
					z.Primary = new(Index)
				}
				err = z.Primary.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Primary")
					return
				}
			}
		case "Indexes":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Indexes")
				return
			}
			if cap(z.Indexes) >= int(zb0002) {
				z.Indexes = (z.Indexes)[:zb0002]
			} else {
				z.Indexes = make([]*Index, zb0002)
			}
			for za0001 := range z.Indexes {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "Indexes", za0001)
						return
					}
					z.Indexes[za0001] = nil
				} else {
					if z.Indexes[za0001] == nil {
						z.Indexes[za0001] = new(Index)
					}
					err = z.Indexes[za0001].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "Indexes", za0001)
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Table) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Seq"
	err = en.Append(0x84, 0xa3, 0x53, 0x65, 0x71)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Seq)
	if err != nil {
		err = msgp.WrapError(err, "Seq")
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "Primary"
	err = en.Append(0xa7, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79)
	if err != nil {
		return
	}
	if z.Primary == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Primary.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Primary")
			return
		}
	}
	// write "Indexes"
	err = en.Append(0xa7, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Indexes)))
	if err != nil {
		err = msgp.WrapError(err, "Indexes")
		return
	}
	for za0001 := range z.Indexes {
		if z.Indexes[za0001] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Indexes[za0001].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "Indexes", za0001)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Table) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Seq"
	o = append(o, 0x84, 0xa3, 0x53, 0x65, 0x71)
	o = msgp.AppendInt(o, z.Seq)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Primary"
	o = append(o, 0xa7, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79)
	if z.Primary == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Primary.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Primary")
			return
		}
	}
	// string "Indexes"
	o = append(o, 0xa7, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Indexes)))
	for za0001 := range z.Indexes {
		if z.Indexes[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Indexes[za0001].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "Indexes", za0001)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Table) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Primary":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Primary = nil
			} else {
				if z.Primary == nil {
					z.Primary = new(Index)
				}
				bts, err = z.Primary.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Primary")
					return
				}
			}
		case "Indexes":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Indexes")
				return
			}
			if cap(z.Indexes) >= int(zb0002) {
				z.Indexes = (z.Indexes)[:zb0002]
			} else {
				z.Indexes = make([]*Index, zb0002)
			}
			for za0001 := range z.Indexes {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Indexes[za0001] = nil
				} else {
					if z.Indexes[za0001] == nil {
						z.Indexes[za0001] = new(Index)
					}
					bts, err = z.Indexes[za0001].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Indexes", za0001)
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Table) Msgsize() (s int) {
	s = 1 + 4 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 8
	if z.Primary == nil {
		s += msgp.NilSize
	} else {
		s += z.Primary.Msgsize()
	}
	s += 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Indexes {
		if z.Indexes[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Indexes[za0001].Msgsize()
		}
	}
	return
}