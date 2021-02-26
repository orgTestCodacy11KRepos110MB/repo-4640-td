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

// ContactsImportedContacts represents TL type `contacts.importedContacts#77d01c3b`.
// Info on succesfully imported contacts.
//
// See https://core.telegram.org/constructor/contacts.importedContacts for reference.
type ContactsImportedContacts struct {
	// List of succesfully imported contacts
	Imported []ImportedContact `tl:"imported"`
	// Popular contacts
	PopularInvites []PopularContact `tl:"popular_invites"`
	// List of contact ids that could not be imported due to system limitation and will need to be imported at a later date.Parameter added in Layer 13¹
	//
	// Links:
	//  1) https://core.telegram.org/api/layers#layer-13
	RetryContacts []int64 `tl:"retry_contacts"`
	// List of users
	Users []UserClass `tl:"users"`
}

// ContactsImportedContactsTypeID is TL type id of ContactsImportedContacts.
const ContactsImportedContactsTypeID = 0x77d01c3b

func (i *ContactsImportedContacts) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Imported == nil) {
		return false
	}
	if !(i.PopularInvites == nil) {
		return false
	}
	if !(i.RetryContacts == nil) {
		return false
	}
	if !(i.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *ContactsImportedContacts) String() string {
	if i == nil {
		return "ContactsImportedContacts(nil)"
	}
	type Alias ContactsImportedContacts
	return fmt.Sprintf("ContactsImportedContacts%+v", Alias(*i))
}

// FillFrom fills ContactsImportedContacts from given interface.
func (i *ContactsImportedContacts) FillFrom(from interface {
	GetImported() (value []ImportedContact)
	GetPopularInvites() (value []PopularContact)
	GetRetryContacts() (value []int64)
	GetUsers() (value []UserClass)
}) {
	i.Imported = from.GetImported()
	i.PopularInvites = from.GetPopularInvites()
	i.RetryContacts = from.GetRetryContacts()
	i.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *ContactsImportedContacts) TypeID() uint32 {
	return ContactsImportedContactsTypeID
}

// TypeName returns name of type in TL schema.
func (i *ContactsImportedContacts) TypeName() string {
	return "contacts.importedContacts"
}

// Encode implements bin.Encoder.
func (i *ContactsImportedContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode contacts.importedContacts#77d01c3b as nil")
	}
	b.PutID(ContactsImportedContactsTypeID)
	b.PutVectorHeader(len(i.Imported))
	for idx, v := range i.Imported {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field imported element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(i.PopularInvites))
	for idx, v := range i.PopularInvites {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field popular_invites element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(i.RetryContacts))
	for _, v := range i.RetryContacts {
		b.PutLong(v)
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetImported returns value of Imported field.
func (i *ContactsImportedContacts) GetImported() (value []ImportedContact) {
	return i.Imported
}

// GetPopularInvites returns value of PopularInvites field.
func (i *ContactsImportedContacts) GetPopularInvites() (value []PopularContact) {
	return i.PopularInvites
}

// GetRetryContacts returns value of RetryContacts field.
func (i *ContactsImportedContacts) GetRetryContacts() (value []int64) {
	return i.RetryContacts
}

// GetUsers returns value of Users field.
func (i *ContactsImportedContacts) GetUsers() (value []UserClass) {
	return i.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (i *ContactsImportedContacts) MapUsers() (value UserClassArray) {
	return UserClassArray(i.Users)
}

// Decode implements bin.Decoder.
func (i *ContactsImportedContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode contacts.importedContacts#77d01c3b to nil")
	}
	if err := b.ConsumeID(ContactsImportedContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field imported: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ImportedContact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field imported: %w", err)
			}
			i.Imported = append(i.Imported, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field popular_invites: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PopularContact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field popular_invites: %w", err)
			}
			i.PopularInvites = append(i.PopularInvites, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field retry_contacts: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field retry_contacts: %w", err)
			}
			i.RetryContacts = append(i.RetryContacts, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsImportedContacts.
var (
	_ bin.Encoder = &ContactsImportedContacts{}
	_ bin.Decoder = &ContactsImportedContacts{}
)
