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

// MessageInteractionCounters represents TL type `messageInteractionCounters#ad4fc9bd`.
// Message interaction counters
//
// See https://core.telegram.org/constructor/messageInteractionCounters for reference.
type MessageInteractionCounters struct {
	// Message ID
	MsgID int `tl:"msg_id"`
	// Views
	Views int `tl:"views"`
	// Number of times this message was forwarded
	Forwards int `tl:"forwards"`
}

// MessageInteractionCountersTypeID is TL type id of MessageInteractionCounters.
const MessageInteractionCountersTypeID = 0xad4fc9bd

func (m *MessageInteractionCounters) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgID == 0) {
		return false
	}
	if !(m.Views == 0) {
		return false
	}
	if !(m.Forwards == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageInteractionCounters) String() string {
	if m == nil {
		return "MessageInteractionCounters(nil)"
	}
	type Alias MessageInteractionCounters
	return fmt.Sprintf("MessageInteractionCounters%+v", Alias(*m))
}

// FillFrom fills MessageInteractionCounters from given interface.
func (m *MessageInteractionCounters) FillFrom(from interface {
	GetMsgID() (value int)
	GetViews() (value int)
	GetForwards() (value int)
}) {
	m.MsgID = from.GetMsgID()
	m.Views = from.GetViews()
	m.Forwards = from.GetForwards()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessageInteractionCounters) TypeID() uint32 {
	return MessageInteractionCountersTypeID
}

// TypeName returns name of type in TL schema.
func (m *MessageInteractionCounters) TypeName() string {
	return "messageInteractionCounters"
}

// Encode implements bin.Encoder.
func (m *MessageInteractionCounters) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageInteractionCounters#ad4fc9bd as nil")
	}
	b.PutID(MessageInteractionCountersTypeID)
	b.PutInt(m.MsgID)
	b.PutInt(m.Views)
	b.PutInt(m.Forwards)
	return nil
}

// GetMsgID returns value of MsgID field.
func (m *MessageInteractionCounters) GetMsgID() (value int) {
	return m.MsgID
}

// GetViews returns value of Views field.
func (m *MessageInteractionCounters) GetViews() (value int) {
	return m.Views
}

// GetForwards returns value of Forwards field.
func (m *MessageInteractionCounters) GetForwards() (value int) {
	return m.Forwards
}

// Decode implements bin.Decoder.
func (m *MessageInteractionCounters) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageInteractionCounters#ad4fc9bd to nil")
	}
	if err := b.ConsumeID(MessageInteractionCountersTypeID); err != nil {
		return fmt.Errorf("unable to decode messageInteractionCounters#ad4fc9bd: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageInteractionCounters#ad4fc9bd: field msg_id: %w", err)
		}
		m.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageInteractionCounters#ad4fc9bd: field views: %w", err)
		}
		m.Views = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageInteractionCounters#ad4fc9bd: field forwards: %w", err)
		}
		m.Forwards = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessageInteractionCounters.
var (
	_ bin.Encoder = &MessageInteractionCounters{}
	_ bin.Decoder = &MessageInteractionCounters{}
)
