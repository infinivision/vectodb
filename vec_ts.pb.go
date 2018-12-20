// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vec_ts.proto

/*
	Package vectodb is a generated protocol buffer package.

	It is generated from these files:
		vec_ts.proto

	It has these top-level messages:
		VecTimestamp
*/
package vectodb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VecTimestamp struct {
	Vec      []float32 `protobuf:"fixed32,1,rep,packed,name=Vec,json=vec" json:"Vec,omitempty"`
	ExpireAt int64     `protobuf:"varint,2,opt,name=ExpireAt,json=expireAt,proto3" json:"ExpireAt,omitempty"`
}

func (m *VecTimestamp) Reset()                    { *m = VecTimestamp{} }
func (m *VecTimestamp) String() string            { return proto.CompactTextString(m) }
func (*VecTimestamp) ProtoMessage()               {}
func (*VecTimestamp) Descriptor() ([]byte, []int) { return fileDescriptorVecTs, []int{0} }

func init() {
	proto.RegisterType((*VecTimestamp)(nil), "vectodb.VecTimestamp")
}
func (m *VecTimestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VecTimestamp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Vec) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintVecTs(dAtA, i, uint64(len(m.Vec)*4))
		for _, num := range m.Vec {
			f1 := math.Float32bits(float32(num))
			encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(f1))
			i += 4
		}
	}
	if m.ExpireAt != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintVecTs(dAtA, i, uint64(m.ExpireAt))
	}
	return i, nil
}

func encodeVarintVecTs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VecTimestamp) Size() (n int) {
	var l int
	_ = l
	if len(m.Vec) > 0 {
		n += 1 + sovVecTs(uint64(len(m.Vec)*4)) + len(m.Vec)*4
	}
	if m.ExpireAt != 0 {
		n += 1 + sovVecTs(uint64(m.ExpireAt))
	}
	return n
}

func sovVecTs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVecTs(x uint64) (n int) {
	return sovVecTs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VecTimestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVecTs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VecTimestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VecTimestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 5 {
				var v uint32
				if (iNdEx + 4) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
				v2 := float32(math.Float32frombits(v))
				m.Vec = append(m.Vec, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVecTs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthVecTs
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					if (iNdEx + 4) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
					iNdEx += 4
					v2 := float32(math.Float32frombits(v))
					m.Vec = append(m.Vec, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Vec", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireAt", wireType)
			}
			m.ExpireAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVecTs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVecTs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVecTs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVecTs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVecTs
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVecTs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVecTs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthVecTs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVecTs
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipVecTs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthVecTs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVecTs   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("vec_ts.proto", fileDescriptorVecTs) }

var fileDescriptorVecTs = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4b, 0x4d, 0x8e,
	0x2f, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4b, 0x4d, 0x2e, 0xc9, 0x4f,
	0x49, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xe9, 0x83, 0x58, 0x10, 0x69, 0x25, 0x1b,
	0x2e, 0x9e, 0xb0, 0xd4, 0xe4, 0x90, 0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x21, 0x01,
	0x2e, 0xe6, 0xb0, 0xd4, 0x64, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xa6, 0x20, 0xe6, 0xb2, 0xd4, 0x64,
	0x21, 0x29, 0x2e, 0x0e, 0xd7, 0x8a, 0x82, 0xcc, 0xa2, 0x54, 0xc7, 0x12, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xe6, 0x20, 0x8e, 0x54, 0x28, 0xdf, 0x49, 0xe4, 0xc4, 0x43, 0x39, 0x86, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36,
	0xb0, 0xd1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xce, 0xf0, 0x91, 0x89, 0x00, 0x00,
	0x00,
}
