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

// MessagesGetAttachedStickersRequest represents TL type `messages.getAttachedStickers#cc5b67cc`.
// Get stickers attached to a photo or video
//
// See https://core.telegram.org/method/messages.getAttachedStickers for reference.
type MessagesGetAttachedStickersRequest struct {
	// Stickered media
	Media InputStickeredMediaClass `tl:"media"`
}

// MessagesGetAttachedStickersRequestTypeID is TL type id of MessagesGetAttachedStickersRequest.
const MessagesGetAttachedStickersRequestTypeID = 0xcc5b67cc

func (g *MessagesGetAttachedStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Media == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAttachedStickersRequest) String() string {
	if g == nil {
		return "MessagesGetAttachedStickersRequest(nil)"
	}
	type Alias MessagesGetAttachedStickersRequest
	return fmt.Sprintf("MessagesGetAttachedStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetAttachedStickersRequest from given interface.
func (g *MessagesGetAttachedStickersRequest) FillFrom(from interface {
	GetMedia() (value InputStickeredMediaClass)
}) {
	g.Media = from.GetMedia()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetAttachedStickersRequest) TypeID() uint32 {
	return MessagesGetAttachedStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (g *MessagesGetAttachedStickersRequest) TypeName() string {
	return "messages.getAttachedStickers"
}

// Encode implements bin.Encoder.
func (g *MessagesGetAttachedStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAttachedStickers#cc5b67cc as nil")
	}
	b.PutID(MessagesGetAttachedStickersRequestTypeID)
	if g.Media == nil {
		return fmt.Errorf("unable to encode messages.getAttachedStickers#cc5b67cc: field media is nil")
	}
	if err := g.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getAttachedStickers#cc5b67cc: field media: %w", err)
	}
	return nil
}

// GetMedia returns value of Media field.
func (g *MessagesGetAttachedStickersRequest) GetMedia() (value InputStickeredMediaClass) {
	return g.Media
}

// Decode implements bin.Decoder.
func (g *MessagesGetAttachedStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAttachedStickers#cc5b67cc to nil")
	}
	if err := b.ConsumeID(MessagesGetAttachedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAttachedStickers#cc5b67cc: %w", err)
	}
	{
		value, err := DecodeInputStickeredMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAttachedStickers#cc5b67cc: field media: %w", err)
		}
		g.Media = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAttachedStickersRequest.
var (
	_ bin.Encoder = &MessagesGetAttachedStickersRequest{}
	_ bin.Decoder = &MessagesGetAttachedStickersRequest{}
)

// MessagesGetAttachedStickers invokes method messages.getAttachedStickers#cc5b67cc returning error if any.
// Get stickers attached to a photo or video
//
// See https://core.telegram.org/method/messages.getAttachedStickers for reference.
func (c *Client) MessagesGetAttachedStickers(ctx context.Context, media InputStickeredMediaClass) ([]StickerSetCoveredClass, error) {
	var result StickerSetCoveredClassVector

	request := &MessagesGetAttachedStickersRequest{
		Media: media,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return []StickerSetCoveredClass(result.Elems), nil
}
