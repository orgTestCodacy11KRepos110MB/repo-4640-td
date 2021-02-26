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

// InvokeWithTakeoutRequest represents TL type `invokeWithTakeout#aca9fd2e`.
// Invoke a method within a takeout session
//
// See https://core.telegram.org/constructor/invokeWithTakeout for reference.
type InvokeWithTakeoutRequest struct {
	// Takeout session ID
	TakeoutID int64 `tl:"takeout_id"`
	// Query
	Query bin.Object `tl:"query"`
}

// InvokeWithTakeoutRequestTypeID is TL type id of InvokeWithTakeoutRequest.
const InvokeWithTakeoutRequestTypeID = 0xaca9fd2e

func (i *InvokeWithTakeoutRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.TakeoutID == 0) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeWithTakeoutRequest) String() string {
	if i == nil {
		return "InvokeWithTakeoutRequest(nil)"
	}
	type Alias InvokeWithTakeoutRequest
	return fmt.Sprintf("InvokeWithTakeoutRequest%+v", Alias(*i))
}

// FillFrom fills InvokeWithTakeoutRequest from given interface.
func (i *InvokeWithTakeoutRequest) FillFrom(from interface {
	GetTakeoutID() (value int64)
	GetQuery() (value bin.Object)
}) {
	i.TakeoutID = from.GetTakeoutID()
	i.Query = from.GetQuery()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InvokeWithTakeoutRequest) TypeID() uint32 {
	return InvokeWithTakeoutRequestTypeID
}

// TypeName returns name of type in TL schema.
func (i *InvokeWithTakeoutRequest) TypeName() string {
	return "invokeWithTakeout"
}

// Encode implements bin.Encoder.
func (i *InvokeWithTakeoutRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithTakeout#aca9fd2e as nil")
	}
	b.PutID(InvokeWithTakeoutRequestTypeID)
	b.PutLong(i.TakeoutID)
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeWithTakeout#aca9fd2e: field query: %w", err)
	}
	return nil
}

// GetTakeoutID returns value of TakeoutID field.
func (i *InvokeWithTakeoutRequest) GetTakeoutID() (value int64) {
	return i.TakeoutID
}

// GetQuery returns value of Query field.
func (i *InvokeWithTakeoutRequest) GetQuery() (value bin.Object) {
	return i.Query
}

// Decode implements bin.Decoder.
func (i *InvokeWithTakeoutRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithTakeout#aca9fd2e to nil")
	}
	if err := b.ConsumeID(InvokeWithTakeoutRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeWithTakeout#aca9fd2e: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode invokeWithTakeout#aca9fd2e: field takeout_id: %w", err)
		}
		i.TakeoutID = value
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeWithTakeout#aca9fd2e: field query: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InvokeWithTakeoutRequest.
var (
	_ bin.Encoder = &InvokeWithTakeoutRequest{}
	_ bin.Decoder = &InvokeWithTakeoutRequest{}
)
