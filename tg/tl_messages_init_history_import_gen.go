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

// MessagesInitHistoryImportRequest represents TL type `messages.initHistoryImport#34090c3b`.
//
// See https://core.telegram.org/method/messages.initHistoryImport for reference.
type MessagesInitHistoryImportRequest struct {
	// Peer field of MessagesInitHistoryImportRequest.
	Peer InputPeerClass `tl:"peer"`
	// File field of MessagesInitHistoryImportRequest.
	File InputFileClass `tl:"file"`
	// MediaCount field of MessagesInitHistoryImportRequest.
	MediaCount int `tl:"media_count"`
}

// MessagesInitHistoryImportRequestTypeID is TL type id of MessagesInitHistoryImportRequest.
const MessagesInitHistoryImportRequestTypeID = 0x34090c3b

func (i *MessagesInitHistoryImportRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}
	if !(i.File == nil) {
		return false
	}
	if !(i.MediaCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *MessagesInitHistoryImportRequest) String() string {
	if i == nil {
		return "MessagesInitHistoryImportRequest(nil)"
	}
	type Alias MessagesInitHistoryImportRequest
	return fmt.Sprintf("MessagesInitHistoryImportRequest%+v", Alias(*i))
}

// FillFrom fills MessagesInitHistoryImportRequest from given interface.
func (i *MessagesInitHistoryImportRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetFile() (value InputFileClass)
	GetMediaCount() (value int)
}) {
	i.Peer = from.GetPeer()
	i.File = from.GetFile()
	i.MediaCount = from.GetMediaCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *MessagesInitHistoryImportRequest) TypeID() uint32 {
	return MessagesInitHistoryImportRequestTypeID
}

// TypeName returns name of type in TL schema.
func (i *MessagesInitHistoryImportRequest) TypeName() string {
	return "messages.initHistoryImport"
}

// Encode implements bin.Encoder.
func (i *MessagesInitHistoryImportRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.initHistoryImport#34090c3b as nil")
	}
	b.PutID(MessagesInitHistoryImportRequestTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode messages.initHistoryImport#34090c3b: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.initHistoryImport#34090c3b: field peer: %w", err)
	}
	if i.File == nil {
		return fmt.Errorf("unable to encode messages.initHistoryImport#34090c3b: field file is nil")
	}
	if err := i.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.initHistoryImport#34090c3b: field file: %w", err)
	}
	b.PutInt(i.MediaCount)
	return nil
}

// GetPeer returns value of Peer field.
func (i *MessagesInitHistoryImportRequest) GetPeer() (value InputPeerClass) {
	return i.Peer
}

// GetFile returns value of File field.
func (i *MessagesInitHistoryImportRequest) GetFile() (value InputFileClass) {
	return i.File
}

// GetMediaCount returns value of MediaCount field.
func (i *MessagesInitHistoryImportRequest) GetMediaCount() (value int) {
	return i.MediaCount
}

// Decode implements bin.Decoder.
func (i *MessagesInitHistoryImportRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.initHistoryImport#34090c3b to nil")
	}
	if err := b.ConsumeID(MessagesInitHistoryImportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.initHistoryImport#34090c3b: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.initHistoryImport#34090c3b: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.initHistoryImport#34090c3b: field file: %w", err)
		}
		i.File = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.initHistoryImport#34090c3b: field media_count: %w", err)
		}
		i.MediaCount = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesInitHistoryImportRequest.
var (
	_ bin.Encoder = &MessagesInitHistoryImportRequest{}
	_ bin.Decoder = &MessagesInitHistoryImportRequest{}
)

// MessagesInitHistoryImport invokes method messages.initHistoryImport#34090c3b returning error if any.
//
// See https://core.telegram.org/method/messages.initHistoryImport for reference.
func (c *Client) MessagesInitHistoryImport(ctx context.Context, request *MessagesInitHistoryImportRequest) (*MessagesHistoryImport, error) {
	var result MessagesHistoryImport

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
