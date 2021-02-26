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

// LangPackString represents TL type `langPackString#cad181f6`.
// Translated localization string
//
// See https://core.telegram.org/constructor/langPackString for reference.
type LangPackString struct {
	// Language key
	Key string `tl:"key"`
	// Value
	Value string `tl:"value"`
}

// LangPackStringTypeID is TL type id of LangPackString.
const LangPackStringTypeID = 0xcad181f6

func (l *LangPackString) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Key == "") {
		return false
	}
	if !(l.Value == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LangPackString) String() string {
	if l == nil {
		return "LangPackString(nil)"
	}
	type Alias LangPackString
	return fmt.Sprintf("LangPackString%+v", Alias(*l))
}

// FillFrom fills LangPackString from given interface.
func (l *LangPackString) FillFrom(from interface {
	GetKey() (value string)
	GetValue() (value string)
}) {
	l.Key = from.GetKey()
	l.Value = from.GetValue()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (l *LangPackString) TypeID() uint32 {
	return LangPackStringTypeID
}

// TypeName returns name of type in TL schema.
func (l *LangPackString) TypeName() string {
	return "langPackString"
}

// Encode implements bin.Encoder.
func (l *LangPackString) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode langPackString#cad181f6 as nil")
	}
	b.PutID(LangPackStringTypeID)
	b.PutString(l.Key)
	b.PutString(l.Value)
	return nil
}

// GetKey returns value of Key field.
func (l *LangPackString) GetKey() (value string) {
	return l.Key
}

// GetValue returns value of Value field.
func (l *LangPackString) GetValue() (value string) {
	return l.Value
}

// Decode implements bin.Decoder.
func (l *LangPackString) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode langPackString#cad181f6 to nil")
	}
	if err := b.ConsumeID(LangPackStringTypeID); err != nil {
		return fmt.Errorf("unable to decode langPackString#cad181f6: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackString#cad181f6: field key: %w", err)
		}
		l.Key = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackString#cad181f6: field value: %w", err)
		}
		l.Value = value
	}
	return nil
}

// construct implements constructor of LangPackStringClass.
func (l LangPackString) construct() LangPackStringClass { return &l }

// Ensuring interfaces in compile-time for LangPackString.
var (
	_ bin.Encoder = &LangPackString{}
	_ bin.Decoder = &LangPackString{}

	_ LangPackStringClass = &LangPackString{}
)

// LangPackStringPluralized represents TL type `langPackStringPluralized#6c47ac9f`.
// A language pack string which has different forms based on the number of some object it mentions. See https://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html¹ for more info
//
// Links:
//  1) https://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html
//
// See https://core.telegram.org/constructor/langPackStringPluralized for reference.
type LangPackStringPluralized struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Localization key
	Key string `tl:"key"`
	// Value for zero objects
	//
	// Use SetZeroValue and GetZeroValue helpers.
	ZeroValue string `tl:"zero_value"`
	// Value for one object
	//
	// Use SetOneValue and GetOneValue helpers.
	OneValue string `tl:"one_value"`
	// Value for two objects
	//
	// Use SetTwoValue and GetTwoValue helpers.
	TwoValue string `tl:"two_value"`
	// Value for a few objects
	//
	// Use SetFewValue and GetFewValue helpers.
	FewValue string `tl:"few_value"`
	// Value for many objects
	//
	// Use SetManyValue and GetManyValue helpers.
	ManyValue string `tl:"many_value"`
	// Default value
	OtherValue string `tl:"other_value"`
}

// LangPackStringPluralizedTypeID is TL type id of LangPackStringPluralized.
const LangPackStringPluralizedTypeID = 0x6c47ac9f

func (l *LangPackStringPluralized) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Flags.Zero()) {
		return false
	}
	if !(l.Key == "") {
		return false
	}
	if !(l.ZeroValue == "") {
		return false
	}
	if !(l.OneValue == "") {
		return false
	}
	if !(l.TwoValue == "") {
		return false
	}
	if !(l.FewValue == "") {
		return false
	}
	if !(l.ManyValue == "") {
		return false
	}
	if !(l.OtherValue == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LangPackStringPluralized) String() string {
	if l == nil {
		return "LangPackStringPluralized(nil)"
	}
	type Alias LangPackStringPluralized
	return fmt.Sprintf("LangPackStringPluralized%+v", Alias(*l))
}

// FillFrom fills LangPackStringPluralized from given interface.
func (l *LangPackStringPluralized) FillFrom(from interface {
	GetKey() (value string)
	GetZeroValue() (value string, ok bool)
	GetOneValue() (value string, ok bool)
	GetTwoValue() (value string, ok bool)
	GetFewValue() (value string, ok bool)
	GetManyValue() (value string, ok bool)
	GetOtherValue() (value string)
}) {
	l.Key = from.GetKey()
	if val, ok := from.GetZeroValue(); ok {
		l.ZeroValue = val
	}

	if val, ok := from.GetOneValue(); ok {
		l.OneValue = val
	}

	if val, ok := from.GetTwoValue(); ok {
		l.TwoValue = val
	}

	if val, ok := from.GetFewValue(); ok {
		l.FewValue = val
	}

	if val, ok := from.GetManyValue(); ok {
		l.ManyValue = val
	}

	l.OtherValue = from.GetOtherValue()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (l *LangPackStringPluralized) TypeID() uint32 {
	return LangPackStringPluralizedTypeID
}

// TypeName returns name of type in TL schema.
func (l *LangPackStringPluralized) TypeName() string {
	return "langPackStringPluralized"
}

// Encode implements bin.Encoder.
func (l *LangPackStringPluralized) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode langPackStringPluralized#6c47ac9f as nil")
	}
	b.PutID(LangPackStringPluralizedTypeID)
	if !(l.ZeroValue == "") {
		l.Flags.Set(0)
	}
	if !(l.OneValue == "") {
		l.Flags.Set(1)
	}
	if !(l.TwoValue == "") {
		l.Flags.Set(2)
	}
	if !(l.FewValue == "") {
		l.Flags.Set(3)
	}
	if !(l.ManyValue == "") {
		l.Flags.Set(4)
	}
	if err := l.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode langPackStringPluralized#6c47ac9f: field flags: %w", err)
	}
	b.PutString(l.Key)
	if l.Flags.Has(0) {
		b.PutString(l.ZeroValue)
	}
	if l.Flags.Has(1) {
		b.PutString(l.OneValue)
	}
	if l.Flags.Has(2) {
		b.PutString(l.TwoValue)
	}
	if l.Flags.Has(3) {
		b.PutString(l.FewValue)
	}
	if l.Flags.Has(4) {
		b.PutString(l.ManyValue)
	}
	b.PutString(l.OtherValue)
	return nil
}

// GetKey returns value of Key field.
func (l *LangPackStringPluralized) GetKey() (value string) {
	return l.Key
}

// SetZeroValue sets value of ZeroValue conditional field.
func (l *LangPackStringPluralized) SetZeroValue(value string) {
	l.Flags.Set(0)
	l.ZeroValue = value
}

// GetZeroValue returns value of ZeroValue conditional field and
// boolean which is true if field was set.
func (l *LangPackStringPluralized) GetZeroValue() (value string, ok bool) {
	if !l.Flags.Has(0) {
		return value, false
	}
	return l.ZeroValue, true
}

// SetOneValue sets value of OneValue conditional field.
func (l *LangPackStringPluralized) SetOneValue(value string) {
	l.Flags.Set(1)
	l.OneValue = value
}

// GetOneValue returns value of OneValue conditional field and
// boolean which is true if field was set.
func (l *LangPackStringPluralized) GetOneValue() (value string, ok bool) {
	if !l.Flags.Has(1) {
		return value, false
	}
	return l.OneValue, true
}

// SetTwoValue sets value of TwoValue conditional field.
func (l *LangPackStringPluralized) SetTwoValue(value string) {
	l.Flags.Set(2)
	l.TwoValue = value
}

// GetTwoValue returns value of TwoValue conditional field and
// boolean which is true if field was set.
func (l *LangPackStringPluralized) GetTwoValue() (value string, ok bool) {
	if !l.Flags.Has(2) {
		return value, false
	}
	return l.TwoValue, true
}

// SetFewValue sets value of FewValue conditional field.
func (l *LangPackStringPluralized) SetFewValue(value string) {
	l.Flags.Set(3)
	l.FewValue = value
}

// GetFewValue returns value of FewValue conditional field and
// boolean which is true if field was set.
func (l *LangPackStringPluralized) GetFewValue() (value string, ok bool) {
	if !l.Flags.Has(3) {
		return value, false
	}
	return l.FewValue, true
}

// SetManyValue sets value of ManyValue conditional field.
func (l *LangPackStringPluralized) SetManyValue(value string) {
	l.Flags.Set(4)
	l.ManyValue = value
}

// GetManyValue returns value of ManyValue conditional field and
// boolean which is true if field was set.
func (l *LangPackStringPluralized) GetManyValue() (value string, ok bool) {
	if !l.Flags.Has(4) {
		return value, false
	}
	return l.ManyValue, true
}

// GetOtherValue returns value of OtherValue field.
func (l *LangPackStringPluralized) GetOtherValue() (value string) {
	return l.OtherValue
}

// Decode implements bin.Decoder.
func (l *LangPackStringPluralized) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode langPackStringPluralized#6c47ac9f to nil")
	}
	if err := b.ConsumeID(LangPackStringPluralizedTypeID); err != nil {
		return fmt.Errorf("unable to decode langPackStringPluralized#6c47ac9f: %w", err)
	}
	{
		if err := l.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode langPackStringPluralized#6c47ac9f: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackStringPluralized#6c47ac9f: field key: %w", err)
		}
		l.Key = value
	}
	if l.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackStringPluralized#6c47ac9f: field zero_value: %w", err)
		}
		l.ZeroValue = value
	}
	if l.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackStringPluralized#6c47ac9f: field one_value: %w", err)
		}
		l.OneValue = value
	}
	if l.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackStringPluralized#6c47ac9f: field two_value: %w", err)
		}
		l.TwoValue = value
	}
	if l.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackStringPluralized#6c47ac9f: field few_value: %w", err)
		}
		l.FewValue = value
	}
	if l.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackStringPluralized#6c47ac9f: field many_value: %w", err)
		}
		l.ManyValue = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackStringPluralized#6c47ac9f: field other_value: %w", err)
		}
		l.OtherValue = value
	}
	return nil
}

// construct implements constructor of LangPackStringClass.
func (l LangPackStringPluralized) construct() LangPackStringClass { return &l }

// Ensuring interfaces in compile-time for LangPackStringPluralized.
var (
	_ bin.Encoder = &LangPackStringPluralized{}
	_ bin.Decoder = &LangPackStringPluralized{}

	_ LangPackStringClass = &LangPackStringPluralized{}
)

// LangPackStringDeleted represents TL type `langPackStringDeleted#2979eeb2`.
// Deleted localization string
//
// See https://core.telegram.org/constructor/langPackStringDeleted for reference.
type LangPackStringDeleted struct {
	// Localization key
	Key string `tl:"key"`
}

// LangPackStringDeletedTypeID is TL type id of LangPackStringDeleted.
const LangPackStringDeletedTypeID = 0x2979eeb2

func (l *LangPackStringDeleted) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Key == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LangPackStringDeleted) String() string {
	if l == nil {
		return "LangPackStringDeleted(nil)"
	}
	type Alias LangPackStringDeleted
	return fmt.Sprintf("LangPackStringDeleted%+v", Alias(*l))
}

// FillFrom fills LangPackStringDeleted from given interface.
func (l *LangPackStringDeleted) FillFrom(from interface {
	GetKey() (value string)
}) {
	l.Key = from.GetKey()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (l *LangPackStringDeleted) TypeID() uint32 {
	return LangPackStringDeletedTypeID
}

// TypeName returns name of type in TL schema.
func (l *LangPackStringDeleted) TypeName() string {
	return "langPackStringDeleted"
}

// Encode implements bin.Encoder.
func (l *LangPackStringDeleted) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode langPackStringDeleted#2979eeb2 as nil")
	}
	b.PutID(LangPackStringDeletedTypeID)
	b.PutString(l.Key)
	return nil
}

// GetKey returns value of Key field.
func (l *LangPackStringDeleted) GetKey() (value string) {
	return l.Key
}

// Decode implements bin.Decoder.
func (l *LangPackStringDeleted) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode langPackStringDeleted#2979eeb2 to nil")
	}
	if err := b.ConsumeID(LangPackStringDeletedTypeID); err != nil {
		return fmt.Errorf("unable to decode langPackStringDeleted#2979eeb2: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackStringDeleted#2979eeb2: field key: %w", err)
		}
		l.Key = value
	}
	return nil
}

// construct implements constructor of LangPackStringClass.
func (l LangPackStringDeleted) construct() LangPackStringClass { return &l }

// Ensuring interfaces in compile-time for LangPackStringDeleted.
var (
	_ bin.Encoder = &LangPackStringDeleted{}
	_ bin.Decoder = &LangPackStringDeleted{}

	_ LangPackStringClass = &LangPackStringDeleted{}
)

// LangPackStringClass represents LangPackString generic type.
//
// See https://core.telegram.org/type/LangPackString for reference.
//
// Example:
//  g, err := tg.DecodeLangPackString(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.LangPackString: // langPackString#cad181f6
//  case *tg.LangPackStringPluralized: // langPackStringPluralized#6c47ac9f
//  case *tg.LangPackStringDeleted: // langPackStringDeleted#2979eeb2
//  default: panic(v)
//  }
type LangPackStringClass interface {
	bin.Encoder
	bin.Decoder
	construct() LangPackStringClass

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

	// Language key
	GetKey() (value string)
}

// DecodeLangPackString implements binary de-serialization for LangPackStringClass.
func DecodeLangPackString(buf *bin.Buffer) (LangPackStringClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case LangPackStringTypeID:
		// Decoding langPackString#cad181f6.
		v := LangPackString{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LangPackStringClass: %w", err)
		}
		return &v, nil
	case LangPackStringPluralizedTypeID:
		// Decoding langPackStringPluralized#6c47ac9f.
		v := LangPackStringPluralized{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LangPackStringClass: %w", err)
		}
		return &v, nil
	case LangPackStringDeletedTypeID:
		// Decoding langPackStringDeleted#2979eeb2.
		v := LangPackStringDeleted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LangPackStringClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode LangPackStringClass: %w", bin.NewUnexpectedID(id))
	}
}

// LangPackString boxes the LangPackStringClass providing a helper.
type LangPackStringBox struct {
	LangPackString LangPackStringClass
}

// Decode implements bin.Decoder for LangPackStringBox.
func (b *LangPackStringBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode LangPackStringBox to nil")
	}
	v, err := DecodeLangPackString(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.LangPackString = v
	return nil
}

// Encode implements bin.Encode for LangPackStringBox.
func (b *LangPackStringBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.LangPackString == nil {
		return fmt.Errorf("unable to encode LangPackStringClass as nil")
	}
	return b.LangPackString.Encode(buf)
}

// LangPackStringClassArray is adapter for slice of LangPackStringClass.
type LangPackStringClassArray []LangPackStringClass

// Sort sorts slice of LangPackStringClass.
func (s LangPackStringClassArray) Sort(less func(a, b LangPackStringClass) bool) LangPackStringClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of LangPackStringClass.
func (s LangPackStringClassArray) SortStable(less func(a, b LangPackStringClass) bool) LangPackStringClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of LangPackStringClass.
func (s LangPackStringClassArray) Retain(keep func(x LangPackStringClass) bool) LangPackStringClassArray {
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
func (s LangPackStringClassArray) First() (v LangPackStringClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s LangPackStringClassArray) Last() (v LangPackStringClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *LangPackStringClassArray) PopFirst() (v LangPackStringClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero LangPackStringClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *LangPackStringClassArray) Pop() (v LangPackStringClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsLangPackString returns copy with only LangPackString constructors.
func (s LangPackStringClassArray) AsLangPackString() (to LangPackStringArray) {
	for _, elem := range s {
		value, ok := elem.(*LangPackString)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsLangPackStringPluralized returns copy with only LangPackStringPluralized constructors.
func (s LangPackStringClassArray) AsLangPackStringPluralized() (to LangPackStringPluralizedArray) {
	for _, elem := range s {
		value, ok := elem.(*LangPackStringPluralized)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsLangPackStringDeleted returns copy with only LangPackStringDeleted constructors.
func (s LangPackStringClassArray) AsLangPackStringDeleted() (to LangPackStringDeletedArray) {
	for _, elem := range s {
		value, ok := elem.(*LangPackStringDeleted)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// LangPackStringArray is adapter for slice of LangPackString.
type LangPackStringArray []LangPackString

// Sort sorts slice of LangPackString.
func (s LangPackStringArray) Sort(less func(a, b LangPackString) bool) LangPackStringArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of LangPackString.
func (s LangPackStringArray) SortStable(less func(a, b LangPackString) bool) LangPackStringArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of LangPackString.
func (s LangPackStringArray) Retain(keep func(x LangPackString) bool) LangPackStringArray {
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
func (s LangPackStringArray) First() (v LangPackString, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s LangPackStringArray) Last() (v LangPackString, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *LangPackStringArray) PopFirst() (v LangPackString, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero LangPackString
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *LangPackStringArray) Pop() (v LangPackString, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// LangPackStringPluralizedArray is adapter for slice of LangPackStringPluralized.
type LangPackStringPluralizedArray []LangPackStringPluralized

// Sort sorts slice of LangPackStringPluralized.
func (s LangPackStringPluralizedArray) Sort(less func(a, b LangPackStringPluralized) bool) LangPackStringPluralizedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of LangPackStringPluralized.
func (s LangPackStringPluralizedArray) SortStable(less func(a, b LangPackStringPluralized) bool) LangPackStringPluralizedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of LangPackStringPluralized.
func (s LangPackStringPluralizedArray) Retain(keep func(x LangPackStringPluralized) bool) LangPackStringPluralizedArray {
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
func (s LangPackStringPluralizedArray) First() (v LangPackStringPluralized, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s LangPackStringPluralizedArray) Last() (v LangPackStringPluralized, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *LangPackStringPluralizedArray) PopFirst() (v LangPackStringPluralized, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero LangPackStringPluralized
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *LangPackStringPluralizedArray) Pop() (v LangPackStringPluralized, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// LangPackStringDeletedArray is adapter for slice of LangPackStringDeleted.
type LangPackStringDeletedArray []LangPackStringDeleted

// Sort sorts slice of LangPackStringDeleted.
func (s LangPackStringDeletedArray) Sort(less func(a, b LangPackStringDeleted) bool) LangPackStringDeletedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of LangPackStringDeleted.
func (s LangPackStringDeletedArray) SortStable(less func(a, b LangPackStringDeleted) bool) LangPackStringDeletedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of LangPackStringDeleted.
func (s LangPackStringDeletedArray) Retain(keep func(x LangPackStringDeleted) bool) LangPackStringDeletedArray {
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
func (s LangPackStringDeletedArray) First() (v LangPackStringDeleted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s LangPackStringDeletedArray) Last() (v LangPackStringDeleted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *LangPackStringDeletedArray) PopFirst() (v LangPackStringDeleted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero LangPackStringDeleted
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *LangPackStringDeletedArray) Pop() (v LangPackStringDeleted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
