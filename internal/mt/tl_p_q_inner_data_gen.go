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

// PQInnerData represents TL type `p_q_inner_data#83c95aec`.
type PQInnerData struct {
	// Pq field of PQInnerData.
	Pq []byte `tl:"pq"`
	// P field of PQInnerData.
	P []byte `tl:"p"`
	// Q field of PQInnerData.
	Q []byte `tl:"q"`
	// Nonce field of PQInnerData.
	Nonce bin.Int128 `tl:"nonce"`
	// ServerNonce field of PQInnerData.
	ServerNonce bin.Int128 `tl:"server_nonce"`
	// NewNonce field of PQInnerData.
	NewNonce bin.Int256 `tl:"new_nonce"`
}

// PQInnerDataTypeID is TL type id of PQInnerData.
const PQInnerDataTypeID = 0x83c95aec

func (p *PQInnerData) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Pq == nil) {
		return false
	}
	if !(p.P == nil) {
		return false
	}
	if !(p.Q == nil) {
		return false
	}
	if !(p.Nonce == bin.Int128{}) {
		return false
	}
	if !(p.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(p.NewNonce == bin.Int256{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PQInnerData) String() string {
	if p == nil {
		return "PQInnerData(nil)"
	}
	type Alias PQInnerData
	return fmt.Sprintf("PQInnerData%+v", Alias(*p))
}

// FillFrom fills PQInnerData from given interface.
func (p *PQInnerData) FillFrom(from interface {
	GetPq() (value []byte)
	GetP() (value []byte)
	GetQ() (value []byte)
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetNewNonce() (value bin.Int256)
}) {
	p.Pq = from.GetPq()
	p.P = from.GetP()
	p.Q = from.GetQ()
	p.Nonce = from.GetNonce()
	p.ServerNonce = from.GetServerNonce()
	p.NewNonce = from.GetNewNonce()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PQInnerData) TypeID() uint32 {
	return PQInnerDataTypeID
}

// TypeName returns name of type in TL schema.
func (p *PQInnerData) TypeName() string {
	return "p_q_inner_data"
}

// Encode implements bin.Encoder.
func (p *PQInnerData) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode p_q_inner_data#83c95aec as nil")
	}
	b.PutID(PQInnerDataTypeID)
	b.PutBytes(p.Pq)
	b.PutBytes(p.P)
	b.PutBytes(p.Q)
	b.PutInt128(p.Nonce)
	b.PutInt128(p.ServerNonce)
	b.PutInt256(p.NewNonce)
	return nil
}

// GetPq returns value of Pq field.
func (p *PQInnerData) GetPq() (value []byte) {
	return p.Pq
}

// GetP returns value of P field.
func (p *PQInnerData) GetP() (value []byte) {
	return p.P
}

// GetQ returns value of Q field.
func (p *PQInnerData) GetQ() (value []byte) {
	return p.Q
}

// GetNonce returns value of Nonce field.
func (p *PQInnerData) GetNonce() (value bin.Int128) {
	return p.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (p *PQInnerData) GetServerNonce() (value bin.Int128) {
	return p.ServerNonce
}

// GetNewNonce returns value of NewNonce field.
func (p *PQInnerData) GetNewNonce() (value bin.Int256) {
	return p.NewNonce
}

// Decode implements bin.Decoder.
func (p *PQInnerData) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode p_q_inner_data#83c95aec to nil")
	}
	if err := b.ConsumeID(PQInnerDataTypeID); err != nil {
		return fmt.Errorf("unable to decode p_q_inner_data#83c95aec: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode p_q_inner_data#83c95aec: field pq: %w", err)
		}
		p.Pq = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode p_q_inner_data#83c95aec: field p: %w", err)
		}
		p.P = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode p_q_inner_data#83c95aec: field q: %w", err)
		}
		p.Q = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode p_q_inner_data#83c95aec: field nonce: %w", err)
		}
		p.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode p_q_inner_data#83c95aec: field server_nonce: %w", err)
		}
		p.ServerNonce = value
	}
	{
		value, err := b.Int256()
		if err != nil {
			return fmt.Errorf("unable to decode p_q_inner_data#83c95aec: field new_nonce: %w", err)
		}
		p.NewNonce = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PQInnerData.
var (
	_ bin.Encoder = &PQInnerData{}
	_ bin.Decoder = &PQInnerData{}
)
