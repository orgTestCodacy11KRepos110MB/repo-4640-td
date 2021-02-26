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

// MessagesDeletePhoneCallHistoryRequest represents TL type `messages.deletePhoneCallHistory#f9cbe409`.
//
// See https://core.telegram.org/method/messages.deletePhoneCallHistory for reference.
type MessagesDeletePhoneCallHistoryRequest struct {
	// Flags field of MessagesDeletePhoneCallHistoryRequest.
	Flags bin.Fields `tl:"flags"`
	// Revoke field of MessagesDeletePhoneCallHistoryRequest.
	Revoke bool `tl:"revoke"`
}

// MessagesDeletePhoneCallHistoryRequestTypeID is TL type id of MessagesDeletePhoneCallHistoryRequest.
const MessagesDeletePhoneCallHistoryRequestTypeID = 0xf9cbe409

func (d *MessagesDeletePhoneCallHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Revoke == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeletePhoneCallHistoryRequest) String() string {
	if d == nil {
		return "MessagesDeletePhoneCallHistoryRequest(nil)"
	}
	type Alias MessagesDeletePhoneCallHistoryRequest
	return fmt.Sprintf("MessagesDeletePhoneCallHistoryRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeletePhoneCallHistoryRequest from given interface.
func (d *MessagesDeletePhoneCallHistoryRequest) FillFrom(from interface {
	GetRevoke() (value bool)
}) {
	d.Revoke = from.GetRevoke()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *MessagesDeletePhoneCallHistoryRequest) TypeID() uint32 {
	return MessagesDeletePhoneCallHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (d *MessagesDeletePhoneCallHistoryRequest) TypeName() string {
	return "messages.deletePhoneCallHistory"
}

// Encode implements bin.Encoder.
func (d *MessagesDeletePhoneCallHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deletePhoneCallHistory#f9cbe409 as nil")
	}
	b.PutID(MessagesDeletePhoneCallHistoryRequestTypeID)
	if !(d.Revoke == false) {
		d.Flags.Set(0)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deletePhoneCallHistory#f9cbe409: field flags: %w", err)
	}
	return nil
}

// SetRevoke sets value of Revoke conditional field.
func (d *MessagesDeletePhoneCallHistoryRequest) SetRevoke(value bool) {
	if value {
		d.Flags.Set(0)
		d.Revoke = true
	} else {
		d.Flags.Unset(0)
		d.Revoke = false
	}
}

// GetRevoke returns value of Revoke conditional field.
func (d *MessagesDeletePhoneCallHistoryRequest) GetRevoke() (value bool) {
	return d.Flags.Has(0)
}

// Decode implements bin.Decoder.
func (d *MessagesDeletePhoneCallHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deletePhoneCallHistory#f9cbe409 to nil")
	}
	if err := b.ConsumeID(MessagesDeletePhoneCallHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deletePhoneCallHistory#f9cbe409: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.deletePhoneCallHistory#f9cbe409: field flags: %w", err)
		}
	}
	d.Revoke = d.Flags.Has(0)
	return nil
}

// Ensuring interfaces in compile-time for MessagesDeletePhoneCallHistoryRequest.
var (
	_ bin.Encoder = &MessagesDeletePhoneCallHistoryRequest{}
	_ bin.Decoder = &MessagesDeletePhoneCallHistoryRequest{}
)

// MessagesDeletePhoneCallHistory invokes method messages.deletePhoneCallHistory#f9cbe409 returning error if any.
//
// See https://core.telegram.org/method/messages.deletePhoneCallHistory for reference.
func (c *Client) MessagesDeletePhoneCallHistory(ctx context.Context, request *MessagesDeletePhoneCallHistoryRequest) (*MessagesAffectedFoundMessages, error) {
	var result MessagesAffectedFoundMessages

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
