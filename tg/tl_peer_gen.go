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

// PeerUser represents TL type `peerUser#9db1bc6d`.
// Chat partner
//
// See https://core.telegram.org/constructor/peerUser for reference.
type PeerUser struct {
	// User identifier
	UserID int `tl:"user_id"`
}

// PeerUserTypeID is TL type id of PeerUser.
const PeerUserTypeID = 0x9db1bc6d

func (p *PeerUser) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerUser) String() string {
	if p == nil {
		return "PeerUser(nil)"
	}
	type Alias PeerUser
	return fmt.Sprintf("PeerUser%+v", Alias(*p))
}

// FillFrom fills PeerUser from given interface.
func (p *PeerUser) FillFrom(from interface {
	GetUserID() (value int)
}) {
	p.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PeerUser) TypeID() uint32 {
	return PeerUserTypeID
}

// TypeName returns name of type in TL schema.
func (p *PeerUser) TypeName() string {
	return "peerUser"
}

// Encode implements bin.Encoder.
func (p *PeerUser) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerUser#9db1bc6d as nil")
	}
	b.PutID(PeerUserTypeID)
	b.PutInt(p.UserID)
	return nil
}

// GetUserID returns value of UserID field.
func (p *PeerUser) GetUserID() (value int) {
	return p.UserID
}

// Decode implements bin.Decoder.
func (p *PeerUser) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerUser#9db1bc6d to nil")
	}
	if err := b.ConsumeID(PeerUserTypeID); err != nil {
		return fmt.Errorf("unable to decode peerUser#9db1bc6d: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerUser#9db1bc6d: field user_id: %w", err)
		}
		p.UserID = value
	}
	return nil
}

// construct implements constructor of PeerClass.
func (p PeerUser) construct() PeerClass { return &p }

// Ensuring interfaces in compile-time for PeerUser.
var (
	_ bin.Encoder = &PeerUser{}
	_ bin.Decoder = &PeerUser{}

	_ PeerClass = &PeerUser{}
)

// PeerChat represents TL type `peerChat#bad0e5bb`.
// Group.
//
// See https://core.telegram.org/constructor/peerChat for reference.
type PeerChat struct {
	// Group identifier
	ChatID int `tl:"chat_id"`
}

// PeerChatTypeID is TL type id of PeerChat.
const PeerChatTypeID = 0xbad0e5bb

func (p *PeerChat) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerChat) String() string {
	if p == nil {
		return "PeerChat(nil)"
	}
	type Alias PeerChat
	return fmt.Sprintf("PeerChat%+v", Alias(*p))
}

// FillFrom fills PeerChat from given interface.
func (p *PeerChat) FillFrom(from interface {
	GetChatID() (value int)
}) {
	p.ChatID = from.GetChatID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PeerChat) TypeID() uint32 {
	return PeerChatTypeID
}

// TypeName returns name of type in TL schema.
func (p *PeerChat) TypeName() string {
	return "peerChat"
}

// Encode implements bin.Encoder.
func (p *PeerChat) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerChat#bad0e5bb as nil")
	}
	b.PutID(PeerChatTypeID)
	b.PutInt(p.ChatID)
	return nil
}

// GetChatID returns value of ChatID field.
func (p *PeerChat) GetChatID() (value int) {
	return p.ChatID
}

// Decode implements bin.Decoder.
func (p *PeerChat) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerChat#bad0e5bb to nil")
	}
	if err := b.ConsumeID(PeerChatTypeID); err != nil {
		return fmt.Errorf("unable to decode peerChat#bad0e5bb: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerChat#bad0e5bb: field chat_id: %w", err)
		}
		p.ChatID = value
	}
	return nil
}

// construct implements constructor of PeerClass.
func (p PeerChat) construct() PeerClass { return &p }

// Ensuring interfaces in compile-time for PeerChat.
var (
	_ bin.Encoder = &PeerChat{}
	_ bin.Decoder = &PeerChat{}

	_ PeerClass = &PeerChat{}
)

// PeerChannel represents TL type `peerChannel#bddde532`.
// Channel/supergroup
//
// See https://core.telegram.org/constructor/peerChannel for reference.
type PeerChannel struct {
	// Channel ID
	ChannelID int `tl:"channel_id"`
}

// PeerChannelTypeID is TL type id of PeerChannel.
const PeerChannelTypeID = 0xbddde532

func (p *PeerChannel) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ChannelID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerChannel) String() string {
	if p == nil {
		return "PeerChannel(nil)"
	}
	type Alias PeerChannel
	return fmt.Sprintf("PeerChannel%+v", Alias(*p))
}

// FillFrom fills PeerChannel from given interface.
func (p *PeerChannel) FillFrom(from interface {
	GetChannelID() (value int)
}) {
	p.ChannelID = from.GetChannelID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PeerChannel) TypeID() uint32 {
	return PeerChannelTypeID
}

// TypeName returns name of type in TL schema.
func (p *PeerChannel) TypeName() string {
	return "peerChannel"
}

// Encode implements bin.Encoder.
func (p *PeerChannel) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerChannel#bddde532 as nil")
	}
	b.PutID(PeerChannelTypeID)
	b.PutInt(p.ChannelID)
	return nil
}

// GetChannelID returns value of ChannelID field.
func (p *PeerChannel) GetChannelID() (value int) {
	return p.ChannelID
}

// Decode implements bin.Decoder.
func (p *PeerChannel) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerChannel#bddde532 to nil")
	}
	if err := b.ConsumeID(PeerChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode peerChannel#bddde532: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerChannel#bddde532: field channel_id: %w", err)
		}
		p.ChannelID = value
	}
	return nil
}

// construct implements constructor of PeerClass.
func (p PeerChannel) construct() PeerClass { return &p }

// Ensuring interfaces in compile-time for PeerChannel.
var (
	_ bin.Encoder = &PeerChannel{}
	_ bin.Decoder = &PeerChannel{}

	_ PeerClass = &PeerChannel{}
)

// PeerClass represents Peer generic type.
//
// See https://core.telegram.org/type/Peer for reference.
//
// Example:
//  g, err := tg.DecodePeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.PeerUser: // peerUser#9db1bc6d
//  case *tg.PeerChat: // peerChat#bad0e5bb
//  case *tg.PeerChannel: // peerChannel#bddde532
//  default: panic(v)
//  }
type PeerClass interface {
	bin.Encoder
	bin.Decoder
	construct() PeerClass

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
}

// AsInput tries to map PeerChat to InputPeerChat.
func (p *PeerChat) AsInput() *InputPeerChat {
	value := new(InputPeerChat)
	value.ChatID = p.GetChatID()

	return value
}

// DecodePeer implements binary de-serialization for PeerClass.
func DecodePeer(buf *bin.Buffer) (PeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PeerUserTypeID:
		// Decoding peerUser#9db1bc6d.
		v := PeerUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PeerClass: %w", err)
		}
		return &v, nil
	case PeerChatTypeID:
		// Decoding peerChat#bad0e5bb.
		v := PeerChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PeerClass: %w", err)
		}
		return &v, nil
	case PeerChannelTypeID:
		// Decoding peerChannel#bddde532.
		v := PeerChannel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// Peer boxes the PeerClass providing a helper.
type PeerBox struct {
	Peer PeerClass
}

// Decode implements bin.Decoder for PeerBox.
func (b *PeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PeerBox to nil")
	}
	v, err := DecodePeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Peer = v
	return nil
}

// Encode implements bin.Encode for PeerBox.
func (b *PeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Peer == nil {
		return fmt.Errorf("unable to encode PeerClass as nil")
	}
	return b.Peer.Encode(buf)
}

// PeerClassArray is adapter for slice of PeerClass.
type PeerClassArray []PeerClass

// Sort sorts slice of PeerClass.
func (s PeerClassArray) Sort(less func(a, b PeerClass) bool) PeerClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PeerClass.
func (s PeerClassArray) SortStable(less func(a, b PeerClass) bool) PeerClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PeerClass.
func (s PeerClassArray) Retain(keep func(x PeerClass) bool) PeerClassArray {
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
func (s PeerClassArray) First() (v PeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PeerClassArray) Last() (v PeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PeerClassArray) PopFirst() (v PeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PeerClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PeerClassArray) Pop() (v PeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPeerUser returns copy with only PeerUser constructors.
func (s PeerClassArray) AsPeerUser() (to PeerUserArray) {
	for _, elem := range s {
		value, ok := elem.(*PeerUser)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPeerChat returns copy with only PeerChat constructors.
func (s PeerClassArray) AsPeerChat() (to PeerChatArray) {
	for _, elem := range s {
		value, ok := elem.(*PeerChat)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPeerChannel returns copy with only PeerChannel constructors.
func (s PeerClassArray) AsPeerChannel() (to PeerChannelArray) {
	for _, elem := range s {
		value, ok := elem.(*PeerChannel)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PeerUserArray is adapter for slice of PeerUser.
type PeerUserArray []PeerUser

// Sort sorts slice of PeerUser.
func (s PeerUserArray) Sort(less func(a, b PeerUser) bool) PeerUserArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PeerUser.
func (s PeerUserArray) SortStable(less func(a, b PeerUser) bool) PeerUserArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PeerUser.
func (s PeerUserArray) Retain(keep func(x PeerUser) bool) PeerUserArray {
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
func (s PeerUserArray) First() (v PeerUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PeerUserArray) Last() (v PeerUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PeerUserArray) PopFirst() (v PeerUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PeerUser
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PeerUserArray) Pop() (v PeerUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PeerChatArray is adapter for slice of PeerChat.
type PeerChatArray []PeerChat

// Sort sorts slice of PeerChat.
func (s PeerChatArray) Sort(less func(a, b PeerChat) bool) PeerChatArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PeerChat.
func (s PeerChatArray) SortStable(less func(a, b PeerChat) bool) PeerChatArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PeerChat.
func (s PeerChatArray) Retain(keep func(x PeerChat) bool) PeerChatArray {
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
func (s PeerChatArray) First() (v PeerChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PeerChatArray) Last() (v PeerChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PeerChatArray) PopFirst() (v PeerChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PeerChat
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PeerChatArray) Pop() (v PeerChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PeerChannelArray is adapter for slice of PeerChannel.
type PeerChannelArray []PeerChannel

// Sort sorts slice of PeerChannel.
func (s PeerChannelArray) Sort(less func(a, b PeerChannel) bool) PeerChannelArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PeerChannel.
func (s PeerChannelArray) SortStable(less func(a, b PeerChannel) bool) PeerChannelArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PeerChannel.
func (s PeerChannelArray) Retain(keep func(x PeerChannel) bool) PeerChannelArray {
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
func (s PeerChannelArray) First() (v PeerChannel, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PeerChannelArray) Last() (v PeerChannel, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PeerChannelArray) PopFirst() (v PeerChannel, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PeerChannel
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PeerChannelArray) Pop() (v PeerChannel, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
