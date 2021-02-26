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

// ServerDHInnerData represents TL type `server_DH_inner_data#b5890dba`.
type ServerDHInnerData struct {
	// Nonce field of ServerDHInnerData.
	Nonce bin.Int128 `tl:"nonce"`
	// ServerNonce field of ServerDHInnerData.
	ServerNonce bin.Int128 `tl:"server_nonce"`
	// G field of ServerDHInnerData.
	G int `tl:"g"`
	// DhPrime field of ServerDHInnerData.
	DhPrime []byte `tl:"dh_prime"`
	// GA field of ServerDHInnerData.
	GA []byte `tl:"g_a"`
	// ServerTime field of ServerDHInnerData.
	ServerTime int `tl:"server_time"`
}

// ServerDHInnerDataTypeID is TL type id of ServerDHInnerData.
const ServerDHInnerDataTypeID = 0xb5890dba

func (s *ServerDHInnerData) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Nonce == bin.Int128{}) {
		return false
	}
	if !(s.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(s.G == 0) {
		return false
	}
	if !(s.DhPrime == nil) {
		return false
	}
	if !(s.GA == nil) {
		return false
	}
	if !(s.ServerTime == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ServerDHInnerData) String() string {
	if s == nil {
		return "ServerDHInnerData(nil)"
	}
	type Alias ServerDHInnerData
	return fmt.Sprintf("ServerDHInnerData%+v", Alias(*s))
}

// FillFrom fills ServerDHInnerData from given interface.
func (s *ServerDHInnerData) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetG() (value int)
	GetDhPrime() (value []byte)
	GetGA() (value []byte)
	GetServerTime() (value int)
}) {
	s.Nonce = from.GetNonce()
	s.ServerNonce = from.GetServerNonce()
	s.G = from.GetG()
	s.DhPrime = from.GetDhPrime()
	s.GA = from.GetGA()
	s.ServerTime = from.GetServerTime()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *ServerDHInnerData) TypeID() uint32 {
	return ServerDHInnerDataTypeID
}

// TypeName returns name of type in TL schema.
func (s *ServerDHInnerData) TypeName() string {
	return "server_DH_inner_data"
}

// Encode implements bin.Encoder.
func (s *ServerDHInnerData) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode server_DH_inner_data#b5890dba as nil")
	}
	b.PutID(ServerDHInnerDataTypeID)
	b.PutInt128(s.Nonce)
	b.PutInt128(s.ServerNonce)
	b.PutInt(s.G)
	b.PutBytes(s.DhPrime)
	b.PutBytes(s.GA)
	b.PutInt(s.ServerTime)
	return nil
}

// GetNonce returns value of Nonce field.
func (s *ServerDHInnerData) GetNonce() (value bin.Int128) {
	return s.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (s *ServerDHInnerData) GetServerNonce() (value bin.Int128) {
	return s.ServerNonce
}

// GetG returns value of G field.
func (s *ServerDHInnerData) GetG() (value int) {
	return s.G
}

// GetDhPrime returns value of DhPrime field.
func (s *ServerDHInnerData) GetDhPrime() (value []byte) {
	return s.DhPrime
}

// GetGA returns value of GA field.
func (s *ServerDHInnerData) GetGA() (value []byte) {
	return s.GA
}

// GetServerTime returns value of ServerTime field.
func (s *ServerDHInnerData) GetServerTime() (value int) {
	return s.ServerTime
}

// Decode implements bin.Decoder.
func (s *ServerDHInnerData) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode server_DH_inner_data#b5890dba to nil")
	}
	if err := b.ConsumeID(ServerDHInnerDataTypeID); err != nil {
		return fmt.Errorf("unable to decode server_DH_inner_data#b5890dba: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_inner_data#b5890dba: field nonce: %w", err)
		}
		s.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_inner_data#b5890dba: field server_nonce: %w", err)
		}
		s.ServerNonce = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_inner_data#b5890dba: field g: %w", err)
		}
		s.G = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_inner_data#b5890dba: field dh_prime: %w", err)
		}
		s.DhPrime = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_inner_data#b5890dba: field g_a: %w", err)
		}
		s.GA = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_inner_data#b5890dba: field server_time: %w", err)
		}
		s.ServerTime = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ServerDHInnerData.
var (
	_ bin.Encoder = &ServerDHInnerData{}
	_ bin.Decoder = &ServerDHInnerData{}
)
