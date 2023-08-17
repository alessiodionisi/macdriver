// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Contact] class.
var ContactClass = _ContactClass{objc.GetClass("CNContact")}

type _ContactClass struct {
	objc.Class
}

// An interface definition for the [Contact] class.
type IContact interface {
	objc.IObject
	IsUnifiedWithContactWithIdentifier(contactIdentifier string) bool
	IsKeyAvailable(key string) bool
	AreKeysAvailable(keyDescriptors []objc.IObject) bool
	Dates() []LabeledValue
	PhoneNumbers() []LabeledValue
	PreviousFamilyName() string
	PhoneticOrganizationName() string
	ContactType() ContactType
	ThumbnailImageData() []byte
	PhoneticGivenName() string
	PhoneticMiddleName() string
	NamePrefix() string
	PostalAddresses() []LabeledValue
	NonGregorianBirthday() foundation.DateComponents
	ContactRelations() []LabeledValue
	Birthday() foundation.DateComponents
	ImageData() []byte
	Note() string
	MiddleName() string
	JobTitle() string
	GivenName() string
	EmailAddresses() []LabeledValue
	InstantMessageAddresses() []LabeledValue
	UrlAddresses() []LabeledValue
	ImageDataAvailable() bool
	OrganizationName() string
	NameSuffix() string
	DepartmentName() string
	PhoneticFamilyName() string
	SocialProfiles() []LabeledValue
	FamilyName() string
	Identifier() string
	Nickname() string
}

// An immutable object that stores information about a single contact, such as the contact's first name, phone numbers, and addresses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact?language=objc
type Contact struct {
	objc.Object
}

func ContactFrom(ptr unsafe.Pointer) Contact {
	return Contact{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContactClass) Alloc() Contact {
	rv := objc.Call[Contact](cc, objc.Sel("alloc"))
	return rv
}

func Contact_Alloc() Contact {
	return ContactClass.Alloc()
}

func (cc _ContactClass) New() Contact {
	rv := objc.Call[Contact](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContact() Contact {
	return ContactClass.New()
}

func (c_ Contact) Init() Contact {
	rv := objc.Call[Contact](c_, objc.Sel("init"))
	return rv
}

// Returns a predicate to find the contacts whose email address matches the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/3020510-predicateforcontactsmatchingemai?language=objc
func (cc _ContactClass) PredicateForContactsMatchingEmailAddress(emailAddress string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](cc, objc.Sel("predicateForContactsMatchingEmailAddress:"), emailAddress)
	return rv
}

// Returns a predicate to find the contacts whose email address matches the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/3020510-predicateforcontactsmatchingemai?language=objc
func Contact_PredicateForContactsMatchingEmailAddress(emailAddress string) foundation.Predicate {
	return ContactClass.PredicateForContactsMatchingEmailAddress(emailAddress)
}

// Returns a comparator to sort contacts with the specified order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403401-comparatorfornamesortorder?language=objc
func (cc _ContactClass) ComparatorForNameSortOrder(sortOrder ContactSortOrder) foundation.Comparator {
	rv := objc.Call[foundation.Comparator](cc, objc.Sel("comparatorForNameSortOrder:"), sortOrder)
	return rv
}

// Returns a comparator to sort contacts with the specified order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403401-comparatorfornamesortorder?language=objc
func Contact_ComparatorForNameSortOrder(sortOrder ContactSortOrder) foundation.Comparator {
	return ContactClass.ComparatorForNameSortOrder(sortOrder)
}

// Returns a Boolean indicating whether the current contact is a unified contact and includes a contact with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403218-isunifiedwithcontactwithidentifi?language=objc
func (c_ Contact) IsUnifiedWithContactWithIdentifier(contactIdentifier string) bool {
	rv := objc.Call[bool](c_, objc.Sel("isUnifiedWithContactWithIdentifier:"), contactIdentifier)
	return rv
}

// Determines whether the contact property value for the specified key is fetched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403210-iskeyavailable?language=objc
func (c_ Contact) IsKeyAvailable(key string) bool {
	rv := objc.Call[bool](c_, objc.Sel("isKeyAvailable:"), key)
	return rv
}

// Returns a predicate to find the contacts in the specified container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403246-predicateforcontactsincontainerw?language=objc
func (cc _ContactClass) PredicateForContactsInContainerWithIdentifier(containerIdentifier string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](cc, objc.Sel("predicateForContactsInContainerWithIdentifier:"), containerIdentifier)
	return rv
}

// Returns a predicate to find the contacts in the specified container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403246-predicateforcontactsincontainerw?language=objc
func Contact_PredicateForContactsInContainerWithIdentifier(containerIdentifier string) foundation.Predicate {
	return ContactClass.PredicateForContactsInContainerWithIdentifier(containerIdentifier)
}

// Returns a predicate to find the contacts whose phone number matches the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/3020511-predicateforcontactsmatchingphon?language=objc
func (cc _ContactClass) PredicateForContactsMatchingPhoneNumber(phoneNumber IPhoneNumber) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](cc, objc.Sel("predicateForContactsMatchingPhoneNumber:"), objc.Ptr(phoneNumber))
	return rv
}

// Returns a predicate to find the contacts whose phone number matches the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/3020511-predicateforcontactsmatchingphon?language=objc
func Contact_PredicateForContactsMatchingPhoneNumber(phoneNumber IPhoneNumber) foundation.Predicate {
	return ContactClass.PredicateForContactsMatchingPhoneNumber(phoneNumber)
}

// Fetches all the keys required for the contact sort comparator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402871-descriptorforallcomparatorkeys?language=objc
func (cc _ContactClass) DescriptorForAllComparatorKeys() objc.Object {
	rv := objc.Call[objc.Object](cc, objc.Sel("descriptorForAllComparatorKeys"))
	return rv
}

// Fetches all the keys required for the contact sort comparator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402871-descriptorforallcomparatorkeys?language=objc
func Contact_DescriptorForAllComparatorKeys() objc.Object {
	return ContactClass.DescriptorForAllComparatorKeys()
}

// Returns a predicate to find the contacts that are members in the specified group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402810-predicateforcontactsingroupwithi?language=objc
func (cc _ContactClass) PredicateForContactsInGroupWithIdentifier(groupIdentifier string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](cc, objc.Sel("predicateForContactsInGroupWithIdentifier:"), groupIdentifier)
	return rv
}

// Returns a predicate to find the contacts that are members in the specified group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402810-predicateforcontactsingroupwithi?language=objc
func Contact_PredicateForContactsInGroupWithIdentifier(groupIdentifier string) foundation.Predicate {
	return ContactClass.PredicateForContactsInGroupWithIdentifier(groupIdentifier)
}

// Returns a predicate to find the contacts matching the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403234-predicateforcontactsmatchingname?language=objc
func (cc _ContactClass) PredicateForContactsMatchingName(name string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](cc, objc.Sel("predicateForContactsMatchingName:"), name)
	return rv
}

// Returns a predicate to find the contacts matching the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403234-predicateforcontactsmatchingname?language=objc
func Contact_PredicateForContactsMatchingName(name string) foundation.Predicate {
	return ContactClass.PredicateForContactsMatchingName(name)
}

// Determines whether all contact property values for the specified keys are fetched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403416-arekeysavailable?language=objc
func (c_ Contact) AreKeysAvailable(keyDescriptors []objc.IObject) bool {
	rv := objc.Call[bool](c_, objc.Sel("areKeysAvailable:"), keyDescriptors)
	return rv
}

// Returns a string containing the localized contact property name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403244-localizedstringforkey?language=objc
func (cc _ContactClass) LocalizedStringForKey(key string) string {
	rv := objc.Call[string](cc, objc.Sel("localizedStringForKey:"), key)
	return rv
}

// Returns a string containing the localized contact property name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403244-localizedstringforkey?language=objc
func Contact_LocalizedStringForKey(key string) string {
	return ContactClass.LocalizedStringForKey(key)
}

// Returns a predicate to find the contacts matching the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402882-predicateforcontactswithidentifi?language=objc
func (cc _ContactClass) PredicateForContactsWithIdentifiers(identifiers []string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](cc, objc.Sel("predicateForContactsWithIdentifiers:"), identifiers)
	return rv
}

// Returns a predicate to find the contacts matching the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402882-predicateforcontactswithidentifi?language=objc
func Contact_PredicateForContactsWithIdentifiers(identifiers []string) foundation.Predicate {
	return ContactClass.PredicateForContactsWithIdentifiers(identifiers)
}

// An array containing labeled Gregorian dates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403196-dates?language=objc
func (c_ Contact) Dates() []LabeledValue {
	rv := objc.Call[[]LabeledValue](c_, objc.Sel("dates"))
	return rv
}

// An array of labeled phone numbers for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403305-phonenumbers?language=objc
func (c_ Contact) PhoneNumbers() []LabeledValue {
	rv := objc.Call[[]LabeledValue](c_, objc.Sel("phoneNumbers"))
	return rv
}

// A string for the previous family name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402878-previousfamilyname?language=objc
func (c_ Contact) PreviousFamilyName() string {
	rv := objc.Call[string](c_, objc.Sel("previousFamilyName"))
	return rv
}

// The phonetic name of the organization associated with the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/2142774-phoneticorganizationname?language=objc
func (c_ Contact) PhoneticOrganizationName() string {
	rv := objc.Call[string](c_, objc.Sel("phoneticOrganizationName"))
	return rv
}

// An enum identifying the contact type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403119-contacttype?language=objc
func (c_ Contact) ContactType() ContactType {
	rv := objc.Call[ContactType](c_, objc.Sel("contactType"))
	return rv
}

// The thumbnail version of the contact’s profile picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402903-thumbnailimagedata?language=objc
func (c_ Contact) ThumbnailImageData() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("thumbnailImageData"))
	return rv
}

// The phonetic given name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403068-phoneticgivenname?language=objc
func (c_ Contact) PhoneticGivenName() string {
	rv := objc.Call[string](c_, objc.Sel("phoneticGivenName"))
	return rv
}

// The phonetic middle name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403031-phoneticmiddlename?language=objc
func (c_ Contact) PhoneticMiddleName() string {
	rv := objc.Call[string](c_, objc.Sel("phoneticMiddleName"))
	return rv
}

// The name prefix of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403002-nameprefix?language=objc
func (c_ Contact) NamePrefix() string {
	rv := objc.Call[string](c_, objc.Sel("namePrefix"))
	return rv
}

// An array of labeled postal addresses for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403066-postaladdresses?language=objc
func (c_ Contact) PostalAddresses() []LabeledValue {
	rv := objc.Call[[]LabeledValue](c_, objc.Sel("postalAddresses"))
	return rv
}

// A date component for the non-Gregorian birthday of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403425-nongregorianbirthday?language=objc
func (c_ Contact) NonGregorianBirthday() foundation.DateComponents {
	rv := objc.Call[foundation.DateComponents](c_, objc.Sel("nonGregorianBirthday"))
	return rv
}

// An array of labeled relations for the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402998-contactrelations?language=objc
func (c_ Contact) ContactRelations() []LabeledValue {
	rv := objc.Call[[]LabeledValue](c_, objc.Sel("contactRelations"))
	return rv
}

// A date component for the Gregorian birthday of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403059-birthday?language=objc
func (c_ Contact) Birthday() foundation.DateComponents {
	rv := objc.Call[foundation.DateComponents](c_, objc.Sel("birthday"))
	return rv
}

// The profile picture of a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402968-imagedata?language=objc
func (c_ Contact) ImageData() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("imageData"))
	return rv
}

// A string containing notes for the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402800-note?language=objc
func (c_ Contact) Note() string {
	rv := objc.Call[string](c_, objc.Sel("note"))
	return rv
}

// The middle name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403040-middlename?language=objc
func (c_ Contact) MiddleName() string {
	rv := objc.Call[string](c_, objc.Sel("middleName"))
	return rv
}

// The contact’s job title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403151-jobtitle?language=objc
func (c_ Contact) JobTitle() string {
	rv := objc.Call[string](c_, objc.Sel("jobTitle"))
	return rv
}

// The given name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403313-givenname?language=objc
func (c_ Contact) GivenName() string {
	rv := objc.Call[string](c_, objc.Sel("givenName"))
	return rv
}

// An array of labeled email addresses for the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403338-emailaddresses?language=objc
func (c_ Contact) EmailAddresses() []LabeledValue {
	rv := objc.Call[[]LabeledValue](c_, objc.Sel("emailAddresses"))
	return rv
}

// An array of labeled IM addresses for the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403147-instantmessageaddresses?language=objc
func (c_ Contact) InstantMessageAddresses() []LabeledValue {
	rv := objc.Call[[]LabeledValue](c_, objc.Sel("instantMessageAddresses"))
	return rv
}

// An array of labeled URL addresses for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403111-urladdresses?language=objc
func (c_ Contact) UrlAddresses() []LabeledValue {
	rv := objc.Call[[]LabeledValue](c_, objc.Sel("urlAddresses"))
	return rv
}

// A Boolean indicating whether a contact has a profile picture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1614589-imagedataavailable?language=objc
func (c_ Contact) ImageDataAvailable() bool {
	rv := objc.Call[bool](c_, objc.Sel("imageDataAvailable"))
	return rv
}

// The name of the organization associated with the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403182-organizationname?language=objc
func (c_ Contact) OrganizationName() string {
	rv := objc.Call[string](c_, objc.Sel("organizationName"))
	return rv
}

// The name suffix of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1402817-namesuffix?language=objc
func (c_ Contact) NameSuffix() string {
	rv := objc.Call[string](c_, objc.Sel("nameSuffix"))
	return rv
}

// The name of the department associated with the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403232-departmentname?language=objc
func (c_ Contact) DepartmentName() string {
	rv := objc.Call[string](c_, objc.Sel("departmentName"))
	return rv
}

// A string for the phonetic family name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403186-phoneticfamilyname?language=objc
func (c_ Contact) PhoneticFamilyName() string {
	rv := objc.Call[string](c_, objc.Sel("phoneticFamilyName"))
	return rv
}

// An array of labeled social profiles for a contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403054-socialprofiles?language=objc
func (c_ Contact) SocialProfiles() []LabeledValue {
	rv := objc.Call[[]LabeledValue](c_, objc.Sel("socialProfiles"))
	return rv
}

// The family name of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403397-familyname?language=objc
func (c_ Contact) FamilyName() string {
	rv := objc.Call[string](c_, objc.Sel("familyName"))
	return rv
}

// A value that uniquely identifies a contact on the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403103-identifier?language=objc
func (c_ Contact) Identifier() string {
	rv := objc.Call[string](c_, objc.Sel("identifier"))
	return rv
}

// The nickname of the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontact/1403299-nickname?language=objc
func (c_ Contact) Nickname() string {
	rv := objc.Call[string](c_, objc.Sel("nickname"))
	return rv
}