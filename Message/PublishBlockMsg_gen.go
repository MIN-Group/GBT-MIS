package Message

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *PublishBlockMsg) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Height":
			z.Height, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "BlockNum":
			z.Block_num, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "Block_num")
				return
			}
		case "Block":
			z.Block, err = dc.ReadBytes(z.Block)
			if err != nil {
				err = msgp.WrapError(err, "Block")
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
func (z *PublishBlockMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Height"
	err = en.Append(0x83, 0xa6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Height)
	if err != nil {
		err = msgp.WrapError(err, "Height")
		return
	}
	// write "BlockNum"
	err = en.Append(0xa8, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Block_num)
	if err != nil {
		err = msgp.WrapError(err, "Block_num")
		return
	}
	// write "Block"
	err = en.Append(0xa5, 0x42, 0x6c, 0x6f, 0x63, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Block)
	if err != nil {
		err = msgp.WrapError(err, "Block")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PublishBlockMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Height"
	o = append(o, 0x83, 0xa6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendInt(o, z.Height)
	// string "BlockNum"
	o = append(o, 0xa8, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d)
	o = msgp.AppendUint32(o, z.Block_num)
	// string "Block"
	o = append(o, 0xa5, 0x42, 0x6c, 0x6f, 0x63, 0x6b)
	o = msgp.AppendBytes(o, z.Block)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PublishBlockMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Height":
			z.Height, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "BlockNum":
			z.Block_num, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Block_num")
				return
			}
		case "Block":
			z.Block, bts, err = msgp.ReadBytesBytes(bts, z.Block)
			if err != nil {
				err = msgp.WrapError(err, "Block")
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
func (z *PublishBlockMsg) Msgsize() (s int) {
	s = 1 + 7 + msgp.IntSize + 9 + msgp.Uint32Size + 6 + msgp.BytesPrefixSize + len(z.Block)
	return
}
