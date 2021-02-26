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

// UserClassVector is a box for Vector<User>
type UserClassVector struct {
	// Elements of Vector<User>
	Elems []UserClass `tl:"Elems"`
}

// UserClassVectorTypeID is TL type id of UserClassVector.
const UserClassVectorTypeID = bin.TypeVector

func (vec *UserClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *UserClassVector) String() string {
	if vec == nil {
		return "UserClassVector(nil)"
	}
	type Alias UserClassVector
	return fmt.Sprintf("UserClassVector%+v", Alias(*vec))
}

// FillFrom fills UserClassVector from given interface.
func (vec *UserClassVector) FillFrom(from interface {
	GetElems() (value []UserClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *UserClassVector) TypeID() uint32 {
	return UserClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (vec *UserClassVector) TypeName() string {
	return ""
}

// Encode implements bin.Encoder.
func (vec *UserClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<User> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<User>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<User>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *UserClassVector) GetElems() (value []UserClass) {
	return vec.Elems
}

// MapElems returns field Elems wrapped in UserClassArray helper.
func (vec *UserClassVector) MapElems() (value UserClassArray) {
	return UserClassArray(vec.Elems)
}

// Decode implements bin.Decoder.
func (vec *UserClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<User> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<User>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<User>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for UserClassVector.
var (
	_ bin.Encoder = &UserClassVector{}
	_ bin.Decoder = &UserClassVector{}
)
