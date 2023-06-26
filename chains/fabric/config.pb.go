// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayer/chains/fabric/config.proto

package fabric

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ChainConfig struct {
	ChainId               string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	WalletLabel           string `protobuf:"bytes,2,opt,name=wallet_label,json=walletLabel,proto3" json:"wallet_label,omitempty"`
	ConnectionProfilePath string `protobuf:"bytes,3,opt,name=connection_profile_path,json=connectionProfilePath,proto3" json:"connection_profile_path,omitempty"`
	Channel               string `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	ChaincodeId           string `protobuf:"bytes,5,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc92cd1bee875088, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

type ProverConfig struct {
	IbcPolicies         []string `protobuf:"bytes,1,rep,name=ibc_policies,json=ibcPolicies,proto3" json:"ibc_policies,omitempty"`
	EndorsementPolicies []string `protobuf:"bytes,2,rep,name=endorsement_policies,json=endorsementPolicies,proto3" json:"endorsement_policies,omitempty"`
	MspConfigPaths      []string `protobuf:"bytes,3,rep,name=msp_config_paths,json=mspConfigPaths,proto3" json:"msp_config_paths,omitempty"`
}

func (m *ProverConfig) Reset()         { *m = ProverConfig{} }
func (m *ProverConfig) String() string { return proto.CompactTextString(m) }
func (*ProverConfig) ProtoMessage()    {}
func (*ProverConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc92cd1bee875088, []int{1}
}
func (m *ProverConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProverConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProverConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProverConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProverConfig.Merge(m, src)
}
func (m *ProverConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProverConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProverConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProverConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainConfig)(nil), "relayer.chains.fabric.config.ChainConfig")
	proto.RegisterType((*ProverConfig)(nil), "relayer.chains.fabric.config.ProverConfig")
}

func init() {
	proto.RegisterFile("relayer/chains/fabric/config.proto", fileDescriptor_bc92cd1bee875088)
}

var fileDescriptor_bc92cd1bee875088 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0xe3, 0xf6, 0xff, 0x01, 0x6e, 0x85, 0x50, 0x28, 0x22, 0x20, 0x14, 0xb5, 0x9d, 0xba,
	0x34, 0x11, 0x42, 0xf0, 0x00, 0x74, 0xaa, 0xc4, 0x10, 0xc1, 0xc6, 0x12, 0x39, 0xce, 0x6d, 0x62,
	0xc9, 0xb1, 0x23, 0x3b, 0x05, 0xf5, 0x25, 0x10, 0x8f, 0xc4, 0xd8, 0xb1, 0x23, 0x23, 0xb4, 0x2f,
	0x82, 0x6c, 0x57, 0x2d, 0x03, 0x5b, 0xee, 0x39, 0xbf, 0xa3, 0x93, 0xeb, 0x8b, 0x87, 0x0a, 0x38,
	0x59, 0x80, 0x8a, 0x69, 0x49, 0x98, 0xd0, 0xf1, 0x8c, 0x64, 0x8a, 0xd1, 0x98, 0x4a, 0x31, 0x63,
	0x45, 0x54, 0x2b, 0xd9, 0x48, 0xff, 0x72, 0xcb, 0x44, 0x8e, 0x89, 0x1c, 0x13, 0x39, 0xe6, 0xa2,
	0x57, 0xc8, 0x42, 0x5a, 0x30, 0x36, 0x5f, 0x2e, 0x33, 0x7c, 0x47, 0xb8, 0x33, 0x31, 0xf8, 0xc4,
	0x52, 0xfe, 0x39, 0x3e, 0xb0, 0xe9, 0x94, 0xe5, 0x01, 0xea, 0xa3, 0xd1, 0xe1, 0xc3, 0x7f, 0x3b,
	0x4f, 0x73, 0x7f, 0x80, 0xbb, 0x2f, 0x84, 0x73, 0x68, 0x52, 0x4e, 0x32, 0xe0, 0x41, 0xcb, 0xda,
	0x1d, 0xa7, 0xdd, 0x1b, 0xc9, 0xbf, 0xc5, 0x67, 0x54, 0x0a, 0x01, 0xb4, 0x61, 0x52, 0xa4, 0xb5,
	0x92, 0x33, 0xc6, 0x21, 0xad, 0x49, 0x53, 0x06, 0x6d, 0x4b, 0x9f, 0xee, 0xed, 0xc4, 0xb9, 0x09,
	0x69, 0x4a, 0x3f, 0xc0, 0xa6, 0x45, 0x08, 0xe0, 0xc1, 0x9f, 0x5d, 0xa9, 0x19, 0x4d, 0xa9, 0xed,
	0xa7, 0x32, 0x07, 0xf3, 0x4f, 0x7f, 0x5d, 0xe9, 0x4e, 0x9b, 0xe6, 0xc3, 0x57, 0x84, 0xbb, 0x89,
	0x92, 0xcf, 0xa0, 0xb6, 0x3b, 0x0c, 0x70, 0x97, 0x65, 0x34, 0xad, 0x25, 0x67, 0x94, 0x81, 0x0e,
	0x50, 0xbf, 0x6d, 0x32, 0x2c, 0xa3, 0xc9, 0x56, 0xf2, 0xaf, 0x70, 0x0f, 0x44, 0x2e, 0x95, 0x86,
	0x0a, 0x44, 0xb3, 0x47, 0x5b, 0x16, 0x3d, 0xf9, 0xe1, 0xed, 0x22, 0x23, 0x7c, 0x5c, 0xe9, 0x3a,
	0x75, 0xaf, 0x69, 0x77, 0xd2, 0x41, 0xdb, 0xe2, 0x47, 0x95, 0xae, 0x5d, 0xb5, 0x59, 0x46, 0xdf,
	0x3d, 0x2e, 0xbf, 0x42, 0x6f, 0xb9, 0x0e, 0xd1, 0x6a, 0x1d, 0xa2, 0xcf, 0x75, 0x88, 0xde, 0x36,
	0xa1, 0xb7, 0xda, 0x84, 0xde, 0xc7, 0x26, 0xf4, 0x9e, 0x6e, 0x0a, 0xd6, 0x94, 0xf3, 0x2c, 0xa2,
	0xb2, 0x8a, 0xcb, 0x45, 0x0d, 0x8a, 0x43, 0x5e, 0x80, 0x1a, 0x73, 0x92, 0xe9, 0x78, 0x31, 0x67,
	0xe3, 0x5f, 0x2f, 0x9d, 0xfd, 0xb3, 0xf7, 0xba, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xf5,
	0xfa, 0xcd, 0x09, 0x02, 0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChaincodeId) > 0 {
		i -= len(m.ChaincodeId)
		copy(dAtA[i:], m.ChaincodeId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ChaincodeId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ConnectionProfilePath) > 0 {
		i -= len(m.ConnectionProfilePath)
		copy(dAtA[i:], m.ConnectionProfilePath)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ConnectionProfilePath)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.WalletLabel) > 0 {
		i -= len(m.WalletLabel)
		copy(dAtA[i:], m.WalletLabel)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.WalletLabel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProverConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProverConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProverConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MspConfigPaths) > 0 {
		for iNdEx := len(m.MspConfigPaths) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.MspConfigPaths[iNdEx])
			copy(dAtA[i:], m.MspConfigPaths[iNdEx])
			i = encodeVarintConfig(dAtA, i, uint64(len(m.MspConfigPaths[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.EndorsementPolicies) > 0 {
		for iNdEx := len(m.EndorsementPolicies) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EndorsementPolicies[iNdEx])
			copy(dAtA[i:], m.EndorsementPolicies[iNdEx])
			i = encodeVarintConfig(dAtA, i, uint64(len(m.EndorsementPolicies[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.IbcPolicies) > 0 {
		for iNdEx := len(m.IbcPolicies) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.IbcPolicies[iNdEx])
			copy(dAtA[i:], m.IbcPolicies[iNdEx])
			i = encodeVarintConfig(dAtA, i, uint64(len(m.IbcPolicies[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.WalletLabel)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ConnectionProfilePath)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ChaincodeId)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *ProverConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IbcPolicies) > 0 {
		for _, s := range m.IbcPolicies {
			l = len(s)
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	if len(m.EndorsementPolicies) > 0 {
		for _, s := range m.EndorsementPolicies {
			l = len(s)
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	if len(m.MspConfigPaths) > 0 {
		for _, s := range m.MspConfigPaths {
			l = len(s)
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletLabel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WalletLabel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionProfilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionProfilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChaincodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChaincodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func (m *ProverConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ProverConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProverConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcPolicies", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcPolicies = append(m.IbcPolicies, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndorsementPolicies", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndorsementPolicies = append(m.EndorsementPolicies, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MspConfigPaths", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MspConfigPaths = append(m.MspConfigPaths, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)