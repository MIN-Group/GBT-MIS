package MetaData

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BlockGroup) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "height":
			z.Height, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "generator":
			z.Generator, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "Generator")
				return
			}
		case "preHash":
			z.PreviousHash, err = dc.ReadBytes(z.PreviousHash)
			if err != nil {
				err = msgp.WrapError(err, "PreviousHash")
				return
			}
		case "merkleRoot":
			z.MerkleRoot, err = dc.ReadBytes(z.MerkleRoot)
			if err != nil {
				err = msgp.WrapError(err, "MerkleRoot")
				return
			}
		case "timestamp":
			z.Timestamp, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "Timestamp")
				return
			}
		case "VoteResult":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "VoteResult")
				return
			}
			if cap(z.VoteResult) >= int(zb0002) {
				z.VoteResult = (z.VoteResult)[:zb0002]
			} else {
				z.VoteResult = make([]uint8, zb0002)
			}
			for za0001 := range z.VoteResult {
				z.VoteResult[za0001], err = dc.ReadUint8()
				if err != nil {
					err = msgp.WrapError(err, "VoteResult", za0001)
					return
				}
			}
		case "BlockHashes":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "BlockHashes")
				return
			}
			if cap(z.BlockHashes) >= int(zb0003) {
				z.BlockHashes = (z.BlockHashes)[:zb0003]
			} else {
				z.BlockHashes = make([][]byte, zb0003)
			}
			for za0002 := range z.BlockHashes {
				z.BlockHashes[za0002], err = dc.ReadBytes(z.BlockHashes[za0002])
				if err != nil {
					err = msgp.WrapError(err, "BlockHashes", za0002)
					return
				}
			}
		case "NextDutyWorker":
			z.NextDutyWorker, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "NextDutyWorker")
				return
			}
		case "voteAggSign":
			z.VoteAggSign, err = dc.ReadBytes(z.VoteAggSign)
			if err != nil {
				err = msgp.WrapError(err, "VoteAggSign")
				return
			}
		case "Sig":
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
func (z *BlockGroup) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "height"
	err = en.Append(0x8a, 0xa6, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Height)
	if err != nil {
		err = msgp.WrapError(err, "Height")
		return
	}
	// write "generator"
	err = en.Append(0xa9, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Generator)
	if err != nil {
		err = msgp.WrapError(err, "Generator")
		return
	}
	// write "preHash"
	err = en.Append(0xa7, 0x70, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.PreviousHash)
	if err != nil {
		err = msgp.WrapError(err, "PreviousHash")
		return
	}
	// write "merkleRoot"
	err = en.Append(0xaa, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.MerkleRoot)
	if err != nil {
		err = msgp.WrapError(err, "MerkleRoot")
		return
	}
	// write "timestamp"
	err = en.Append(0xa9, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Timestamp)
	if err != nil {
		err = msgp.WrapError(err, "Timestamp")
		return
	}
	// write "VoteResult"
	err = en.Append(0xaa, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.VoteResult)))
	if err != nil {
		err = msgp.WrapError(err, "VoteResult")
		return
	}
	for za0001 := range z.VoteResult {
		err = en.WriteUint8(z.VoteResult[za0001])
		if err != nil {
			err = msgp.WrapError(err, "VoteResult", za0001)
			return
		}
	}
	// write "BlockHashes"
	err = en.Append(0xab, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.BlockHashes)))
	if err != nil {
		err = msgp.WrapError(err, "BlockHashes")
		return
	}
	for za0002 := range z.BlockHashes {
		err = en.WriteBytes(z.BlockHashes[za0002])
		if err != nil {
			err = msgp.WrapError(err, "BlockHashes", za0002)
			return
		}
	}
	// write "NextDutyWorker"
	err = en.Append(0xae, 0x4e, 0x65, 0x78, 0x74, 0x44, 0x75, 0x74, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.NextDutyWorker)
	if err != nil {
		err = msgp.WrapError(err, "NextDutyWorker")
		return
	}
	// write "voteAggSign"
	err = en.Append(0xab, 0x76, 0x6f, 0x74, 0x65, 0x41, 0x67, 0x67, 0x53, 0x69, 0x67, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.VoteAggSign)
	if err != nil {
		err = msgp.WrapError(err, "VoteAggSign")
		return
	}
	// write "Sig"
	err = en.Append(0xa3, 0x53, 0x69, 0x67)
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
func (z *BlockGroup) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "height"
	o = append(o, 0x8a, 0xa6, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendInt(o, z.Height)
	// string "generator"
	o = append(o, 0xa9, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72)
	o = msgp.AppendUint32(o, z.Generator)
	// string "preHash"
	o = append(o, 0xa7, 0x70, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68)
	o = msgp.AppendBytes(o, z.PreviousHash)
	// string "merkleRoot"
	o = append(o, 0xaa, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74)
	o = msgp.AppendBytes(o, z.MerkleRoot)
	// string "timestamp"
	o = append(o, 0xa9, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70)
	o = msgp.AppendFloat64(o, z.Timestamp)
	// string "VoteResult"
	o = append(o, 0xaa, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.VoteResult)))
	for za0001 := range z.VoteResult {
		o = msgp.AppendUint8(o, z.VoteResult[za0001])
	}
	// string "BlockHashes"
	o = append(o, 0xab, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.BlockHashes)))
	for za0002 := range z.BlockHashes {
		o = msgp.AppendBytes(o, z.BlockHashes[za0002])
	}
	// string "NextDutyWorker"
	o = append(o, 0xae, 0x4e, 0x65, 0x78, 0x74, 0x44, 0x75, 0x74, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72)
	o = msgp.AppendUint32(o, z.NextDutyWorker)
	// string "voteAggSign"
	o = append(o, 0xab, 0x76, 0x6f, 0x74, 0x65, 0x41, 0x67, 0x67, 0x53, 0x69, 0x67, 0x6e)
	o = msgp.AppendBytes(o, z.VoteAggSign)
	// string "Sig"
	o = append(o, 0xa3, 0x53, 0x69, 0x67)
	o = msgp.AppendString(o, z.Sig)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BlockGroup) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "height":
			z.Height, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "generator":
			z.Generator, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Generator")
				return
			}
		case "preHash":
			z.PreviousHash, bts, err = msgp.ReadBytesBytes(bts, z.PreviousHash)
			if err != nil {
				err = msgp.WrapError(err, "PreviousHash")
				return
			}
		case "merkleRoot":
			z.MerkleRoot, bts, err = msgp.ReadBytesBytes(bts, z.MerkleRoot)
			if err != nil {
				err = msgp.WrapError(err, "MerkleRoot")
				return
			}
		case "timestamp":
			z.Timestamp, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Timestamp")
				return
			}
		case "VoteResult":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "VoteResult")
				return
			}
			if cap(z.VoteResult) >= int(zb0002) {
				z.VoteResult = (z.VoteResult)[:zb0002]
			} else {
				z.VoteResult = make([]uint8, zb0002)
			}
			for za0001 := range z.VoteResult {
				z.VoteResult[za0001], bts, err = msgp.ReadUint8Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "VoteResult", za0001)
					return
				}
			}
		case "BlockHashes":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BlockHashes")
				return
			}
			if cap(z.BlockHashes) >= int(zb0003) {
				z.BlockHashes = (z.BlockHashes)[:zb0003]
			} else {
				z.BlockHashes = make([][]byte, zb0003)
			}
			for za0002 := range z.BlockHashes {
				z.BlockHashes[za0002], bts, err = msgp.ReadBytesBytes(bts, z.BlockHashes[za0002])
				if err != nil {
					err = msgp.WrapError(err, "BlockHashes", za0002)
					return
				}
			}
		case "NextDutyWorker":
			z.NextDutyWorker, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NextDutyWorker")
				return
			}
		case "voteAggSign":
			z.VoteAggSign, bts, err = msgp.ReadBytesBytes(bts, z.VoteAggSign)
			if err != nil {
				err = msgp.WrapError(err, "VoteAggSign")
				return
			}
		case "Sig":
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
func (z *BlockGroup) Msgsize() (s int) {
	s = 1 + 7 + msgp.IntSize + 10 + msgp.Uint32Size + 8 + msgp.BytesPrefixSize + len(z.PreviousHash) + 11 + msgp.BytesPrefixSize + len(z.MerkleRoot) + 10 + msgp.Float64Size + 11 + msgp.ArrayHeaderSize + (len(z.VoteResult) * (msgp.Uint8Size)) + 12 + msgp.ArrayHeaderSize
	for za0002 := range z.BlockHashes {
		s += msgp.BytesPrefixSize + len(z.BlockHashes[za0002])
	}
	s += 15 + msgp.Uint32Size + 12 + msgp.BytesPrefixSize + len(z.VoteAggSign) + 4 + msgp.StringPrefixSize + len(z.Sig)
	return
}
