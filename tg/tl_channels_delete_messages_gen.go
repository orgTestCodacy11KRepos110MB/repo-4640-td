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

// ChannelsDeleteMessagesRequest represents TL type `channels.deleteMessages#84c1fd4e`.
// Delete messages in a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.deleteMessages for reference.
type ChannelsDeleteMessagesRequest struct {
	// Channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass `tl:"channel"`
	// IDs of messages to delete
	ID []int `tl:"id"`
}

// ChannelsDeleteMessagesRequestTypeID is TL type id of ChannelsDeleteMessagesRequest.
const ChannelsDeleteMessagesRequestTypeID = 0x84c1fd4e

func (d *ChannelsDeleteMessagesRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Channel == nil) {
		return false
	}
	if !(d.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteMessagesRequest) String() string {
	if d == nil {
		return "ChannelsDeleteMessagesRequest(nil)"
	}
	type Alias ChannelsDeleteMessagesRequest
	return fmt.Sprintf("ChannelsDeleteMessagesRequest%+v", Alias(*d))
}

// FillFrom fills ChannelsDeleteMessagesRequest from given interface.
func (d *ChannelsDeleteMessagesRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetID() (value []int)
}) {
	d.Channel = from.GetChannel()
	d.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *ChannelsDeleteMessagesRequest) TypeID() uint32 {
	return ChannelsDeleteMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (d *ChannelsDeleteMessagesRequest) TypeName() string {
	return "channels.deleteMessages"
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteMessagesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteMessages#84c1fd4e as nil")
	}
	b.PutID(ChannelsDeleteMessagesRequestTypeID)
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteMessages#84c1fd4e: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteMessages#84c1fd4e: field channel: %w", err)
	}
	b.PutVectorHeader(len(d.ID))
	for _, v := range d.ID {
		b.PutInt(v)
	}
	return nil
}

// GetChannel returns value of Channel field.
func (d *ChannelsDeleteMessagesRequest) GetChannel() (value InputChannelClass) {
	return d.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (d *ChannelsDeleteMessagesRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return d.Channel.AsNotEmpty()
}

// GetID returns value of ID field.
func (d *ChannelsDeleteMessagesRequest) GetID() (value []int) {
	return d.ID
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteMessagesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteMessages#84c1fd4e to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: field id: %w", err)
			}
			d.ID = append(d.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsDeleteMessagesRequest.
var (
	_ bin.Encoder = &ChannelsDeleteMessagesRequest{}
	_ bin.Decoder = &ChannelsDeleteMessagesRequest{}
)

// ChannelsDeleteMessages invokes method channels.deleteMessages#84c1fd4e returning error if any.
// Delete messages in a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  403 MESSAGE_DELETE_FORBIDDEN: You can't delete one of the messages you tried to delete, most likely because it is a service message.
//  400 MSG_ID_INVALID: Invalid message ID provided
//
// See https://core.telegram.org/method/channels.deleteMessages for reference.
// Can be used by bots.
func (c *Client) ChannelsDeleteMessages(ctx context.Context, request *ChannelsDeleteMessagesRequest) (*MessagesAffectedMessages, error) {
	var result MessagesAffectedMessages

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
