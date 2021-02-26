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

// WebAuthorization represents TL type `webAuthorization#cac943f2`.
// Represents a bot logged in using the Telegram login widget¹
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/constructor/webAuthorization for reference.
type WebAuthorization struct {
	// Authorization hash
	Hash int64 `tl:"hash"`
	// Bot ID
	BotID int `tl:"bot_id"`
	// The domain name of the website on which the user has logged in.
	Domain string `tl:"domain"`
	// Browser user-agent
	Browser string `tl:"browser"`
	// Platform
	Platform string `tl:"platform"`
	// When was the web session created
	DateCreated int `tl:"date_created"`
	// When was the web session last active
	DateActive int `tl:"date_active"`
	// IP address
	IP string `tl:"ip"`
	// Region, determined from IP address
	Region string `tl:"region"`
}

// WebAuthorizationTypeID is TL type id of WebAuthorization.
const WebAuthorizationTypeID = 0xcac943f2

func (w *WebAuthorization) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Hash == 0) {
		return false
	}
	if !(w.BotID == 0) {
		return false
	}
	if !(w.Domain == "") {
		return false
	}
	if !(w.Browser == "") {
		return false
	}
	if !(w.Platform == "") {
		return false
	}
	if !(w.DateCreated == 0) {
		return false
	}
	if !(w.DateActive == 0) {
		return false
	}
	if !(w.IP == "") {
		return false
	}
	if !(w.Region == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebAuthorization) String() string {
	if w == nil {
		return "WebAuthorization(nil)"
	}
	type Alias WebAuthorization
	return fmt.Sprintf("WebAuthorization%+v", Alias(*w))
}

// FillFrom fills WebAuthorization from given interface.
func (w *WebAuthorization) FillFrom(from interface {
	GetHash() (value int64)
	GetBotID() (value int)
	GetDomain() (value string)
	GetBrowser() (value string)
	GetPlatform() (value string)
	GetDateCreated() (value int)
	GetDateActive() (value int)
	GetIP() (value string)
	GetRegion() (value string)
}) {
	w.Hash = from.GetHash()
	w.BotID = from.GetBotID()
	w.Domain = from.GetDomain()
	w.Browser = from.GetBrowser()
	w.Platform = from.GetPlatform()
	w.DateCreated = from.GetDateCreated()
	w.DateActive = from.GetDateActive()
	w.IP = from.GetIP()
	w.Region = from.GetRegion()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (w *WebAuthorization) TypeID() uint32 {
	return WebAuthorizationTypeID
}

// TypeName returns name of type in TL schema.
func (w *WebAuthorization) TypeName() string {
	return "webAuthorization"
}

// Encode implements bin.Encoder.
func (w *WebAuthorization) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webAuthorization#cac943f2 as nil")
	}
	b.PutID(WebAuthorizationTypeID)
	b.PutLong(w.Hash)
	b.PutInt(w.BotID)
	b.PutString(w.Domain)
	b.PutString(w.Browser)
	b.PutString(w.Platform)
	b.PutInt(w.DateCreated)
	b.PutInt(w.DateActive)
	b.PutString(w.IP)
	b.PutString(w.Region)
	return nil
}

// GetHash returns value of Hash field.
func (w *WebAuthorization) GetHash() (value int64) {
	return w.Hash
}

// GetBotID returns value of BotID field.
func (w *WebAuthorization) GetBotID() (value int) {
	return w.BotID
}

// GetDomain returns value of Domain field.
func (w *WebAuthorization) GetDomain() (value string) {
	return w.Domain
}

// GetBrowser returns value of Browser field.
func (w *WebAuthorization) GetBrowser() (value string) {
	return w.Browser
}

// GetPlatform returns value of Platform field.
func (w *WebAuthorization) GetPlatform() (value string) {
	return w.Platform
}

// GetDateCreated returns value of DateCreated field.
func (w *WebAuthorization) GetDateCreated() (value int) {
	return w.DateCreated
}

// GetDateActive returns value of DateActive field.
func (w *WebAuthorization) GetDateActive() (value int) {
	return w.DateActive
}

// GetIP returns value of IP field.
func (w *WebAuthorization) GetIP() (value string) {
	return w.IP
}

// GetRegion returns value of Region field.
func (w *WebAuthorization) GetRegion() (value string) {
	return w.Region
}

// Decode implements bin.Decoder.
func (w *WebAuthorization) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webAuthorization#cac943f2 to nil")
	}
	if err := b.ConsumeID(WebAuthorizationTypeID); err != nil {
		return fmt.Errorf("unable to decode webAuthorization#cac943f2: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field hash: %w", err)
		}
		w.Hash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field bot_id: %w", err)
		}
		w.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field domain: %w", err)
		}
		w.Domain = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field browser: %w", err)
		}
		w.Browser = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field platform: %w", err)
		}
		w.Platform = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field date_created: %w", err)
		}
		w.DateCreated = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field date_active: %w", err)
		}
		w.DateActive = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field ip: %w", err)
		}
		w.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webAuthorization#cac943f2: field region: %w", err)
		}
		w.Region = value
	}
	return nil
}

// Ensuring interfaces in compile-time for WebAuthorization.
var (
	_ bin.Encoder = &WebAuthorization{}
	_ bin.Decoder = &WebAuthorization{}
)
