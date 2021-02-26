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

// InputMessageID represents TL type `inputMessageID#a676a322`.
// Message by ID
//
// See https://core.telegram.org/constructor/inputMessageID for reference.
type InputMessageID struct {
	// Message ID
	ID int `tl:"id"`
}

// InputMessageIDTypeID is TL type id of InputMessageID.
const InputMessageIDTypeID = 0xa676a322

func (i *InputMessageID) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputMessageID) String() string {
	if i == nil {
		return "InputMessageID(nil)"
	}
	type Alias InputMessageID
	return fmt.Sprintf("InputMessageID%+v", Alias(*i))
}

// FillFrom fills InputMessageID from given interface.
func (i *InputMessageID) FillFrom(from interface {
	GetID() (value int)
}) {
	i.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputMessageID) TypeID() uint32 {
	return InputMessageIDTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputMessageID) TypeName() string {
	return "inputMessageID"
}

// Encode implements bin.Encoder.
func (i *InputMessageID) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageID#a676a322 as nil")
	}
	b.PutID(InputMessageIDTypeID)
	b.PutInt(i.ID)
	return nil
}

// GetID returns value of ID field.
func (i *InputMessageID) GetID() (value int) {
	return i.ID
}

// Decode implements bin.Decoder.
func (i *InputMessageID) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageID#a676a322 to nil")
	}
	if err := b.ConsumeID(InputMessageIDTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessageID#a676a322: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageID#a676a322: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// construct implements constructor of InputMessageClass.
func (i InputMessageID) construct() InputMessageClass { return &i }

// Ensuring interfaces in compile-time for InputMessageID.
var (
	_ bin.Encoder = &InputMessageID{}
	_ bin.Decoder = &InputMessageID{}

	_ InputMessageClass = &InputMessageID{}
)

// InputMessageReplyTo represents TL type `inputMessageReplyTo#bad88395`.
// Message to which the specified message replies to
//
// See https://core.telegram.org/constructor/inputMessageReplyTo for reference.
type InputMessageReplyTo struct {
	// ID of the message that replies to the message we need
	ID int `tl:"id"`
}

// InputMessageReplyToTypeID is TL type id of InputMessageReplyTo.
const InputMessageReplyToTypeID = 0xbad88395

func (i *InputMessageReplyTo) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputMessageReplyTo) String() string {
	if i == nil {
		return "InputMessageReplyTo(nil)"
	}
	type Alias InputMessageReplyTo
	return fmt.Sprintf("InputMessageReplyTo%+v", Alias(*i))
}

// FillFrom fills InputMessageReplyTo from given interface.
func (i *InputMessageReplyTo) FillFrom(from interface {
	GetID() (value int)
}) {
	i.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputMessageReplyTo) TypeID() uint32 {
	return InputMessageReplyToTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputMessageReplyTo) TypeName() string {
	return "inputMessageReplyTo"
}

// Encode implements bin.Encoder.
func (i *InputMessageReplyTo) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageReplyTo#bad88395 as nil")
	}
	b.PutID(InputMessageReplyToTypeID)
	b.PutInt(i.ID)
	return nil
}

// GetID returns value of ID field.
func (i *InputMessageReplyTo) GetID() (value int) {
	return i.ID
}

// Decode implements bin.Decoder.
func (i *InputMessageReplyTo) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageReplyTo#bad88395 to nil")
	}
	if err := b.ConsumeID(InputMessageReplyToTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessageReplyTo#bad88395: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageReplyTo#bad88395: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// construct implements constructor of InputMessageClass.
func (i InputMessageReplyTo) construct() InputMessageClass { return &i }

// Ensuring interfaces in compile-time for InputMessageReplyTo.
var (
	_ bin.Encoder = &InputMessageReplyTo{}
	_ bin.Decoder = &InputMessageReplyTo{}

	_ InputMessageClass = &InputMessageReplyTo{}
)

// InputMessagePinned represents TL type `inputMessagePinned#86872538`.
// Pinned message
//
// See https://core.telegram.org/constructor/inputMessagePinned for reference.
type InputMessagePinned struct {
}

// InputMessagePinnedTypeID is TL type id of InputMessagePinned.
const InputMessagePinnedTypeID = 0x86872538

func (i *InputMessagePinned) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputMessagePinned) String() string {
	if i == nil {
		return "InputMessagePinned(nil)"
	}
	type Alias InputMessagePinned
	return fmt.Sprintf("InputMessagePinned%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputMessagePinned) TypeID() uint32 {
	return InputMessagePinnedTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputMessagePinned) TypeName() string {
	return "inputMessagePinned"
}

// Encode implements bin.Encoder.
func (i *InputMessagePinned) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagePinned#86872538 as nil")
	}
	b.PutID(InputMessagePinnedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagePinned) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagePinned#86872538 to nil")
	}
	if err := b.ConsumeID(InputMessagePinnedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagePinned#86872538: %w", err)
	}
	return nil
}

// construct implements constructor of InputMessageClass.
func (i InputMessagePinned) construct() InputMessageClass { return &i }

// Ensuring interfaces in compile-time for InputMessagePinned.
var (
	_ bin.Encoder = &InputMessagePinned{}
	_ bin.Decoder = &InputMessagePinned{}

	_ InputMessageClass = &InputMessagePinned{}
)

// InputMessageCallbackQuery represents TL type `inputMessageCallbackQuery#acfa1a7e`.
// Used by bots for fetching information about the message that originated a callback query
//
// See https://core.telegram.org/constructor/inputMessageCallbackQuery for reference.
type InputMessageCallbackQuery struct {
	// Message ID
	ID int `tl:"id"`
	// Callback query ID
	QueryID int64 `tl:"query_id"`
}

// InputMessageCallbackQueryTypeID is TL type id of InputMessageCallbackQuery.
const InputMessageCallbackQueryTypeID = 0xacfa1a7e

func (i *InputMessageCallbackQuery) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.QueryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputMessageCallbackQuery) String() string {
	if i == nil {
		return "InputMessageCallbackQuery(nil)"
	}
	type Alias InputMessageCallbackQuery
	return fmt.Sprintf("InputMessageCallbackQuery%+v", Alias(*i))
}

// FillFrom fills InputMessageCallbackQuery from given interface.
func (i *InputMessageCallbackQuery) FillFrom(from interface {
	GetID() (value int)
	GetQueryID() (value int64)
}) {
	i.ID = from.GetID()
	i.QueryID = from.GetQueryID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputMessageCallbackQuery) TypeID() uint32 {
	return InputMessageCallbackQueryTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputMessageCallbackQuery) TypeName() string {
	return "inputMessageCallbackQuery"
}

// Encode implements bin.Encoder.
func (i *InputMessageCallbackQuery) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageCallbackQuery#acfa1a7e as nil")
	}
	b.PutID(InputMessageCallbackQueryTypeID)
	b.PutInt(i.ID)
	b.PutLong(i.QueryID)
	return nil
}

// GetID returns value of ID field.
func (i *InputMessageCallbackQuery) GetID() (value int) {
	return i.ID
}

// GetQueryID returns value of QueryID field.
func (i *InputMessageCallbackQuery) GetQueryID() (value int64) {
	return i.QueryID
}

// Decode implements bin.Decoder.
func (i *InputMessageCallbackQuery) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageCallbackQuery#acfa1a7e to nil")
	}
	if err := b.ConsumeID(InputMessageCallbackQueryTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessageCallbackQuery#acfa1a7e: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageCallbackQuery#acfa1a7e: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageCallbackQuery#acfa1a7e: field query_id: %w", err)
		}
		i.QueryID = value
	}
	return nil
}

// construct implements constructor of InputMessageClass.
func (i InputMessageCallbackQuery) construct() InputMessageClass { return &i }

// Ensuring interfaces in compile-time for InputMessageCallbackQuery.
var (
	_ bin.Encoder = &InputMessageCallbackQuery{}
	_ bin.Decoder = &InputMessageCallbackQuery{}

	_ InputMessageClass = &InputMessageCallbackQuery{}
)

// InputMessageClass represents InputMessage generic type.
//
// See https://core.telegram.org/type/InputMessage for reference.
//
// Example:
//  g, err := tg.DecodeInputMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputMessageID: // inputMessageID#a676a322
//  case *tg.InputMessageReplyTo: // inputMessageReplyTo#bad88395
//  case *tg.InputMessagePinned: // inputMessagePinned#86872538
//  case *tg.InputMessageCallbackQuery: // inputMessageCallbackQuery#acfa1a7e
//  default: panic(v)
//  }
type InputMessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputMessageClass

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

// DecodeInputMessage implements binary de-serialization for InputMessageClass.
func DecodeInputMessage(buf *bin.Buffer) (InputMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputMessageIDTypeID:
		// Decoding inputMessageID#a676a322.
		v := InputMessageID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageClass: %w", err)
		}
		return &v, nil
	case InputMessageReplyToTypeID:
		// Decoding inputMessageReplyTo#bad88395.
		v := InputMessageReplyTo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageClass: %w", err)
		}
		return &v, nil
	case InputMessagePinnedTypeID:
		// Decoding inputMessagePinned#86872538.
		v := InputMessagePinned{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageClass: %w", err)
		}
		return &v, nil
	case InputMessageCallbackQueryTypeID:
		// Decoding inputMessageCallbackQuery#acfa1a7e.
		v := InputMessageCallbackQuery{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputMessage boxes the InputMessageClass providing a helper.
type InputMessageBox struct {
	InputMessage InputMessageClass
}

// Decode implements bin.Decoder for InputMessageBox.
func (b *InputMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputMessageBox to nil")
	}
	v, err := DecodeInputMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputMessage = v
	return nil
}

// Encode implements bin.Encode for InputMessageBox.
func (b *InputMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputMessage == nil {
		return fmt.Errorf("unable to encode InputMessageClass as nil")
	}
	return b.InputMessage.Encode(buf)
}

// InputMessageClassArray is adapter for slice of InputMessageClass.
type InputMessageClassArray []InputMessageClass

// Sort sorts slice of InputMessageClass.
func (s InputMessageClassArray) Sort(less func(a, b InputMessageClass) bool) InputMessageClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMessageClass.
func (s InputMessageClassArray) SortStable(less func(a, b InputMessageClass) bool) InputMessageClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMessageClass.
func (s InputMessageClassArray) Retain(keep func(x InputMessageClass) bool) InputMessageClassArray {
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
func (s InputMessageClassArray) First() (v InputMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMessageClassArray) Last() (v InputMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMessageClassArray) PopFirst() (v InputMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMessageClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMessageClassArray) Pop() (v InputMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputMessageID returns copy with only InputMessageID constructors.
func (s InputMessageClassArray) AsInputMessageID() (to InputMessageIDArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMessageID)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMessageReplyTo returns copy with only InputMessageReplyTo constructors.
func (s InputMessageClassArray) AsInputMessageReplyTo() (to InputMessageReplyToArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMessageReplyTo)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputMessageCallbackQuery returns copy with only InputMessageCallbackQuery constructors.
func (s InputMessageClassArray) AsInputMessageCallbackQuery() (to InputMessageCallbackQueryArray) {
	for _, elem := range s {
		value, ok := elem.(*InputMessageCallbackQuery)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputMessageIDArray is adapter for slice of InputMessageID.
type InputMessageIDArray []InputMessageID

// Sort sorts slice of InputMessageID.
func (s InputMessageIDArray) Sort(less func(a, b InputMessageID) bool) InputMessageIDArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMessageID.
func (s InputMessageIDArray) SortStable(less func(a, b InputMessageID) bool) InputMessageIDArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMessageID.
func (s InputMessageIDArray) Retain(keep func(x InputMessageID) bool) InputMessageIDArray {
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
func (s InputMessageIDArray) First() (v InputMessageID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMessageIDArray) Last() (v InputMessageID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMessageIDArray) PopFirst() (v InputMessageID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMessageID
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMessageIDArray) Pop() (v InputMessageID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputMessageID by ID.
func (s InputMessageIDArray) SortByID() InputMessageIDArray {
	return s.Sort(func(a, b InputMessageID) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputMessageID by ID.
func (s InputMessageIDArray) SortStableByID() InputMessageIDArray {
	return s.SortStable(func(a, b InputMessageID) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputMessageIDArray) FillMap(to map[int]InputMessageID) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputMessageIDArray) ToMap() map[int]InputMessageID {
	r := make(map[int]InputMessageID, len(s))
	s.FillMap(r)
	return r
}

// InputMessageReplyToArray is adapter for slice of InputMessageReplyTo.
type InputMessageReplyToArray []InputMessageReplyTo

// Sort sorts slice of InputMessageReplyTo.
func (s InputMessageReplyToArray) Sort(less func(a, b InputMessageReplyTo) bool) InputMessageReplyToArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMessageReplyTo.
func (s InputMessageReplyToArray) SortStable(less func(a, b InputMessageReplyTo) bool) InputMessageReplyToArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMessageReplyTo.
func (s InputMessageReplyToArray) Retain(keep func(x InputMessageReplyTo) bool) InputMessageReplyToArray {
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
func (s InputMessageReplyToArray) First() (v InputMessageReplyTo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMessageReplyToArray) Last() (v InputMessageReplyTo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMessageReplyToArray) PopFirst() (v InputMessageReplyTo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMessageReplyTo
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMessageReplyToArray) Pop() (v InputMessageReplyTo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputMessageReplyTo by ID.
func (s InputMessageReplyToArray) SortByID() InputMessageReplyToArray {
	return s.Sort(func(a, b InputMessageReplyTo) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputMessageReplyTo by ID.
func (s InputMessageReplyToArray) SortStableByID() InputMessageReplyToArray {
	return s.SortStable(func(a, b InputMessageReplyTo) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputMessageReplyToArray) FillMap(to map[int]InputMessageReplyTo) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputMessageReplyToArray) ToMap() map[int]InputMessageReplyTo {
	r := make(map[int]InputMessageReplyTo, len(s))
	s.FillMap(r)
	return r
}

// InputMessageCallbackQueryArray is adapter for slice of InputMessageCallbackQuery.
type InputMessageCallbackQueryArray []InputMessageCallbackQuery

// Sort sorts slice of InputMessageCallbackQuery.
func (s InputMessageCallbackQueryArray) Sort(less func(a, b InputMessageCallbackQuery) bool) InputMessageCallbackQueryArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputMessageCallbackQuery.
func (s InputMessageCallbackQueryArray) SortStable(less func(a, b InputMessageCallbackQuery) bool) InputMessageCallbackQueryArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputMessageCallbackQuery.
func (s InputMessageCallbackQueryArray) Retain(keep func(x InputMessageCallbackQuery) bool) InputMessageCallbackQueryArray {
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
func (s InputMessageCallbackQueryArray) First() (v InputMessageCallbackQuery, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputMessageCallbackQueryArray) Last() (v InputMessageCallbackQuery, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputMessageCallbackQueryArray) PopFirst() (v InputMessageCallbackQuery, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputMessageCallbackQuery
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputMessageCallbackQueryArray) Pop() (v InputMessageCallbackQuery, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputMessageCallbackQuery by ID.
func (s InputMessageCallbackQueryArray) SortByID() InputMessageCallbackQueryArray {
	return s.Sort(func(a, b InputMessageCallbackQuery) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputMessageCallbackQuery by ID.
func (s InputMessageCallbackQueryArray) SortStableByID() InputMessageCallbackQueryArray {
	return s.SortStable(func(a, b InputMessageCallbackQuery) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputMessageCallbackQueryArray) FillMap(to map[int]InputMessageCallbackQuery) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputMessageCallbackQueryArray) ToMap() map[int]InputMessageCallbackQuery {
	r := make(map[int]InputMessageCallbackQuery, len(s))
	s.FillMap(r)
	return r
}
