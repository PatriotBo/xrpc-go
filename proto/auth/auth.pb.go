<<<<<<< HEAD
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth/auth.proto
=======
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906

package auth

import (
<<<<<<< HEAD
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
=======
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
<<<<<<< HEAD
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AuthReq struct {
	Uid                  *int32   `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Name                 *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
=======
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AuthReq struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

<<<<<<< HEAD
func (m *AuthReq) Reset()      { *m = AuthReq{} }
func (*AuthReq) ProtoMessage() {}
func (*AuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_712ec48c1eaf43a2, []int{0}
}
func (m *AuthReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
=======
func (m *AuthReq) Reset()         { *m = AuthReq{} }
func (m *AuthReq) String() string { return proto.CompactTextString(m) }
func (*AuthReq) ProtoMessage()    {}
func (*AuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *AuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReq.Unmarshal(m, b)
}
func (m *AuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReq.Marshal(b, m, deterministic)
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
}
func (m *AuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReq.Merge(m, src)
}
func (m *AuthReq) XXX_Size() int {
<<<<<<< HEAD
	return m.Size()
=======
	return xxx_messageInfo_AuthReq.Size(m)
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
}
func (m *AuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReq proto.InternalMessageInfo

func (m *AuthReq) GetUid() int32 {
<<<<<<< HEAD
	if m != nil && m.Uid != nil {
		return *m.Uid
=======
	if m != nil {
		return m.Uid
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
	}
	return 0
}

func (m *AuthReq) GetName() string {
<<<<<<< HEAD
	if m != nil && m.Name != nil {
		return *m.Name
=======
	if m != nil {
		return m.Name
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
	}
	return ""
}

type AuthResp struct {
<<<<<<< HEAD
	Code                 *int32   `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg                  *string  `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
=======
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

<<<<<<< HEAD
func (m *AuthResp) Reset()      { *m = AuthResp{} }
func (*AuthResp) ProtoMessage() {}
func (*AuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_712ec48c1eaf43a2, []int{1}
}
func (m *AuthResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
=======
func (m *AuthResp) Reset()         { *m = AuthResp{} }
func (m *AuthResp) String() string { return proto.CompactTextString(m) }
func (*AuthResp) ProtoMessage()    {}
func (*AuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResp.Unmarshal(m, b)
}
func (m *AuthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResp.Marshal(b, m, deterministic)
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
}
func (m *AuthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResp.Merge(m, src)
}
func (m *AuthResp) XXX_Size() int {
<<<<<<< HEAD
	return m.Size()
=======
	return xxx_messageInfo_AuthResp.Size(m)
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
}
func (m *AuthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResp proto.InternalMessageInfo

func (m *AuthResp) GetCode() int32 {
<<<<<<< HEAD
	if m != nil && m.Code != nil {
		return *m.Code
=======
	if m != nil {
		return m.Code
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
	}
	return 0
}

func (m *AuthResp) GetMsg() string {
<<<<<<< HEAD
	if m != nil && m.Msg != nil {
		return *m.Msg
=======
	if m != nil {
		return m.Msg
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthReq)(nil), "auth.AuthReq")
	proto.RegisterType((*AuthResp)(nil), "auth.AuthResp")
}

<<<<<<< HEAD
func init() { proto.RegisterFile("auth/auth.proto", fileDescriptor_712ec48c1eaf43a2) }

var fileDescriptor_712ec48c1eaf43a2 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0x2d, 0xc9,
	0xd0, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x94, 0x6e, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x7a, 0xbe, 0x3e, 0x58,
	0x32, 0xa9, 0x34, 0x0d, 0xcc, 0x03, 0x73, 0xc0, 0x2c, 0x88, 0x26, 0x25, 0x7d, 0x2e, 0x76, 0xc7,
	0xd2, 0x92, 0x8c, 0xa0, 0xd4, 0x42, 0x21, 0x01, 0x2e, 0xe6, 0xd2, 0xcc, 0x14, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xd6, 0x20, 0x10, 0x53, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x49,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x32, 0xe0, 0xe2, 0x80, 0x68, 0x28, 0x2e, 0x00, 0xc9,
	0x27, 0xe7, 0xa7, 0xa4, 0x42, 0xb5, 0x80, 0xd9, 0x20, 0x53, 0x72, 0x8b, 0xd3, 0xa1, 0x5a, 0x40,
	0x4c, 0x27, 0x9d, 0x1b, 0x0f, 0xe5, 0x18, 0x1e, 0x3c, 0x94, 0x63, 0xfc, 0xf0, 0x50, 0x8e, 0xf1,
	0xc7, 0x43, 0x39, 0xc6, 0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0xee, 0x78, 0x24, 0xc7,
	0x78, 0xe0, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x08, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xa0, 0x3c, 0xde, 0xd7, 0x00, 0x00, 0x00,
}

func (this *AuthReq) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AuthReq)
	if !ok {
		that2, ok := that.(AuthReq)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *AuthReq")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *AuthReq but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *AuthReq but is not nil && this == nil")
	}
	if this.Uid != nil && that1.Uid != nil {
		if *this.Uid != *that1.Uid {
			return fmt.Errorf("Uid this(%v) Not Equal that(%v)", *this.Uid, *that1.Uid)
		}
	} else if this.Uid != nil {
		return fmt.Errorf("this.Uid == nil && that.Uid != nil")
	} else if that1.Uid != nil {
		return fmt.Errorf("Uid this(%v) Not Equal that(%v)", this.Uid, that1.Uid)
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return fmt.Errorf("Name this(%v) Not Equal that(%v)", *this.Name, *that1.Name)
		}
	} else if this.Name != nil {
		return fmt.Errorf("this.Name == nil && that.Name != nil")
	} else if that1.Name != nil {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *AuthReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthReq)
	if !ok {
		that2, ok := that.(AuthReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Uid != nil && that1.Uid != nil {
		if *this.Uid != *that1.Uid {
			return false
		}
	} else if this.Uid != nil {
		return false
	} else if that1.Uid != nil {
		return false
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return false
		}
	} else if this.Name != nil {
		return false
	} else if that1.Name != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AuthResp) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AuthResp)
	if !ok {
		that2, ok := that.(AuthResp)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *AuthResp")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *AuthResp but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *AuthResp but is not nil && this == nil")
	}
	if this.Code != nil && that1.Code != nil {
		if *this.Code != *that1.Code {
			return fmt.Errorf("Code this(%v) Not Equal that(%v)", *this.Code, *that1.Code)
		}
	} else if this.Code != nil {
		return fmt.Errorf("this.Code == nil && that.Code != nil")
	} else if that1.Code != nil {
		return fmt.Errorf("Code this(%v) Not Equal that(%v)", this.Code, that1.Code)
	}
	if this.Msg != nil && that1.Msg != nil {
		if *this.Msg != *that1.Msg {
			return fmt.Errorf("Msg this(%v) Not Equal that(%v)", *this.Msg, *that1.Msg)
		}
	} else if this.Msg != nil {
		return fmt.Errorf("this.Msg == nil && that.Msg != nil")
	} else if that1.Msg != nil {
		return fmt.Errorf("Msg this(%v) Not Equal that(%v)", this.Msg, that1.Msg)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *AuthResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthResp)
	if !ok {
		that2, ok := that.(AuthResp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != nil && that1.Code != nil {
		if *this.Code != *that1.Code {
			return false
		}
	} else if this.Code != nil {
		return false
	} else if that1.Code != nil {
		return false
	}
	if this.Msg != nil && that1.Msg != nil {
		if *this.Msg != *that1.Msg {
			return false
		}
	} else if this.Msg != nil {
		return false
	} else if that1.Msg != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AuthReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&auth.AuthReq{")
	if this.Uid != nil {
		s = append(s, "Uid: "+valueToGoStringAuth(this.Uid, "int32")+",\n")
	}
	if this.Name != nil {
		s = append(s, "Name: "+valueToGoStringAuth(this.Name, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AuthResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&auth.AuthResp{")
	if this.Code != nil {
		s = append(s, "Code: "+valueToGoStringAuth(this.Code, "int32")+",\n")
	}
	if this.Msg != nil {
		s = append(s, "Msg: "+valueToGoStringAuth(this.Msg, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAuth(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AuthReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Name != nil {
		i -= len(*m.Name)
		copy(dAtA[i:], *m.Name)
		i = encodeVarintAuth(dAtA, i, uint64(len(*m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Uid != nil {
		i = encodeVarintAuth(dAtA, i, uint64(*m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AuthResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Msg != nil {
		i -= len(*m.Msg)
		copy(dAtA[i:], *m.Msg)
		i = encodeVarintAuth(dAtA, i, uint64(len(*m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != nil {
		i = encodeVarintAuth(dAtA, i, uint64(*m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedAuthReq(r randyAuth, easy bool) *AuthReq {
	this := &AuthReq{}
	if r.Intn(5) != 0 {
		v1 := int32(r.Int31())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Uid = &v1
	}
	if r.Intn(5) != 0 {
		v2 := string(randStringAuth(r))
		this.Name = &v2
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedAuth(r, 3)
	}
	return this
}

func NewPopulatedAuthResp(r randyAuth, easy bool) *AuthResp {
	this := &AuthResp{}
	if r.Intn(5) != 0 {
		v3 := int32(r.Int31())
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		this.Code = &v3
	}
	if r.Intn(5) != 0 {
		v4 := string(randStringAuth(r))
		this.Msg = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedAuth(r, 3)
	}
	return this
}

type randyAuth interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAuth(r randyAuth) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAuth(r randyAuth) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneAuth(r)
	}
	return string(tmps)
}
func randUnrecognizedAuth(r randyAuth, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAuth(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAuth(dAtA []byte, r randyAuth, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAuth(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateAuth(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateAuth(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAuth(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAuth(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAuth(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAuth(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *AuthReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != nil {
		n += 1 + sovAuth(uint64(*m.Uid))
	}
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AuthResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != nil {
		n += 1 + sovAuth(uint64(*m.Code))
	}
	if m.Msg != nil {
		l = len(*m.Msg)
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAuth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AuthReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AuthReq{`,
		`Uid:` + valueToStringAuth(this.Uid) + `,`,
		`Name:` + valueToStringAuth(this.Name) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AuthResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AuthResp{`,
		`Code:` + valueToStringAuth(this.Code) + `,`,
		`Msg:` + valueToStringAuth(this.Msg) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAuth(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AuthReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: AuthReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Uid = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: AuthResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Code = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Msg = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
				return 0, ErrInvalidLengthAuth
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAuth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuth
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
				next, err := skipAuth(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAuth
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
	ErrInvalidLengthAuth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth   = fmt.Errorf("proto: integer overflow")
)
=======
func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xf4, 0xb9, 0xd8, 0x1d, 0x4b,
	0x4b, 0x32, 0x82, 0x52, 0x0b, 0x85, 0x04, 0xb8, 0x98, 0x4b, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x58, 0x83, 0x40, 0x4c, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x80, 0x8b, 0x03, 0xa2, 0xa1, 0xb8, 0x00, 0x24, 0x9f,
	0x9c, 0x9f, 0x92, 0x0a, 0xd5, 0x02, 0x66, 0x83, 0x4c, 0xc9, 0x2d, 0x4e, 0x87, 0x6a, 0x01, 0x31,
	0x93, 0xd8, 0xc0, 0xf6, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xb9, 0x87, 0xeb, 0x7d,
	0x00, 0x00, 0x00,
}
>>>>>>> 2d50a11b6136d4b7181c584f5d92b9d6546e4906
