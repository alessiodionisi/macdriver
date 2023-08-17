// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccessibilityCustomAction] class.
var AccessibilityCustomActionClass = _AccessibilityCustomActionClass{objc.GetClass("NSAccessibilityCustomAction")}

type _AccessibilityCustomActionClass struct {
	objc.Class
}

// An interface definition for the [AccessibilityCustomAction] class.
type IAccessibilityCustomAction interface {
	objc.IObject
	Target() objc.Object
	SetTarget(value objc.IObject)
	Name() string
	SetName(value string)
	Selector() objc.Selector
	SetSelector(value objc.Selector)
	Handler() func() bool
	SetHandler(value func() bool)
}

// A custom action to perform on an accessible object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction?language=objc
type AccessibilityCustomAction struct {
	objc.Object
}

func AccessibilityCustomActionFrom(ptr unsafe.Pointer) AccessibilityCustomAction {
	return AccessibilityCustomAction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ AccessibilityCustomAction) InitWithNameHandler(name string, handler func() bool) AccessibilityCustomAction {
	rv := objc.Call[AccessibilityCustomAction](a_, objc.Sel("initWithName:handler:"), name, handler)
	return rv
}

// Creates a custom action object with the specified name and handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction/2870120-initwithname?language=objc
func AccessibilityCustomAction_InitWithNameHandler(name string, handler func() bool) AccessibilityCustomAction {
	return AccessibilityCustomActionClass.Alloc().InitWithNameHandler(name, handler)
}

func (ac _AccessibilityCustomActionClass) Alloc() AccessibilityCustomAction {
	rv := objc.Call[AccessibilityCustomAction](ac, objc.Sel("alloc"))
	return rv
}

func AccessibilityCustomAction_Alloc() AccessibilityCustomAction {
	return AccessibilityCustomActionClass.Alloc()
}

func (ac _AccessibilityCustomActionClass) New() AccessibilityCustomAction {
	rv := objc.Call[AccessibilityCustomAction](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccessibilityCustomAction() AccessibilityCustomAction {
	return AccessibilityCustomActionClass.New()
}

func (a_ AccessibilityCustomAction) Init() AccessibilityCustomAction {
	rv := objc.Call[AccessibilityCustomAction](a_, objc.Sel("init"))
	return rv
}

// The object that performs the action through a selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction/2870105-target?language=objc
func (a_ AccessibilityCustomAction) Target() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("target"))
	return rv
}

// The object that performs the action through a selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction/2870105-target?language=objc
func (a_ AccessibilityCustomAction) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setTarget:"), value)
}

// A localized name that describes the action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction/2870118-name?language=objc
func (a_ AccessibilityCustomAction) Name() string {
	rv := objc.Call[string](a_, objc.Sel("name"))
	return rv
}

// A localized name that describes the action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction/2870118-name?language=objc
func (a_ AccessibilityCustomAction) SetName(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setName:"), value)
}

// The method to call on the target to perform the action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction/2870110-selector?language=objc
func (a_ AccessibilityCustomAction) Selector() objc.Selector {
	rv := objc.Call[objc.Selector](a_, objc.Sel("selector"))
	return rv
}

// The method to call on the target to perform the action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction/2870110-selector?language=objc
func (a_ AccessibilityCustomAction) SetSelector(value objc.Selector) {
	objc.Call[objc.Void](a_, objc.Sel("setSelector:"), value)
}

// The closure that handles the execution of the action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction/2870157-handler?language=objc
func (a_ AccessibilityCustomAction) Handler() func() bool {
	rv := objc.Call[func() bool](a_, objc.Sel("handler"))
	return rv
}

// The closure that handles the execution of the action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomaction/2870157-handler?language=objc
func (a_ AccessibilityCustomAction) SetHandler(value func() bool) {
	objc.Call[objc.Void](a_, objc.Sel("setHandler:"), value)
}