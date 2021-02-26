// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints

// PhoneConnection represents TL type `phoneConnection#9d4c17c0`.
// Identifies an endpoint that can be used to connect to the other user in a phone call
//
// See https://core.telegram.org/constructor/phoneConnection for reference.
type PhoneConnection struct {
	// Endpoint ID
	ID int64 `tl:"id"`
	// IP address of endpoint
	IP string `tl:"ip"`
	// IPv6 address of endpoint
	Ipv6 string `tl:"ipv6"`
	// Port ID
	Port int `tl:"port"`
	// Our peer tag
	PeerTag []byte `tl:"peer_tag"`
}

// PhoneConnectionTypeID is TL type id of PhoneConnection.
const PhoneConnectionTypeID = 0x9d4c17c0

func (p *PhoneConnection) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.IP == "") {
		return false
	}
	if !(p.Ipv6 == "") {
		return false
	}
	if !(p.Port == 0) {
		return false
	}
	if !(p.PeerTag == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneConnection) String() string {
	if p == nil {
		return "PhoneConnection(nil)"
	}
	type Alias PhoneConnection
	return fmt.Sprintf("PhoneConnection%+v", Alias(*p))
}

// FillFrom fills PhoneConnection from given interface.
func (p *PhoneConnection) FillFrom(from interface {
	GetID() (value int64)
	GetIP() (value string)
	GetIpv6() (value string)
	GetPort() (value int)
	GetPeerTag() (value []byte)
}) {
	p.ID = from.GetID()
	p.IP = from.GetIP()
	p.Ipv6 = from.GetIpv6()
	p.Port = from.GetPort()
	p.PeerTag = from.GetPeerTag()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhoneConnection) TypeID() uint32 {
	return PhoneConnectionTypeID
}

// TypeName returns name of type in TL schema.
func (p *PhoneConnection) TypeName() string {
	return "phoneConnection"
}

// Encode implements bin.Encoder.
func (p *PhoneConnection) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneConnection#9d4c17c0 as nil")
	}
	b.PutID(PhoneConnectionTypeID)
	b.PutLong(p.ID)
	b.PutString(p.IP)
	b.PutString(p.Ipv6)
	b.PutInt(p.Port)
	b.PutBytes(p.PeerTag)
	return nil
}

// GetID returns value of ID field.
func (p *PhoneConnection) GetID() (value int64) {
	return p.ID
}

// GetIP returns value of IP field.
func (p *PhoneConnection) GetIP() (value string) {
	return p.IP
}

// GetIpv6 returns value of Ipv6 field.
func (p *PhoneConnection) GetIpv6() (value string) {
	return p.Ipv6
}

// GetPort returns value of Port field.
func (p *PhoneConnection) GetPort() (value int) {
	return p.Port
}

// GetPeerTag returns value of PeerTag field.
func (p *PhoneConnection) GetPeerTag() (value []byte) {
	return p.PeerTag
}

// Decode implements bin.Decoder.
func (p *PhoneConnection) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneConnection#9d4c17c0 to nil")
	}
	if err := b.ConsumeID(PhoneConnectionTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field ip: %w", err)
		}
		p.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field ipv6: %w", err)
		}
		p.Ipv6 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field port: %w", err)
		}
		p.Port = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field peer_tag: %w", err)
		}
		p.PeerTag = value
	}
	return nil
}

// construct implements constructor of PhoneConnectionClass.
func (p PhoneConnection) construct() PhoneConnectionClass { return &p }

// Ensuring interfaces in compile-time for PhoneConnection.
var (
	_ bin.Encoder = &PhoneConnection{}
	_ bin.Decoder = &PhoneConnection{}

	_ PhoneConnectionClass = &PhoneConnection{}
)

// PhoneConnectionWebrtc represents TL type `phoneConnectionWebrtc#635fe375`.
// WebRTC connection parameters
//
// See https://core.telegram.org/constructor/phoneConnectionWebrtc for reference.
type PhoneConnectionWebrtc struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Whether this is a TURN endpoint
	Turn bool `tl:"turn"`
	// Whether this is a STUN endpoint
	Stun bool `tl:"stun"`
	// Endpoint ID
	ID int64 `tl:"id"`
	// IP address
	IP string `tl:"ip"`
	// IPv6 address
	Ipv6 string `tl:"ipv6"`
	// Port
	Port int `tl:"port"`
	// Username
	Username string `tl:"username"`
	// Password
	Password string `tl:"password"`
}

// PhoneConnectionWebrtcTypeID is TL type id of PhoneConnectionWebrtc.
const PhoneConnectionWebrtcTypeID = 0x635fe375

func (p *PhoneConnectionWebrtc) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Turn == false) {
		return false
	}
	if !(p.Stun == false) {
		return false
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.IP == "") {
		return false
	}
	if !(p.Ipv6 == "") {
		return false
	}
	if !(p.Port == 0) {
		return false
	}
	if !(p.Username == "") {
		return false
	}
	if !(p.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneConnectionWebrtc) String() string {
	if p == nil {
		return "PhoneConnectionWebrtc(nil)"
	}
	type Alias PhoneConnectionWebrtc
	return fmt.Sprintf("PhoneConnectionWebrtc%+v", Alias(*p))
}

// FillFrom fills PhoneConnectionWebrtc from given interface.
func (p *PhoneConnectionWebrtc) FillFrom(from interface {
	GetTurn() (value bool)
	GetStun() (value bool)
	GetID() (value int64)
	GetIP() (value string)
	GetIpv6() (value string)
	GetPort() (value int)
	GetUsername() (value string)
	GetPassword() (value string)
}) {
	p.Turn = from.GetTurn()
	p.Stun = from.GetStun()
	p.ID = from.GetID()
	p.IP = from.GetIP()
	p.Ipv6 = from.GetIpv6()
	p.Port = from.GetPort()
	p.Username = from.GetUsername()
	p.Password = from.GetPassword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhoneConnectionWebrtc) TypeID() uint32 {
	return PhoneConnectionWebrtcTypeID
}

// TypeName returns name of type in TL schema.
func (p *PhoneConnectionWebrtc) TypeName() string {
	return "phoneConnectionWebrtc"
}

// Encode implements bin.Encoder.
func (p *PhoneConnectionWebrtc) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneConnectionWebrtc#635fe375 as nil")
	}
	b.PutID(PhoneConnectionWebrtcTypeID)
	if !(p.Turn == false) {
		p.Flags.Set(0)
	}
	if !(p.Stun == false) {
		p.Flags.Set(1)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneConnectionWebrtc#635fe375: field flags: %w", err)
	}
	b.PutLong(p.ID)
	b.PutString(p.IP)
	b.PutString(p.Ipv6)
	b.PutInt(p.Port)
	b.PutString(p.Username)
	b.PutString(p.Password)
	return nil
}

// SetTurn sets value of Turn conditional field.
func (p *PhoneConnectionWebrtc) SetTurn(value bool) {
	if value {
		p.Flags.Set(0)
		p.Turn = true
	} else {
		p.Flags.Unset(0)
		p.Turn = false
	}
}

// GetTurn returns value of Turn conditional field.
func (p *PhoneConnectionWebrtc) GetTurn() (value bool) {
	return p.Flags.Has(0)
}

// SetStun sets value of Stun conditional field.
func (p *PhoneConnectionWebrtc) SetStun(value bool) {
	if value {
		p.Flags.Set(1)
		p.Stun = true
	} else {
		p.Flags.Unset(1)
		p.Stun = false
	}
}

// GetStun returns value of Stun conditional field.
func (p *PhoneConnectionWebrtc) GetStun() (value bool) {
	return p.Flags.Has(1)
}

// GetID returns value of ID field.
func (p *PhoneConnectionWebrtc) GetID() (value int64) {
	return p.ID
}

// GetIP returns value of IP field.
func (p *PhoneConnectionWebrtc) GetIP() (value string) {
	return p.IP
}

// GetIpv6 returns value of Ipv6 field.
func (p *PhoneConnectionWebrtc) GetIpv6() (value string) {
	return p.Ipv6
}

// GetPort returns value of Port field.
func (p *PhoneConnectionWebrtc) GetPort() (value int) {
	return p.Port
}

// GetUsername returns value of Username field.
func (p *PhoneConnectionWebrtc) GetUsername() (value string) {
	return p.Username
}

// GetPassword returns value of Password field.
func (p *PhoneConnectionWebrtc) GetPassword() (value string) {
	return p.Password
}

// Decode implements bin.Decoder.
func (p *PhoneConnectionWebrtc) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneConnectionWebrtc#635fe375 to nil")
	}
	if err := b.ConsumeID(PhoneConnectionWebrtcTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field flags: %w", err)
		}
	}
	p.Turn = p.Flags.Has(0)
	p.Stun = p.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field ip: %w", err)
		}
		p.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field ipv6: %w", err)
		}
		p.Ipv6 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field port: %w", err)
		}
		p.Port = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field username: %w", err)
		}
		p.Username = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field password: %w", err)
		}
		p.Password = value
	}
	return nil
}

// construct implements constructor of PhoneConnectionClass.
func (p PhoneConnectionWebrtc) construct() PhoneConnectionClass { return &p }

// Ensuring interfaces in compile-time for PhoneConnectionWebrtc.
var (
	_ bin.Encoder = &PhoneConnectionWebrtc{}
	_ bin.Decoder = &PhoneConnectionWebrtc{}

	_ PhoneConnectionClass = &PhoneConnectionWebrtc{}
)

// PhoneConnectionClass represents PhoneConnection generic type.
//
// See https://core.telegram.org/type/PhoneConnection for reference.
//
// Example:
//  g, err := tg.DecodePhoneConnection(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.PhoneConnection: // phoneConnection#9d4c17c0
//  case *tg.PhoneConnectionWebrtc: // phoneConnectionWebrtc#635fe375
//  default: panic(v)
//  }
type PhoneConnectionClass interface {
	bin.Encoder
	bin.Decoder
	construct() PhoneConnectionClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Endpoint ID
	GetID() (value int64)
	// IP address of endpoint
	GetIP() (value string)
	// IPv6 address of endpoint
	GetIpv6() (value string)
	// Port ID
	GetPort() (value int)
}

// DecodePhoneConnection implements binary de-serialization for PhoneConnectionClass.
func DecodePhoneConnection(buf *bin.Buffer) (PhoneConnectionClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhoneConnectionTypeID:
		// Decoding phoneConnection#9d4c17c0.
		v := PhoneConnection{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneConnectionClass: %w", err)
		}
		return &v, nil
	case PhoneConnectionWebrtcTypeID:
		// Decoding phoneConnectionWebrtc#635fe375.
		v := PhoneConnectionWebrtc{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneConnectionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhoneConnectionClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhoneConnection boxes the PhoneConnectionClass providing a helper.
type PhoneConnectionBox struct {
	PhoneConnection PhoneConnectionClass
}

// Decode implements bin.Decoder for PhoneConnectionBox.
func (b *PhoneConnectionBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhoneConnectionBox to nil")
	}
	v, err := DecodePhoneConnection(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PhoneConnection = v
	return nil
}

// Encode implements bin.Encode for PhoneConnectionBox.
func (b *PhoneConnectionBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PhoneConnection == nil {
		return fmt.Errorf("unable to encode PhoneConnectionClass as nil")
	}
	return b.PhoneConnection.Encode(buf)
}

// PhoneConnectionClassArray is adapter for slice of PhoneConnectionClass.
type PhoneConnectionClassArray []PhoneConnectionClass

// Sort sorts slice of PhoneConnectionClass.
func (s PhoneConnectionClassArray) Sort(less func(a, b PhoneConnectionClass) bool) PhoneConnectionClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneConnectionClass.
func (s PhoneConnectionClassArray) SortStable(less func(a, b PhoneConnectionClass) bool) PhoneConnectionClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneConnectionClass.
func (s PhoneConnectionClassArray) Retain(keep func(x PhoneConnectionClass) bool) PhoneConnectionClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneConnectionClassArray) First() (v PhoneConnectionClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneConnectionClassArray) Last() (v PhoneConnectionClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneConnectionClassArray) PopFirst() (v PhoneConnectionClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneConnectionClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneConnectionClassArray) Pop() (v PhoneConnectionClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPhoneConnection returns copy with only PhoneConnection constructors.
func (s PhoneConnectionClassArray) AsPhoneConnection() (to PhoneConnectionArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneConnection)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPhoneConnectionWebrtc returns copy with only PhoneConnectionWebrtc constructors.
func (s PhoneConnectionClassArray) AsPhoneConnectionWebrtc() (to PhoneConnectionWebrtcArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneConnectionWebrtc)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PhoneConnectionArray is adapter for slice of PhoneConnection.
type PhoneConnectionArray []PhoneConnection

// Sort sorts slice of PhoneConnection.
func (s PhoneConnectionArray) Sort(less func(a, b PhoneConnection) bool) PhoneConnectionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneConnection.
func (s PhoneConnectionArray) SortStable(less func(a, b PhoneConnection) bool) PhoneConnectionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneConnection.
func (s PhoneConnectionArray) Retain(keep func(x PhoneConnection) bool) PhoneConnectionArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneConnectionArray) First() (v PhoneConnection, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneConnectionArray) Last() (v PhoneConnection, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneConnectionArray) PopFirst() (v PhoneConnection, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneConnection
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneConnectionArray) Pop() (v PhoneConnection, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PhoneConnectionWebrtcArray is adapter for slice of PhoneConnectionWebrtc.
type PhoneConnectionWebrtcArray []PhoneConnectionWebrtc

// Sort sorts slice of PhoneConnectionWebrtc.
func (s PhoneConnectionWebrtcArray) Sort(less func(a, b PhoneConnectionWebrtc) bool) PhoneConnectionWebrtcArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneConnectionWebrtc.
func (s PhoneConnectionWebrtcArray) SortStable(less func(a, b PhoneConnectionWebrtc) bool) PhoneConnectionWebrtcArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneConnectionWebrtc.
func (s PhoneConnectionWebrtcArray) Retain(keep func(x PhoneConnectionWebrtc) bool) PhoneConnectionWebrtcArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneConnectionWebrtcArray) First() (v PhoneConnectionWebrtc, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneConnectionWebrtcArray) Last() (v PhoneConnectionWebrtc, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneConnectionWebrtcArray) PopFirst() (v PhoneConnectionWebrtc, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneConnectionWebrtc
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneConnectionWebrtcArray) Pop() (v PhoneConnectionWebrtc, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
