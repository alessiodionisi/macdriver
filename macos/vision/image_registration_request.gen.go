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

// The class instance for the [ImageRegistrationRequest] class.
var ImageRegistrationRequestClass = _ImageRegistrationRequestClass{objc.GetClass("VNImageRegistrationRequest")}

type _ImageRegistrationRequestClass struct {
	objc.Class
}

// An interface definition for the [ImageRegistrationRequest] class.
type IImageRegistrationRequest interface {
	ITargetedImageRequest
}

// The abstract superclass for image analysis requests that align images according to their content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnimageregistrationrequest?language=objc
type ImageRegistrationRequest struct {
	TargetedImageRequest
}

func ImageRegistrationRequestFrom(ptr unsafe.Pointer) ImageRegistrationRequest {
	return ImageRegistrationRequest{
		TargetedImageRequest: TargetedImageRequestFrom(ptr),
	}
}

func (ic _ImageRegistrationRequestClass) Alloc() ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](ic, objc.Sel("alloc"))
	return rv
}

func (ic _ImageRegistrationRequestClass) New() ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageRegistrationRequest() ImageRegistrationRequest {
	return ImageRegistrationRequestClass.New()
}

func (i_ ImageRegistrationRequest) Init() ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageRegistrationRequest) InitWithTargetedCGImageOrientationOptions(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	return rv
}

// Creates a new request targeting a Core Graphics image of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923444-initwithtargetedcgimage?language=objc
func NewImageRegistrationRequestWithTargetedCGImageOrientationOptions(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOrientationOptions(cgImage, orientation, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return rv
}

// Creates a new request targeting a Core Graphics image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923452-initwithtargetedcgimage?language=objc
func NewImageRegistrationRequestWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOptions(cgImage, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.IURL, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), objc.Ptr(imageURL), options, completionHandler)
	return rv
}

// Creates a new request targeting an image at the specified URL, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923453-initwithtargetedimageurl?language=objc
func NewImageRegistrationRequestWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.IURL, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOptionsCompletionHandler(imageURL, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return rv
}

// Creates a new request targeting an image as raw data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923460-initwithtargetedimagedata?language=objc
func NewImageRegistrationRequestWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOptions(imageData, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, completionHandler)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923446-initwithtargetedcvpixelbuffer?language=objc
func NewImageRegistrationRequestWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, completionHandler)
	return rv
}

// Creates a new request with a completion handler that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571275-initwithtargetedcmsamplebuffer?language=objc
func NewImageRegistrationRequestWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageDataOrientationOptionsCompletionHandler(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting a raw data image of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923443-initwithtargetedimagedata?language=objc
func NewImageRegistrationRequestWithTargetedImageDataOrientationOptionsCompletionHandler(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOrientationOptionsCompletionHandler(imageData, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCMSampleBufferOrientationOptions(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return rv
}

// Creates a new request that targets an image of a known orientation in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571276-initwithtargetedcmsamplebuffer?language=objc
func NewImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptions(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOrientationOptions(sampleBuffer, orientation, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923445-initwithtargetedcvpixelbuffer?language=objc
func NewImageRegistrationRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOptions(pixelBuffer, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageDataOptionsCompletionHandler(imageData []byte, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, completionHandler)
	return rv
}

// Creates a new request targeting an image as raw data, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923455-initwithtargetedimagedata?language=objc
func NewImageRegistrationRequestWithTargetedImageDataOptionsCompletionHandler(imageData []byte, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOptionsCompletionHandler(imageData, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.IImage, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), objc.Ptr(ciImage), options, completionHandler)
	return rv
}

// Creates a new request targeting a CIImage, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923454-initwithtargetedciimage?language=objc
func NewImageRegistrationRequestWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.IImage, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOptionsCompletionHandler(ciImage, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCIImageOrientationOptions(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCIImage:orientation:options:"), objc.Ptr(ciImage), orientation, options)
	return rv
}

// Creates a new request targeting a CIImage of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923459-initwithtargetedciimage?language=objc
func NewImageRegistrationRequestWithTargetedCIImageOrientationOptions(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOrientationOptions(ciImage, orientation, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageURLOrientationOptions(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageURL:orientation:options:"), objc.Ptr(imageURL), orientation, options)
	return rv
}

// Creates a new request targeting an image of known orientation, at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923456-initwithtargetedimageurl?language=objc
func NewImageRegistrationRequestWithTargetedImageURLOrientationOptions(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOrientationOptions(imageURL, orientation, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), objc.Ptr(ciImage), orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting a CIImage of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923451-initwithtargetedciimage?language=objc
func NewImageRegistrationRequestWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, completionHandler)
	return rv
}

// Creates a new request targeting a Core Graphics image, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923448-initwithtargetedcgimage?language=objc
func NewImageRegistrationRequestWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOptionsCompletionHandler(cgImage, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageDataOrientationOptions(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	return rv
}

// Creates a new request targeting a raw data image of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923441-initwithtargetedimagedata?language=objc
func NewImageRegistrationRequestWithTargetedImageDataOrientationOptions(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOrientationOptions(imageData, orientation, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a new request that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571274-initwithtargetedcmsamplebuffer?language=objc
func NewImageRegistrationRequestWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOptions(sampleBuffer, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), objc.Ptr(imageURL), orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting an image of known orientation, at the specified URL, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923457-initwithtargetedimageurl?language=objc
func NewImageRegistrationRequestWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, completionHandler)
	return rv
}

// Creates a new request with a completion handler that targets an image of a known orientation in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571277-initwithtargetedcmsamplebuffer?language=objc
func NewImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCIImage:options:"), objc.Ptr(ciImage), options)
	return rv
}

// Creates a new request targeting a CIImage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923447-initwithtargetedciimage?language=objc
func NewImageRegistrationRequestWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOptions(ciImage, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting a Core Graphics image of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923450-initwithtargetedcgimage?language=objc
func NewImageRegistrationRequestWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923442-initwithtargetedcvpixelbuffer?language=objc
func NewImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOrientationOptions(pixelBuffer, orientation, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923449-initwithtargetedcvpixelbuffer?language=objc
func NewImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithTargetedImageURL:options:"), objc.Ptr(imageURL), options)
	return rv
}

// Creates a new request targeting an image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923458-initwithtargetedimageurl?language=objc
func NewImageRegistrationRequestWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOptions(imageURL, options)
	instance.Autorelease()
	return instance
}

func (i_ ImageRegistrationRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	rv := objc.Call[ImageRegistrationRequest](i_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewImageRegistrationRequestWithCompletionHandler(completionHandler RequestCompletionHandler) ImageRegistrationRequest {
	instance := ImageRegistrationRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}