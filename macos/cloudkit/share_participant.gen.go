// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ShareParticipant] class.
var ShareParticipantClass = _ShareParticipantClass{objc.GetClass("CKShareParticipant")}

type _ShareParticipantClass struct {
	objc.Class
}

// An interface definition for the [ShareParticipant] class.
type IShareParticipant interface {
	objc.IObject
	Role() ShareParticipantRole
	SetRole(value ShareParticipantRole)
	Permission() ShareParticipantPermission
	SetPermission(value ShareParticipantPermission)
	UserIdentity() UserIdentity
	AcceptanceStatus() ShareParticipantAcceptanceStatus
}

// An object that describes a user’s participation in a share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipant?language=objc
type ShareParticipant struct {
	objc.Object
}

func ShareParticipantFrom(ptr unsafe.Pointer) ShareParticipant {
	return ShareParticipant{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _ShareParticipantClass) Alloc() ShareParticipant {
	rv := objc.Call[ShareParticipant](sc, objc.Sel("alloc"))
	return rv
}

func ShareParticipant_Alloc() ShareParticipant {
	return ShareParticipantClass.Alloc()
}

func (sc _ShareParticipantClass) New() ShareParticipant {
	rv := objc.Call[ShareParticipant](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewShareParticipant() ShareParticipant {
	return ShareParticipantClass.New()
}

func (s_ ShareParticipant) Init() ShareParticipant {
	rv := objc.Call[ShareParticipant](s_, objc.Sel("init"))
	return rv
}

// The participant’s role for the share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipant/2980667-role?language=objc
func (s_ ShareParticipant) Role() ShareParticipantRole {
	rv := objc.Call[ShareParticipantRole](s_, objc.Sel("role"))
	return rv
}

// The participant’s role for the share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipant/2980667-role?language=objc
func (s_ ShareParticipant) SetRole(value ShareParticipantRole) {
	objc.Call[objc.Void](s_, objc.Sel("setRole:"), value)
}

// The participant’s permission level for the share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipant/1640433-permission?language=objc
func (s_ ShareParticipant) Permission() ShareParticipantPermission {
	rv := objc.Call[ShareParticipantPermission](s_, objc.Sel("permission"))
	return rv
}

// The participant’s permission level for the share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipant/1640433-permission?language=objc
func (s_ ShareParticipant) SetPermission(value ShareParticipantPermission) {
	objc.Call[objc.Void](s_, objc.Sel("setPermission:"), value)
}

// The identity of the participant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipant/1640488-useridentity?language=objc
func (s_ ShareParticipant) UserIdentity() UserIdentity {
	rv := objc.Call[UserIdentity](s_, objc.Sel("userIdentity"))
	return rv
}

// The current state of the user’s acceptance of the share. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckshareparticipant/1640395-acceptancestatus?language=objc
func (s_ ShareParticipant) AcceptanceStatus() ShareParticipantAcceptanceStatus {
	rv := objc.Call[ShareParticipantAcceptanceStatus](s_, objc.Sel("acceptanceStatus"))
	return rv
}