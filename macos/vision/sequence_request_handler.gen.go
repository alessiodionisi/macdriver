// Code generated by DarwinKit. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/imageio"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SequenceRequestHandler] class.
var SequenceRequestHandlerClass = _SequenceRequestHandlerClass{objc.GetClass("VNSequenceRequestHandler")}

type _SequenceRequestHandlerClass struct {
	objc.Class
}

// An interface definition for the [SequenceRequestHandler] class.
type ISequenceRequestHandler interface {
	objc.IObject
	PerformRequestsOnCVPixelBufferError(requests []IRequest, pixelBuffer corevideo.PixelBufferRef, error foundation.IError) bool
	PerformRequestsOnImageDataError(requests []IRequest, imageData []byte, error foundation.IError) bool
	PerformRequestsOnCVPixelBufferOrientationError(requests []IRequest, pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool
	PerformRequestsOnCGImageOrientationError(requests []IRequest, image coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool
	PerformRequestsOnCIImageOrientationError(requests []IRequest, image coreimage.IImage, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool
	PerformRequestsOnCMSampleBufferOrientationError(requests []IRequest, sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool
	PerformRequestsOnCGImageError(requests []IRequest, image coregraphics.ImageRef, error foundation.IError) bool
	PerformRequestsOnImageDataOrientationError(requests []IRequest, imageData []byte, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool
	PerformRequestsOnCIImageError(requests []IRequest, image coreimage.IImage, error foundation.IError) bool
	PerformRequestsOnImageURLError(requests []IRequest, imageURL foundation.IURL, error foundation.IError) bool
	PerformRequestsOnImageURLOrientationError(requests []IRequest, imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool
	PerformRequestsOnCMSampleBufferError(requests []IRequest, sampleBuffer coremedia.SampleBufferRef, error foundation.IError) bool
}

// An object that processes image analysis requests for each frame in a sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler?language=objc
type SequenceRequestHandler struct {
	objc.Object
}

func SequenceRequestHandlerFrom(ptr unsafe.Pointer) SequenceRequestHandler {
	return SequenceRequestHandler{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SequenceRequestHandler) Init() SequenceRequestHandler {
	rv := objc.Call[SequenceRequestHandler](s_, objc.Sel("init"))
	return rv
}

func (sc _SequenceRequestHandlerClass) Alloc() SequenceRequestHandler {
	rv := objc.Call[SequenceRequestHandler](sc, objc.Sel("alloc"))
	return rv
}

func (sc _SequenceRequestHandlerClass) New() SequenceRequestHandler {
	rv := objc.Call[SequenceRequestHandler](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSequenceRequestHandler() SequenceRequestHandler {
	return SequenceRequestHandlerClass.New()
}

// Schedules one or more Vision requests to be performed on a Core Video pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2880307-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnCVPixelBufferError(requests []IRequest, pixelBuffer corevideo.PixelBufferRef, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onCVPixelBuffer:error:"), requests, pixelBuffer, objc.Ptr(error))
	return rv
}

// Schedules one or more Vision requests to be performed on raw image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2880302-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnImageDataError(requests []IRequest, imageData []byte, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onImageData:error:"), requests, imageData, objc.Ptr(error))
	return rv
}

// Schedules one or more Vision requests to be performed on a Core Video pixel buffer with known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2880308-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnCVPixelBufferOrientationError(requests []IRequest, pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onCVPixelBuffer:orientation:error:"), requests, pixelBuffer, orientation, objc.Ptr(error))
	return rv
}

// Schedules one or more Vision requests to be performed on a Core Graphics image with known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2880298-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnCGImageOrientationError(requests []IRequest, image coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onCGImage:orientation:error:"), requests, image, orientation, objc.Ptr(error))
	return rv
}

// Schedules one or more Vision requests to be performed on CIImage data with known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2880301-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnCIImageOrientationError(requests []IRequest, image coreimage.IImage, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onCIImage:orientation:error:"), requests, objc.Ptr(image), orientation, objc.Ptr(error))
	return rv
}

// Performs one or more requests on an image of a specified orientation contained within a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/3571273-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnCMSampleBufferOrientationError(requests []IRequest, sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onCMSampleBuffer:orientation:error:"), requests, sampleBuffer, orientation, objc.Ptr(error))
	return rv
}

// Schedules Vision requests to be performed on a Core Graphics image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2880300-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnCGImageError(requests []IRequest, image coregraphics.ImageRef, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onCGImage:error:"), requests, image, objc.Ptr(error))
	return rv
}

// Schedules one or more Vision requests to be performed on raw data containing an image with known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2881930-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnImageDataOrientationError(requests []IRequest, imageData []byte, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onImageData:orientation:error:"), requests, imageData, orientation, objc.Ptr(error))
	return rv
}

// Schedules one or more Vision requests to be performed on  CIImage data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2880305-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnCIImageError(requests []IRequest, image coreimage.IImage, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onCIImage:error:"), requests, objc.Ptr(image), objc.Ptr(error))
	return rv
}

// Schedules one or more Vision requests to be performed on an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2880299-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnImageURLError(requests []IRequest, imageURL foundation.IURL, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onImageURL:error:"), requests, objc.Ptr(imageURL), objc.Ptr(error))
	return rv
}

// Schedules one or more Vision requests to be performed on an image with known orientation, at a specific URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/2881927-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnImageURLOrientationError(requests []IRequest, imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onImageURL:orientation:error:"), requests, objc.Ptr(imageURL), orientation, objc.Ptr(error))
	return rv
}

// Performs one or more requests on an image contained within a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnsequencerequesthandler/3571272-performrequests?language=objc
func (s_ SequenceRequestHandler) PerformRequestsOnCMSampleBufferError(requests []IRequest, sampleBuffer coremedia.SampleBufferRef, error foundation.IError) bool {
	rv := objc.Call[bool](s_, objc.Sel("performRequests:onCMSampleBuffer:error:"), requests, sampleBuffer, objc.Ptr(error))
	return rv
}