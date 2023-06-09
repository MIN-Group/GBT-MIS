package MetaData

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *CreatAccount) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "address":
			z.Address, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Address")
				return
			}
		case "pubkey":
			z.Pubkey, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Pubkey")
				return
			}
		case "timestamp":
			z.Timestamp, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Timestamp")
				return
			}
		case "sig":
			z.Sig, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Sig")
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
func (z *CreatAccount) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "address"
	err = en.Append(0x84, 0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	if err != nil {
		return
	}
	err = en.WriteString(z.Address)
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	// write "pubkey"
	err = en.Append(0xa6, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79)
	if err != nil {
		return
	}
	err = en.WriteString(z.Pubkey)
	if err != nil {
		err = msgp.WrapError(err, "Pubkey")
		return
	}
	// write "timestamp"
	err = en.Append(0xa9, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return
	}
	err = en.WriteString(z.Timestamp)
	if err != nil {
		err = msgp.WrapError(err, "Timestamp")
		return
	}
	// write "sig"
	err = en.Append(0xa3, 0x73, 0x69, 0x67)
	if err != nil {
		return
	}
	err = en.WriteString(z.Sig)
	if err != nil {
		err = msgp.WrapError(err, "Sig")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CreatAccount) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "address"
	o = append(o, 0x84, 0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	o = msgp.AppendString(o, z.Address)
	// string "pubkey"
	o = append(o, 0xa6, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Pubkey)
	// string "timestamp"
	o = append(o, 0xa9, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendString(o, z.Timestamp)
	// string "sig"
	o = append(o, 0xa3, 0x73, 0x69, 0x67)
	o = msgp.AppendString(o, z.Sig)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CreatAccount) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "address":
			z.Address, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Address")
				return
			}
		case "pubkey":
			z.Pubkey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Pubkey")
				return
			}
		case "timestamp":
			z.Timestamp, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Timestamp")
				return
			}
		case "sig":
			z.Sig, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Sig")
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
func (z *CreatAccount) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Address) + 7 + msgp.StringPrefixSize + len(z.Pubkey) + 10 + msgp.StringPrefixSize + len(z.Timestamp) + 4 + msgp.StringPrefixSize + len(z.Sig)
	return
}
