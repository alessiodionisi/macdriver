// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A depth and stencil state object that specifies the depth and stencil configuration and operations used in a render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencilstate?language=objc
type PDepthStencilState interface {
	// optional
	Label() string
	HasLabel() bool

	// optional
	Device() DeviceObject
	HasDevice() bool
}

// ensure impl type implements protocol interface
var _ PDepthStencilState = (*DepthStencilStateObject)(nil)

// A concrete type for the [PDepthStencilState] protocol.
type DepthStencilStateObject struct {
	objc.Object
}

func (d_ DepthStencilStateObject) HasLabel() bool {
	return d_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies this object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencilstate/1462472-label?language=objc
func (d_ DepthStencilStateObject) Label() string {
	rv := objc.Call[string](d_, objc.Sel("label"))
	return rv
}

func (d_ DepthStencilStateObject) HasDevice() bool {
	return d_.RespondsToSelector(objc.Sel("device"))
}

// The device from which this state object was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencilstate/1462453-device?language=objc
func (d_ DepthStencilStateObject) Device() DeviceObject {
	rv := objc.Call[DeviceObject](d_, objc.Sel("device"))
	return rv
}