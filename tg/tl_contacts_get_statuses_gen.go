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

// ContactsGetStatusesRequest represents TL type `contacts.getStatuses#c4a353ee`.
// Returns the list of contact statuses.
//
// See https://core.telegram.org/method/contacts.getStatuses for reference.
type ContactsGetStatusesRequest struct {
}

// ContactsGetStatusesRequestTypeID is TL type id of ContactsGetStatusesRequest.
const ContactsGetStatusesRequestTypeID = 0xc4a353ee

func (g *ContactsGetStatusesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetStatusesRequest) String() string {
	if g == nil {
		return "ContactsGetStatusesRequest(nil)"
	}
	type Alias ContactsGetStatusesRequest
	return fmt.Sprintf("ContactsGetStatusesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *ContactsGetStatusesRequest) TypeID() uint32 {
	return ContactsGetStatusesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (g *ContactsGetStatusesRequest) TypeName() string {
	return "contacts.getStatuses"
}

// Encode implements bin.Encoder.
func (g *ContactsGetStatusesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getStatuses#c4a353ee as nil")
	}
	b.PutID(ContactsGetStatusesRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetStatusesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getStatuses#c4a353ee to nil")
	}
	if err := b.ConsumeID(ContactsGetStatusesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getStatuses#c4a353ee: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetStatusesRequest.
var (
	_ bin.Encoder = &ContactsGetStatusesRequest{}
	_ bin.Decoder = &ContactsGetStatusesRequest{}
)

// ContactsGetStatuses invokes method contacts.getStatuses#c4a353ee returning error if any.
// Returns the list of contact statuses.
//
// See https://core.telegram.org/method/contacts.getStatuses for reference.
func (c *Client) ContactsGetStatuses(ctx context.Context) ([]ContactStatus, error) {
	var result ContactStatusVector

	request := &ContactsGetStatusesRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return []ContactStatus(result.Elems), nil
}
