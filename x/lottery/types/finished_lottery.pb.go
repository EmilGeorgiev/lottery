// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lottery/finished_lottery.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FinishedLottery struct {
	Index           string            `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Winner          string            `protobuf:"bytes,2,opt,name=winner,proto3" json:"winner,omitempty"`
	Reward          uint64            `protobuf:"varint,3,opt,name=reward,proto3" json:"reward,omitempty"`
	EnterLotteryTxs []*EnterLotteryTx `protobuf:"bytes,4,rep,name=enter_lottery_txs,json=enterLotteryTxs,proto3" json:"enter_lottery_txs,omitempty"`
	WinnerIndex     uint64            `protobuf:"varint,5,opt,name=winner_index,json=winnerIndex,proto3" json:"winner_index,omitempty"`
}

func (m *FinishedLottery) Reset()         { *m = FinishedLottery{} }
func (m *FinishedLottery) String() string { return proto.CompactTextString(m) }
func (*FinishedLottery) ProtoMessage()    {}
func (*FinishedLottery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20ea49878d19341, []int{0}
}
func (m *FinishedLottery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FinishedLottery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FinishedLottery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FinishedLottery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinishedLottery.Merge(m, src)
}
func (m *FinishedLottery) XXX_Size() int {
	return m.Size()
}
func (m *FinishedLottery) XXX_DiscardUnknown() {
	xxx_messageInfo_FinishedLottery.DiscardUnknown(m)
}

var xxx_messageInfo_FinishedLottery proto.InternalMessageInfo

func (m *FinishedLottery) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *FinishedLottery) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *FinishedLottery) GetReward() uint64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *FinishedLottery) GetEnterLotteryTxs() []*EnterLotteryTx {
	if m != nil {
		return m.EnterLotteryTxs
	}
	return nil
}

func (m *FinishedLottery) GetWinnerIndex() uint64 {
	if m != nil {
		return m.WinnerIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*FinishedLottery)(nil), "emilgeorgiev.lottery.lottery.FinishedLottery")
}

func init() { proto.RegisterFile("lottery/finished_lottery.proto", fileDescriptor_b20ea49878d19341) }

var fileDescriptor_b20ea49878d19341 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0xc9, 0x2f, 0x29,
	0x49, 0x2d, 0xaa, 0xd4, 0x4f, 0xcb, 0xcc, 0xcb, 0x2c, 0xce, 0x48, 0x4d, 0x89, 0x87, 0x0a, 0xe8,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xc9, 0xa4, 0xe6, 0x66, 0xe6, 0xa4, 0xa7, 0xe6, 0x17, 0xa5,
	0x67, 0xa6, 0x96, 0xe9, 0xc1, 0xe4, 0xa0, 0xb4, 0x94, 0x28, 0x4c, 0x37, 0x8a, 0x26, 0xa5, 0x2b,
	0x8c, 0x5c, 0xfc, 0x6e, 0x50, 0xf3, 0x7c, 0x20, 0x32, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x79, 0x29,
	0xa9, 0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x18, 0x17, 0x5b, 0x79,
	0x66, 0x5e, 0x5e, 0x6a, 0x91, 0x04, 0x13, 0x58, 0x18, 0xca, 0x03, 0x89, 0x17, 0xa5, 0x96, 0x27,
	0x16, 0xa5, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0xb0, 0x04, 0x41, 0x79, 0x42, 0x11, 0x5c, 0x82, 0xa9,
	0x79, 0x25, 0xa9, 0x45, 0x30, 0x57, 0xc6, 0x97, 0x54, 0x14, 0x4b, 0xb0, 0x28, 0x30, 0x6b, 0x70,
	0x1b, 0xe9, 0xe8, 0xe1, 0x73, 0xaa, 0x9e, 0x2b, 0x48, 0x1b, 0xd4, 0x31, 0x21, 0x15, 0x41, 0xfc,
	0xa9, 0x28, 0xfc, 0x62, 0x21, 0x45, 0x2e, 0x1e, 0x88, 0xdd, 0xf1, 0x10, 0x67, 0xb2, 0x82, 0xed,
	0xe5, 0x86, 0x88, 0x79, 0x82, 0x84, 0x9c, 0x3c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e,
	0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58,
	0x8e, 0x21, 0x4a, 0x3f, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0x35,
	0x37, 0x33, 0xc7, 0x1d, 0xea, 0x0a, 0x58, 0xb8, 0xe8, 0x57, 0xc0, 0x59, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0xe0, 0x80, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xa2, 0xb9, 0x72,
	0x7f, 0x01, 0x00, 0x00,
}

func (m *FinishedLottery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FinishedLottery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FinishedLottery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WinnerIndex != 0 {
		i = encodeVarintFinishedLottery(dAtA, i, uint64(m.WinnerIndex))
		i--
		dAtA[i] = 0x28
	}
	if len(m.EnterLotteryTxs) > 0 {
		for iNdEx := len(m.EnterLotteryTxs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EnterLotteryTxs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFinishedLottery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Reward != 0 {
		i = encodeVarintFinishedLottery(dAtA, i, uint64(m.Reward))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintFinishedLottery(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintFinishedLottery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFinishedLottery(dAtA []byte, offset int, v uint64) int {
	offset -= sovFinishedLottery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FinishedLottery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovFinishedLottery(uint64(l))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovFinishedLottery(uint64(l))
	}
	if m.Reward != 0 {
		n += 1 + sovFinishedLottery(uint64(m.Reward))
	}
	if len(m.EnterLotteryTxs) > 0 {
		for _, e := range m.EnterLotteryTxs {
			l = e.Size()
			n += 1 + l + sovFinishedLottery(uint64(l))
		}
	}
	if m.WinnerIndex != 0 {
		n += 1 + sovFinishedLottery(uint64(m.WinnerIndex))
	}
	return n
}

func sovFinishedLottery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFinishedLottery(x uint64) (n int) {
	return sovFinishedLottery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FinishedLottery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFinishedLottery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FinishedLottery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FinishedLottery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinishedLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFinishedLottery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFinishedLottery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinishedLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFinishedLottery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFinishedLottery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			m.Reward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinishedLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnterLotteryTxs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinishedLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFinishedLottery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFinishedLottery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnterLotteryTxs = append(m.EnterLotteryTxs, &EnterLotteryTx{})
			if err := m.EnterLotteryTxs[len(m.EnterLotteryTxs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerIndex", wireType)
			}
			m.WinnerIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinishedLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WinnerIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFinishedLottery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFinishedLottery
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
func skipFinishedLottery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFinishedLottery
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
					return 0, ErrIntOverflowFinishedLottery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFinishedLottery
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
			if length < 0 {
				return 0, ErrInvalidLengthFinishedLottery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFinishedLottery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFinishedLottery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFinishedLottery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFinishedLottery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFinishedLottery = fmt.Errorf("proto: unexpected end of group")
)
