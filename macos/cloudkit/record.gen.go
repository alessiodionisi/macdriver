// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Record] class.
var RecordClass = _RecordClass{objc.GetClass("CKRecord")}

type _RecordClass struct {
	objc.Class
}

// An interface definition for the [Record] class.
type IRecord interface {
	objc.IObject
	ObjectForKey(key RecordFieldKey) RecordValueWrapper
	ObjectForKeyedSubscript(key RecordFieldKey) RecordValueWrapper
	SetObjectForKeyedSubscript(object PRecordValue, key RecordFieldKey)
	SetObjectObjectForKeyedSubscript(objectObject objc.IObject, key RecordFieldKey)
	AllTokens() []string
	SetParentReferenceFromRecordID(parentRecordID IRecordID)
	SetParentReferenceFromRecord(parentRecord IRecord)
	EncodeSystemFieldsWithCoder(coder foundation.ICoder)
	AllKeys() []RecordFieldKey
	ChangedKeys() []RecordFieldKey
	RecordID() RecordID
	Share() Reference
	Parent() Reference
	SetParent(value IReference)
	LastModifiedUserRecordID() RecordID
	ModificationDate() foundation.Date
	RecordChangeTag() string
	CreatorUserRecordID() RecordID
	EncryptedValues() RecordKeyValueSettingWrapper
	CreationDate() foundation.Date
	RecordType() RecordType
}

// A collection of key-value pairs that store your app’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord?language=objc
type Record struct {
	objc.Object
}

func RecordFrom(ptr unsafe.Pointer) Record {
	return Record{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ Record) InitWithRecordTypeZoneID(recordType RecordType, zoneID IRecordZoneID) Record {
	rv := objc.Call[Record](r_, objc.Sel("initWithRecordType:zoneID:"), recordType, objc.Ptr(zoneID))
	return rv
}

// Creates a record in the specified zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462202-initwithrecordtype?language=objc
func Record_InitWithRecordTypeZoneID(recordType RecordType, zoneID IRecordZoneID) Record {
	return RecordClass.Alloc().InitWithRecordTypeZoneID(recordType, zoneID)
}

func (rc _RecordClass) Alloc() Record {
	rv := objc.Call[Record](rc, objc.Sel("alloc"))
	return rv
}

func Record_Alloc() Record {
	return RecordClass.Alloc()
}

func (rc _RecordClass) New() Record {
	rv := objc.Call[Record](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecord() Record {
	return RecordClass.New()
}

func (r_ Record) Init() Record {
	rv := objc.Call[Record](r_, objc.Sel("init"))
	return rv
}

// Returns the object that the record stores for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462216-objectforkey?language=objc
func (r_ Record) ObjectForKey(key RecordFieldKey) RecordValueWrapper {
	rv := objc.Call[RecordValueWrapper](r_, objc.Sel("objectForKey:"), key)
	return rv
}

// Returns the object that the record stores for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462210-objectforkeyedsubscript?language=objc
func (r_ Record) ObjectForKeyedSubscript(key RecordFieldKey) RecordValueWrapper {
	rv := objc.Call[RecordValueWrapper](r_, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}

// Stores an object in the record using the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462221-setobject?language=objc
func (r_ Record) SetObjectForKeyedSubscript(object PRecordValue, key RecordFieldKey) {
	po0 := objc.WrapAsProtocol("CKRecordValue", object)
	objc.Call[objc.Void](r_, objc.Sel("setObject:forKeyedSubscript:"), po0, key)
}

// Stores an object in the record using the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462221-setobject?language=objc
func (r_ Record) SetObjectObjectForKeyedSubscript(objectObject objc.IObject, key RecordFieldKey) {
	objc.Call[objc.Void](r_, objc.Sel("setObject:forKeyedSubscript:"), objc.Ptr(objectObject), key)
}

// Returns an array of strings to use for full-text searches of the field’s string-based values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462199-alltokens?language=objc
func (r_ Record) AllTokens() []string {
	rv := objc.Call[[]string](r_, objc.Sel("allTokens"))
	return rv
}

// Creates and sets a reference object for a parent from the parent’s record ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1690508-setparentreferencefromrecordid?language=objc
func (r_ Record) SetParentReferenceFromRecordID(parentRecordID IRecordID) {
	objc.Call[objc.Void](r_, objc.Sel("setParentReferenceFromRecordID:"), objc.Ptr(parentRecordID))
}

// Creates and sets a reference object for a parent from its record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1690507-setparentreferencefromrecord?language=objc
func (r_ Record) SetParentReferenceFromRecord(parentRecord IRecord) {
	objc.Call[objc.Void](r_, objc.Sel("setParentReferenceFromRecord:"), objc.Ptr(parentRecord))
}

// Encodes the record’s system fields using the specified archiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462200-encodesystemfieldswithcoder?language=objc
func (r_ Record) EncodeSystemFieldsWithCoder(coder foundation.ICoder) {
	objc.Call[objc.Void](r_, objc.Sel("encodeSystemFieldsWithCoder:"), objc.Ptr(coder))
}

// Returns an array of the record’s keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462220-allkeys?language=objc
func (r_ Record) AllKeys() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](r_, objc.Sel("allKeys"))
	return rv
}

// Returns an array of keys with recent changes to their values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462197-changedkeys?language=objc
func (r_ Record) ChangedKeys() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](r_, objc.Sel("changedKeys"))
	return rv
}

// The unique ID of the record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462229-recordid?language=objc
func (r_ Record) RecordID() RecordID {
	rv := objc.Call[RecordID](r_, objc.Sel("recordID"))
	return rv
}

// A reference to the share object that determines the share status of the record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1640378-share?language=objc
func (r_ Record) Share() Reference {
	rv := objc.Call[Reference](r_, objc.Sel("share"))
	return rv
}

// A reference to the record’s parent record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1640527-parent?language=objc
func (r_ Record) Parent() Reference {
	rv := objc.Call[Reference](r_, objc.Sel("parent"))
	return rv
}

// A reference to the record’s parent record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1640527-parent?language=objc
func (r_ Record) SetParent(value IReference) {
	objc.Call[objc.Void](r_, objc.Sel("setParent:"), objc.Ptr(value))
}

// The ID of the user who most recently modified the record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462212-lastmodifieduserrecordid?language=objc
func (r_ Record) LastModifiedUserRecordID() RecordID {
	rv := objc.Call[RecordID](r_, objc.Sel("lastModifiedUserRecordID"))
	return rv
}

// The most recent time that CloudKit saved the record to the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462227-modificationdate?language=objc
func (r_ Record) ModificationDate() foundation.Date {
	rv := objc.Call[foundation.Date](r_, objc.Sel("modificationDate"))
	return rv
}

// The server change token for the record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462195-recordchangetag?language=objc
func (r_ Record) RecordChangeTag() string {
	rv := objc.Call[string](r_, objc.Sel("recordChangeTag"))
	return rv
}

// The ID of the user who creates the record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462208-creatoruserrecordid?language=objc
func (r_ Record) CreatorUserRecordID() RecordID {
	rv := objc.Call[RecordID](r_, objc.Sel("creatorUserRecordID"))
	return rv
}

// An object that manages the record’s encrypted key-value pairs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/3746821-encryptedvalues?language=objc
func (r_ Record) EncryptedValues() RecordKeyValueSettingWrapper {
	rv := objc.Call[RecordKeyValueSettingWrapper](r_, objc.Sel("encryptedValues"))
	return rv
}

// The time when CloudKit first saves the record to the server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462223-creationdate?language=objc
func (r_ Record) CreationDate() foundation.Date {
	rv := objc.Call[foundation.Date](r_, objc.Sel("creationDate"))
	return rv
}

// The value that your app defines to identify the type of record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/1462206-recordtype?language=objc
func (r_ Record) RecordType() RecordType {
	rv := objc.Call[RecordType](r_, objc.Sel("recordType"))
	return rv
}