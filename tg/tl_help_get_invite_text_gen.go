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

// HelpGetInviteTextRequest represents TL type `help.getInviteText#4d392343`.
// Returns localized text of a text message with an invitation.
//
// See https://core.telegram.org/method/help.getInviteText for reference.
type HelpGetInviteTextRequest struct {
}

// HelpGetInviteTextRequestTypeID is TL type id of HelpGetInviteTextRequest.
const HelpGetInviteTextRequestTypeID = 0x4d392343

func (g *HelpGetInviteTextRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetInviteTextRequest) String() string {
	if g == nil {
		return "HelpGetInviteTextRequest(nil)"
	}
	type Alias HelpGetInviteTextRequest
	return fmt.Sprintf("HelpGetInviteTextRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *HelpGetInviteTextRequest) TypeID() uint32 {
	return HelpGetInviteTextRequestTypeID
}

// TypeName returns name of type in TL schema.
func (g *HelpGetInviteTextRequest) TypeName() string {
	return "help.getInviteText"
}

// Encode implements bin.Encoder.
func (g *HelpGetInviteTextRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getInviteText#4d392343 as nil")
	}
	b.PutID(HelpGetInviteTextRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetInviteTextRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getInviteText#4d392343 to nil")
	}
	if err := b.ConsumeID(HelpGetInviteTextRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getInviteText#4d392343: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetInviteTextRequest.
var (
	_ bin.Encoder = &HelpGetInviteTextRequest{}
	_ bin.Decoder = &HelpGetInviteTextRequest{}
)

// HelpGetInviteText invokes method help.getInviteText#4d392343 returning error if any.
// Returns localized text of a text message with an invitation.
//
// See https://core.telegram.org/method/help.getInviteText for reference.
func (c *Client) HelpGetInviteText(ctx context.Context) (*HelpInviteText, error) {
	var result HelpInviteText

	request := &HelpGetInviteTextRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
