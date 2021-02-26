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

// AccountGetContentSettingsRequest represents TL type `account.getContentSettings#8b9b4dae`.
// Get sensitive content settings
//
// See https://core.telegram.org/method/account.getContentSettings for reference.
type AccountGetContentSettingsRequest struct {
}

// AccountGetContentSettingsRequestTypeID is TL type id of AccountGetContentSettingsRequest.
const AccountGetContentSettingsRequestTypeID = 0x8b9b4dae

func (g *AccountGetContentSettingsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetContentSettingsRequest) String() string {
	if g == nil {
		return "AccountGetContentSettingsRequest(nil)"
	}
	type Alias AccountGetContentSettingsRequest
	return fmt.Sprintf("AccountGetContentSettingsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *AccountGetContentSettingsRequest) TypeID() uint32 {
	return AccountGetContentSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (g *AccountGetContentSettingsRequest) TypeName() string {
	return "account.getContentSettings"
}

// Encode implements bin.Encoder.
func (g *AccountGetContentSettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getContentSettings#8b9b4dae as nil")
	}
	b.PutID(AccountGetContentSettingsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetContentSettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getContentSettings#8b9b4dae to nil")
	}
	if err := b.ConsumeID(AccountGetContentSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getContentSettings#8b9b4dae: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetContentSettingsRequest.
var (
	_ bin.Encoder = &AccountGetContentSettingsRequest{}
	_ bin.Decoder = &AccountGetContentSettingsRequest{}
)

// AccountGetContentSettings invokes method account.getContentSettings#8b9b4dae returning error if any.
// Get sensitive content settings
//
// See https://core.telegram.org/method/account.getContentSettings for reference.
func (c *Client) AccountGetContentSettings(ctx context.Context) (*AccountContentSettings, error) {
	var result AccountContentSettings

	request := &AccountGetContentSettingsRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
