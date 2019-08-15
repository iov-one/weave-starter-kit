// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/codec.proto

package app

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	custom "github.com/iov-one/weave-starter-kit/x/custom"
	migration "github.com/iov-one/weave/migration"
	cash "github.com/iov-one/weave/x/cash"
	currency "github.com/iov-one/weave/x/currency"
	multisig "github.com/iov-one/weave/x/multisig"
	sigs "github.com/iov-one/weave/x/sigs"
	validators "github.com/iov-one/weave/x/validators"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Tx contains the message
// When extending Tx, follow the rules:
// - Range 1-50 is reserved for middlewares,
// - Range 51-inf is reserved for different message types,
// - Keep the same numbers for the same message types in weave based applications to
//   sustain compatibility between blockchains. For example, FeeInfo field is used by
//   both and indexed at first position. Skip unused fields (leave index unused or
//   comment out for clarity).
// When there is a gap in message sequence numbers - that most likely means some
// old fields got deprecated. This is done to maintain binary compatibility.
type Tx struct {
	// Enables coin.GetFees()
	Fees *cash.FeeInfo `protobuf:"bytes,1,opt,name=fees,proto3" json:"fees,omitempty"`
	//StdSignature represents the signature, the identity of the signer
	// (the Pubkey), and a sequence number to prevent replay attacks.
	Signatures []*sigs.StdSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	// ID of a multisig contract.
	Multisig [][]byte `protobuf:"bytes,4,rep,name=multisig,proto3" json:"multisig,omitempty"`
	// sum defines over all allowed messages on this chain.
	//
	// Types that are valid to be assigned to Sum:
	//	*Tx_CashSendMsg
	//	*Tx_CreateContractMsg
	//	*Tx_UpdateContractMsg
	//	*Tx_ValidatorsApplyDiffMsg
	//	*Tx_CurrencyCreateMsg
	//	*Tx_MigrationUpgradeSchemaMsg
	//	*Tx_CustomCreateTimedStateMsg
	//	*Tx_CreateStateMsg
	Sum isTx_Sum `protobuf_oneof:"sum"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_e43b82f4f03f64b8, []int{0}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(m, src)
}
func (m *Tx) XXX_Size() int {
	return m.Size()
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

type isTx_Sum interface {
	isTx_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Tx_CashSendMsg struct {
	CashSendMsg *cash.SendMsg `protobuf:"bytes,51,opt,name=cash_send_msg,json=cashSendMsg,proto3,oneof"`
}
type Tx_CreateContractMsg struct {
	CreateContractMsg *multisig.CreateMsg `protobuf:"bytes,56,opt,name=create_contract_msg,json=createContractMsg,proto3,oneof"`
}
type Tx_UpdateContractMsg struct {
	UpdateContractMsg *multisig.UpdateMsg `protobuf:"bytes,57,opt,name=update_contract_msg,json=updateContractMsg,proto3,oneof"`
}
type Tx_ValidatorsApplyDiffMsg struct {
	ValidatorsApplyDiffMsg *validators.ApplyDiffMsg `protobuf:"bytes,58,opt,name=validators_apply_diff_msg,json=validatorsApplyDiffMsg,proto3,oneof"`
}
type Tx_CurrencyCreateMsg struct {
	CurrencyCreateMsg *currency.CreateMsg `protobuf:"bytes,59,opt,name=currency_create_msg,json=currencyCreateMsg,proto3,oneof"`
}
type Tx_MigrationUpgradeSchemaMsg struct {
	MigrationUpgradeSchemaMsg *migration.UpgradeSchemaMsg `protobuf:"bytes,69,opt,name=migration_upgrade_schema_msg,json=migrationUpgradeSchemaMsg,proto3,oneof"`
}
type Tx_CustomCreateTimedStateMsg struct {
	CustomCreateTimedStateMsg *custom.CreateTimedStateMsg `protobuf:"bytes,100,opt,name=custom_create_timed_state_msg,json=customCreateTimedStateMsg,proto3,oneof"`
}
type Tx_CreateStateMsg struct {
	CreateStateMsg *custom.CreateStateMsg `protobuf:"bytes,101,opt,name=create_state_msg,json=createStateMsg,proto3,oneof"`
}

func (*Tx_CashSendMsg) isTx_Sum()               {}
func (*Tx_CreateContractMsg) isTx_Sum()         {}
func (*Tx_UpdateContractMsg) isTx_Sum()         {}
func (*Tx_ValidatorsApplyDiffMsg) isTx_Sum()    {}
func (*Tx_CurrencyCreateMsg) isTx_Sum()         {}
func (*Tx_MigrationUpgradeSchemaMsg) isTx_Sum() {}
func (*Tx_CustomCreateTimedStateMsg) isTx_Sum() {}
func (*Tx_CreateStateMsg) isTx_Sum()            {}

func (m *Tx) GetSum() isTx_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Tx) GetFees() *cash.FeeInfo {
	if m != nil {
		return m.Fees
	}
	return nil
}

func (m *Tx) GetSignatures() []*sigs.StdSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Tx) GetMultisig() [][]byte {
	if m != nil {
		return m.Multisig
	}
	return nil
}

func (m *Tx) GetCashSendMsg() *cash.SendMsg {
	if x, ok := m.GetSum().(*Tx_CashSendMsg); ok {
		return x.CashSendMsg
	}
	return nil
}

func (m *Tx) GetCreateContractMsg() *multisig.CreateMsg {
	if x, ok := m.GetSum().(*Tx_CreateContractMsg); ok {
		return x.CreateContractMsg
	}
	return nil
}

func (m *Tx) GetUpdateContractMsg() *multisig.UpdateMsg {
	if x, ok := m.GetSum().(*Tx_UpdateContractMsg); ok {
		return x.UpdateContractMsg
	}
	return nil
}

func (m *Tx) GetValidatorsApplyDiffMsg() *validators.ApplyDiffMsg {
	if x, ok := m.GetSum().(*Tx_ValidatorsApplyDiffMsg); ok {
		return x.ValidatorsApplyDiffMsg
	}
	return nil
}

func (m *Tx) GetCurrencyCreateMsg() *currency.CreateMsg {
	if x, ok := m.GetSum().(*Tx_CurrencyCreateMsg); ok {
		return x.CurrencyCreateMsg
	}
	return nil
}

func (m *Tx) GetMigrationUpgradeSchemaMsg() *migration.UpgradeSchemaMsg {
	if x, ok := m.GetSum().(*Tx_MigrationUpgradeSchemaMsg); ok {
		return x.MigrationUpgradeSchemaMsg
	}
	return nil
}

func (m *Tx) GetCustomCreateTimedStateMsg() *custom.CreateTimedStateMsg {
	if x, ok := m.GetSum().(*Tx_CustomCreateTimedStateMsg); ok {
		return x.CustomCreateTimedStateMsg
	}
	return nil
}

func (m *Tx) GetCreateStateMsg() *custom.CreateStateMsg {
	if x, ok := m.GetSum().(*Tx_CreateStateMsg); ok {
		return x.CreateStateMsg
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Tx) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Tx_OneofMarshaler, _Tx_OneofUnmarshaler, _Tx_OneofSizer, []interface{}{
		(*Tx_CashSendMsg)(nil),
		(*Tx_CreateContractMsg)(nil),
		(*Tx_UpdateContractMsg)(nil),
		(*Tx_ValidatorsApplyDiffMsg)(nil),
		(*Tx_CurrencyCreateMsg)(nil),
		(*Tx_MigrationUpgradeSchemaMsg)(nil),
		(*Tx_CustomCreateTimedStateMsg)(nil),
		(*Tx_CreateStateMsg)(nil),
	}
}

func _Tx_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Tx)
	// sum
	switch x := m.Sum.(type) {
	case *Tx_CashSendMsg:
		_ = b.EncodeVarint(51<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CashSendMsg); err != nil {
			return err
		}
	case *Tx_CreateContractMsg:
		_ = b.EncodeVarint(56<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateContractMsg); err != nil {
			return err
		}
	case *Tx_UpdateContractMsg:
		_ = b.EncodeVarint(57<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateContractMsg); err != nil {
			return err
		}
	case *Tx_ValidatorsApplyDiffMsg:
		_ = b.EncodeVarint(58<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ValidatorsApplyDiffMsg); err != nil {
			return err
		}
	case *Tx_CurrencyCreateMsg:
		_ = b.EncodeVarint(59<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CurrencyCreateMsg); err != nil {
			return err
		}
	case *Tx_MigrationUpgradeSchemaMsg:
		_ = b.EncodeVarint(69<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MigrationUpgradeSchemaMsg); err != nil {
			return err
		}
	case *Tx_CustomCreateTimedStateMsg:
		_ = b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CustomCreateTimedStateMsg); err != nil {
			return err
		}
	case *Tx_CreateStateMsg:
		_ = b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateStateMsg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Tx.Sum has unexpected type %T", x)
	}
	return nil
}

func _Tx_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Tx)
	switch tag {
	case 51: // sum.cash_send_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cash.SendMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_CashSendMsg{msg}
		return true, err
	case 56: // sum.create_contract_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(multisig.CreateMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_CreateContractMsg{msg}
		return true, err
	case 57: // sum.update_contract_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(multisig.UpdateMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_UpdateContractMsg{msg}
		return true, err
	case 58: // sum.validators_apply_diff_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(validators.ApplyDiffMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_ValidatorsApplyDiffMsg{msg}
		return true, err
	case 59: // sum.currency_create_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(currency.CreateMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_CurrencyCreateMsg{msg}
		return true, err
	case 69: // sum.migration_upgrade_schema_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(migration.UpgradeSchemaMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_MigrationUpgradeSchemaMsg{msg}
		return true, err
	case 100: // sum.custom_create_timed_state_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(custom.CreateTimedStateMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_CustomCreateTimedStateMsg{msg}
		return true, err
	case 101: // sum.create_state_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(custom.CreateStateMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_CreateStateMsg{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Tx_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Tx)
	// sum
	switch x := m.Sum.(type) {
	case *Tx_CashSendMsg:
		s := proto.Size(x.CashSendMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_CreateContractMsg:
		s := proto.Size(x.CreateContractMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_UpdateContractMsg:
		s := proto.Size(x.UpdateContractMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_ValidatorsApplyDiffMsg:
		s := proto.Size(x.ValidatorsApplyDiffMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_CurrencyCreateMsg:
		s := proto.Size(x.CurrencyCreateMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_MigrationUpgradeSchemaMsg:
		s := proto.Size(x.MigrationUpgradeSchemaMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_CustomCreateTimedStateMsg:
		s := proto.Size(x.CustomCreateTimedStateMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_CreateStateMsg:
		s := proto.Size(x.CreateStateMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Tx)(nil), "app.Tx")
}

func init() { proto.RegisterFile("app/codec.proto", fileDescriptor_e43b82f4f03f64b8) }

var fileDescriptor_e43b82f4f03f64b8 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0xdb, 0x75, 0x43, 0xc8, 0x65, 0xbc, 0x78, 0x68, 0x4a, 0x0b, 0x44, 0x85, 0x53, 0x25,
	0x84, 0x23, 0xda, 0x0b, 0x2f, 0x27, 0x3a, 0x86, 0xe0, 0xc0, 0xa5, 0x5d, 0xaf, 0x58, 0x9e, 0xed,
	0xb8, 0x96, 0x9a, 0x38, 0x8a, 0x9d, 0xb2, 0x7d, 0x0b, 0x3e, 0x16, 0xc7, 0x1d, 0x39, 0xa2, 0x96,
	0x0f, 0x82, 0xe2, 0x38, 0x69, 0x52, 0xd8, 0xb4, 0x5b, 0xfc, 0xfc, 0x7f, 0xcf, 0xcf, 0x8f, 0x93,
	0x18, 0x3c, 0x20, 0x49, 0x12, 0x50, 0xc5, 0x38, 0x45, 0x49, 0xaa, 0x8c, 0x82, 0x1d, 0x92, 0x24,
	0x7d, 0x24, 0xa4, 0x59, 0x64, 0xe7, 0x88, 0xaa, 0x28, 0x90, 0x6a, 0xf5, 0x4a, 0xc5, 0x3c, 0xf8,
	0xce, 0xc9, 0x8a, 0x07, 0x91, 0x14, 0x29, 0x31, 0x52, 0xc5, 0xf5, 0xa6, 0xfe, 0xcb, 0x6b, 0xf9,
	0x8b, 0x80, 0x12, 0xbd, 0x68, 0xc0, 0xc1, 0x4d, 0x70, 0x96, 0xa6, 0x3c, 0xa6, 0x97, 0xb7, 0x6e,
	0x88, 0xb2, 0xa5, 0x91, 0x5a, 0x8a, 0x5b, 0x8f, 0xa3, 0xa5, 0xd0, 0x0d, 0xf8, 0xf5, 0x0d, 0xf0,
	0x8a, 0x2c, 0x25, 0x23, 0x46, 0xa5, 0xcd, 0x96, 0xc7, 0x42, 0x09, 0x65, 0x1f, 0x83, 0xfc, 0xa9,
	0xac, 0xe6, 0xe3, 0x6b, 0xa3, 0xa2, 0x3a, 0xfb, 0xe2, 0xcf, 0x01, 0xd8, 0x3b, 0xbb, 0x80, 0xcf,
	0xc1, 0x7e, 0xc8, 0xb9, 0xf6, 0xda, 0x83, 0xf6, 0xb0, 0x3b, 0x3a, 0x44, 0xf9, 0x5b, 0x41, 0x9f,
	0x38, 0xff, 0x12, 0x87, 0x6a, 0x6a, 0x23, 0x38, 0x02, 0x40, 0x4b, 0x11, 0x13, 0x93, 0xa5, 0x5c,
	0x7b, 0x7b, 0x83, 0xce, 0xb0, 0x3b, 0x82, 0x28, 0x9f, 0x17, 0xcd, 0x0c, 0x9b, 0x95, 0xd1, 0xb4,
	0x46, 0xc1, 0x3e, 0xb8, 0x5b, 0xbe, 0x01, 0x6f, 0x7f, 0xd0, 0x19, 0xde, 0x9b, 0x56, 0x6b, 0x38,
	0x06, 0x87, 0xf9, 0x2e, 0x58, 0xf3, 0x98, 0xe1, 0x48, 0x0b, 0x6f, 0x5c, 0xdf, 0x7b, 0xc6, 0x63,
	0xf6, 0x55, 0x8b, 0xcf, 0xad, 0x69, 0x37, 0x5f, 0xbb, 0x25, 0x3c, 0x05, 0x47, 0x34, 0xe5, 0xc4,
	0x70, 0x4c, 0x55, 0x6c, 0x52, 0x42, 0x8d, 0x6d, 0x7d, 0x63, 0x5b, 0x8f, 0x50, 0x29, 0x47, 0x27,
	0x16, 0x2a, 0x04, 0x8f, 0x8a, 0x8e, 0x13, 0xd7, 0xe0, 0x34, 0x59, 0xc2, 0xfe, 0xd1, 0xbc, 0xdd,
	0xd5, 0xcc, 0x2d, 0xe4, 0x34, 0x45, 0x47, 0x5d, 0x33, 0x07, 0xbd, 0xed, 0x27, 0xc0, 0x24, 0x49,
	0x96, 0x97, 0x98, 0xc9, 0x30, 0xb4, 0xb2, 0x77, 0x56, 0xe6, 0xa1, 0x2d, 0x81, 0x3e, 0xe4, 0xc4,
	0x47, 0x19, 0x86, 0x85, 0xf1, 0x78, 0x1b, 0xd5, 0x13, 0x7b, 0x48, 0xf7, 0xa3, 0x61, 0x77, 0xda,
	0x5c, 0xf8, 0xde, 0x4d, 0x57, 0x66, 0x3b, 0x87, 0x74, 0xd5, 0xaa, 0x08, 0xbf, 0x81, 0xa7, 0xd5,
	0x75, 0xc0, 0x59, 0x22, 0x52, 0xc2, 0x38, 0xd6, 0x74, 0xc1, 0x23, 0x62, 0x7d, 0xa7, 0xd6, 0xf7,
	0x04, 0x55, 0x10, 0x9a, 0x17, 0xd0, 0xcc, 0x32, 0x85, 0xb7, 0x57, 0xa5, 0xbb, 0x21, 0xc4, 0xe0,
	0x59, 0xf1, 0x43, 0x95, 0x43, 0x1a, 0x19, 0x71, 0x86, 0xb5, 0x29, 0x07, 0x66, 0x6e, 0x83, 0x82,
	0x72, 0xe3, 0x9e, 0xe5, 0xd0, 0xcc, 0x54, 0x83, 0xf7, 0x8a, 0xf4, 0x3f, 0x21, 0x9c, 0x80, 0x87,
	0xce, 0xbc, 0x75, 0x72, 0xeb, 0x3c, 0x6e, 0x3a, 0x6b, 0xba, 0xfb, 0xb4, 0x51, 0x99, 0x1c, 0x80,
	0x8e, 0xce, 0xa2, 0x89, 0xf7, 0x73, 0xed, 0xb7, 0xaf, 0xd6, 0x7e, 0xfb, 0xf7, 0xda, 0x6f, 0xff,
	0xd8, 0xf8, 0xad, 0xab, 0x8d, 0xdf, 0xfa, 0xb5, 0xf1, 0x5b, 0xe7, 0x77, 0xec, 0x3d, 0x18, 0xff,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xb4, 0x98, 0x5c, 0x6a, 0x04, 0x00, 0x00,
}

func (m *Tx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tx) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Fees != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Fees.Size()))
		n1, err := m.Fees.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Signatures) > 0 {
		for _, msg := range m.Signatures {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Multisig) > 0 {
		for _, b := range m.Multisig {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCodec(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if m.Sum != nil {
		nn2, err := m.Sum.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	return i, nil
}

func (m *Tx_CashSendMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CashSendMsg != nil {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.CashSendMsg.Size()))
		n3, err := m.CashSendMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *Tx_CreateContractMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CreateContractMsg != nil {
		dAtA[i] = 0xc2
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.CreateContractMsg.Size()))
		n4, err := m.CreateContractMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *Tx_UpdateContractMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.UpdateContractMsg != nil {
		dAtA[i] = 0xca
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.UpdateContractMsg.Size()))
		n5, err := m.UpdateContractMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *Tx_ValidatorsApplyDiffMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ValidatorsApplyDiffMsg != nil {
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.ValidatorsApplyDiffMsg.Size()))
		n6, err := m.ValidatorsApplyDiffMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *Tx_CurrencyCreateMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CurrencyCreateMsg != nil {
		dAtA[i] = 0xda
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.CurrencyCreateMsg.Size()))
		n7, err := m.CurrencyCreateMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func (m *Tx_MigrationUpgradeSchemaMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.MigrationUpgradeSchemaMsg != nil {
		dAtA[i] = 0xaa
		i++
		dAtA[i] = 0x4
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.MigrationUpgradeSchemaMsg.Size()))
		n8, err := m.MigrationUpgradeSchemaMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}
func (m *Tx_CustomCreateTimedStateMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CustomCreateTimedStateMsg != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.CustomCreateTimedStateMsg.Size()))
		n9, err := m.CustomCreateTimedStateMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	return i, nil
}
func (m *Tx_CreateStateMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CreateStateMsg != nil {
		dAtA[i] = 0xaa
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.CreateStateMsg.Size()))
		n10, err := m.CreateStateMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	return i, nil
}
func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Tx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Fees != nil {
		l = m.Fees.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, e := range m.Signatures {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if len(m.Multisig) > 0 {
		for _, b := range m.Multisig {
			l = len(b)
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Tx_CashSendMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CashSendMsg != nil {
		l = m.CashSendMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_CreateContractMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreateContractMsg != nil {
		l = m.CreateContractMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_UpdateContractMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpdateContractMsg != nil {
		l = m.UpdateContractMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_ValidatorsApplyDiffMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ValidatorsApplyDiffMsg != nil {
		l = m.ValidatorsApplyDiffMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_CurrencyCreateMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrencyCreateMsg != nil {
		l = m.CurrencyCreateMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_MigrationUpgradeSchemaMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MigrationUpgradeSchemaMsg != nil {
		l = m.MigrationUpgradeSchemaMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_CustomCreateTimedStateMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CustomCreateTimedStateMsg != nil {
		l = m.CustomCreateTimedStateMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_CreateStateMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreateStateMsg != nil {
		l = m.CreateStateMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Tx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fees == nil {
				m.Fees = &cash.FeeInfo{}
			}
			if err := m.Fees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, &sigs.StdSignature{})
			if err := m.Signatures[len(m.Signatures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multisig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Multisig = append(m.Multisig, make([]byte, postIndex-iNdEx))
			copy(m.Multisig[len(m.Multisig)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 51:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CashSendMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &cash.SendMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_CashSendMsg{v}
			iNdEx = postIndex
		case 56:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateContractMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &multisig.CreateMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_CreateContractMsg{v}
			iNdEx = postIndex
		case 57:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateContractMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &multisig.UpdateMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_UpdateContractMsg{v}
			iNdEx = postIndex
		case 58:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorsApplyDiffMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &validators.ApplyDiffMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_ValidatorsApplyDiffMsg{v}
			iNdEx = postIndex
		case 59:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyCreateMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &currency.CreateMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_CurrencyCreateMsg{v}
			iNdEx = postIndex
		case 69:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MigrationUpgradeSchemaMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &migration.UpgradeSchemaMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_MigrationUpgradeSchemaMsg{v}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomCreateTimedStateMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &custom.CreateTimedStateMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_CustomCreateTimedStateMsg{v}
			iNdEx = postIndex
		case 101:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateStateMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &custom.CreateStateMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_CreateStateMsg{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
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
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCodec
				}
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
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)
