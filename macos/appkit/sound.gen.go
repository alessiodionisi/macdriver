// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Sound] class.
var SoundClass = _SoundClass{objc.GetClass("NSSound")}

type _SoundClass struct {
	objc.Class
}

// An interface definition for the [Sound] class.
type ISound interface {
	objc.IObject
	WriteToPasteboard(pasteboard IPasteboard)
	SetName(string_ SoundName) bool
	Stop() bool
	Pause() bool
	Play() bool
	Resume() bool
	Volume() float64
	SetVolume(value float64)
	PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier
	SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier)
	Name() SoundName
	Loops() bool
	SetLoops(value bool)
	Delegate() SoundDelegateWrapper
	SetDelegate(value PSoundDelegate)
	SetDelegateObject(valueObject objc.IObject)
	IsPlaying() bool
	Duration() foundation.TimeInterval
	CurrentTime() foundation.TimeInterval
	SetCurrentTime(value foundation.TimeInterval)
}

// A simple interface for loading and playing audio files. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound?language=objc
type Sound struct {
	objc.Object
}

func SoundFrom(ptr unsafe.Pointer) Sound {
	return Sound{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ Sound) InitWithPasteboard(pasteboard IPasteboard) Sound {
	rv := objc.Call[Sound](s_, objc.Sel("initWithPasteboard:"), objc.Ptr(pasteboard))
	return rv
}

// Initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by NSSound. NSSound expects the data to have a proper magic number, sound header, and data for the formats it supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477294-initwithpasteboard?language=objc
func Sound_InitWithPasteboard(pasteboard IPasteboard) Sound {
	return SoundClass.Alloc().InitWithPasteboard(pasteboard)
}

func (s_ Sound) InitWithContentsOfFileByReference(path string, byRef bool) Sound {
	rv := objc.Call[Sound](s_, objc.Sel("initWithContentsOfFile:byReference:"), path, byRef)
	return rv
}

// Initializes the receiver with the audio data located at a given filepath. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477274-initwithcontentsoffile?language=objc
func Sound_InitWithContentsOfFileByReference(path string, byRef bool) Sound {
	return SoundClass.Alloc().InitWithContentsOfFileByReference(path, byRef)
}

func (s_ Sound) InitWithData(data []byte) Sound {
	rv := objc.Call[Sound](s_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes the receiver with a given audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477292-initwithdata?language=objc
func Sound_InitWithData(data []byte) Sound {
	return SoundClass.Alloc().InitWithData(data)
}

func (s_ Sound) InitWithContentsOfURLByReference(url foundation.IURL, byRef bool) Sound {
	rv := objc.Call[Sound](s_, objc.Sel("initWithContentsOfURL:byReference:"), objc.Ptr(url), byRef)
	return rv
}

// Initializes the receiver with the audio data located at a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477288-initwithcontentsofurl?language=objc
func Sound_InitWithContentsOfURLByReference(url foundation.IURL, byRef bool) Sound {
	return SoundClass.Alloc().InitWithContentsOfURLByReference(url, byRef)
}

func (sc _SoundClass) Alloc() Sound {
	rv := objc.Call[Sound](sc, objc.Sel("alloc"))
	return rv
}

func Sound_Alloc() Sound {
	return SoundClass.Alloc()
}

func (sc _SoundClass) New() Sound {
	rv := objc.Call[Sound](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSound() Sound {
	return SoundClass.New()
}

func (s_ Sound) Init() Sound {
	rv := objc.Call[Sound](s_, objc.Sel("init"))
	return rv
}

// Writes the receiver’s data to a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477338-writetopasteboard?language=objc
func (s_ Sound) WriteToPasteboard(pasteboard IPasteboard) {
	objc.Call[objc.Void](s_, objc.Sel("writeToPasteboard:"), objc.Ptr(pasteboard))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477286-setname?language=objc
func (s_ Sound) SetName(string_ SoundName) bool {
	rv := objc.Call[bool](s_, objc.Sel("setName:"), string_)
	return rv
}

// Indicates whether the receiver can create an instance of itself from the data in a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477276-caninitwithpasteboard?language=objc
func (sc _SoundClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.Call[bool](sc, objc.Sel("canInitWithPasteboard:"), objc.Ptr(pasteboard))
	return rv
}

// Indicates whether the receiver can create an instance of itself from the data in a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477276-caninitwithpasteboard?language=objc
func Sound_CanInitWithPasteboard(pasteboard IPasteboard) bool {
	return SoundClass.CanInitWithPasteboard(pasteboard)
}

// Concludes audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477282-stop?language=objc
func (s_ Sound) Stop() bool {
	rv := objc.Call[bool](s_, objc.Sel("stop"))
	return rv
}

// Pauses audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477307-pause?language=objc
func (s_ Sound) Pause() bool {
	rv := objc.Call[bool](s_, objc.Sel("pause"))
	return rv
}

// Initiates audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477322-play?language=objc
func (s_ Sound) Play() bool {
	rv := objc.Call[bool](s_, objc.Sel("play"))
	return rv
}

// Resumes audio playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477336-resume?language=objc
func (s_ Sound) Resume() bool {
	rv := objc.Call[bool](s_, objc.Sel("resume"))
	return rv
}

// Returns the NSSound instance associated with a given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477318-soundnamed?language=objc
func (sc _SoundClass) SoundNamed(name SoundName) Sound {
	rv := objc.Call[Sound](sc, objc.Sel("soundNamed:"), name)
	return rv
}

// Returns the NSSound instance associated with a given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477318-soundnamed?language=objc
func Sound_SoundNamed(name SoundName) Sound {
	return SoundClass.SoundNamed(name)
}

// The volume of the sound. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477315-volume?language=objc
func (s_ Sound) Volume() float64 {
	rv := objc.Call[float64](s_, objc.Sel("volume"))
	return rv
}

// The volume of the sound. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477315-volume?language=objc
func (s_ Sound) SetVolume(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setVolume:"), value)
}

// Identifies the sound’s output device [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477284-playbackdeviceidentifier?language=objc
func (s_ Sound) PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier {
	rv := objc.Call[SoundPlaybackDeviceIdentifier](s_, objc.Sel("playbackDeviceIdentifier"))
	return rv
}

// Identifies the sound’s output device [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477284-playbackdeviceidentifier?language=objc
func (s_ Sound) SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier) {
	objc.Call[objc.Void](s_, objc.Sel("setPlaybackDeviceIdentifier:"), value)
}

// The name assigned to the sound. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477296-name?language=objc
func (s_ Sound) Name() SoundName {
	rv := objc.Call[SoundName](s_, objc.Sel("name"))
	return rv
}

// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477311-loops?language=objc
func (s_ Sound) Loops() bool {
	rv := objc.Call[bool](s_, objc.Sel("loops"))
	return rv
}

// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477311-loops?language=objc
func (s_ Sound) SetLoops(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setLoops:"), value)
}

// The sound’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc
func (s_ Sound) Delegate() SoundDelegateWrapper {
	rv := objc.Call[SoundDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// The sound’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc
func (s_ Sound) SetDelegate(value PSoundDelegate) {
	po0 := objc.WrapAsProtocol("NSSoundDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The sound’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477300-delegate?language=objc
func (s_ Sound) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A Boolean that indicates whether the sound is playing its audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477302-playing?language=objc
func (s_ Sound) IsPlaying() bool {
	rv := objc.Call[bool](s_, objc.Sel("isPlaying"))
	return rv
}

// The duration of the sound, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477313-duration?language=objc
func (s_ Sound) Duration() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("duration"))
	return rv
}

// Provides the file types the NSSound class understands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477290-soundunfilteredtypes?language=objc
func (sc _SoundClass) SoundUnfilteredTypes() []string {
	rv := objc.Call[[]string](sc, objc.Sel("soundUnfilteredTypes"))
	return rv
}

// Provides the file types the NSSound class understands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477290-soundunfilteredtypes?language=objc
func Sound_SoundUnfilteredTypes() []string {
	return SoundClass.SoundUnfilteredTypes()
}

// The sound’s playback progress, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477320-currenttime?language=objc
func (s_ Sound) CurrentTime() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](s_, objc.Sel("currentTime"))
	return rv
}

// The sound’s playback progress, in seconds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/1477320-currenttime?language=objc
func (s_ Sound) SetCurrentTime(value foundation.TimeInterval) {
	objc.Call[objc.Void](s_, objc.Sel("setCurrentTime:"), value)
}