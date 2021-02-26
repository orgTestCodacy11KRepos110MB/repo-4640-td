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

// AccountUnregisterDeviceRequest represents TL type `account.unregisterDevice#3076c4bf`.
// Deletes a device by its token, stops sending PUSH-notifications to it.
//
// See https://core.telegram.org/method/account.unregisterDevice for reference.
type AccountUnregisterDeviceRequest struct {
	// Device token type.Possible values:1 - APNS (device token for apple push)2 - FCM (firebase token for google firebase)3 - MPNS (channel URI for microsoft push)4 - Simple push (endpoint for firefox's simple push API)5 - Ubuntu phone (token for ubuntu push)6 - Blackberry (token for blackberry push)7 - Unused8 - WNS (windows push)9 - APNS VoIP (token for apple push VoIP)10 - Web push (web push, see below)11 - MPNS VoIP (token for microsoft push VoIP)12 - Tizen (token for tizen push)For 10 web push, the token must be a JSON-encoded object containing the keys described in PUSH updates¹
	//
	// Links:
	//  1) https://core.telegram.org/api/push-updates
	TokenType int `tl:"token_type"`
	// Device token
	Token string `tl:"token"`
	// List of user identifiers of other users currently using the client
	OtherUids []int `tl:"other_uids"`
}

// AccountUnregisterDeviceRequestTypeID is TL type id of AccountUnregisterDeviceRequest.
const AccountUnregisterDeviceRequestTypeID = 0x3076c4bf

func (u *AccountUnregisterDeviceRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.TokenType == 0) {
		return false
	}
	if !(u.Token == "") {
		return false
	}
	if !(u.OtherUids == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUnregisterDeviceRequest) String() string {
	if u == nil {
		return "AccountUnregisterDeviceRequest(nil)"
	}
	type Alias AccountUnregisterDeviceRequest
	return fmt.Sprintf("AccountUnregisterDeviceRequest%+v", Alias(*u))
}

// FillFrom fills AccountUnregisterDeviceRequest from given interface.
func (u *AccountUnregisterDeviceRequest) FillFrom(from interface {
	GetTokenType() (value int)
	GetToken() (value string)
	GetOtherUids() (value []int)
}) {
	u.TokenType = from.GetTokenType()
	u.Token = from.GetToken()
	u.OtherUids = from.GetOtherUids()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (u *AccountUnregisterDeviceRequest) TypeID() uint32 {
	return AccountUnregisterDeviceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (u *AccountUnregisterDeviceRequest) TypeName() string {
	return "account.unregisterDevice"
}

// Encode implements bin.Encoder.
func (u *AccountUnregisterDeviceRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.unregisterDevice#3076c4bf as nil")
	}
	b.PutID(AccountUnregisterDeviceRequestTypeID)
	b.PutInt(u.TokenType)
	b.PutString(u.Token)
	b.PutVectorHeader(len(u.OtherUids))
	for _, v := range u.OtherUids {
		b.PutInt(v)
	}
	return nil
}

// GetTokenType returns value of TokenType field.
func (u *AccountUnregisterDeviceRequest) GetTokenType() (value int) {
	return u.TokenType
}

// GetToken returns value of Token field.
func (u *AccountUnregisterDeviceRequest) GetToken() (value string) {
	return u.Token
}

// GetOtherUids returns value of OtherUids field.
func (u *AccountUnregisterDeviceRequest) GetOtherUids() (value []int) {
	return u.OtherUids
}

// Decode implements bin.Decoder.
func (u *AccountUnregisterDeviceRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.unregisterDevice#3076c4bf to nil")
	}
	if err := b.ConsumeID(AccountUnregisterDeviceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: field token_type: %w", err)
		}
		u.TokenType = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: field token: %w", err)
		}
		u.Token = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: field other_uids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: field other_uids: %w", err)
			}
			u.OtherUids = append(u.OtherUids, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUnregisterDeviceRequest.
var (
	_ bin.Encoder = &AccountUnregisterDeviceRequest{}
	_ bin.Decoder = &AccountUnregisterDeviceRequest{}
)

// AccountUnregisterDevice invokes method account.unregisterDevice#3076c4bf returning error if any.
// Deletes a device by its token, stops sending PUSH-notifications to it.
//
// Possible errors:
//  400 TOKEN_INVALID: The provided token is invalid
//
// See https://core.telegram.org/method/account.unregisterDevice for reference.
func (c *Client) AccountUnregisterDevice(ctx context.Context, request *AccountUnregisterDeviceRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
