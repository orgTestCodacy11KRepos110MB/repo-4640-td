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

// LangPackDifference represents TL type `langPackDifference#f385c1f6`.
// Changes to the app's localization pack
//
// See https://core.telegram.org/constructor/langPackDifference for reference.
type LangPackDifference struct {
	// Language code
	LangCode string `tl:"lang_code"`
	// Previous version number
	FromVersion int `tl:"from_version"`
	// New version number
	Version int `tl:"version"`
	// Localized strings
	Strings []LangPackStringClass `tl:"strings"`
}

// LangPackDifferenceTypeID is TL type id of LangPackDifference.
const LangPackDifferenceTypeID = 0xf385c1f6

func (l *LangPackDifference) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.LangCode == "") {
		return false
	}
	if !(l.FromVersion == 0) {
		return false
	}
	if !(l.Version == 0) {
		return false
	}
	if !(l.Strings == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LangPackDifference) String() string {
	if l == nil {
		return "LangPackDifference(nil)"
	}
	type Alias LangPackDifference
	return fmt.Sprintf("LangPackDifference%+v", Alias(*l))
}

// FillFrom fills LangPackDifference from given interface.
func (l *LangPackDifference) FillFrom(from interface {
	GetLangCode() (value string)
	GetFromVersion() (value int)
	GetVersion() (value int)
	GetStrings() (value []LangPackStringClass)
}) {
	l.LangCode = from.GetLangCode()
	l.FromVersion = from.GetFromVersion()
	l.Version = from.GetVersion()
	l.Strings = from.GetStrings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (l *LangPackDifference) TypeID() uint32 {
	return LangPackDifferenceTypeID
}

// TypeName returns name of type in TL schema.
func (l *LangPackDifference) TypeName() string {
	return "langPackDifference"
}

// Encode implements bin.Encoder.
func (l *LangPackDifference) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode langPackDifference#f385c1f6 as nil")
	}
	b.PutID(LangPackDifferenceTypeID)
	b.PutString(l.LangCode)
	b.PutInt(l.FromVersion)
	b.PutInt(l.Version)
	b.PutVectorHeader(len(l.Strings))
	for idx, v := range l.Strings {
		if v == nil {
			return fmt.Errorf("unable to encode langPackDifference#f385c1f6: field strings element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode langPackDifference#f385c1f6: field strings element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetLangCode returns value of LangCode field.
func (l *LangPackDifference) GetLangCode() (value string) {
	return l.LangCode
}

// GetFromVersion returns value of FromVersion field.
func (l *LangPackDifference) GetFromVersion() (value int) {
	return l.FromVersion
}

// GetVersion returns value of Version field.
func (l *LangPackDifference) GetVersion() (value int) {
	return l.Version
}

// GetStrings returns value of Strings field.
func (l *LangPackDifference) GetStrings() (value []LangPackStringClass) {
	return l.Strings
}

// MapStrings returns field Strings wrapped in LangPackStringClassArray helper.
func (l *LangPackDifference) MapStrings() (value LangPackStringClassArray) {
	return LangPackStringClassArray(l.Strings)
}

// Decode implements bin.Decoder.
func (l *LangPackDifference) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode langPackDifference#f385c1f6 to nil")
	}
	if err := b.ConsumeID(LangPackDifferenceTypeID); err != nil {
		return fmt.Errorf("unable to decode langPackDifference#f385c1f6: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field lang_code: %w", err)
		}
		l.LangCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field from_version: %w", err)
		}
		l.FromVersion = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field version: %w", err)
		}
		l.Version = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field strings: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeLangPackString(b)
			if err != nil {
				return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field strings: %w", err)
			}
			l.Strings = append(l.Strings, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for LangPackDifference.
var (
	_ bin.Encoder = &LangPackDifference{}
	_ bin.Decoder = &LangPackDifference{}
)
