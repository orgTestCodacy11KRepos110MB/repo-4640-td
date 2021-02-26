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

// PageListOrderedItemText represents TL type `pageListOrderedItemText#5e068047`.
// Ordered list of text items
//
// See https://core.telegram.org/constructor/pageListOrderedItemText for reference.
type PageListOrderedItemText struct {
	// Number of element within ordered list
	Num string `tl:"num"`
	// Text
	Text RichTextClass `tl:"text"`
}

// PageListOrderedItemTextTypeID is TL type id of PageListOrderedItemText.
const PageListOrderedItemTextTypeID = 0x5e068047

func (p *PageListOrderedItemText) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Num == "") {
		return false
	}
	if !(p.Text == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageListOrderedItemText) String() string {
	if p == nil {
		return "PageListOrderedItemText(nil)"
	}
	type Alias PageListOrderedItemText
	return fmt.Sprintf("PageListOrderedItemText%+v", Alias(*p))
}

// FillFrom fills PageListOrderedItemText from given interface.
func (p *PageListOrderedItemText) FillFrom(from interface {
	GetNum() (value string)
	GetText() (value RichTextClass)
}) {
	p.Num = from.GetNum()
	p.Text = from.GetText()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PageListOrderedItemText) TypeID() uint32 {
	return PageListOrderedItemTextTypeID
}

// TypeName returns name of type in TL schema.
func (p *PageListOrderedItemText) TypeName() string {
	return "pageListOrderedItemText"
}

// Encode implements bin.Encoder.
func (p *PageListOrderedItemText) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListOrderedItemText#5e068047 as nil")
	}
	b.PutID(PageListOrderedItemTextTypeID)
	b.PutString(p.Num)
	if p.Text == nil {
		return fmt.Errorf("unable to encode pageListOrderedItemText#5e068047: field text is nil")
	}
	if err := p.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageListOrderedItemText#5e068047: field text: %w", err)
	}
	return nil
}

// GetNum returns value of Num field.
func (p *PageListOrderedItemText) GetNum() (value string) {
	return p.Num
}

// GetText returns value of Text field.
func (p *PageListOrderedItemText) GetText() (value RichTextClass) {
	return p.Text
}

// Decode implements bin.Decoder.
func (p *PageListOrderedItemText) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListOrderedItemText#5e068047 to nil")
	}
	if err := b.ConsumeID(PageListOrderedItemTextTypeID); err != nil {
		return fmt.Errorf("unable to decode pageListOrderedItemText#5e068047: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageListOrderedItemText#5e068047: field num: %w", err)
		}
		p.Num = value
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageListOrderedItemText#5e068047: field text: %w", err)
		}
		p.Text = value
	}
	return nil
}

// construct implements constructor of PageListOrderedItemClass.
func (p PageListOrderedItemText) construct() PageListOrderedItemClass { return &p }

// Ensuring interfaces in compile-time for PageListOrderedItemText.
var (
	_ bin.Encoder = &PageListOrderedItemText{}
	_ bin.Decoder = &PageListOrderedItemText{}

	_ PageListOrderedItemClass = &PageListOrderedItemText{}
)

// PageListOrderedItemBlocks represents TL type `pageListOrderedItemBlocks#98dd8936`.
// Ordered list of IV¹ blocks
//
// Links:
//  1) https://instantview.telegram.org
//
// See https://core.telegram.org/constructor/pageListOrderedItemBlocks for reference.
type PageListOrderedItemBlocks struct {
	// Number of element within ordered list
	Num string `tl:"num"`
	// Item contents
	Blocks []PageBlockClass `tl:"blocks"`
}

// PageListOrderedItemBlocksTypeID is TL type id of PageListOrderedItemBlocks.
const PageListOrderedItemBlocksTypeID = 0x98dd8936

func (p *PageListOrderedItemBlocks) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Num == "") {
		return false
	}
	if !(p.Blocks == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageListOrderedItemBlocks) String() string {
	if p == nil {
		return "PageListOrderedItemBlocks(nil)"
	}
	type Alias PageListOrderedItemBlocks
	return fmt.Sprintf("PageListOrderedItemBlocks%+v", Alias(*p))
}

// FillFrom fills PageListOrderedItemBlocks from given interface.
func (p *PageListOrderedItemBlocks) FillFrom(from interface {
	GetNum() (value string)
	GetBlocks() (value []PageBlockClass)
}) {
	p.Num = from.GetNum()
	p.Blocks = from.GetBlocks()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PageListOrderedItemBlocks) TypeID() uint32 {
	return PageListOrderedItemBlocksTypeID
}

// TypeName returns name of type in TL schema.
func (p *PageListOrderedItemBlocks) TypeName() string {
	return "pageListOrderedItemBlocks"
}

// Encode implements bin.Encoder.
func (p *PageListOrderedItemBlocks) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListOrderedItemBlocks#98dd8936 as nil")
	}
	b.PutID(PageListOrderedItemBlocksTypeID)
	b.PutString(p.Num)
	b.PutVectorHeader(len(p.Blocks))
	for idx, v := range p.Blocks {
		if v == nil {
			return fmt.Errorf("unable to encode pageListOrderedItemBlocks#98dd8936: field blocks element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode pageListOrderedItemBlocks#98dd8936: field blocks element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetNum returns value of Num field.
func (p *PageListOrderedItemBlocks) GetNum() (value string) {
	return p.Num
}

// GetBlocks returns value of Blocks field.
func (p *PageListOrderedItemBlocks) GetBlocks() (value []PageBlockClass) {
	return p.Blocks
}

// MapBlocks returns field Blocks wrapped in PageBlockClassArray helper.
func (p *PageListOrderedItemBlocks) MapBlocks() (value PageBlockClassArray) {
	return PageBlockClassArray(p.Blocks)
}

// Decode implements bin.Decoder.
func (p *PageListOrderedItemBlocks) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListOrderedItemBlocks#98dd8936 to nil")
	}
	if err := b.ConsumeID(PageListOrderedItemBlocksTypeID); err != nil {
		return fmt.Errorf("unable to decode pageListOrderedItemBlocks#98dd8936: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageListOrderedItemBlocks#98dd8936: field num: %w", err)
		}
		p.Num = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode pageListOrderedItemBlocks#98dd8936: field blocks: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePageBlock(b)
			if err != nil {
				return fmt.Errorf("unable to decode pageListOrderedItemBlocks#98dd8936: field blocks: %w", err)
			}
			p.Blocks = append(p.Blocks, value)
		}
	}
	return nil
}

// construct implements constructor of PageListOrderedItemClass.
func (p PageListOrderedItemBlocks) construct() PageListOrderedItemClass { return &p }

// Ensuring interfaces in compile-time for PageListOrderedItemBlocks.
var (
	_ bin.Encoder = &PageListOrderedItemBlocks{}
	_ bin.Decoder = &PageListOrderedItemBlocks{}

	_ PageListOrderedItemClass = &PageListOrderedItemBlocks{}
)

// PageListOrderedItemClass represents PageListOrderedItem generic type.
//
// See https://core.telegram.org/type/PageListOrderedItem for reference.
//
// Example:
//  g, err := tg.DecodePageListOrderedItem(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.PageListOrderedItemText: // pageListOrderedItemText#5e068047
//  case *tg.PageListOrderedItemBlocks: // pageListOrderedItemBlocks#98dd8936
//  default: panic(v)
//  }
type PageListOrderedItemClass interface {
	bin.Encoder
	bin.Decoder
	construct() PageListOrderedItemClass

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

	// Number of element within ordered list
	GetNum() (value string)
}

// DecodePageListOrderedItem implements binary de-serialization for PageListOrderedItemClass.
func DecodePageListOrderedItem(buf *bin.Buffer) (PageListOrderedItemClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PageListOrderedItemTextTypeID:
		// Decoding pageListOrderedItemText#5e068047.
		v := PageListOrderedItemText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageListOrderedItemClass: %w", err)
		}
		return &v, nil
	case PageListOrderedItemBlocksTypeID:
		// Decoding pageListOrderedItemBlocks#98dd8936.
		v := PageListOrderedItemBlocks{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageListOrderedItemClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PageListOrderedItemClass: %w", bin.NewUnexpectedID(id))
	}
}

// PageListOrderedItem boxes the PageListOrderedItemClass providing a helper.
type PageListOrderedItemBox struct {
	PageListOrderedItem PageListOrderedItemClass
}

// Decode implements bin.Decoder for PageListOrderedItemBox.
func (b *PageListOrderedItemBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PageListOrderedItemBox to nil")
	}
	v, err := DecodePageListOrderedItem(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PageListOrderedItem = v
	return nil
}

// Encode implements bin.Encode for PageListOrderedItemBox.
func (b *PageListOrderedItemBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PageListOrderedItem == nil {
		return fmt.Errorf("unable to encode PageListOrderedItemClass as nil")
	}
	return b.PageListOrderedItem.Encode(buf)
}

// PageListOrderedItemClassArray is adapter for slice of PageListOrderedItemClass.
type PageListOrderedItemClassArray []PageListOrderedItemClass

// Sort sorts slice of PageListOrderedItemClass.
func (s PageListOrderedItemClassArray) Sort(less func(a, b PageListOrderedItemClass) bool) PageListOrderedItemClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PageListOrderedItemClass.
func (s PageListOrderedItemClassArray) SortStable(less func(a, b PageListOrderedItemClass) bool) PageListOrderedItemClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PageListOrderedItemClass.
func (s PageListOrderedItemClassArray) Retain(keep func(x PageListOrderedItemClass) bool) PageListOrderedItemClassArray {
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
func (s PageListOrderedItemClassArray) First() (v PageListOrderedItemClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PageListOrderedItemClassArray) Last() (v PageListOrderedItemClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PageListOrderedItemClassArray) PopFirst() (v PageListOrderedItemClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PageListOrderedItemClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PageListOrderedItemClassArray) Pop() (v PageListOrderedItemClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPageListOrderedItemText returns copy with only PageListOrderedItemText constructors.
func (s PageListOrderedItemClassArray) AsPageListOrderedItemText() (to PageListOrderedItemTextArray) {
	for _, elem := range s {
		value, ok := elem.(*PageListOrderedItemText)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPageListOrderedItemBlocks returns copy with only PageListOrderedItemBlocks constructors.
func (s PageListOrderedItemClassArray) AsPageListOrderedItemBlocks() (to PageListOrderedItemBlocksArray) {
	for _, elem := range s {
		value, ok := elem.(*PageListOrderedItemBlocks)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PageListOrderedItemTextArray is adapter for slice of PageListOrderedItemText.
type PageListOrderedItemTextArray []PageListOrderedItemText

// Sort sorts slice of PageListOrderedItemText.
func (s PageListOrderedItemTextArray) Sort(less func(a, b PageListOrderedItemText) bool) PageListOrderedItemTextArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PageListOrderedItemText.
func (s PageListOrderedItemTextArray) SortStable(less func(a, b PageListOrderedItemText) bool) PageListOrderedItemTextArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PageListOrderedItemText.
func (s PageListOrderedItemTextArray) Retain(keep func(x PageListOrderedItemText) bool) PageListOrderedItemTextArray {
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
func (s PageListOrderedItemTextArray) First() (v PageListOrderedItemText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PageListOrderedItemTextArray) Last() (v PageListOrderedItemText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PageListOrderedItemTextArray) PopFirst() (v PageListOrderedItemText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PageListOrderedItemText
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PageListOrderedItemTextArray) Pop() (v PageListOrderedItemText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PageListOrderedItemBlocksArray is adapter for slice of PageListOrderedItemBlocks.
type PageListOrderedItemBlocksArray []PageListOrderedItemBlocks

// Sort sorts slice of PageListOrderedItemBlocks.
func (s PageListOrderedItemBlocksArray) Sort(less func(a, b PageListOrderedItemBlocks) bool) PageListOrderedItemBlocksArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PageListOrderedItemBlocks.
func (s PageListOrderedItemBlocksArray) SortStable(less func(a, b PageListOrderedItemBlocks) bool) PageListOrderedItemBlocksArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PageListOrderedItemBlocks.
func (s PageListOrderedItemBlocksArray) Retain(keep func(x PageListOrderedItemBlocks) bool) PageListOrderedItemBlocksArray {
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
func (s PageListOrderedItemBlocksArray) First() (v PageListOrderedItemBlocks, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PageListOrderedItemBlocksArray) Last() (v PageListOrderedItemBlocks, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PageListOrderedItemBlocksArray) PopFirst() (v PageListOrderedItemBlocks, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PageListOrderedItemBlocks
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PageListOrderedItemBlocksArray) Pop() (v PageListOrderedItemBlocks, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
