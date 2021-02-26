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

// SecurePasswordKdfAlgoUnknown represents TL type `securePasswordKdfAlgoUnknown#4a8537`.
// Unknown KDF algo (most likely the client has to be updated)
//
// See https://core.telegram.org/constructor/securePasswordKdfAlgoUnknown for reference.
type SecurePasswordKdfAlgoUnknown struct {
}

// SecurePasswordKdfAlgoUnknownTypeID is TL type id of SecurePasswordKdfAlgoUnknown.
const SecurePasswordKdfAlgoUnknownTypeID = 0x4a8537

func (s *SecurePasswordKdfAlgoUnknown) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecurePasswordKdfAlgoUnknown) String() string {
	if s == nil {
		return "SecurePasswordKdfAlgoUnknown(nil)"
	}
	type Alias SecurePasswordKdfAlgoUnknown
	return fmt.Sprintf("SecurePasswordKdfAlgoUnknown%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecurePasswordKdfAlgoUnknown) TypeID() uint32 {
	return SecurePasswordKdfAlgoUnknownTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecurePasswordKdfAlgoUnknown) TypeName() string {
	return "securePasswordKdfAlgoUnknown"
}

// Encode implements bin.Encoder.
func (s *SecurePasswordKdfAlgoUnknown) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode securePasswordKdfAlgoUnknown#4a8537 as nil")
	}
	b.PutID(SecurePasswordKdfAlgoUnknownTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecurePasswordKdfAlgoUnknown) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode securePasswordKdfAlgoUnknown#4a8537 to nil")
	}
	if err := b.ConsumeID(SecurePasswordKdfAlgoUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode securePasswordKdfAlgoUnknown#4a8537: %w", err)
	}
	return nil
}

// construct implements constructor of SecurePasswordKdfAlgoClass.
func (s SecurePasswordKdfAlgoUnknown) construct() SecurePasswordKdfAlgoClass { return &s }

// Ensuring interfaces in compile-time for SecurePasswordKdfAlgoUnknown.
var (
	_ bin.Encoder = &SecurePasswordKdfAlgoUnknown{}
	_ bin.Decoder = &SecurePasswordKdfAlgoUnknown{}

	_ SecurePasswordKdfAlgoClass = &SecurePasswordKdfAlgoUnknown{}
)

// SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000 represents TL type `securePasswordKdfAlgoPBKDF2HMACSHA512iter100000#bbf2dda0`.
// PBKDF2 with SHA512 and 100000 iterations KDF algo
//
// See https://core.telegram.org/constructor/securePasswordKdfAlgoPBKDF2HMACSHA512iter100000 for reference.
type SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000 struct {
	// Salt
	Salt []byte `tl:"salt"`
}

// SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000TypeID is TL type id of SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000.
const SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000TypeID = 0xbbf2dda0

func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Salt == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) String() string {
	if s == nil {
		return "SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000(nil)"
	}
	type Alias SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000
	return fmt.Sprintf("SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000%+v", Alias(*s))
}

// FillFrom fills SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000 from given interface.
func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) FillFrom(from interface {
	GetSalt() (value []byte)
}) {
	s.Salt = from.GetSalt()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) TypeID() uint32 {
	return SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000TypeID
}

// TypeName returns name of type in TL schema.
func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) TypeName() string {
	return "securePasswordKdfAlgoPBKDF2HMACSHA512iter100000"
}

// Encode implements bin.Encoder.
func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode securePasswordKdfAlgoPBKDF2HMACSHA512iter100000#bbf2dda0 as nil")
	}
	b.PutID(SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000TypeID)
	b.PutBytes(s.Salt)
	return nil
}

// GetSalt returns value of Salt field.
func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) GetSalt() (value []byte) {
	return s.Salt
}

// Decode implements bin.Decoder.
func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode securePasswordKdfAlgoPBKDF2HMACSHA512iter100000#bbf2dda0 to nil")
	}
	if err := b.ConsumeID(SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000TypeID); err != nil {
		return fmt.Errorf("unable to decode securePasswordKdfAlgoPBKDF2HMACSHA512iter100000#bbf2dda0: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode securePasswordKdfAlgoPBKDF2HMACSHA512iter100000#bbf2dda0: field salt: %w", err)
		}
		s.Salt = value
	}
	return nil
}

// construct implements constructor of SecurePasswordKdfAlgoClass.
func (s SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) construct() SecurePasswordKdfAlgoClass {
	return &s
}

// Ensuring interfaces in compile-time for SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000.
var (
	_ bin.Encoder = &SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000{}
	_ bin.Decoder = &SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000{}

	_ SecurePasswordKdfAlgoClass = &SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000{}
)

// SecurePasswordKdfAlgoSHA512 represents TL type `securePasswordKdfAlgoSHA512#86471d92`.
// SHA512 KDF algo
//
// See https://core.telegram.org/constructor/securePasswordKdfAlgoSHA512 for reference.
type SecurePasswordKdfAlgoSHA512 struct {
	// Salt
	Salt []byte `tl:"salt"`
}

// SecurePasswordKdfAlgoSHA512TypeID is TL type id of SecurePasswordKdfAlgoSHA512.
const SecurePasswordKdfAlgoSHA512TypeID = 0x86471d92

func (s *SecurePasswordKdfAlgoSHA512) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Salt == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecurePasswordKdfAlgoSHA512) String() string {
	if s == nil {
		return "SecurePasswordKdfAlgoSHA512(nil)"
	}
	type Alias SecurePasswordKdfAlgoSHA512
	return fmt.Sprintf("SecurePasswordKdfAlgoSHA512%+v", Alias(*s))
}

// FillFrom fills SecurePasswordKdfAlgoSHA512 from given interface.
func (s *SecurePasswordKdfAlgoSHA512) FillFrom(from interface {
	GetSalt() (value []byte)
}) {
	s.Salt = from.GetSalt()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecurePasswordKdfAlgoSHA512) TypeID() uint32 {
	return SecurePasswordKdfAlgoSHA512TypeID
}

// TypeName returns name of type in TL schema.
func (s *SecurePasswordKdfAlgoSHA512) TypeName() string {
	return "securePasswordKdfAlgoSHA512"
}

// Encode implements bin.Encoder.
func (s *SecurePasswordKdfAlgoSHA512) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode securePasswordKdfAlgoSHA512#86471d92 as nil")
	}
	b.PutID(SecurePasswordKdfAlgoSHA512TypeID)
	b.PutBytes(s.Salt)
	return nil
}

// GetSalt returns value of Salt field.
func (s *SecurePasswordKdfAlgoSHA512) GetSalt() (value []byte) {
	return s.Salt
}

// Decode implements bin.Decoder.
func (s *SecurePasswordKdfAlgoSHA512) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode securePasswordKdfAlgoSHA512#86471d92 to nil")
	}
	if err := b.ConsumeID(SecurePasswordKdfAlgoSHA512TypeID); err != nil {
		return fmt.Errorf("unable to decode securePasswordKdfAlgoSHA512#86471d92: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode securePasswordKdfAlgoSHA512#86471d92: field salt: %w", err)
		}
		s.Salt = value
	}
	return nil
}

// construct implements constructor of SecurePasswordKdfAlgoClass.
func (s SecurePasswordKdfAlgoSHA512) construct() SecurePasswordKdfAlgoClass { return &s }

// Ensuring interfaces in compile-time for SecurePasswordKdfAlgoSHA512.
var (
	_ bin.Encoder = &SecurePasswordKdfAlgoSHA512{}
	_ bin.Decoder = &SecurePasswordKdfAlgoSHA512{}

	_ SecurePasswordKdfAlgoClass = &SecurePasswordKdfAlgoSHA512{}
)

// SecurePasswordKdfAlgoClass represents SecurePasswordKdfAlgo generic type.
//
// See https://core.telegram.org/type/SecurePasswordKdfAlgo for reference.
//
// Example:
//  g, err := tg.DecodeSecurePasswordKdfAlgo(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.SecurePasswordKdfAlgoUnknown: // securePasswordKdfAlgoUnknown#4a8537
//  case *tg.SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000: // securePasswordKdfAlgoPBKDF2HMACSHA512iter100000#bbf2dda0
//  case *tg.SecurePasswordKdfAlgoSHA512: // securePasswordKdfAlgoSHA512#86471d92
//  default: panic(v)
//  }
type SecurePasswordKdfAlgoClass interface {
	bin.Encoder
	bin.Decoder
	construct() SecurePasswordKdfAlgoClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeSecurePasswordKdfAlgo implements binary de-serialization for SecurePasswordKdfAlgoClass.
func DecodeSecurePasswordKdfAlgo(buf *bin.Buffer) (SecurePasswordKdfAlgoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecurePasswordKdfAlgoUnknownTypeID:
		// Decoding securePasswordKdfAlgoUnknown#4a8537.
		v := SecurePasswordKdfAlgoUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecurePasswordKdfAlgoClass: %w", err)
		}
		return &v, nil
	case SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000TypeID:
		// Decoding securePasswordKdfAlgoPBKDF2HMACSHA512iter100000#bbf2dda0.
		v := SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecurePasswordKdfAlgoClass: %w", err)
		}
		return &v, nil
	case SecurePasswordKdfAlgoSHA512TypeID:
		// Decoding securePasswordKdfAlgoSHA512#86471d92.
		v := SecurePasswordKdfAlgoSHA512{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecurePasswordKdfAlgoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecurePasswordKdfAlgoClass: %w", bin.NewUnexpectedID(id))
	}
}

// SecurePasswordKdfAlgo boxes the SecurePasswordKdfAlgoClass providing a helper.
type SecurePasswordKdfAlgoBox struct {
	SecurePasswordKdfAlgo SecurePasswordKdfAlgoClass
}

// Decode implements bin.Decoder for SecurePasswordKdfAlgoBox.
func (b *SecurePasswordKdfAlgoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SecurePasswordKdfAlgoBox to nil")
	}
	v, err := DecodeSecurePasswordKdfAlgo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SecurePasswordKdfAlgo = v
	return nil
}

// Encode implements bin.Encode for SecurePasswordKdfAlgoBox.
func (b *SecurePasswordKdfAlgoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SecurePasswordKdfAlgo == nil {
		return fmt.Errorf("unable to encode SecurePasswordKdfAlgoClass as nil")
	}
	return b.SecurePasswordKdfAlgo.Encode(buf)
}

// SecurePasswordKdfAlgoClassArray is adapter for slice of SecurePasswordKdfAlgoClass.
type SecurePasswordKdfAlgoClassArray []SecurePasswordKdfAlgoClass

// Sort sorts slice of SecurePasswordKdfAlgoClass.
func (s SecurePasswordKdfAlgoClassArray) Sort(less func(a, b SecurePasswordKdfAlgoClass) bool) SecurePasswordKdfAlgoClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecurePasswordKdfAlgoClass.
func (s SecurePasswordKdfAlgoClassArray) SortStable(less func(a, b SecurePasswordKdfAlgoClass) bool) SecurePasswordKdfAlgoClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecurePasswordKdfAlgoClass.
func (s SecurePasswordKdfAlgoClassArray) Retain(keep func(x SecurePasswordKdfAlgoClass) bool) SecurePasswordKdfAlgoClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecurePasswordKdfAlgoClassArray) First() (v SecurePasswordKdfAlgoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecurePasswordKdfAlgoClassArray) Last() (v SecurePasswordKdfAlgoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecurePasswordKdfAlgoClassArray) PopFirst() (v SecurePasswordKdfAlgoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecurePasswordKdfAlgoClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecurePasswordKdfAlgoClassArray) Pop() (v SecurePasswordKdfAlgoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsSecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000 returns copy with only SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000 constructors.
func (s SecurePasswordKdfAlgoClassArray) AsSecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000() (to SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array) {
	for _, elem := range s {
		value, ok := elem.(*SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecurePasswordKdfAlgoSHA512 returns copy with only SecurePasswordKdfAlgoSHA512 constructors.
func (s SecurePasswordKdfAlgoClassArray) AsSecurePasswordKdfAlgoSHA512() (to SecurePasswordKdfAlgoSHA512Array) {
	for _, elem := range s {
		value, ok := elem.(*SecurePasswordKdfAlgoSHA512)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array is adapter for slice of SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000.
type SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array []SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000

// Sort sorts slice of SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000.
func (s SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array) Sort(less func(a, b SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) bool) SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000.
func (s SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array) SortStable(less func(a, b SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) bool) SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000.
func (s SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array) Retain(keep func(x SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) bool) SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array) First() (v SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array) Last() (v SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array) PopFirst() (v SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000Array) Pop() (v SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecurePasswordKdfAlgoSHA512Array is adapter for slice of SecurePasswordKdfAlgoSHA512.
type SecurePasswordKdfAlgoSHA512Array []SecurePasswordKdfAlgoSHA512

// Sort sorts slice of SecurePasswordKdfAlgoSHA512.
func (s SecurePasswordKdfAlgoSHA512Array) Sort(less func(a, b SecurePasswordKdfAlgoSHA512) bool) SecurePasswordKdfAlgoSHA512Array {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecurePasswordKdfAlgoSHA512.
func (s SecurePasswordKdfAlgoSHA512Array) SortStable(less func(a, b SecurePasswordKdfAlgoSHA512) bool) SecurePasswordKdfAlgoSHA512Array {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecurePasswordKdfAlgoSHA512.
func (s SecurePasswordKdfAlgoSHA512Array) Retain(keep func(x SecurePasswordKdfAlgoSHA512) bool) SecurePasswordKdfAlgoSHA512Array {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecurePasswordKdfAlgoSHA512Array) First() (v SecurePasswordKdfAlgoSHA512, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecurePasswordKdfAlgoSHA512Array) Last() (v SecurePasswordKdfAlgoSHA512, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecurePasswordKdfAlgoSHA512Array) PopFirst() (v SecurePasswordKdfAlgoSHA512, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecurePasswordKdfAlgoSHA512
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecurePasswordKdfAlgoSHA512Array) Pop() (v SecurePasswordKdfAlgoSHA512, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
