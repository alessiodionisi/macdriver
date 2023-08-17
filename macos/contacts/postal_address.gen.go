// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PostalAddress] class.
var PostalAddressClass = _PostalAddressClass{objc.GetClass("CNPostalAddress")}

type _PostalAddressClass struct {
	objc.Class
}

// An interface definition for the [PostalAddress] class.
type IPostalAddress interface {
	objc.IObject
	State() string
	SubLocality() string
	Street() string
	City() string
	PostalCode() string
	ISOCountryCode() string
	SubAdministrativeArea() string
	Country() string
}

// An immutable representation of the postal address for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress?language=objc
type PostalAddress struct {
	objc.Object
}

func PostalAddressFrom(ptr unsafe.Pointer) PostalAddress {
	return PostalAddress{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PostalAddressClass) Alloc() PostalAddress {
	rv := objc.Call[PostalAddress](pc, objc.Sel("alloc"))
	return rv
}

func PostalAddress_Alloc() PostalAddress {
	return PostalAddressClass.Alloc()
}

func (pc _PostalAddressClass) New() PostalAddress {
	rv := objc.Call[PostalAddress](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPostalAddress() PostalAddress {
	return PostalAddressClass.New()
}

func (p_ PostalAddress) Init() PostalAddress {
	rv := objc.Call[PostalAddress](p_, objc.Sel("init"))
	return rv
}

// Returns the localized name for the property associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/1402938-localizedstringforkey?language=objc
func (pc _PostalAddressClass) LocalizedStringForKey(key string) string {
	rv := objc.Call[string](pc, objc.Sel("localizedStringForKey:"), key)
	return rv
}

// Returns the localized name for the property associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/1402938-localizedstringforkey?language=objc
func PostalAddress_LocalizedStringForKey(key string) string {
	return PostalAddressClass.LocalizedStringForKey(key)
}

// The state name in a postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/1402886-state?language=objc
func (p_ PostalAddress) State() string {
	rv := objc.Call[string](p_, objc.Sel("state"))
	return rv
}

// Additional information associated with the location, typically defined at the city or town level, in a postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/2799049-sublocality?language=objc
func (p_ PostalAddress) SubLocality() string {
	rv := objc.Call[string](p_, objc.Sel("subLocality"))
	return rv
}

// The street name in a postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/1403414-street?language=objc
func (p_ PostalAddress) Street() string {
	rv := objc.Call[string](p_, objc.Sel("street"))
	return rv
}

// The city name in a postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/1403034-city?language=objc
func (p_ PostalAddress) City() string {
	rv := objc.Call[string](p_, objc.Sel("city"))
	return rv
}

// The postal code in a postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/1402983-postalcode?language=objc
func (p_ PostalAddress) PostalCode() string {
	rv := objc.Call[string](p_, objc.Sel("postalCode"))
	return rv
}

// The ISO country code for the country or region in a postal address, using the ISO 3166-1 alpha-2 standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/1403396-isocountrycode?language=objc
func (p_ PostalAddress) ISOCountryCode() string {
	rv := objc.Call[string](p_, objc.Sel("ISOCountryCode"))
	return rv
}

// The subadministrative area (such as a county or other region) in a postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/2799048-subadministrativearea?language=objc
func (p_ PostalAddress) SubAdministrativeArea() string {
	rv := objc.Call[string](p_, objc.Sel("subAdministrativeArea"))
	return rv
}

// The country or region name in a postal address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnpostaladdress/1402880-country?language=objc
func (p_ PostalAddress) Country() string {
	rv := objc.Call[string](p_, objc.Sel("country"))
	return rv
}