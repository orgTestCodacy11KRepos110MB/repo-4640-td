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

// MessagesArchivedStickers represents TL type `messages.archivedStickers#4fcba9c8`.
// Archived stickersets
//
// See https://core.telegram.org/constructor/messages.archivedStickers for reference.
type MessagesArchivedStickers struct {
	// Number of archived stickers
	Count int `tl:"count"`
	// Archived stickersets
	Sets []StickerSetCoveredClass `tl:"sets"`
}

// MessagesArchivedStickersTypeID is TL type id of MessagesArchivedStickers.
const MessagesArchivedStickersTypeID = 0x4fcba9c8

func (a *MessagesArchivedStickers) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Count == 0) {
		return false
	}
	if !(a.Sets == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesArchivedStickers) String() string {
	if a == nil {
		return "MessagesArchivedStickers(nil)"
	}
	type Alias MessagesArchivedStickers
	return fmt.Sprintf("MessagesArchivedStickers%+v", Alias(*a))
}

// FillFrom fills MessagesArchivedStickers from given interface.
func (a *MessagesArchivedStickers) FillFrom(from interface {
	GetCount() (value int)
	GetSets() (value []StickerSetCoveredClass)
}) {
	a.Count = from.GetCount()
	a.Sets = from.GetSets()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (a *MessagesArchivedStickers) TypeID() uint32 {
	return MessagesArchivedStickersTypeID
}

// TypeName returns name of type in TL schema.
func (a *MessagesArchivedStickers) TypeName() string {
	return "messages.archivedStickers"
}

// Encode implements bin.Encoder.
func (a *MessagesArchivedStickers) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.archivedStickers#4fcba9c8 as nil")
	}
	b.PutID(MessagesArchivedStickersTypeID)
	b.PutInt(a.Count)
	b.PutVectorHeader(len(a.Sets))
	for idx, v := range a.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.archivedStickers#4fcba9c8: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.archivedStickers#4fcba9c8: field sets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (a *MessagesArchivedStickers) GetCount() (value int) {
	return a.Count
}

// GetSets returns value of Sets field.
func (a *MessagesArchivedStickers) GetSets() (value []StickerSetCoveredClass) {
	return a.Sets
}

// MapSets returns field Sets wrapped in StickerSetCoveredClassArray helper.
func (a *MessagesArchivedStickers) MapSets() (value StickerSetCoveredClassArray) {
	return StickerSetCoveredClassArray(a.Sets)
}

// Decode implements bin.Decoder.
func (a *MessagesArchivedStickers) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.archivedStickers#4fcba9c8 to nil")
	}
	if err := b.ConsumeID(MessagesArchivedStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.archivedStickers#4fcba9c8: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.archivedStickers#4fcba9c8: field count: %w", err)
		}
		a.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.archivedStickers#4fcba9c8: field sets: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.archivedStickers#4fcba9c8: field sets: %w", err)
			}
			a.Sets = append(a.Sets, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesArchivedStickers.
var (
	_ bin.Encoder = &MessagesArchivedStickers{}
	_ bin.Decoder = &MessagesArchivedStickers{}
)
