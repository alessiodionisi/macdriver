// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SplitViewController] class.
var SplitViewControllerClass = _SplitViewControllerClass{objc.GetClass("NSSplitViewController")}

type _SplitViewControllerClass struct {
	objc.Class
}

// An interface definition for the [SplitViewController] class.
type ISplitViewController interface {
	IViewController
	RemoveSplitViewItem(splitViewItem ISplitViewItem)
	ToggleSidebar(sender objc.IObject) objc.Object
	AddSplitViewItem(splitViewItem ISplitViewItem)
	InsertSplitViewItemAtIndex(splitViewItem ISplitViewItem, index int)
	ValidateUserInterfaceItem(item PValidatedUserInterfaceItem) bool
	ValidateUserInterfaceItemObject(itemObject objc.IObject) bool
	SplitViewItemForViewController(viewController IViewController) SplitViewItem
	SplitViewItems() []SplitViewItem
	SetSplitViewItems(value []ISplitViewItem)
	SplitView() SplitView
	SetSplitView(value ISplitView)
	MinimumThicknessForInlineSidebars() float64
	SetMinimumThicknessForInlineSidebars(value float64)
}

// An object that manages an array of adjacent child views, and has a split view object for managing dividers between those views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller?language=objc
type SplitViewController struct {
	ViewController
}

func SplitViewControllerFrom(ptr unsafe.Pointer) SplitViewController {
	return SplitViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}

func (sc _SplitViewControllerClass) Alloc() SplitViewController {
	rv := objc.Call[SplitViewController](sc, objc.Sel("alloc"))
	return rv
}

func SplitViewController_Alloc() SplitViewController {
	return SplitViewControllerClass.Alloc()
}

func (sc _SplitViewControllerClass) New() SplitViewController {
	rv := objc.Call[SplitViewController](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSplitViewController() SplitViewController {
	return SplitViewControllerClass.New()
}

func (s_ SplitViewController) Init() SplitViewController {
	rv := objc.Call[SplitViewController](s_, objc.Sel("init"))
	return rv
}

func (s_ SplitViewController) InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) SplitViewController {
	rv := objc.Call[SplitViewController](s_, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, objc.Ptr(nibBundleOrNil))
	return rv
}

// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434481-initwithnibname?language=objc
func SplitViewController_InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) SplitViewController {
	return SplitViewControllerClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
}

// Removes a specified split view item from the split view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388903-removesplitviewitem?language=objc
func (s_ SplitViewController) RemoveSplitViewItem(splitViewItem ISplitViewItem) {
	objc.Call[objc.Void](s_, objc.Sel("removeSplitViewItem:"), objc.Ptr(splitViewItem))
}

// Collapses or expands the first sidebar in the split view controller using an animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388905-togglesidebar?language=objc
func (s_ SplitViewController) ToggleSidebar(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("toggleSidebar:"), sender)
	return rv
}

// Adds a split view item to the end of the array of split view items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388928-addsplitviewitem?language=objc
func (s_ SplitViewController) AddSplitViewItem(splitViewItem ISplitViewItem) {
	objc.Call[objc.Void](s_, objc.Sel("addSplitViewItem:"), objc.Ptr(splitViewItem))
}

// Adds a split view item to the array of split view items at the specified index position. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388886-insertsplitviewitem?language=objc
func (s_ SplitViewController) InsertSplitViewItemAtIndex(splitViewItem ISplitViewItem, index int) {
	objc.Call[objc.Void](s_, objc.Sel("insertSplitViewItem:atIndex:"), objc.Ptr(splitViewItem), index)
}

// Returns a Boolean value that indicates whether to enable the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/2881936-validateuserinterfaceitem?language=objc
func (s_ SplitViewController) ValidateUserInterfaceItem(item PValidatedUserInterfaceItem) bool {
	po0 := objc.WrapAsProtocol("NSValidatedUserInterfaceItem", item)
	rv := objc.Call[bool](s_, objc.Sel("validateUserInterfaceItem:"), po0)
	return rv
}

// Returns a Boolean value that indicates whether to enable the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/2881936-validateuserinterfaceitem?language=objc
func (s_ SplitViewController) ValidateUserInterfaceItemObject(itemObject objc.IObject) bool {
	rv := objc.Call[bool](s_, objc.Sel("validateUserInterfaceItem:"), objc.Ptr(itemObject))
	return rv
}

// Returns the corresponding split view item for the specified child view controller of the split view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388857-splitviewitemforviewcontroller?language=objc
func (s_ SplitViewController) SplitViewItemForViewController(viewController IViewController) SplitViewItem {
	rv := objc.Call[SplitViewItem](s_, objc.Sel("splitViewItemForViewController:"), objc.Ptr(viewController))
	return rv
}

// The array of split view items that correspond to the split view controller’s child view controllers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388876-splitviewitems?language=objc
func (s_ SplitViewController) SplitViewItems() []SplitViewItem {
	rv := objc.Call[[]SplitViewItem](s_, objc.Sel("splitViewItems"))
	return rv
}

// The array of split view items that correspond to the split view controller’s child view controllers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388876-splitviewitems?language=objc
func (s_ SplitViewController) SetSplitViewItems(value []ISplitViewItem) {
	objc.Call[objc.Void](s_, objc.Sel("setSplitViewItems:"), value)
}

// The split view that the split view controller manages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388907-splitview?language=objc
func (s_ SplitViewController) SplitView() SplitView {
	rv := objc.Call[SplitView](s_, objc.Sel("splitView"))
	return rv
}

// The split view that the split view controller manages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388907-splitview?language=objc
func (s_ SplitViewController) SetSplitView(value ISplitView) {
	objc.Call[objc.Void](s_, objc.Sel("setSplitView:"), objc.Ptr(value))
}

// The minimum thickness for a sidebar before it automatically collapses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388863-minimumthicknessforinlinesidebar?language=objc
func (s_ SplitViewController) MinimumThicknessForInlineSidebars() float64 {
	rv := objc.Call[float64](s_, objc.Sel("minimumThicknessForInlineSidebars"))
	return rv
}

// The minimum thickness for a sidebar before it automatically collapses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewcontroller/1388863-minimumthicknessforinlinesidebar?language=objc
func (s_ SplitViewController) SetMinimumThicknessForInlineSidebars(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMinimumThicknessForInlineSidebars:"), value)
}