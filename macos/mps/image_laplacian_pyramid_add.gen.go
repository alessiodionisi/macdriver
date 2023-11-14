// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageLaplacianPyramidAdd] class.
var ImageLaplacianPyramidAddClass = _ImageLaplacianPyramidAddClass{objc.GetClass("MPSImageLaplacianPyramidAdd")}

type _ImageLaplacianPyramidAddClass struct {
	objc.Class
}

// An interface definition for the [ImageLaplacianPyramidAdd] class.
type IImageLaplacianPyramidAdd interface {
	IImageLaplacianPyramid
}

// A filter that convolves an image with an additive Laplacian pyramid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramidadd?language=objc
type ImageLaplacianPyramidAdd struct {
	ImageLaplacianPyramid
}

func ImageLaplacianPyramidAddFrom(ptr unsafe.Pointer) ImageLaplacianPyramidAdd {
	return ImageLaplacianPyramidAdd{
		ImageLaplacianPyramid: ImageLaplacianPyramidFrom(ptr),
	}
}

func (ic _ImageLaplacianPyramidAddClass) Alloc() ImageLaplacianPyramidAdd {
	rv := objc.Call[ImageLaplacianPyramidAdd](ic, objc.Sel("alloc"))
	return rv
}

func (ic _ImageLaplacianPyramidAddClass) New() ImageLaplacianPyramidAdd {
	rv := objc.Call[ImageLaplacianPyramidAdd](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageLaplacianPyramidAdd() ImageLaplacianPyramidAdd {
	return ImageLaplacianPyramidAddClass.New()
}

func (i_ ImageLaplacianPyramidAdd) Init() ImageLaplacianPyramidAdd {
	rv := objc.Call[ImageLaplacianPyramidAdd](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageLaplacianPyramidAdd) InitWithDevice(device metal.PDevice) ImageLaplacianPyramidAdd {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramidAdd](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a downwards 5-tap image pyramid with the default filter kernel and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648935-initwithdevice?language=objc
func NewImageLaplacianPyramidAddWithDevice(device metal.PDevice) ImageLaplacianPyramidAdd {
	instance := ImageLaplacianPyramidAddClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageLaplacianPyramidAdd) InitWithDeviceCenterWeight(device metal.PDevice, centerWeight float64) ImageLaplacianPyramidAdd {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramidAdd](i_, objc.Sel("initWithDevice:centerWeight:"), po0, centerWeight)
	return rv
}

// Initialize a downwards 5-tap image pyramid with a central weight parameter and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648889-initwithdevice?language=objc
func NewImageLaplacianPyramidAddWithDeviceCenterWeight(device metal.PDevice, centerWeight float64) ImageLaplacianPyramidAdd {
	instance := ImageLaplacianPyramidAddClass.Alloc().InitWithDeviceCenterWeight(device, centerWeight)
	instance.Autorelease()
	return instance
}

func (i_ ImageLaplacianPyramidAdd) InitWithDeviceKernelWidthKernelHeightWeights(device metal.PDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float64) ImageLaplacianPyramidAdd {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramidAdd](i_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), po0, kernelWidth, kernelHeight, kernelWeights)
	return rv
}

// Initialize a downwards n-tap image pyramid with a custom filter kernel and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648821-initwithdevice?language=objc
func NewImageLaplacianPyramidAddWithDeviceKernelWidthKernelHeightWeights(device metal.PDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float64) ImageLaplacianPyramidAdd {
	instance := ImageLaplacianPyramidAddClass.Alloc().InitWithDeviceKernelWidthKernelHeightWeights(device, kernelWidth, kernelHeight, kernelWeights)
	instance.Autorelease()
	return instance
}

func (i_ ImageLaplacianPyramidAdd) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLaplacianPyramidAdd {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramidAdd](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageLaplacianPyramidAdd_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLaplacianPyramidAdd {
	instance := ImageLaplacianPyramidAddClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}