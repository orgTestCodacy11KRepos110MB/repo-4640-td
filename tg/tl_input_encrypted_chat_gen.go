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

// InputEncryptedChat represents TL type `inputEncryptedChat#f141b5e1`.
// Creates an encrypted chat.
//
// See https://core.telegram.org/constructor/inputEncryptedChat for reference.
type InputEncryptedChat struct {
	// Chat ID
	ChatID int `tl:"chat_id"`
	// Checking sum from constructor encryptedChat¹, encryptedChatWaiting² or encryptedChatRequested³
	//
	// Links:
	//  1) https://core.telegram.org/constructor/encryptedChat
	//  2) https://core.telegram.org/constructor/encryptedChatWaiting
	//  3) https://core.telegram.org/constructor/encryptedChatRequested
	AccessHash int64 `tl:"access_hash"`
}

// InputEncryptedChatTypeID is TL type id of InputEncryptedChat.
const InputEncryptedChatTypeID = 0xf141b5e1

func (i *InputEncryptedChat) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChatID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputEncryptedChat) String() string {
	if i == nil {
		return "InputEncryptedChat(nil)"
	}
	type Alias InputEncryptedChat
	return fmt.Sprintf("InputEncryptedChat%+v", Alias(*i))
}

// FillFrom fills InputEncryptedChat from given interface.
func (i *InputEncryptedChat) FillFrom(from interface {
	GetChatID() (value int)
	GetAccessHash() (value int64)
}) {
	i.ChatID = from.GetChatID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputEncryptedChat) TypeID() uint32 {
	return InputEncryptedChatTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputEncryptedChat) TypeName() string {
	return "inputEncryptedChat"
}

// Encode implements bin.Encoder.
func (i *InputEncryptedChat) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedChat#f141b5e1 as nil")
	}
	b.PutID(InputEncryptedChatTypeID)
	b.PutInt(i.ChatID)
	b.PutLong(i.AccessHash)
	return nil
}

// GetChatID returns value of ChatID field.
func (i *InputEncryptedChat) GetChatID() (value int) {
	return i.ChatID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputEncryptedChat) GetAccessHash() (value int64) {
	return i.AccessHash
}

// Decode implements bin.Decoder.
func (i *InputEncryptedChat) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedChat#f141b5e1 to nil")
	}
	if err := b.ConsumeID(InputEncryptedChatTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedChat#f141b5e1: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedChat#f141b5e1: field chat_id: %w", err)
		}
		i.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedChat#f141b5e1: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputEncryptedChat.
var (
	_ bin.Encoder = &InputEncryptedChat{}
	_ bin.Decoder = &InputEncryptedChat{}
)
