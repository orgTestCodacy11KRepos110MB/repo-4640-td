// Code generated by gotdgen, DO NOT EDIT.

package mt

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// HTTPWaitRequest represents TL type `http_wait#9299359f`.
type HTTPWaitRequest struct {
	// MaxDelay field of HTTPWaitRequest.
	MaxDelay int
	// WaitAfter field of HTTPWaitRequest.
	WaitAfter int
	// MaxWait field of HTTPWaitRequest.
	MaxWait int
}

// HTTPWaitRequestTypeID is TL type id of HTTPWaitRequest.
const HTTPWaitRequestTypeID = 0x9299359f

// String implements fmt.Stringer.
func (h *HTTPWaitRequest) String() string {
	if h == nil {
		return "HTTPWaitRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HTTPWaitRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tMaxDelay: ")
	sb.WriteString(fmt.Sprint(h.MaxDelay))
	sb.WriteString(",\n")
	sb.WriteString("\tWaitAfter: ")
	sb.WriteString(fmt.Sprint(h.WaitAfter))
	sb.WriteString(",\n")
	sb.WriteString("\tMaxWait: ")
	sb.WriteString(fmt.Sprint(h.MaxWait))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (h *HTTPWaitRequest) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode http_wait#9299359f as nil")
	}
	b.PutID(HTTPWaitRequestTypeID)
	b.PutInt(h.MaxDelay)
	b.PutInt(h.WaitAfter)
	b.PutInt(h.MaxWait)
	return nil
}

// Decode implements bin.Decoder.
func (h *HTTPWaitRequest) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode http_wait#9299359f to nil")
	}
	if err := b.ConsumeID(HTTPWaitRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode http_wait#9299359f: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode http_wait#9299359f: field max_delay: %w", err)
		}
		h.MaxDelay = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode http_wait#9299359f: field wait_after: %w", err)
		}
		h.WaitAfter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode http_wait#9299359f: field max_wait: %w", err)
		}
		h.MaxWait = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HTTPWaitRequest.
var (
	_ bin.Encoder = &HTTPWaitRequest{}
	_ bin.Decoder = &HTTPWaitRequest{}
)