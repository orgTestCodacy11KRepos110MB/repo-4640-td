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

// AccountSaveAutoDownloadSettingsRequest represents TL type `account.saveAutoDownloadSettings#76f36233`.
// Change media autodownload settings
//
// See https://core.telegram.org/method/account.saveAutoDownloadSettings for reference.
type AccountSaveAutoDownloadSettingsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Whether to save settings in the low data usage preset
	Low bool `tl:"low"`
	// Whether to save settings in the high data usage preset
	High bool `tl:"high"`
	// Media autodownload settings
	Settings AutoDownloadSettings `tl:"settings"`
}

// AccountSaveAutoDownloadSettingsRequestTypeID is TL type id of AccountSaveAutoDownloadSettingsRequest.
const AccountSaveAutoDownloadSettingsRequestTypeID = 0x76f36233

func (s *AccountSaveAutoDownloadSettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Low == false) {
		return false
	}
	if !(s.High == false) {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSaveAutoDownloadSettingsRequest) String() string {
	if s == nil {
		return "AccountSaveAutoDownloadSettingsRequest(nil)"
	}
	type Alias AccountSaveAutoDownloadSettingsRequest
	return fmt.Sprintf("AccountSaveAutoDownloadSettingsRequest%+v", Alias(*s))
}

// FillFrom fills AccountSaveAutoDownloadSettingsRequest from given interface.
func (s *AccountSaveAutoDownloadSettingsRequest) FillFrom(from interface {
	GetLow() (value bool)
	GetHigh() (value bool)
	GetSettings() (value AutoDownloadSettings)
}) {
	s.Low = from.GetLow()
	s.High = from.GetHigh()
	s.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AccountSaveAutoDownloadSettingsRequest) TypeID() uint32 {
	return AccountSaveAutoDownloadSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (s *AccountSaveAutoDownloadSettingsRequest) TypeName() string {
	return "account.saveAutoDownloadSettings"
}

// Encode implements bin.Encoder.
func (s *AccountSaveAutoDownloadSettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveAutoDownloadSettings#76f36233 as nil")
	}
	b.PutID(AccountSaveAutoDownloadSettingsRequestTypeID)
	if !(s.Low == false) {
		s.Flags.Set(0)
	}
	if !(s.High == false) {
		s.Flags.Set(1)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveAutoDownloadSettings#76f36233: field flags: %w", err)
	}
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveAutoDownloadSettings#76f36233: field settings: %w", err)
	}
	return nil
}

// SetLow sets value of Low conditional field.
func (s *AccountSaveAutoDownloadSettingsRequest) SetLow(value bool) {
	if value {
		s.Flags.Set(0)
		s.Low = true
	} else {
		s.Flags.Unset(0)
		s.Low = false
	}
}

// GetLow returns value of Low conditional field.
func (s *AccountSaveAutoDownloadSettingsRequest) GetLow() (value bool) {
	return s.Flags.Has(0)
}

// SetHigh sets value of High conditional field.
func (s *AccountSaveAutoDownloadSettingsRequest) SetHigh(value bool) {
	if value {
		s.Flags.Set(1)
		s.High = true
	} else {
		s.Flags.Unset(1)
		s.High = false
	}
}

// GetHigh returns value of High conditional field.
func (s *AccountSaveAutoDownloadSettingsRequest) GetHigh() (value bool) {
	return s.Flags.Has(1)
}

// GetSettings returns value of Settings field.
func (s *AccountSaveAutoDownloadSettingsRequest) GetSettings() (value AutoDownloadSettings) {
	return s.Settings
}

// Decode implements bin.Decoder.
func (s *AccountSaveAutoDownloadSettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveAutoDownloadSettings#76f36233 to nil")
	}
	if err := b.ConsumeID(AccountSaveAutoDownloadSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveAutoDownloadSettings#76f36233: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.saveAutoDownloadSettings#76f36233: field flags: %w", err)
		}
	}
	s.Low = s.Flags.Has(0)
	s.High = s.Flags.Has(1)
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.saveAutoDownloadSettings#76f36233: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSaveAutoDownloadSettingsRequest.
var (
	_ bin.Encoder = &AccountSaveAutoDownloadSettingsRequest{}
	_ bin.Decoder = &AccountSaveAutoDownloadSettingsRequest{}
)

// AccountSaveAutoDownloadSettings invokes method account.saveAutoDownloadSettings#76f36233 returning error if any.
// Change media autodownload settings
//
// See https://core.telegram.org/method/account.saveAutoDownloadSettings for reference.
func (c *Client) AccountSaveAutoDownloadSettings(ctx context.Context, request *AccountSaveAutoDownloadSettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
