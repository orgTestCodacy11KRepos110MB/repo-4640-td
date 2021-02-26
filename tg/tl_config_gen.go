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

// Config represents TL type `config#330b4067`.
// Current configuration
//
// See https://core.telegram.org/constructor/config for reference.
type Config struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Whether phone calls can be used
	PhonecallsEnabled bool `tl:"phonecalls_enabled"`
	// Whether the client should use P2P by default for phone calls with contacts
	DefaultP2PContacts bool `tl:"default_p2p_contacts"`
	// Whether the client should preload featured stickers
	PreloadFeaturedStickers bool `tl:"preload_featured_stickers"`
	// Whether the client should ignore phone entities¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	IgnorePhoneEntities bool `tl:"ignore_phone_entities"`
	// Whether incoming private messages can be deleted for both participants
	RevokePmInbox bool `tl:"revoke_pm_inbox"`
	// Indicates that telegram is probably censored by governments/ISPs in the current region
	BlockedMode bool `tl:"blocked_mode"`
	// Whether pfs¹ was used
	//
	// Links:
	//  1) https://core.telegram.org/api/pfs
	PFSEnabled bool `tl:"pfs_enabled"`
	// Current date at the server
	Date int `tl:"date"`
	// Expiration date of this config: when it expires it'll have to be refetched using help.getConfig¹
	//
	// Links:
	//  1) https://core.telegram.org/method/help.getConfig
	Expires int `tl:"expires"`
	// Whether we're connected to the test DCs
	TestMode bool `tl:"test_mode"`
	// ID of the DC that returned the reply
	ThisDC int `tl:"this_dc"`
	// DC IP list
	DCOptions []DcOption `tl:"dc_options"`
	// Domain name for fetching encrypted DC list from DNS TXT record
	DCTxtDomainName string `tl:"dc_txt_domain_name"`
	// Maximum member count for normal groups¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	ChatSizeMax int `tl:"chat_size_max"`
	// Maximum member count for supergroups¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	MegagroupSizeMax int `tl:"megagroup_size_max"`
	// Maximum number of messages that can be forwarded at once using messages.forwardMessages¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.forwardMessages
	ForwardedCountMax int `tl:"forwarded_count_max"`
	// The client should update its online status¹ every N milliseconds
	//
	// Links:
	//  1) https://core.telegram.org/method/account.updateStatus
	OnlineUpdatePeriodMs int `tl:"online_update_period_ms"`
	// Delay before offline status needs to be sent to the server
	OfflineBlurTimeoutMs int `tl:"offline_blur_timeout_ms"`
	// Time without any user activity after which it should be treated offline
	OfflineIdleTimeoutMs int `tl:"offline_idle_timeout_ms"`
	// If we are offline, but were online from some other client in last online_cloud_timeout_ms milliseconds after we had gone offline, then delay offline notification for notify_cloud_delay_ms milliseconds.
	OnlineCloudTimeoutMs int `tl:"online_cloud_timeout_ms"`
	// If we are offline, but online from some other client then delay sending the offline notification for notify_cloud_delay_ms milliseconds.
	NotifyCloudDelayMs int `tl:"notify_cloud_delay_ms"`
	// If some other client is online, then delay notification for notification_default_delay_ms milliseconds
	NotifyDefaultDelayMs int `tl:"notify_default_delay_ms"`
	// Not for client use
	PushChatPeriodMs int `tl:"push_chat_period_ms"`
	// Not for client use
	PushChatLimit int `tl:"push_chat_limit"`
	// Maximum count of saved gifs
	SavedGifsLimit int `tl:"saved_gifs_limit"`
	// Only messages with age smaller than the one specified can be edited
	EditTimeLimit int `tl:"edit_time_limit"`
	// Only channel/supergroup messages with age smaller than the specified can be deleted
	RevokeTimeLimit int `tl:"revoke_time_limit"`
	// Only private messages with age smaller than the specified can be deleted
	RevokePmTimeLimit int `tl:"revoke_pm_time_limit"`
	// Exponential decay rate for computing top peer rating¹
	//
	// Links:
	//  1) https://core.telegram.org/api/top-rating
	RatingEDecay int `tl:"rating_e_decay"`
	// Maximum number of recent stickers
	StickersRecentLimit int `tl:"stickers_recent_limit"`
	// Maximum number of faved stickers
	StickersFavedLimit int `tl:"stickers_faved_limit"`
	// Indicates that round videos (video notes) and voice messages sent in channels and older than the specified period must be marked as read
	ChannelsReadMediaPeriod int `tl:"channels_read_media_period"`
	// Temporary passport¹ sessions
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetTmpSessions and GetTmpSessions helpers.
	TmpSessions int `tl:"tmp_sessions"`
	// Maximum count of pinned dialogs
	PinnedDialogsCountMax int `tl:"pinned_dialogs_count_max"`
	// Maximum count of dialogs per folder
	PinnedInfolderCountMax int `tl:"pinned_infolder_count_max"`
	// Maximum allowed outgoing ring time in VoIP calls: if the user we're calling doesn't reply within the specified time (in milliseconds), we should hang up the call
	CallReceiveTimeoutMs int `tl:"call_receive_timeout_ms"`
	// Maximum allowed incoming ring time in VoIP calls: if the current user doesn't reply within the specified time (in milliseconds), the call will be automatically refused
	CallRingTimeoutMs int `tl:"call_ring_timeout_ms"`
	// VoIP connection timeout: if the instance of libtgvoip on the other side of the call doesn't connect to our instance of libtgvoip within the specified time (in milliseconds), the call must be aborted
	CallConnectTimeoutMs int `tl:"call_connect_timeout_ms"`
	// If during a VoIP call a packet isn't received for the specified period of time, the call must be aborted
	CallPacketTimeoutMs int `tl:"call_packet_timeout_ms"`
	// The domain to use to parse in-app links.For example t.me indicates that t.me/username links should parsed to @username, t.me/addsticker/name should be parsed to the appropriate stickerset and so on...
	MeURLPrefix string `tl:"me_url_prefix"`
	// URL to use to auto-update the current app
	//
	// Use SetAutoupdateURLPrefix and GetAutoupdateURLPrefix helpers.
	AutoupdateURLPrefix string `tl:"autoupdate_url_prefix"`
	// Username of the bot to use to search for GIFs
	//
	// Use SetGifSearchUsername and GetGifSearchUsername helpers.
	GifSearchUsername string `tl:"gif_search_username"`
	// Username of the bot to use to search for venues
	//
	// Use SetVenueSearchUsername and GetVenueSearchUsername helpers.
	VenueSearchUsername string `tl:"venue_search_username"`
	// Username of the bot to use for image search
	//
	// Use SetImgSearchUsername and GetImgSearchUsername helpers.
	ImgSearchUsername string `tl:"img_search_username"`
	// ID of the map provider to use for venues
	//
	// Use SetStaticMapsProvider and GetStaticMapsProvider helpers.
	StaticMapsProvider string `tl:"static_maps_provider"`
	// Maximum length of caption (length in utf8 codepoints)
	CaptionLengthMax int `tl:"caption_length_max"`
	// Maximum length of messages (length in utf8 codepoints)
	MessageLengthMax int `tl:"message_length_max"`
	// DC ID to use to download webfiles¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files
	WebfileDCID int `tl:"webfile_dc_id"`
	// Suggested language code
	//
	// Use SetSuggestedLangCode and GetSuggestedLangCode helpers.
	SuggestedLangCode string `tl:"suggested_lang_code"`
	// Language pack version
	//
	// Use SetLangPackVersion and GetLangPackVersion helpers.
	LangPackVersion int `tl:"lang_pack_version"`
	// Basic language pack version
	//
	// Use SetBaseLangPackVersion and GetBaseLangPackVersion helpers.
	BaseLangPackVersion int `tl:"base_lang_pack_version"`
}

// ConfigTypeID is TL type id of Config.
const ConfigTypeID = 0x330b4067

func (c *Config) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.PhonecallsEnabled == false) {
		return false
	}
	if !(c.DefaultP2PContacts == false) {
		return false
	}
	if !(c.PreloadFeaturedStickers == false) {
		return false
	}
	if !(c.IgnorePhoneEntities == false) {
		return false
	}
	if !(c.RevokePmInbox == false) {
		return false
	}
	if !(c.BlockedMode == false) {
		return false
	}
	if !(c.PFSEnabled == false) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.Expires == 0) {
		return false
	}
	if !(c.TestMode == false) {
		return false
	}
	if !(c.ThisDC == 0) {
		return false
	}
	if !(c.DCOptions == nil) {
		return false
	}
	if !(c.DCTxtDomainName == "") {
		return false
	}
	if !(c.ChatSizeMax == 0) {
		return false
	}
	if !(c.MegagroupSizeMax == 0) {
		return false
	}
	if !(c.ForwardedCountMax == 0) {
		return false
	}
	if !(c.OnlineUpdatePeriodMs == 0) {
		return false
	}
	if !(c.OfflineBlurTimeoutMs == 0) {
		return false
	}
	if !(c.OfflineIdleTimeoutMs == 0) {
		return false
	}
	if !(c.OnlineCloudTimeoutMs == 0) {
		return false
	}
	if !(c.NotifyCloudDelayMs == 0) {
		return false
	}
	if !(c.NotifyDefaultDelayMs == 0) {
		return false
	}
	if !(c.PushChatPeriodMs == 0) {
		return false
	}
	if !(c.PushChatLimit == 0) {
		return false
	}
	if !(c.SavedGifsLimit == 0) {
		return false
	}
	if !(c.EditTimeLimit == 0) {
		return false
	}
	if !(c.RevokeTimeLimit == 0) {
		return false
	}
	if !(c.RevokePmTimeLimit == 0) {
		return false
	}
	if !(c.RatingEDecay == 0) {
		return false
	}
	if !(c.StickersRecentLimit == 0) {
		return false
	}
	if !(c.StickersFavedLimit == 0) {
		return false
	}
	if !(c.ChannelsReadMediaPeriod == 0) {
		return false
	}
	if !(c.TmpSessions == 0) {
		return false
	}
	if !(c.PinnedDialogsCountMax == 0) {
		return false
	}
	if !(c.PinnedInfolderCountMax == 0) {
		return false
	}
	if !(c.CallReceiveTimeoutMs == 0) {
		return false
	}
	if !(c.CallRingTimeoutMs == 0) {
		return false
	}
	if !(c.CallConnectTimeoutMs == 0) {
		return false
	}
	if !(c.CallPacketTimeoutMs == 0) {
		return false
	}
	if !(c.MeURLPrefix == "") {
		return false
	}
	if !(c.AutoupdateURLPrefix == "") {
		return false
	}
	if !(c.GifSearchUsername == "") {
		return false
	}
	if !(c.VenueSearchUsername == "") {
		return false
	}
	if !(c.ImgSearchUsername == "") {
		return false
	}
	if !(c.StaticMapsProvider == "") {
		return false
	}
	if !(c.CaptionLengthMax == 0) {
		return false
	}
	if !(c.MessageLengthMax == 0) {
		return false
	}
	if !(c.WebfileDCID == 0) {
		return false
	}
	if !(c.SuggestedLangCode == "") {
		return false
	}
	if !(c.LangPackVersion == 0) {
		return false
	}
	if !(c.BaseLangPackVersion == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *Config) String() string {
	if c == nil {
		return "Config(nil)"
	}
	type Alias Config
	return fmt.Sprintf("Config%+v", Alias(*c))
}

// FillFrom fills Config from given interface.
func (c *Config) FillFrom(from interface {
	GetPhonecallsEnabled() (value bool)
	GetDefaultP2PContacts() (value bool)
	GetPreloadFeaturedStickers() (value bool)
	GetIgnorePhoneEntities() (value bool)
	GetRevokePmInbox() (value bool)
	GetBlockedMode() (value bool)
	GetPFSEnabled() (value bool)
	GetDate() (value int)
	GetExpires() (value int)
	GetTestMode() (value bool)
	GetThisDC() (value int)
	GetDCOptions() (value []DcOption)
	GetDCTxtDomainName() (value string)
	GetChatSizeMax() (value int)
	GetMegagroupSizeMax() (value int)
	GetForwardedCountMax() (value int)
	GetOnlineUpdatePeriodMs() (value int)
	GetOfflineBlurTimeoutMs() (value int)
	GetOfflineIdleTimeoutMs() (value int)
	GetOnlineCloudTimeoutMs() (value int)
	GetNotifyCloudDelayMs() (value int)
	GetNotifyDefaultDelayMs() (value int)
	GetPushChatPeriodMs() (value int)
	GetPushChatLimit() (value int)
	GetSavedGifsLimit() (value int)
	GetEditTimeLimit() (value int)
	GetRevokeTimeLimit() (value int)
	GetRevokePmTimeLimit() (value int)
	GetRatingEDecay() (value int)
	GetStickersRecentLimit() (value int)
	GetStickersFavedLimit() (value int)
	GetChannelsReadMediaPeriod() (value int)
	GetTmpSessions() (value int, ok bool)
	GetPinnedDialogsCountMax() (value int)
	GetPinnedInfolderCountMax() (value int)
	GetCallReceiveTimeoutMs() (value int)
	GetCallRingTimeoutMs() (value int)
	GetCallConnectTimeoutMs() (value int)
	GetCallPacketTimeoutMs() (value int)
	GetMeURLPrefix() (value string)
	GetAutoupdateURLPrefix() (value string, ok bool)
	GetGifSearchUsername() (value string, ok bool)
	GetVenueSearchUsername() (value string, ok bool)
	GetImgSearchUsername() (value string, ok bool)
	GetStaticMapsProvider() (value string, ok bool)
	GetCaptionLengthMax() (value int)
	GetMessageLengthMax() (value int)
	GetWebfileDCID() (value int)
	GetSuggestedLangCode() (value string, ok bool)
	GetLangPackVersion() (value int, ok bool)
	GetBaseLangPackVersion() (value int, ok bool)
}) {
	c.PhonecallsEnabled = from.GetPhonecallsEnabled()
	c.DefaultP2PContacts = from.GetDefaultP2PContacts()
	c.PreloadFeaturedStickers = from.GetPreloadFeaturedStickers()
	c.IgnorePhoneEntities = from.GetIgnorePhoneEntities()
	c.RevokePmInbox = from.GetRevokePmInbox()
	c.BlockedMode = from.GetBlockedMode()
	c.PFSEnabled = from.GetPFSEnabled()
	c.Date = from.GetDate()
	c.Expires = from.GetExpires()
	c.TestMode = from.GetTestMode()
	c.ThisDC = from.GetThisDC()
	c.DCOptions = from.GetDCOptions()
	c.DCTxtDomainName = from.GetDCTxtDomainName()
	c.ChatSizeMax = from.GetChatSizeMax()
	c.MegagroupSizeMax = from.GetMegagroupSizeMax()
	c.ForwardedCountMax = from.GetForwardedCountMax()
	c.OnlineUpdatePeriodMs = from.GetOnlineUpdatePeriodMs()
	c.OfflineBlurTimeoutMs = from.GetOfflineBlurTimeoutMs()
	c.OfflineIdleTimeoutMs = from.GetOfflineIdleTimeoutMs()
	c.OnlineCloudTimeoutMs = from.GetOnlineCloudTimeoutMs()
	c.NotifyCloudDelayMs = from.GetNotifyCloudDelayMs()
	c.NotifyDefaultDelayMs = from.GetNotifyDefaultDelayMs()
	c.PushChatPeriodMs = from.GetPushChatPeriodMs()
	c.PushChatLimit = from.GetPushChatLimit()
	c.SavedGifsLimit = from.GetSavedGifsLimit()
	c.EditTimeLimit = from.GetEditTimeLimit()
	c.RevokeTimeLimit = from.GetRevokeTimeLimit()
	c.RevokePmTimeLimit = from.GetRevokePmTimeLimit()
	c.RatingEDecay = from.GetRatingEDecay()
	c.StickersRecentLimit = from.GetStickersRecentLimit()
	c.StickersFavedLimit = from.GetStickersFavedLimit()
	c.ChannelsReadMediaPeriod = from.GetChannelsReadMediaPeriod()
	if val, ok := from.GetTmpSessions(); ok {
		c.TmpSessions = val
	}

	c.PinnedDialogsCountMax = from.GetPinnedDialogsCountMax()
	c.PinnedInfolderCountMax = from.GetPinnedInfolderCountMax()
	c.CallReceiveTimeoutMs = from.GetCallReceiveTimeoutMs()
	c.CallRingTimeoutMs = from.GetCallRingTimeoutMs()
	c.CallConnectTimeoutMs = from.GetCallConnectTimeoutMs()
	c.CallPacketTimeoutMs = from.GetCallPacketTimeoutMs()
	c.MeURLPrefix = from.GetMeURLPrefix()
	if val, ok := from.GetAutoupdateURLPrefix(); ok {
		c.AutoupdateURLPrefix = val
	}

	if val, ok := from.GetGifSearchUsername(); ok {
		c.GifSearchUsername = val
	}

	if val, ok := from.GetVenueSearchUsername(); ok {
		c.VenueSearchUsername = val
	}

	if val, ok := from.GetImgSearchUsername(); ok {
		c.ImgSearchUsername = val
	}

	if val, ok := from.GetStaticMapsProvider(); ok {
		c.StaticMapsProvider = val
	}

	c.CaptionLengthMax = from.GetCaptionLengthMax()
	c.MessageLengthMax = from.GetMessageLengthMax()
	c.WebfileDCID = from.GetWebfileDCID()
	if val, ok := from.GetSuggestedLangCode(); ok {
		c.SuggestedLangCode = val
	}

	if val, ok := from.GetLangPackVersion(); ok {
		c.LangPackVersion = val
	}

	if val, ok := from.GetBaseLangPackVersion(); ok {
		c.BaseLangPackVersion = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *Config) TypeID() uint32 {
	return ConfigTypeID
}

// TypeName returns name of type in TL schema.
func (c *Config) TypeName() string {
	return "config"
}

// Encode implements bin.Encoder.
func (c *Config) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode config#330b4067 as nil")
	}
	b.PutID(ConfigTypeID)
	if !(c.PhonecallsEnabled == false) {
		c.Flags.Set(1)
	}
	if !(c.DefaultP2PContacts == false) {
		c.Flags.Set(3)
	}
	if !(c.PreloadFeaturedStickers == false) {
		c.Flags.Set(4)
	}
	if !(c.IgnorePhoneEntities == false) {
		c.Flags.Set(5)
	}
	if !(c.RevokePmInbox == false) {
		c.Flags.Set(6)
	}
	if !(c.BlockedMode == false) {
		c.Flags.Set(8)
	}
	if !(c.PFSEnabled == false) {
		c.Flags.Set(13)
	}
	if !(c.TmpSessions == 0) {
		c.Flags.Set(0)
	}
	if !(c.AutoupdateURLPrefix == "") {
		c.Flags.Set(7)
	}
	if !(c.GifSearchUsername == "") {
		c.Flags.Set(9)
	}
	if !(c.VenueSearchUsername == "") {
		c.Flags.Set(10)
	}
	if !(c.ImgSearchUsername == "") {
		c.Flags.Set(11)
	}
	if !(c.StaticMapsProvider == "") {
		c.Flags.Set(12)
	}
	if !(c.SuggestedLangCode == "") {
		c.Flags.Set(2)
	}
	if !(c.LangPackVersion == 0) {
		c.Flags.Set(2)
	}
	if !(c.BaseLangPackVersion == 0) {
		c.Flags.Set(2)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode config#330b4067: field flags: %w", err)
	}
	b.PutInt(c.Date)
	b.PutInt(c.Expires)
	b.PutBool(c.TestMode)
	b.PutInt(c.ThisDC)
	b.PutVectorHeader(len(c.DCOptions))
	for idx, v := range c.DCOptions {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode config#330b4067: field dc_options element with index %d: %w", idx, err)
		}
	}
	b.PutString(c.DCTxtDomainName)
	b.PutInt(c.ChatSizeMax)
	b.PutInt(c.MegagroupSizeMax)
	b.PutInt(c.ForwardedCountMax)
	b.PutInt(c.OnlineUpdatePeriodMs)
	b.PutInt(c.OfflineBlurTimeoutMs)
	b.PutInt(c.OfflineIdleTimeoutMs)
	b.PutInt(c.OnlineCloudTimeoutMs)
	b.PutInt(c.NotifyCloudDelayMs)
	b.PutInt(c.NotifyDefaultDelayMs)
	b.PutInt(c.PushChatPeriodMs)
	b.PutInt(c.PushChatLimit)
	b.PutInt(c.SavedGifsLimit)
	b.PutInt(c.EditTimeLimit)
	b.PutInt(c.RevokeTimeLimit)
	b.PutInt(c.RevokePmTimeLimit)
	b.PutInt(c.RatingEDecay)
	b.PutInt(c.StickersRecentLimit)
	b.PutInt(c.StickersFavedLimit)
	b.PutInt(c.ChannelsReadMediaPeriod)
	if c.Flags.Has(0) {
		b.PutInt(c.TmpSessions)
	}
	b.PutInt(c.PinnedDialogsCountMax)
	b.PutInt(c.PinnedInfolderCountMax)
	b.PutInt(c.CallReceiveTimeoutMs)
	b.PutInt(c.CallRingTimeoutMs)
	b.PutInt(c.CallConnectTimeoutMs)
	b.PutInt(c.CallPacketTimeoutMs)
	b.PutString(c.MeURLPrefix)
	if c.Flags.Has(7) {
		b.PutString(c.AutoupdateURLPrefix)
	}
	if c.Flags.Has(9) {
		b.PutString(c.GifSearchUsername)
	}
	if c.Flags.Has(10) {
		b.PutString(c.VenueSearchUsername)
	}
	if c.Flags.Has(11) {
		b.PutString(c.ImgSearchUsername)
	}
	if c.Flags.Has(12) {
		b.PutString(c.StaticMapsProvider)
	}
	b.PutInt(c.CaptionLengthMax)
	b.PutInt(c.MessageLengthMax)
	b.PutInt(c.WebfileDCID)
	if c.Flags.Has(2) {
		b.PutString(c.SuggestedLangCode)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.LangPackVersion)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.BaseLangPackVersion)
	}
	return nil
}

// SetPhonecallsEnabled sets value of PhonecallsEnabled conditional field.
func (c *Config) SetPhonecallsEnabled(value bool) {
	if value {
		c.Flags.Set(1)
		c.PhonecallsEnabled = true
	} else {
		c.Flags.Unset(1)
		c.PhonecallsEnabled = false
	}
}

// GetPhonecallsEnabled returns value of PhonecallsEnabled conditional field.
func (c *Config) GetPhonecallsEnabled() (value bool) {
	return c.Flags.Has(1)
}

// SetDefaultP2PContacts sets value of DefaultP2PContacts conditional field.
func (c *Config) SetDefaultP2PContacts(value bool) {
	if value {
		c.Flags.Set(3)
		c.DefaultP2PContacts = true
	} else {
		c.Flags.Unset(3)
		c.DefaultP2PContacts = false
	}
}

// GetDefaultP2PContacts returns value of DefaultP2PContacts conditional field.
func (c *Config) GetDefaultP2PContacts() (value bool) {
	return c.Flags.Has(3)
}

// SetPreloadFeaturedStickers sets value of PreloadFeaturedStickers conditional field.
func (c *Config) SetPreloadFeaturedStickers(value bool) {
	if value {
		c.Flags.Set(4)
		c.PreloadFeaturedStickers = true
	} else {
		c.Flags.Unset(4)
		c.PreloadFeaturedStickers = false
	}
}

// GetPreloadFeaturedStickers returns value of PreloadFeaturedStickers conditional field.
func (c *Config) GetPreloadFeaturedStickers() (value bool) {
	return c.Flags.Has(4)
}

// SetIgnorePhoneEntities sets value of IgnorePhoneEntities conditional field.
func (c *Config) SetIgnorePhoneEntities(value bool) {
	if value {
		c.Flags.Set(5)
		c.IgnorePhoneEntities = true
	} else {
		c.Flags.Unset(5)
		c.IgnorePhoneEntities = false
	}
}

// GetIgnorePhoneEntities returns value of IgnorePhoneEntities conditional field.
func (c *Config) GetIgnorePhoneEntities() (value bool) {
	return c.Flags.Has(5)
}

// SetRevokePmInbox sets value of RevokePmInbox conditional field.
func (c *Config) SetRevokePmInbox(value bool) {
	if value {
		c.Flags.Set(6)
		c.RevokePmInbox = true
	} else {
		c.Flags.Unset(6)
		c.RevokePmInbox = false
	}
}

// GetRevokePmInbox returns value of RevokePmInbox conditional field.
func (c *Config) GetRevokePmInbox() (value bool) {
	return c.Flags.Has(6)
}

// SetBlockedMode sets value of BlockedMode conditional field.
func (c *Config) SetBlockedMode(value bool) {
	if value {
		c.Flags.Set(8)
		c.BlockedMode = true
	} else {
		c.Flags.Unset(8)
		c.BlockedMode = false
	}
}

// GetBlockedMode returns value of BlockedMode conditional field.
func (c *Config) GetBlockedMode() (value bool) {
	return c.Flags.Has(8)
}

// SetPFSEnabled sets value of PFSEnabled conditional field.
func (c *Config) SetPFSEnabled(value bool) {
	if value {
		c.Flags.Set(13)
		c.PFSEnabled = true
	} else {
		c.Flags.Unset(13)
		c.PFSEnabled = false
	}
}

// GetPFSEnabled returns value of PFSEnabled conditional field.
func (c *Config) GetPFSEnabled() (value bool) {
	return c.Flags.Has(13)
}

// GetDate returns value of Date field.
func (c *Config) GetDate() (value int) {
	return c.Date
}

// GetExpires returns value of Expires field.
func (c *Config) GetExpires() (value int) {
	return c.Expires
}

// GetTestMode returns value of TestMode field.
func (c *Config) GetTestMode() (value bool) {
	return c.TestMode
}

// GetThisDC returns value of ThisDC field.
func (c *Config) GetThisDC() (value int) {
	return c.ThisDC
}

// GetDCOptions returns value of DCOptions field.
func (c *Config) GetDCOptions() (value []DcOption) {
	return c.DCOptions
}

// GetDCTxtDomainName returns value of DCTxtDomainName field.
func (c *Config) GetDCTxtDomainName() (value string) {
	return c.DCTxtDomainName
}

// GetChatSizeMax returns value of ChatSizeMax field.
func (c *Config) GetChatSizeMax() (value int) {
	return c.ChatSizeMax
}

// GetMegagroupSizeMax returns value of MegagroupSizeMax field.
func (c *Config) GetMegagroupSizeMax() (value int) {
	return c.MegagroupSizeMax
}

// GetForwardedCountMax returns value of ForwardedCountMax field.
func (c *Config) GetForwardedCountMax() (value int) {
	return c.ForwardedCountMax
}

// GetOnlineUpdatePeriodMs returns value of OnlineUpdatePeriodMs field.
func (c *Config) GetOnlineUpdatePeriodMs() (value int) {
	return c.OnlineUpdatePeriodMs
}

// GetOfflineBlurTimeoutMs returns value of OfflineBlurTimeoutMs field.
func (c *Config) GetOfflineBlurTimeoutMs() (value int) {
	return c.OfflineBlurTimeoutMs
}

// GetOfflineIdleTimeoutMs returns value of OfflineIdleTimeoutMs field.
func (c *Config) GetOfflineIdleTimeoutMs() (value int) {
	return c.OfflineIdleTimeoutMs
}

// GetOnlineCloudTimeoutMs returns value of OnlineCloudTimeoutMs field.
func (c *Config) GetOnlineCloudTimeoutMs() (value int) {
	return c.OnlineCloudTimeoutMs
}

// GetNotifyCloudDelayMs returns value of NotifyCloudDelayMs field.
func (c *Config) GetNotifyCloudDelayMs() (value int) {
	return c.NotifyCloudDelayMs
}

// GetNotifyDefaultDelayMs returns value of NotifyDefaultDelayMs field.
func (c *Config) GetNotifyDefaultDelayMs() (value int) {
	return c.NotifyDefaultDelayMs
}

// GetPushChatPeriodMs returns value of PushChatPeriodMs field.
func (c *Config) GetPushChatPeriodMs() (value int) {
	return c.PushChatPeriodMs
}

// GetPushChatLimit returns value of PushChatLimit field.
func (c *Config) GetPushChatLimit() (value int) {
	return c.PushChatLimit
}

// GetSavedGifsLimit returns value of SavedGifsLimit field.
func (c *Config) GetSavedGifsLimit() (value int) {
	return c.SavedGifsLimit
}

// GetEditTimeLimit returns value of EditTimeLimit field.
func (c *Config) GetEditTimeLimit() (value int) {
	return c.EditTimeLimit
}

// GetRevokeTimeLimit returns value of RevokeTimeLimit field.
func (c *Config) GetRevokeTimeLimit() (value int) {
	return c.RevokeTimeLimit
}

// GetRevokePmTimeLimit returns value of RevokePmTimeLimit field.
func (c *Config) GetRevokePmTimeLimit() (value int) {
	return c.RevokePmTimeLimit
}

// GetRatingEDecay returns value of RatingEDecay field.
func (c *Config) GetRatingEDecay() (value int) {
	return c.RatingEDecay
}

// GetStickersRecentLimit returns value of StickersRecentLimit field.
func (c *Config) GetStickersRecentLimit() (value int) {
	return c.StickersRecentLimit
}

// GetStickersFavedLimit returns value of StickersFavedLimit field.
func (c *Config) GetStickersFavedLimit() (value int) {
	return c.StickersFavedLimit
}

// GetChannelsReadMediaPeriod returns value of ChannelsReadMediaPeriod field.
func (c *Config) GetChannelsReadMediaPeriod() (value int) {
	return c.ChannelsReadMediaPeriod
}

// SetTmpSessions sets value of TmpSessions conditional field.
func (c *Config) SetTmpSessions(value int) {
	c.Flags.Set(0)
	c.TmpSessions = value
}

// GetTmpSessions returns value of TmpSessions conditional field and
// boolean which is true if field was set.
func (c *Config) GetTmpSessions() (value int, ok bool) {
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.TmpSessions, true
}

// GetPinnedDialogsCountMax returns value of PinnedDialogsCountMax field.
func (c *Config) GetPinnedDialogsCountMax() (value int) {
	return c.PinnedDialogsCountMax
}

// GetPinnedInfolderCountMax returns value of PinnedInfolderCountMax field.
func (c *Config) GetPinnedInfolderCountMax() (value int) {
	return c.PinnedInfolderCountMax
}

// GetCallReceiveTimeoutMs returns value of CallReceiveTimeoutMs field.
func (c *Config) GetCallReceiveTimeoutMs() (value int) {
	return c.CallReceiveTimeoutMs
}

// GetCallRingTimeoutMs returns value of CallRingTimeoutMs field.
func (c *Config) GetCallRingTimeoutMs() (value int) {
	return c.CallRingTimeoutMs
}

// GetCallConnectTimeoutMs returns value of CallConnectTimeoutMs field.
func (c *Config) GetCallConnectTimeoutMs() (value int) {
	return c.CallConnectTimeoutMs
}

// GetCallPacketTimeoutMs returns value of CallPacketTimeoutMs field.
func (c *Config) GetCallPacketTimeoutMs() (value int) {
	return c.CallPacketTimeoutMs
}

// GetMeURLPrefix returns value of MeURLPrefix field.
func (c *Config) GetMeURLPrefix() (value string) {
	return c.MeURLPrefix
}

// SetAutoupdateURLPrefix sets value of AutoupdateURLPrefix conditional field.
func (c *Config) SetAutoupdateURLPrefix(value string) {
	c.Flags.Set(7)
	c.AutoupdateURLPrefix = value
}

// GetAutoupdateURLPrefix returns value of AutoupdateURLPrefix conditional field and
// boolean which is true if field was set.
func (c *Config) GetAutoupdateURLPrefix() (value string, ok bool) {
	if !c.Flags.Has(7) {
		return value, false
	}
	return c.AutoupdateURLPrefix, true
}

// SetGifSearchUsername sets value of GifSearchUsername conditional field.
func (c *Config) SetGifSearchUsername(value string) {
	c.Flags.Set(9)
	c.GifSearchUsername = value
}

// GetGifSearchUsername returns value of GifSearchUsername conditional field and
// boolean which is true if field was set.
func (c *Config) GetGifSearchUsername() (value string, ok bool) {
	if !c.Flags.Has(9) {
		return value, false
	}
	return c.GifSearchUsername, true
}

// SetVenueSearchUsername sets value of VenueSearchUsername conditional field.
func (c *Config) SetVenueSearchUsername(value string) {
	c.Flags.Set(10)
	c.VenueSearchUsername = value
}

// GetVenueSearchUsername returns value of VenueSearchUsername conditional field and
// boolean which is true if field was set.
func (c *Config) GetVenueSearchUsername() (value string, ok bool) {
	if !c.Flags.Has(10) {
		return value, false
	}
	return c.VenueSearchUsername, true
}

// SetImgSearchUsername sets value of ImgSearchUsername conditional field.
func (c *Config) SetImgSearchUsername(value string) {
	c.Flags.Set(11)
	c.ImgSearchUsername = value
}

// GetImgSearchUsername returns value of ImgSearchUsername conditional field and
// boolean which is true if field was set.
func (c *Config) GetImgSearchUsername() (value string, ok bool) {
	if !c.Flags.Has(11) {
		return value, false
	}
	return c.ImgSearchUsername, true
}

// SetStaticMapsProvider sets value of StaticMapsProvider conditional field.
func (c *Config) SetStaticMapsProvider(value string) {
	c.Flags.Set(12)
	c.StaticMapsProvider = value
}

// GetStaticMapsProvider returns value of StaticMapsProvider conditional field and
// boolean which is true if field was set.
func (c *Config) GetStaticMapsProvider() (value string, ok bool) {
	if !c.Flags.Has(12) {
		return value, false
	}
	return c.StaticMapsProvider, true
}

// GetCaptionLengthMax returns value of CaptionLengthMax field.
func (c *Config) GetCaptionLengthMax() (value int) {
	return c.CaptionLengthMax
}

// GetMessageLengthMax returns value of MessageLengthMax field.
func (c *Config) GetMessageLengthMax() (value int) {
	return c.MessageLengthMax
}

// GetWebfileDCID returns value of WebfileDCID field.
func (c *Config) GetWebfileDCID() (value int) {
	return c.WebfileDCID
}

// SetSuggestedLangCode sets value of SuggestedLangCode conditional field.
func (c *Config) SetSuggestedLangCode(value string) {
	c.Flags.Set(2)
	c.SuggestedLangCode = value
}

// GetSuggestedLangCode returns value of SuggestedLangCode conditional field and
// boolean which is true if field was set.
func (c *Config) GetSuggestedLangCode() (value string, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.SuggestedLangCode, true
}

// SetLangPackVersion sets value of LangPackVersion conditional field.
func (c *Config) SetLangPackVersion(value int) {
	c.Flags.Set(2)
	c.LangPackVersion = value
}

// GetLangPackVersion returns value of LangPackVersion conditional field and
// boolean which is true if field was set.
func (c *Config) GetLangPackVersion() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.LangPackVersion, true
}

// SetBaseLangPackVersion sets value of BaseLangPackVersion conditional field.
func (c *Config) SetBaseLangPackVersion(value int) {
	c.Flags.Set(2)
	c.BaseLangPackVersion = value
}

// GetBaseLangPackVersion returns value of BaseLangPackVersion conditional field and
// boolean which is true if field was set.
func (c *Config) GetBaseLangPackVersion() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.BaseLangPackVersion, true
}

// Decode implements bin.Decoder.
func (c *Config) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode config#330b4067 to nil")
	}
	if err := b.ConsumeID(ConfigTypeID); err != nil {
		return fmt.Errorf("unable to decode config#330b4067: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field flags: %w", err)
		}
	}
	c.PhonecallsEnabled = c.Flags.Has(1)
	c.DefaultP2PContacts = c.Flags.Has(3)
	c.PreloadFeaturedStickers = c.Flags.Has(4)
	c.IgnorePhoneEntities = c.Flags.Has(5)
	c.RevokePmInbox = c.Flags.Has(6)
	c.BlockedMode = c.Flags.Has(8)
	c.PFSEnabled = c.Flags.Has(13)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field expires: %w", err)
		}
		c.Expires = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field test_mode: %w", err)
		}
		c.TestMode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field this_dc: %w", err)
		}
		c.ThisDC = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field dc_options: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value DcOption
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode config#330b4067: field dc_options: %w", err)
			}
			c.DCOptions = append(c.DCOptions, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field dc_txt_domain_name: %w", err)
		}
		c.DCTxtDomainName = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field chat_size_max: %w", err)
		}
		c.ChatSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field megagroup_size_max: %w", err)
		}
		c.MegagroupSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field forwarded_count_max: %w", err)
		}
		c.ForwardedCountMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field online_update_period_ms: %w", err)
		}
		c.OnlineUpdatePeriodMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field offline_blur_timeout_ms: %w", err)
		}
		c.OfflineBlurTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field offline_idle_timeout_ms: %w", err)
		}
		c.OfflineIdleTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field online_cloud_timeout_ms: %w", err)
		}
		c.OnlineCloudTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field notify_cloud_delay_ms: %w", err)
		}
		c.NotifyCloudDelayMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field notify_default_delay_ms: %w", err)
		}
		c.NotifyDefaultDelayMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field push_chat_period_ms: %w", err)
		}
		c.PushChatPeriodMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field push_chat_limit: %w", err)
		}
		c.PushChatLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field saved_gifs_limit: %w", err)
		}
		c.SavedGifsLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field edit_time_limit: %w", err)
		}
		c.EditTimeLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field revoke_time_limit: %w", err)
		}
		c.RevokeTimeLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field revoke_pm_time_limit: %w", err)
		}
		c.RevokePmTimeLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field rating_e_decay: %w", err)
		}
		c.RatingEDecay = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field stickers_recent_limit: %w", err)
		}
		c.StickersRecentLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field stickers_faved_limit: %w", err)
		}
		c.StickersFavedLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field channels_read_media_period: %w", err)
		}
		c.ChannelsReadMediaPeriod = value
	}
	if c.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field tmp_sessions: %w", err)
		}
		c.TmpSessions = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field pinned_dialogs_count_max: %w", err)
		}
		c.PinnedDialogsCountMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field pinned_infolder_count_max: %w", err)
		}
		c.PinnedInfolderCountMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_receive_timeout_ms: %w", err)
		}
		c.CallReceiveTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_ring_timeout_ms: %w", err)
		}
		c.CallRingTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_connect_timeout_ms: %w", err)
		}
		c.CallConnectTimeoutMs = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field call_packet_timeout_ms: %w", err)
		}
		c.CallPacketTimeoutMs = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field me_url_prefix: %w", err)
		}
		c.MeURLPrefix = value
	}
	if c.Flags.Has(7) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field autoupdate_url_prefix: %w", err)
		}
		c.AutoupdateURLPrefix = value
	}
	if c.Flags.Has(9) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field gif_search_username: %w", err)
		}
		c.GifSearchUsername = value
	}
	if c.Flags.Has(10) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field venue_search_username: %w", err)
		}
		c.VenueSearchUsername = value
	}
	if c.Flags.Has(11) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field img_search_username: %w", err)
		}
		c.ImgSearchUsername = value
	}
	if c.Flags.Has(12) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field static_maps_provider: %w", err)
		}
		c.StaticMapsProvider = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field caption_length_max: %w", err)
		}
		c.CaptionLengthMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field message_length_max: %w", err)
		}
		c.MessageLengthMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field webfile_dc_id: %w", err)
		}
		c.WebfileDCID = value
	}
	if c.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field suggested_lang_code: %w", err)
		}
		c.SuggestedLangCode = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field lang_pack_version: %w", err)
		}
		c.LangPackVersion = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode config#330b4067: field base_lang_pack_version: %w", err)
		}
		c.BaseLangPackVersion = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Config.
var (
	_ bin.Encoder = &Config{}
	_ bin.Decoder = &Config{}
)
