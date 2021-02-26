// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// ClientDHInnerData represents TL type `client_DH_inner_data#6643b654`.
type ClientDHInnerData struct {
	// Nonce field of ClientDHInnerData.
	Nonce bin.Int128 `tl:"nonce"`
	// ServerNonce field of ClientDHInnerData.
	ServerNonce bin.Int128 `tl:"server_nonce"`
	// RetryID field of ClientDHInnerData.
	RetryID int64 `tl:"retry_id"`
	// GB field of ClientDHInnerData.
	GB []byte `tl:"g_b"`
}

// ClientDHInnerDataTypeID is TL type id of ClientDHInnerData.
const ClientDHInnerDataTypeID = 0x6643b654

func (c *ClientDHInnerData) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Nonce == bin.Int128{}) {
		return false
	}
	if !(c.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(c.RetryID == 0) {
		return false
	}
	if !(c.GB == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ClientDHInnerData) String() string {
	if c == nil {
		return "ClientDHInnerData(nil)"
	}
	type Alias ClientDHInnerData
	return fmt.Sprintf("ClientDHInnerData%+v", Alias(*c))
}

// FillFrom fills ClientDHInnerData from given interface.
func (c *ClientDHInnerData) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetRetryID() (value int64)
	GetGB() (value []byte)
}) {
	c.Nonce = from.GetNonce()
	c.ServerNonce = from.GetServerNonce()
	c.RetryID = from.GetRetryID()
	c.GB = from.GetGB()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ClientDHInnerData) TypeID() uint32 {
	return ClientDHInnerDataTypeID
}

// TypeName returns name of type in TL schema.
func (c *ClientDHInnerData) TypeName() string {
	return "client_DH_inner_data"
}

// Encode implements bin.Encoder.
func (c *ClientDHInnerData) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode client_DH_inner_data#6643b654 as nil")
	}
	b.PutID(ClientDHInnerDataTypeID)
	b.PutInt128(c.Nonce)
	b.PutInt128(c.ServerNonce)
	b.PutLong(c.RetryID)
	b.PutBytes(c.GB)
	return nil
}

// GetNonce returns value of Nonce field.
func (c *ClientDHInnerData) GetNonce() (value bin.Int128) {
	return c.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (c *ClientDHInnerData) GetServerNonce() (value bin.Int128) {
	return c.ServerNonce
}

// GetRetryID returns value of RetryID field.
func (c *ClientDHInnerData) GetRetryID() (value int64) {
	return c.RetryID
}

// GetGB returns value of GB field.
func (c *ClientDHInnerData) GetGB() (value []byte) {
	return c.GB
}

// Decode implements bin.Decoder.
func (c *ClientDHInnerData) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode client_DH_inner_data#6643b654 to nil")
	}
	if err := b.ConsumeID(ClientDHInnerDataTypeID); err != nil {
		return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field nonce: %w", err)
		}
		c.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field server_nonce: %w", err)
		}
		c.ServerNonce = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field retry_id: %w", err)
		}
		c.RetryID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field g_b: %w", err)
		}
		c.GB = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ClientDHInnerData.
var (
	_ bin.Encoder = &ClientDHInnerData{}
	_ bin.Decoder = &ClientDHInnerData{}
)
