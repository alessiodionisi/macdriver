// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewUpdateItem] class.
var CollectionViewUpdateItemClass = _CollectionViewUpdateItemClass{objc.GetClass("NSCollectionViewUpdateItem")}

type _CollectionViewUpdateItemClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewUpdateItem] class.
type ICollectionViewUpdateItem interface {
	objc.IObject
	UpdateAction() CollectionUpdateAction
	IndexPathBeforeUpdate() foundation.IndexPath
	IndexPathAfterUpdate() foundation.IndexPath
}

// A description of a single change to make to an item in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewupdateitem?language=objc
type CollectionViewUpdateItem struct {
	objc.Object
}

func CollectionViewUpdateItemFrom(ptr unsafe.Pointer) CollectionViewUpdateItem {
	return CollectionViewUpdateItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionViewUpdateItemClass) Alloc() CollectionViewUpdateItem {
	rv := objc.Call[CollectionViewUpdateItem](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewUpdateItem_Alloc() CollectionViewUpdateItem {
	return CollectionViewUpdateItemClass.Alloc()
}

func (cc _CollectionViewUpdateItemClass) New() CollectionViewUpdateItem {
	rv := objc.Call[CollectionViewUpdateItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewUpdateItem() CollectionViewUpdateItem {
	return CollectionViewUpdateItemClass.New()
}

func (c_ CollectionViewUpdateItem) Init() CollectionViewUpdateItem {
	rv := objc.Call[CollectionViewUpdateItem](c_, objc.Sel("init"))
	return rv
}

// The action being performed on the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewupdateitem/1534857-updateaction?language=objc
func (c_ CollectionViewUpdateItem) UpdateAction() CollectionUpdateAction {
	rv := objc.Call[CollectionUpdateAction](c_, objc.Sel("updateAction"))
	return rv
}

// The index path of the item before the update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewupdateitem/1534798-indexpathbeforeupdate?language=objc
func (c_ CollectionViewUpdateItem) IndexPathBeforeUpdate() foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](c_, objc.Sel("indexPathBeforeUpdate"))
	return rv
}

// The index path of the item after the update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewupdateitem/1530431-indexpathafterupdate?language=objc
func (c_ CollectionViewUpdateItem) IndexPathAfterUpdate() foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](c_, objc.Sel("indexPathAfterUpdate"))
	return rv
}