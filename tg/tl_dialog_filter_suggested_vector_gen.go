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

// DialogFilterSuggestedVector is a box for Vector<DialogFilterSuggested>
type DialogFilterSuggestedVector struct {
	// Elements of Vector<DialogFilterSuggested>
	Elems []DialogFilterSuggested `tl:"Elems"`
}

// DialogFilterSuggestedVectorTypeID is TL type id of DialogFilterSuggestedVector.
const DialogFilterSuggestedVectorTypeID = bin.TypeVector

func (vec *DialogFilterSuggestedVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *DialogFilterSuggestedVector) String() string {
	if vec == nil {
		return "DialogFilterSuggestedVector(nil)"
	}
	type Alias DialogFilterSuggestedVector
	return fmt.Sprintf("DialogFilterSuggestedVector%+v", Alias(*vec))
}

// FillFrom fills DialogFilterSuggestedVector from given interface.
func (vec *DialogFilterSuggestedVector) FillFrom(from interface {
	GetElems() (value []DialogFilterSuggested)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *DialogFilterSuggestedVector) TypeID() uint32 {
	return DialogFilterSuggestedVectorTypeID
}

// TypeName returns name of type in TL schema.
func (vec *DialogFilterSuggestedVector) TypeName() string {
	return ""
}

// Encode implements bin.Encoder.
func (vec *DialogFilterSuggestedVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<DialogFilterSuggested> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<DialogFilterSuggested>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *DialogFilterSuggestedVector) GetElems() (value []DialogFilterSuggested) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *DialogFilterSuggestedVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<DialogFilterSuggested> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<DialogFilterSuggested>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value DialogFilterSuggested
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<DialogFilterSuggested>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for DialogFilterSuggestedVector.
var (
	_ bin.Encoder = &DialogFilterSuggestedVector{}
	_ bin.Decoder = &DialogFilterSuggestedVector{}
)
