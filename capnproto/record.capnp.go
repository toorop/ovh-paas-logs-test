package main

// AUTO GENERATED - DO NOT EDIT

import (
	math "math"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Record struct{ capnp.Struct }

// Record_TypeID is the unique identifier for the type Record.
const Record_TypeID = 0xe1068a6aee02baba

func NewRecord(s *capnp.Segment) (Record, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 8})
	return Record{st}, err
}

func NewRootRecord(s *capnp.Segment) (Record, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 8})
	return Record{st}, err
}

func ReadRootRecord(msg *capnp.Message) (Record, error) {
	root, err := msg.RootPtr()
	return Record{root.Struct()}, err
}

func (s Record) String() string {
	str, _ := text.Marshal(0xe1068a6aee02baba, s.Struct)
	return str
}

func (s Record) Ts() float64 {
	return math.Float64frombits(s.Struct.Uint64(0))
}

func (s Record) SetTs(v float64) {
	s.Struct.SetUint64(0, math.Float64bits(v))
}

func (s Record) Hostname() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Record) HasHostname() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Record) HostnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Record) SetHostname(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Record) Facility() uint8 {
	return s.Struct.Uint8(8)
}

func (s Record) SetFacility(v uint8) {
	s.Struct.SetUint8(8, v)
}

func (s Record) Severity() uint8 {
	return s.Struct.Uint8(9)
}

func (s Record) SetSeverity(v uint8) {
	s.Struct.SetUint8(9, v)
}

func (s Record) Appname() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Record) HasAppname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Record) AppnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Record) SetAppname(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Record) Procid() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Record) HasProcid() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Record) ProcidBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Record) SetProcid(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s Record) Msgid() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s Record) HasMsgid() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Record) MsgidBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s Record) SetMsgid(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, t.List.ToPtr())
}

func (s Record) Msg() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s Record) HasMsg() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Record) MsgBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s Record) SetMsg(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, t.List.ToPtr())
}

func (s Record) FullMsg() (string, error) {
	p, err := s.Struct.Ptr(5)
	return p.Text(), err
}

func (s Record) HasFullMsg() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Record) FullMsgBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return p.TextBytes(), err
}

func (s Record) SetFullMsg(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(5, t.List.ToPtr())
}

func (s Record) SdId() (string, error) {
	p, err := s.Struct.Ptr(6)
	return p.Text(), err
}

func (s Record) HasSdId() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s Record) SdIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(6)
	return p.TextBytes(), err
}

func (s Record) SetSdId(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(6, t.List.ToPtr())
}

func (s Record) Pairs() (Pair_List, error) {
	p, err := s.Struct.Ptr(7)
	return Pair_List{List: p.List()}, err
}

func (s Record) HasPairs() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s Record) SetPairs(v Pair_List) error {
	return s.Struct.SetPtr(7, v.List.ToPtr())
}

// NewPairs sets the pairs field to a newly
// allocated Pair_List, preferring placement in s's segment.
func (s Record) NewPairs(n int32) (Pair_List, error) {
	l, err := NewPair_List(s.Struct.Segment(), n)
	if err != nil {
		return Pair_List{}, err
	}
	err = s.Struct.SetPtr(7, l.List.ToPtr())
	return l, err
}

// Record_List is a list of Record.
type Record_List struct{ capnp.List }

// NewRecord creates a new list of Record.
func NewRecord_List(s *capnp.Segment, sz int32) (Record_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 8}, sz)
	return Record_List{l}, err
}

func (s Record_List) At(i int) Record { return Record{s.List.Struct(i)} }

func (s Record_List) Set(i int, v Record) error { return s.List.SetStruct(i, v.Struct) }

// Record_Promise is a wrapper for a Record promised by a client call.
type Record_Promise struct{ *capnp.Pipeline }

func (p Record_Promise) Struct() (Record, error) {
	s, err := p.Pipeline.Struct()
	return Record{s}, err
}

type Pair struct{ capnp.Struct }
type Pair_value Pair
type Pair_value_Which uint16

const (
	Pair_value_Which_string Pair_value_Which = 0
	Pair_value_Which_bool   Pair_value_Which = 1
	Pair_value_Which_f64    Pair_value_Which = 2
	Pair_value_Which_i64    Pair_value_Which = 3
	Pair_value_Which_u64    Pair_value_Which = 4
	Pair_value_Which_null   Pair_value_Which = 5
)

func (w Pair_value_Which) String() string {
	const s = "stringboolf64i64u64null"
	switch w {
	case Pair_value_Which_string:
		return s[0:6]
	case Pair_value_Which_bool:
		return s[6:10]
	case Pair_value_Which_f64:
		return s[10:13]
	case Pair_value_Which_i64:
		return s[13:16]
	case Pair_value_Which_u64:
		return s[16:19]
	case Pair_value_Which_null:
		return s[19:23]

	}
	return "Pair_value_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Pair_TypeID is the unique identifier for the type Pair.
const Pair_TypeID = 0xb4a68d38a716233d

func NewPair(s *capnp.Segment) (Pair, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Pair{st}, err
}

func NewRootPair(s *capnp.Segment) (Pair, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Pair{st}, err
}

func ReadRootPair(msg *capnp.Message) (Pair, error) {
	root, err := msg.RootPtr()
	return Pair{root.Struct()}, err
}

func (s Pair) String() string {
	str, _ := text.Marshal(0xb4a68d38a716233d, s.Struct)
	return str
}

func (s Pair) Key() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Pair) HasKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Pair) KeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Pair) SetKey(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Pair) Value() Pair_value { return Pair_value(s) }

func (s Pair_value) Which() Pair_value_Which {
	return Pair_value_Which(s.Struct.Uint16(0))
}
func (s Pair_value) String() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Pair_value) HasString() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Pair_value) StringBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Pair_value) SetString(v string) error {
	s.Struct.SetUint16(0, 0)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Pair_value) Bool() bool {
	return s.Struct.Bit(16)
}

func (s Pair_value) SetBool(v bool) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetBit(16, v)
}

func (s Pair_value) F64() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Pair_value) SetF64(v float64) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Pair_value) I64() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Pair_value) SetI64(v int64) {
	s.Struct.SetUint16(0, 3)
	s.Struct.SetUint64(8, uint64(v))
}

func (s Pair_value) U64() uint64 {
	return s.Struct.Uint64(8)
}

func (s Pair_value) SetU64(v uint64) {
	s.Struct.SetUint16(0, 4)
	s.Struct.SetUint64(8, v)
}

func (s Pair_value) SetNull() {
	s.Struct.SetUint16(0, 5)

}

// Pair_List is a list of Pair.
type Pair_List struct{ capnp.List }

// NewPair creates a new list of Pair.
func NewPair_List(s *capnp.Segment, sz int32) (Pair_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return Pair_List{l}, err
}

func (s Pair_List) At(i int) Pair { return Pair{s.List.Struct(i)} }

func (s Pair_List) Set(i int, v Pair) error { return s.List.SetStruct(i, v.Struct) }

// Pair_Promise is a wrapper for a Pair promised by a client call.
type Pair_Promise struct{ *capnp.Pipeline }

func (p Pair_Promise) Struct() (Pair, error) {
	s, err := p.Pipeline.Struct()
	return Pair{s}, err
}

func (p Pair_Promise) Value() Pair_value_Promise { return Pair_value_Promise{p.Pipeline} }

// Pair_value_Promise is a wrapper for a Pair_value promised by a client call.
type Pair_value_Promise struct{ *capnp.Pipeline }

func (p Pair_value_Promise) Struct() (Pair_value, error) {
	s, err := p.Pipeline.Struct()
	return Pair_value{s}, err
}

const schema_e5fbb19e9e69110b = "x\xda\\\xd3_k#U\x18\x06\xf0\xe7y\xcf\xfcI" +
	"\xb6\xbbv\xc7\x09\"\xb2AY\x15\xdc\xb2\xbbtwC" +
	"Y\xcbJ\xabThK\x0b9N\x8b\x0a\xde\xc4$\xad" +
	"\xa3I&\xce$\x95\x82\xd0\xfb\x827\xc5\x0b\x0bZ\xaa" +
	"T\xad`\xc1B\xbd\xe9\xa5\xdf\xc4\x0bo\xfd\x00V\x1d" +
	"yG\xd2\xb0\x9d\xab9\xbfy\xcf\xcb9\xcf93\xfd" +
	"9\xe7\xe5\x81\x9b\x08`o\xb9^\xee\xec\xad\x9eL-" +
	"\xef\xef\xc2>K\xc9\xdfx\xf9\xb9\x1f\x1f\x7f\xf1\xc3\x19" +
	"\xd6\xc5\xa7\x07\x84U\xfe\x06>\xaa\xf2]\x82\xe3\xaf\xf6" +
	"\x06%\x9f\x08\xe2\x83\x83\xd3\x8b?\xe0\x8a\x0f\x84\xb1\xec" +
	"\x85\x9f\x16o]\xf9\x0c\xcc\xcf\xcf\xe5\xcf\x8fw\xbd\xdf" +
	"\xaf\x16\x97\xb4\xe4D\xbe\x0b\x7f\xd5\xe2G\xa7\x92\x08\xee" +
	"\xe5i\xbb\x99\xa4\xad\xfbMi\xf4{\xfd\xd9z#N" +
	"\xefo5:\xc36\xec\xf3\xc6\xb9\x95\xe7\xac\x90@\xb0" +
	"?\x0b\xd8/\x0d\xed\xa1\xb0\xca\x7f\xf3\x9b\x15\x0a\x10|" +
	"3\x05\xd8\xaf\x0c\xed\x91\xb0*\xffh\xb5\x01\x82oo" +
	"\x03\xf6kC{,\xac\x9a\xbf\x95\x1d \xf8^\xf9\xd0" +
	"\xd0\xfe,\xac:\x17\xca.\x10\xfc\xa4|dh\x7f\x11" +
	"\xdep\xff\xca+\x1a@p\xa2\xad\x8f\x0d\xed\x99p." +
	"\x1b\xa4qo\x93\xd7!\xbc\x0eN~\x98$\x1d\x12B" +
	"\x82\xfe\xc6L\x8d\x13\x10N\x80~<S\xa3\x0b\xa1\x0b" +
	"\xfa\xc3\x99\x1a\xcb\x10\x96\xc1\xc9\xde\xb0\xd3\x81w\xb9Y" +
	"\x8e6\xcb\xb4N\xda\x92q\x00G\xb7yG\x97\xf2\x8a" +
	"\xa1\x9d\x16\xea3>\xa5\xe0\xdeC\x88\xffI{{\xb4" +
	"\x88\x17\x8b\x98\xae\xb4|\xa7\xdd\xf4\x93\xb4\xa5M\xef\x8e" +
	"\x9a\x86\xaf\xf2\x05 z\x89\x86\xd1]j\xdf\"\xd2\xf0" +
	"\x0e\x97\x81\xe85\xe5\x1a\x85\x81\x94\x8aL\xc3\x07\x85O" +
	"\xab?Q7\xe5\"\xd4\xf0\xf5\xc2\x1f\xab/\xa8;\xff" +
	"\xa7\x1a\xbe\xc9\xb7\x80\xe8\x89\xfa\xa2\xba+E\xac\xe1\xdb" +
	"\x9c\x05\xa2y\xf5\x15u\xcf\x14\xc1\x86K|\x08D\x0b" +
	"\xeauu\xdf\xa9P\xaf\xc6*o\x03\xd1\xa2\xfa\x9az" +
	"\xc9\xad\xb0\x04\x84\xb6\xe8\xbf\xa2\xfe\x9ez\xd9\xab\xb0\x0c" +
	"\x84\xeb\x9c\x02\xa2\xba\xfa\x07\xea\xd7\xfc\x0a\xaf\x01\xe1\xfb" +
	"E\xff5\xf5>\x85f\x90\x8d\x8e'\xff(\xc9\x06\xbd" +
	"F\xb7\x0d`\x14c\xbe\xd1h\xc6\x9dx\xb0\xad\xe6A" +
	"\xe8\x81y\xd6\xdej\xa7O\xdbN\xa3\xdf\xd7\xa9\xa3y" +
	"s\xfd4i\xc6\xad\xcb\xd3\xe8f\x9b\xe3\x91\xdf\xcd." +
	"/\xcb\xce\xc6\xb0\xd3Y\x1d\x8f'\xb3\xd6\xd2xZ\xbf" +
	"\x11\xa7\x19\x9f\x01\xeb\x86\xbc9\xfe\xd1@\xc5\xff\x02\x00" +
	"\x00\xff\xff\xfa\xe7\xaf\x15"

func init() {
	schemas.Register(schema_e5fbb19e9e69110b,
		0x8a994a2aad4d9204,
		0xb4a68d38a716233d,
		0xe1068a6aee02baba)
}
