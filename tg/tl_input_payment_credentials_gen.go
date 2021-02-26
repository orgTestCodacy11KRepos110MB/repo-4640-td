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

// InputPaymentCredentialsSaved represents TL type `inputPaymentCredentialsSaved#c10eb2cf`.
// Saved payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentialsSaved for reference.
type InputPaymentCredentialsSaved struct {
	// Credential ID
	ID string `tl:"id"`
	// Temporary password
	TmpPassword []byte `tl:"tmp_password"`
}

// InputPaymentCredentialsSavedTypeID is TL type id of InputPaymentCredentialsSaved.
const InputPaymentCredentialsSavedTypeID = 0xc10eb2cf

func (i *InputPaymentCredentialsSaved) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == "") {
		return false
	}
	if !(i.TmpPassword == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentialsSaved) String() string {
	if i == nil {
		return "InputPaymentCredentialsSaved(nil)"
	}
	type Alias InputPaymentCredentialsSaved
	return fmt.Sprintf("InputPaymentCredentialsSaved%+v", Alias(*i))
}

// FillFrom fills InputPaymentCredentialsSaved from given interface.
func (i *InputPaymentCredentialsSaved) FillFrom(from interface {
	GetID() (value string)
	GetTmpPassword() (value []byte)
}) {
	i.ID = from.GetID()
	i.TmpPassword = from.GetTmpPassword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPaymentCredentialsSaved) TypeID() uint32 {
	return InputPaymentCredentialsSavedTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputPaymentCredentialsSaved) TypeName() string {
	return "inputPaymentCredentialsSaved"
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsSaved) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsSaved#c10eb2cf as nil")
	}
	b.PutID(InputPaymentCredentialsSavedTypeID)
	b.PutString(i.ID)
	b.PutBytes(i.TmpPassword)
	return nil
}

// GetID returns value of ID field.
func (i *InputPaymentCredentialsSaved) GetID() (value string) {
	return i.ID
}

// GetTmpPassword returns value of TmpPassword field.
func (i *InputPaymentCredentialsSaved) GetTmpPassword() (value []byte) {
	return i.TmpPassword
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsSaved) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsSaved#c10eb2cf to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsSavedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: field tmp_password: %w", err)
		}
		i.TmpPassword = value
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsSaved) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsSaved.
var (
	_ bin.Encoder = &InputPaymentCredentialsSaved{}
	_ bin.Decoder = &InputPaymentCredentialsSaved{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsSaved{}
)

// InputPaymentCredentials represents TL type `inputPaymentCredentials#3417d728`.
// Payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentials for reference.
type InputPaymentCredentials struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Save payment credential for future use
	Save bool `tl:"save"`
	// Payment credentials
	Data DataJSON `tl:"data"`
}

// InputPaymentCredentialsTypeID is TL type id of InputPaymentCredentials.
const InputPaymentCredentialsTypeID = 0x3417d728

func (i *InputPaymentCredentials) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Save == false) {
		return false
	}
	if !(i.Data.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentials) String() string {
	if i == nil {
		return "InputPaymentCredentials(nil)"
	}
	type Alias InputPaymentCredentials
	return fmt.Sprintf("InputPaymentCredentials%+v", Alias(*i))
}

// FillFrom fills InputPaymentCredentials from given interface.
func (i *InputPaymentCredentials) FillFrom(from interface {
	GetSave() (value bool)
	GetData() (value DataJSON)
}) {
	i.Save = from.GetSave()
	i.Data = from.GetData()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPaymentCredentials) TypeID() uint32 {
	return InputPaymentCredentialsTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputPaymentCredentials) TypeName() string {
	return "inputPaymentCredentials"
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentials) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentials#3417d728 as nil")
	}
	b.PutID(InputPaymentCredentialsTypeID)
	if !(i.Save == false) {
		i.Flags.Set(0)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentials#3417d728: field flags: %w", err)
	}
	if err := i.Data.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentials#3417d728: field data: %w", err)
	}
	return nil
}

// SetSave sets value of Save conditional field.
func (i *InputPaymentCredentials) SetSave(value bool) {
	if value {
		i.Flags.Set(0)
		i.Save = true
	} else {
		i.Flags.Unset(0)
		i.Save = false
	}
}

// GetSave returns value of Save conditional field.
func (i *InputPaymentCredentials) GetSave() (value bool) {
	return i.Flags.Has(0)
}

// GetData returns value of Data field.
func (i *InputPaymentCredentials) GetData() (value DataJSON) {
	return i.Data
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentials) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentials#3417d728 to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: field flags: %w", err)
		}
	}
	i.Save = i.Flags.Has(0)
	{
		if err := i.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: field data: %w", err)
		}
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentials) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentials.
var (
	_ bin.Encoder = &InputPaymentCredentials{}
	_ bin.Decoder = &InputPaymentCredentials{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentials{}
)

// InputPaymentCredentialsApplePay represents TL type `inputPaymentCredentialsApplePay#aa1c39f`.
// Apple pay payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentialsApplePay for reference.
type InputPaymentCredentialsApplePay struct {
	// Payment data
	PaymentData DataJSON `tl:"payment_data"`
}

// InputPaymentCredentialsApplePayTypeID is TL type id of InputPaymentCredentialsApplePay.
const InputPaymentCredentialsApplePayTypeID = 0xaa1c39f

func (i *InputPaymentCredentialsApplePay) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.PaymentData.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentialsApplePay) String() string {
	if i == nil {
		return "InputPaymentCredentialsApplePay(nil)"
	}
	type Alias InputPaymentCredentialsApplePay
	return fmt.Sprintf("InputPaymentCredentialsApplePay%+v", Alias(*i))
}

// FillFrom fills InputPaymentCredentialsApplePay from given interface.
func (i *InputPaymentCredentialsApplePay) FillFrom(from interface {
	GetPaymentData() (value DataJSON)
}) {
	i.PaymentData = from.GetPaymentData()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPaymentCredentialsApplePay) TypeID() uint32 {
	return InputPaymentCredentialsApplePayTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputPaymentCredentialsApplePay) TypeName() string {
	return "inputPaymentCredentialsApplePay"
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsApplePay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsApplePay#aa1c39f as nil")
	}
	b.PutID(InputPaymentCredentialsApplePayTypeID)
	if err := i.PaymentData.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentialsApplePay#aa1c39f: field payment_data: %w", err)
	}
	return nil
}

// GetPaymentData returns value of PaymentData field.
func (i *InputPaymentCredentialsApplePay) GetPaymentData() (value DataJSON) {
	return i.PaymentData
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsApplePay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsApplePay#aa1c39f to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsApplePayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsApplePay#aa1c39f: %w", err)
	}
	{
		if err := i.PaymentData.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsApplePay#aa1c39f: field payment_data: %w", err)
		}
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsApplePay) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsApplePay.
var (
	_ bin.Encoder = &InputPaymentCredentialsApplePay{}
	_ bin.Decoder = &InputPaymentCredentialsApplePay{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsApplePay{}
)

// InputPaymentCredentialsGooglePay represents TL type `inputPaymentCredentialsGooglePay#8ac32801`.
//
// See https://core.telegram.org/constructor/inputPaymentCredentialsGooglePay for reference.
type InputPaymentCredentialsGooglePay struct {
	// PaymentToken field of InputPaymentCredentialsGooglePay.
	PaymentToken DataJSON `tl:"payment_token"`
}

// InputPaymentCredentialsGooglePayTypeID is TL type id of InputPaymentCredentialsGooglePay.
const InputPaymentCredentialsGooglePayTypeID = 0x8ac32801

func (i *InputPaymentCredentialsGooglePay) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.PaymentToken.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentialsGooglePay) String() string {
	if i == nil {
		return "InputPaymentCredentialsGooglePay(nil)"
	}
	type Alias InputPaymentCredentialsGooglePay
	return fmt.Sprintf("InputPaymentCredentialsGooglePay%+v", Alias(*i))
}

// FillFrom fills InputPaymentCredentialsGooglePay from given interface.
func (i *InputPaymentCredentialsGooglePay) FillFrom(from interface {
	GetPaymentToken() (value DataJSON)
}) {
	i.PaymentToken = from.GetPaymentToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPaymentCredentialsGooglePay) TypeID() uint32 {
	return InputPaymentCredentialsGooglePayTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputPaymentCredentialsGooglePay) TypeName() string {
	return "inputPaymentCredentialsGooglePay"
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsGooglePay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsGooglePay#8ac32801 as nil")
	}
	b.PutID(InputPaymentCredentialsGooglePayTypeID)
	if err := i.PaymentToken.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentialsGooglePay#8ac32801: field payment_token: %w", err)
	}
	return nil
}

// GetPaymentToken returns value of PaymentToken field.
func (i *InputPaymentCredentialsGooglePay) GetPaymentToken() (value DataJSON) {
	return i.PaymentToken
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsGooglePay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsGooglePay#8ac32801 to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsGooglePayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsGooglePay#8ac32801: %w", err)
	}
	{
		if err := i.PaymentToken.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsGooglePay#8ac32801: field payment_token: %w", err)
		}
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsGooglePay) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsGooglePay.
var (
	_ bin.Encoder = &InputPaymentCredentialsGooglePay{}
	_ bin.Decoder = &InputPaymentCredentialsGooglePay{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsGooglePay{}
)

// InputPaymentCredentialsClass represents InputPaymentCredentials generic type.
//
// See https://core.telegram.org/type/InputPaymentCredentials for reference.
//
// Example:
//  g, err := tg.DecodeInputPaymentCredentials(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputPaymentCredentialsSaved: // inputPaymentCredentialsSaved#c10eb2cf
//  case *tg.InputPaymentCredentials: // inputPaymentCredentials#3417d728
//  case *tg.InputPaymentCredentialsApplePay: // inputPaymentCredentialsApplePay#aa1c39f
//  case *tg.InputPaymentCredentialsGooglePay: // inputPaymentCredentialsGooglePay#8ac32801
//  default: panic(v)
//  }
type InputPaymentCredentialsClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputPaymentCredentialsClass

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

// DecodeInputPaymentCredentials implements binary de-serialization for InputPaymentCredentialsClass.
func DecodeInputPaymentCredentials(buf *bin.Buffer) (InputPaymentCredentialsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPaymentCredentialsSavedTypeID:
		// Decoding inputPaymentCredentialsSaved#c10eb2cf.
		v := InputPaymentCredentialsSaved{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsTypeID:
		// Decoding inputPaymentCredentials#3417d728.
		v := InputPaymentCredentials{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsApplePayTypeID:
		// Decoding inputPaymentCredentialsApplePay#aa1c39f.
		v := InputPaymentCredentialsApplePay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsGooglePayTypeID:
		// Decoding inputPaymentCredentialsGooglePay#8ac32801.
		v := InputPaymentCredentialsGooglePay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPaymentCredentials boxes the InputPaymentCredentialsClass providing a helper.
type InputPaymentCredentialsBox struct {
	InputPaymentCredentials InputPaymentCredentialsClass
}

// Decode implements bin.Decoder for InputPaymentCredentialsBox.
func (b *InputPaymentCredentialsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPaymentCredentialsBox to nil")
	}
	v, err := DecodeInputPaymentCredentials(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPaymentCredentials = v
	return nil
}

// Encode implements bin.Encode for InputPaymentCredentialsBox.
func (b *InputPaymentCredentialsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPaymentCredentials == nil {
		return fmt.Errorf("unable to encode InputPaymentCredentialsClass as nil")
	}
	return b.InputPaymentCredentials.Encode(buf)
}

// InputPaymentCredentialsClassArray is adapter for slice of InputPaymentCredentialsClass.
type InputPaymentCredentialsClassArray []InputPaymentCredentialsClass

// Sort sorts slice of InputPaymentCredentialsClass.
func (s InputPaymentCredentialsClassArray) Sort(less func(a, b InputPaymentCredentialsClass) bool) InputPaymentCredentialsClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPaymentCredentialsClass.
func (s InputPaymentCredentialsClassArray) SortStable(less func(a, b InputPaymentCredentialsClass) bool) InputPaymentCredentialsClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPaymentCredentialsClass.
func (s InputPaymentCredentialsClassArray) Retain(keep func(x InputPaymentCredentialsClass) bool) InputPaymentCredentialsClassArray {
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
func (s InputPaymentCredentialsClassArray) First() (v InputPaymentCredentialsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPaymentCredentialsClassArray) Last() (v InputPaymentCredentialsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsClassArray) PopFirst() (v InputPaymentCredentialsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPaymentCredentialsClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsClassArray) Pop() (v InputPaymentCredentialsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputPaymentCredentialsSaved returns copy with only InputPaymentCredentialsSaved constructors.
func (s InputPaymentCredentialsClassArray) AsInputPaymentCredentialsSaved() (to InputPaymentCredentialsSavedArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPaymentCredentialsSaved)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPaymentCredentials returns copy with only InputPaymentCredentials constructors.
func (s InputPaymentCredentialsClassArray) AsInputPaymentCredentials() (to InputPaymentCredentialsArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPaymentCredentials)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPaymentCredentialsApplePay returns copy with only InputPaymentCredentialsApplePay constructors.
func (s InputPaymentCredentialsClassArray) AsInputPaymentCredentialsApplePay() (to InputPaymentCredentialsApplePayArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPaymentCredentialsApplePay)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputPaymentCredentialsGooglePay returns copy with only InputPaymentCredentialsGooglePay constructors.
func (s InputPaymentCredentialsClassArray) AsInputPaymentCredentialsGooglePay() (to InputPaymentCredentialsGooglePayArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPaymentCredentialsGooglePay)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputPaymentCredentialsSavedArray is adapter for slice of InputPaymentCredentialsSaved.
type InputPaymentCredentialsSavedArray []InputPaymentCredentialsSaved

// Sort sorts slice of InputPaymentCredentialsSaved.
func (s InputPaymentCredentialsSavedArray) Sort(less func(a, b InputPaymentCredentialsSaved) bool) InputPaymentCredentialsSavedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPaymentCredentialsSaved.
func (s InputPaymentCredentialsSavedArray) SortStable(less func(a, b InputPaymentCredentialsSaved) bool) InputPaymentCredentialsSavedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPaymentCredentialsSaved.
func (s InputPaymentCredentialsSavedArray) Retain(keep func(x InputPaymentCredentialsSaved) bool) InputPaymentCredentialsSavedArray {
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
func (s InputPaymentCredentialsSavedArray) First() (v InputPaymentCredentialsSaved, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPaymentCredentialsSavedArray) Last() (v InputPaymentCredentialsSaved, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsSavedArray) PopFirst() (v InputPaymentCredentialsSaved, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPaymentCredentialsSaved
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsSavedArray) Pop() (v InputPaymentCredentialsSaved, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputPaymentCredentialsArray is adapter for slice of InputPaymentCredentials.
type InputPaymentCredentialsArray []InputPaymentCredentials

// Sort sorts slice of InputPaymentCredentials.
func (s InputPaymentCredentialsArray) Sort(less func(a, b InputPaymentCredentials) bool) InputPaymentCredentialsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPaymentCredentials.
func (s InputPaymentCredentialsArray) SortStable(less func(a, b InputPaymentCredentials) bool) InputPaymentCredentialsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPaymentCredentials.
func (s InputPaymentCredentialsArray) Retain(keep func(x InputPaymentCredentials) bool) InputPaymentCredentialsArray {
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
func (s InputPaymentCredentialsArray) First() (v InputPaymentCredentials, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPaymentCredentialsArray) Last() (v InputPaymentCredentials, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsArray) PopFirst() (v InputPaymentCredentials, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPaymentCredentials
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsArray) Pop() (v InputPaymentCredentials, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputPaymentCredentialsApplePayArray is adapter for slice of InputPaymentCredentialsApplePay.
type InputPaymentCredentialsApplePayArray []InputPaymentCredentialsApplePay

// Sort sorts slice of InputPaymentCredentialsApplePay.
func (s InputPaymentCredentialsApplePayArray) Sort(less func(a, b InputPaymentCredentialsApplePay) bool) InputPaymentCredentialsApplePayArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPaymentCredentialsApplePay.
func (s InputPaymentCredentialsApplePayArray) SortStable(less func(a, b InputPaymentCredentialsApplePay) bool) InputPaymentCredentialsApplePayArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPaymentCredentialsApplePay.
func (s InputPaymentCredentialsApplePayArray) Retain(keep func(x InputPaymentCredentialsApplePay) bool) InputPaymentCredentialsApplePayArray {
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
func (s InputPaymentCredentialsApplePayArray) First() (v InputPaymentCredentialsApplePay, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPaymentCredentialsApplePayArray) Last() (v InputPaymentCredentialsApplePay, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsApplePayArray) PopFirst() (v InputPaymentCredentialsApplePay, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPaymentCredentialsApplePay
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsApplePayArray) Pop() (v InputPaymentCredentialsApplePay, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputPaymentCredentialsGooglePayArray is adapter for slice of InputPaymentCredentialsGooglePay.
type InputPaymentCredentialsGooglePayArray []InputPaymentCredentialsGooglePay

// Sort sorts slice of InputPaymentCredentialsGooglePay.
func (s InputPaymentCredentialsGooglePayArray) Sort(less func(a, b InputPaymentCredentialsGooglePay) bool) InputPaymentCredentialsGooglePayArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPaymentCredentialsGooglePay.
func (s InputPaymentCredentialsGooglePayArray) SortStable(less func(a, b InputPaymentCredentialsGooglePay) bool) InputPaymentCredentialsGooglePayArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPaymentCredentialsGooglePay.
func (s InputPaymentCredentialsGooglePayArray) Retain(keep func(x InputPaymentCredentialsGooglePay) bool) InputPaymentCredentialsGooglePayArray {
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
func (s InputPaymentCredentialsGooglePayArray) First() (v InputPaymentCredentialsGooglePay, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPaymentCredentialsGooglePayArray) Last() (v InputPaymentCredentialsGooglePay, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsGooglePayArray) PopFirst() (v InputPaymentCredentialsGooglePay, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPaymentCredentialsGooglePay
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPaymentCredentialsGooglePayArray) Pop() (v InputPaymentCredentialsGooglePay, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
