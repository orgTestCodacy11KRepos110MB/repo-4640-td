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

// AccountInstallWallPaperRequest represents TL type `account.installWallPaper#feed5769`.
// Install wallpaper
//
// See https://core.telegram.org/method/account.installWallPaper for reference.
type AccountInstallWallPaperRequest struct {
	// Wallpaper to install
	Wallpaper InputWallPaperClass `tl:"wallpaper"`
	// Wallpaper settings
	Settings WallPaperSettings `tl:"settings"`
}

// AccountInstallWallPaperRequestTypeID is TL type id of AccountInstallWallPaperRequest.
const AccountInstallWallPaperRequestTypeID = 0xfeed5769

func (i *AccountInstallWallPaperRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Wallpaper == nil) {
		return false
	}
	if !(i.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AccountInstallWallPaperRequest) String() string {
	if i == nil {
		return "AccountInstallWallPaperRequest(nil)"
	}
	type Alias AccountInstallWallPaperRequest
	return fmt.Sprintf("AccountInstallWallPaperRequest%+v", Alias(*i))
}

// FillFrom fills AccountInstallWallPaperRequest from given interface.
func (i *AccountInstallWallPaperRequest) FillFrom(from interface {
	GetWallpaper() (value InputWallPaperClass)
	GetSettings() (value WallPaperSettings)
}) {
	i.Wallpaper = from.GetWallpaper()
	i.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *AccountInstallWallPaperRequest) TypeID() uint32 {
	return AccountInstallWallPaperRequestTypeID
}

// TypeName returns name of type in TL schema.
func (i *AccountInstallWallPaperRequest) TypeName() string {
	return "account.installWallPaper"
}

// Encode implements bin.Encoder.
func (i *AccountInstallWallPaperRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.installWallPaper#feed5769 as nil")
	}
	b.PutID(AccountInstallWallPaperRequestTypeID)
	if i.Wallpaper == nil {
		return fmt.Errorf("unable to encode account.installWallPaper#feed5769: field wallpaper is nil")
	}
	if err := i.Wallpaper.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.installWallPaper#feed5769: field wallpaper: %w", err)
	}
	if err := i.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.installWallPaper#feed5769: field settings: %w", err)
	}
	return nil
}

// GetWallpaper returns value of Wallpaper field.
func (i *AccountInstallWallPaperRequest) GetWallpaper() (value InputWallPaperClass) {
	return i.Wallpaper
}

// GetSettings returns value of Settings field.
func (i *AccountInstallWallPaperRequest) GetSettings() (value WallPaperSettings) {
	return i.Settings
}

// Decode implements bin.Decoder.
func (i *AccountInstallWallPaperRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.installWallPaper#feed5769 to nil")
	}
	if err := b.ConsumeID(AccountInstallWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.installWallPaper#feed5769: %w", err)
	}
	{
		value, err := DecodeInputWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.installWallPaper#feed5769: field wallpaper: %w", err)
		}
		i.Wallpaper = value
	}
	{
		if err := i.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.installWallPaper#feed5769: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountInstallWallPaperRequest.
var (
	_ bin.Encoder = &AccountInstallWallPaperRequest{}
	_ bin.Decoder = &AccountInstallWallPaperRequest{}
)

// AccountInstallWallPaper invokes method account.installWallPaper#feed5769 returning error if any.
// Install wallpaper
//
// See https://core.telegram.org/method/account.installWallPaper for reference.
func (c *Client) AccountInstallWallPaper(ctx context.Context, request *AccountInstallWallPaperRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
