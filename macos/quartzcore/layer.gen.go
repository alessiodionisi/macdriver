// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Layer] class.
var LayerClass = _LayerClass{objc.GetClass("CALayer")}

type _LayerClass struct {
	objc.Class
}

// An interface definition for the [Layer] class.
type ILayer interface {
	objc.IObject
	AddConstraint(c IConstraint)
	AddSublayer(layer ILayer)
	ReplaceSublayerWith(oldLayer ILayer, newLayer ILayer)
	RenderInContext(ctx coregraphics.ContextRef)
	LayoutSublayers()
	HitTest(p coregraphics.Point) Layer
	NeedsLayout() bool
	ResizeWithOldSuperlayerSize(size coregraphics.Size)
	ScrollRectToVisible(r coregraphics.Rect)
	ActionForKey(event string) ActionWrapper
	ScrollPoint(p coregraphics.Point)
	ContainsPoint(p coregraphics.Point) bool
	RemoveFromSuperlayer()
	ConvertRectFromLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect
	InsertSublayerAtIndex(layer ILayer, idx int)
	RemoveAllAnimations()
	AnimationKeys() []string
	AffineTransform() coregraphics.AffineTransform
	Display()
	DrawInContext(ctx coregraphics.ContextRef)
	AnimationForKey(key string) Animation
	ConvertTimeFromLayer(t corefoundation.TimeInterval, l ILayer) corefoundation.TimeInterval
	ResizeSublayersWithOldSize(size coregraphics.Size)
	ContentsAreFlipped() bool
	ShouldArchiveValueForKey(key string) bool
	SetAffineTransform(m coregraphics.AffineTransform)
	SetNeedsLayout()
	RemoveAnimationForKey(key string)
	PreferredFrameSize() coregraphics.Size
	SetNeedsDisplayInRect(r coregraphics.Rect)
	SetNeedsDisplay()
	LayoutIfNeeded()
	AddAnimationForKey(anim IAnimation, key string)
	DisplayIfNeeded()
	NeedsDisplay() bool
	ConvertPointToLayer(p coregraphics.Point, l ILayer) coregraphics.Point
	IsGeometryFlipped() bool
	SetGeometryFlipped(value bool)
	Opacity() float64
	SetOpacity(value float64)
	VisibleRect() coregraphics.Rect
	ShadowPath() unsafe.Pointer
	SetShadowPath(value unsafe.Pointer)
	IsHidden() bool
	SetHidden(value bool)
	AnchorPoint() coregraphics.Point
	SetAnchorPoint(value coregraphics.Point)
	Name() string
	SetName(value string)
	ContentsGravity() LayerContentsGravity
	SetContentsGravity(value LayerContentsGravity)
	MinificationFilter() LayerContentsFilter
	SetMinificationFilter(value LayerContentsFilter)
	ContentsScale() float64
	SetContentsScale(value float64)
	ShadowOffset() coregraphics.Size
	SetShadowOffset(value coregraphics.Size)
	ShouldRasterize() bool
	SetShouldRasterize(value bool)
	ZPosition() float64
	SetZPosition(value float64)
	RasterizationScale() float64
	SetRasterizationScale(value float64)
	AllowsGroupOpacity() bool
	SetAllowsGroupOpacity(value bool)
	ContentsRect() coregraphics.Rect
	SetContentsRect(value coregraphics.Rect)
	AnchorPointZ() float64
	SetAnchorPointZ(value float64)
	Style() foundation.Dictionary
	SetStyle(value foundation.Dictionary)
	Sublayers() []Layer
	SetSublayers(value []ILayer)
	CornerRadius() float64
	SetCornerRadius(value float64)
	Bounds() coregraphics.Rect
	SetBounds(value coregraphics.Rect)
	SublayerTransform() Transform3D
	SetSublayerTransform(value Transform3D)
	Delegate() LayerDelegateWrapper
	SetDelegate(value PLayerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ShadowRadius() float64
	SetShadowRadius(value float64)
	Position() coregraphics.Point
	SetPosition(value coregraphics.Point)
	BackgroundColor() coregraphics.ColorRef
	SetBackgroundColor(value coregraphics.ColorRef)
	Superlayer() Layer
	IsOpaque() bool
	SetOpaque(value bool)
	LayoutManager() LayoutManagerWrapper
	SetLayoutManager(value PLayoutManager)
	SetLayoutManagerObject(valueObject objc.IObject)
	DrawsAsynchronously() bool
	SetDrawsAsynchronously(value bool)
	MaskedCorners() CornerMask
	SetMaskedCorners(value CornerMask)
	IsDoubleSided() bool
	SetDoubleSided(value bool)
	Mask() Layer
	SetMask(value ILayer)
	MasksToBounds() bool
	SetMasksToBounds(value bool)
	ShadowOpacity() float64
	SetShadowOpacity(value float64)
	Contents() objc.Object
	SetContents(value objc.IObject)
	Filters() []objc.Object
	SetFilters(value []objc.IObject)
	CornerCurve() LayerCornerCurve
	SetCornerCurve(value LayerCornerCurve)
	Frame() coregraphics.Rect
	SetFrame(value coregraphics.Rect)
	ShadowColor() coregraphics.ColorRef
	SetShadowColor(value coregraphics.ColorRef)
	ContentsFormat() LayerContentsFormat
	SetContentsFormat(value LayerContentsFormat)
	MagnificationFilter() LayerContentsFilter
	SetMagnificationFilter(value LayerContentsFilter)
	CompositingFilter() objc.Object
	SetCompositingFilter(value objc.IObject)
	Transform() Transform3D
	SetTransform(value Transform3D)
	AllowsEdgeAntialiasing() bool
	SetAllowsEdgeAntialiasing(value bool)
	BorderColor() coregraphics.ColorRef
	SetBorderColor(value coregraphics.ColorRef)
	NeedsDisplayOnBoundsChange() bool
	SetNeedsDisplayOnBoundsChange(value bool)
	ContentsCenter() coregraphics.Rect
	SetContentsCenter(value coregraphics.Rect)
	EdgeAntialiasingMask() EdgeAntialiasingMask
	SetEdgeAntialiasingMask(value EdgeAntialiasingMask)
	AutoresizingMask() AutoresizingMask
	SetAutoresizingMask(value AutoresizingMask)
	BackgroundFilters() []objc.Object
	SetBackgroundFilters(value []objc.IObject)
	BorderWidth() float64
	SetBorderWidth(value float64)
	Constraints() []Constraint
	SetConstraints(value []IConstraint)
	Actions() foundation.Dictionary
	SetActions(value foundation.Dictionary)
	MinificationFilterBias() float64
	SetMinificationFilterBias(value float64)
}

// An object that manages image-based content and allows you to perform animations on that content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer?language=objc
type Layer struct {
	objc.Object
}

func LayerFrom(ptr unsafe.Pointer) Layer {
	return Layer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LayerClass) Layer() Layer {
	rv := objc.Call[Layer](lc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func Layer_Layer() Layer {
	return LayerClass.Layer()
}

func (l_ Layer) InitWithLayer(layer objc.IObject) Layer {
	rv := objc.Call[Layer](l_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func Layer_InitWithLayer(layer objc.IObject) Layer {
	return LayerClass.Alloc().InitWithLayer(layer)
}

func (l_ Layer) Init() Layer {
	rv := objc.Call[Layer](l_, objc.Sel("init"))
	return rv
}

func (l_ Layer) ModelLayer() Layer {
	rv := objc.Call[Layer](l_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func Layer_ModelLayer() Layer {
	return LayerClass.Alloc().ModelLayer()
}

func (l_ Layer) PresentationLayer() Layer {
	rv := objc.Call[Layer](l_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func Layer_PresentationLayer() Layer {
	return LayerClass.Alloc().PresentationLayer()
}

func (lc _LayerClass) Alloc() Layer {
	rv := objc.Call[Layer](lc, objc.Sel("alloc"))
	return rv
}

func Layer_Alloc() Layer {
	return LayerClass.Alloc()
}

func (lc _LayerClass) New() Layer {
	rv := objc.Call[Layer](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLayer() Layer {
	return LayerClass.New()
}

// Adds the specified constraint to the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1521899-addconstraint?language=objc
func (l_ Layer) AddConstraint(c IConstraint) {
	objc.Call[objc.Void](l_, objc.Sel("addConstraint:"), objc.Ptr(c))
}

// Appends the layer to the layer’s list of sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410833-addsublayer?language=objc
func (l_ Layer) AddSublayer(layer ILayer) {
	objc.Call[objc.Void](l_, objc.Sel("addSublayer:"), objc.Ptr(layer))
}

// Replaces the specified sublayer with a different layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410820-replacesublayer?language=objc
func (l_ Layer) ReplaceSublayerWith(oldLayer ILayer, newLayer ILayer) {
	objc.Call[objc.Void](l_, objc.Sel("replaceSublayer:with:"), objc.Ptr(oldLayer), objc.Ptr(newLayer))
}

// Renders the layer and its sublayers into the specified context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410909-renderincontext?language=objc
func (l_ Layer) RenderInContext(ctx coregraphics.ContextRef) {
	objc.Call[objc.Void](l_, objc.Sel("renderInContext:"), ctx)
}

// Tells the layer to update its layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410935-layoutsublayers?language=objc
func (l_ Layer) LayoutSublayers() {
	objc.Call[objc.Void](l_, objc.Sel("layoutSublayers"))
}

// Returns the farthest descendant of the receiver in the layer hierarchy (including itself) that contains the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410972-hittest?language=objc
func (l_ Layer) HitTest(p coregraphics.Point) Layer {
	rv := objc.Call[Layer](l_, objc.Sel("hitTest:"), p)
	return rv
}

// Returns a Boolean indicating whether the layer has been marked as needing a layout update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410956-needslayout?language=objc
func (l_ Layer) NeedsLayout() bool {
	rv := objc.Call[bool](l_, objc.Sel("needsLayout"))
	return rv
}

// Informs the receiver that the size of its superlayer changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410894-resizewitholdsuperlayersize?language=objc
func (l_ Layer) ResizeWithOldSuperlayerSize(size coregraphics.Size) {
	objc.Call[objc.Void](l_, objc.Sel("resizeWithOldSuperlayerSize:"), size)
}

// Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified rectangle becomes visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1522139-scrollrecttovisible?language=objc
func (l_ Layer) ScrollRectToVisible(r coregraphics.Rect) {
	objc.Call[objc.Void](l_, objc.Sel("scrollRectToVisible:"), r)
}

// Returns the action object assigned to the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410844-actionforkey?language=objc
func (l_ Layer) ActionForKey(event string) ActionWrapper {
	rv := objc.Call[ActionWrapper](l_, objc.Sel("actionForKey:"), event)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/3152597-cornercurveexpansionfactor?language=objc
func (lc _LayerClass) CornerCurveExpansionFactor(curve LayerCornerCurve) float64 {
	rv := objc.Call[float64](lc, objc.Sel("cornerCurveExpansionFactor:"), curve)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/3152597-cornercurveexpansionfactor?language=objc
func Layer_CornerCurveExpansionFactor(curve LayerCornerCurve) float64 {
	return LayerClass.CornerCurveExpansionFactor(curve)
}

// Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified point lies at the origin of the scroll layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1522202-scrollpoint?language=objc
func (l_ Layer) ScrollPoint(p coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("scrollPoint:"), p)
}

// Returns whether the receiver contains a specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410857-containspoint?language=objc
func (l_ Layer) ContainsPoint(p coregraphics.Point) bool {
	rv := objc.Call[bool](l_, objc.Sel("containsPoint:"), p)
	return rv
}

// Returns a Boolean indicating whether changes to the specified key require the layer to be redisplayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410769-needsdisplayforkey?language=objc
func (lc _LayerClass) NeedsDisplayForKey(key string) bool {
	rv := objc.Call[bool](lc, objc.Sel("needsDisplayForKey:"), key)
	return rv
}

// Returns a Boolean indicating whether changes to the specified key require the layer to be redisplayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410769-needsdisplayforkey?language=objc
func Layer_NeedsDisplayForKey(key string) bool {
	return LayerClass.NeedsDisplayForKey(key)
}

// Detaches the layer from its parent layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410767-removefromsuperlayer?language=objc
func (l_ Layer) RemoveFromSuperlayer() {
	objc.Call[objc.Void](l_, objc.Sel("removeFromSuperlayer"))
}

// Converts the rectangle from the specified layer’s coordinate system to the receiver’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410948-convertrect?language=objc
func (l_ Layer) ConvertRectFromLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("convertRect:fromLayer:"), r, objc.Ptr(l))
	return rv
}

// Specifies the default value associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410886-defaultvalueforkey?language=objc
func (lc _LayerClass) DefaultValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](lc, objc.Sel("defaultValueForKey:"), key)
	return rv
}

// Specifies the default value associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410886-defaultvalueforkey?language=objc
func Layer_DefaultValueForKey(key string) objc.Object {
	return LayerClass.DefaultValueForKey(key)
}

// Inserts the specified layer into the receiver’s list of sublayers at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410944-insertsublayer?language=objc
func (l_ Layer) InsertSublayerAtIndex(layer ILayer, idx int) {
	objc.Call[objc.Void](l_, objc.Sel("insertSublayer:atIndex:"), objc.Ptr(layer), idx)
}

// Remove all animations attached to the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410810-removeallanimations?language=objc
func (l_ Layer) RemoveAllAnimations() {
	objc.Call[objc.Void](l_, objc.Sel("removeAllAnimations"))
}

// Returns an array of strings that identify the animations currently attached to the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410937-animationkeys?language=objc
func (l_ Layer) AnimationKeys() []string {
	rv := objc.Call[[]string](l_, objc.Sel("animationKeys"))
	return rv
}

// Returns an affine version of the layer’s transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410882-affinetransform?language=objc
func (l_ Layer) AffineTransform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](l_, objc.Sel("affineTransform"))
	return rv
}

// Reloads the content of this layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410926-display?language=objc
func (l_ Layer) Display() {
	objc.Call[objc.Void](l_, objc.Sel("display"))
}

// Draws the layer’s content using the specified graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410757-drawincontext?language=objc
func (l_ Layer) DrawInContext(ctx coregraphics.ContextRef) {
	objc.Call[objc.Void](l_, objc.Sel("drawInContext:"), ctx)
}

// Returns the animation object with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410808-animationforkey?language=objc
func (l_ Layer) AnimationForKey(key string) Animation {
	rv := objc.Call[Animation](l_, objc.Sel("animationForKey:"), key)
	return rv
}

// Converts the time interval from the specified layer’s time space to the receiver’s time space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410821-converttime?language=objc
func (l_ Layer) ConvertTimeFromLayer(t corefoundation.TimeInterval, l ILayer) corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](l_, objc.Sel("convertTime:fromLayer:"), t, objc.Ptr(l))
	return rv
}

// Informs the receiver’s sublayers that the receiver’s size has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410929-resizesublayerswitholdsize?language=objc
func (l_ Layer) ResizeSublayersWithOldSize(size coregraphics.Size) {
	objc.Call[objc.Void](l_, objc.Sel("resizeSublayersWithOldSize:"), size)
}

// Initializes a layer with a remote client ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1522119-layerwithremoteclientid?language=objc
func (lc _LayerClass) LayerWithRemoteClientId(client_id uint32) Layer {
	rv := objc.Call[Layer](lc, objc.Sel("layerWithRemoteClientId:"), client_id)
	return rv
}

// Initializes a layer with a remote client ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1522119-layerwithremoteclientid?language=objc
func Layer_LayerWithRemoteClientId(client_id uint32) Layer {
	return LayerClass.LayerWithRemoteClientId(client_id)
}

// Returns a Boolean indicating whether the layer content is implicitly flipped when rendered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410777-contentsareflipped?language=objc
func (l_ Layer) ContentsAreFlipped() bool {
	rv := objc.Call[bool](l_, objc.Sel("contentsAreFlipped"))
	return rv
}

// Returns a Boolean indicating whether the value of the specified key should be archived. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410753-shouldarchivevalueforkey?language=objc
func (l_ Layer) ShouldArchiveValueForKey(key string) bool {
	rv := objc.Call[bool](l_, objc.Sel("shouldArchiveValueForKey:"), key)
	return rv
}

// Sets the layer’s transform to the specified affine transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410868-setaffinetransform?language=objc
func (l_ Layer) SetAffineTransform(m coregraphics.AffineTransform) {
	objc.Call[objc.Void](l_, objc.Sel("setAffineTransform:"), m)
}

// Invalidates the layer’s layout and marks it as needing an update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410946-setneedslayout?language=objc
func (l_ Layer) SetNeedsLayout() {
	objc.Call[objc.Void](l_, objc.Sel("setNeedsLayout"))
}

// Remove the animation object with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410939-removeanimationforkey?language=objc
func (l_ Layer) RemoveAnimationForKey(key string) {
	objc.Call[objc.Void](l_, objc.Sel("removeAnimationForKey:"), key)
}

// Returns the preferred size of the layer in the coordinate space of its superlayer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410980-preferredframesize?language=objc
func (l_ Layer) PreferredFrameSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](l_, objc.Sel("preferredFrameSize"))
	return rv
}

// Returns the default action for the current class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410954-defaultactionforkey?language=objc
func (lc _LayerClass) DefaultActionForKey(event string) ActionWrapper {
	rv := objc.Call[ActionWrapper](lc, objc.Sel("defaultActionForKey:"), event)
	return rv
}

// Returns the default action for the current class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410954-defaultactionforkey?language=objc
func Layer_DefaultActionForKey(event string) ActionWrapper {
	return LayerClass.DefaultActionForKey(event)
}

// Marks the region within the specified rectangle as needing to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410800-setneedsdisplayinrect?language=objc
func (l_ Layer) SetNeedsDisplayInRect(r coregraphics.Rect) {
	objc.Call[objc.Void](l_, objc.Sel("setNeedsDisplayInRect:"), r)
}

// Marks the layer’s contents as needing to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410855-setneedsdisplay?language=objc
func (l_ Layer) SetNeedsDisplay() {
	objc.Call[objc.Void](l_, objc.Sel("setNeedsDisplay"))
}

// Recalculate the receiver’s layout, if required. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410873-layoutifneeded?language=objc
func (l_ Layer) LayoutIfNeeded() {
	objc.Call[objc.Void](l_, objc.Sel("layoutIfNeeded"))
}

// Add the specified animation object to the layer’s render tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410848-addanimation?language=objc
func (l_ Layer) AddAnimationForKey(anim IAnimation, key string) {
	objc.Call[objc.Void](l_, objc.Sel("addAnimation:forKey:"), objc.Ptr(anim), key)
}

// Initiates the update process for a layer if it is currently marked as needing an update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410813-displayifneeded?language=objc
func (l_ Layer) DisplayIfNeeded() {
	objc.Call[objc.Void](l_, objc.Sel("displayIfNeeded"))
}

// Returns a Boolean indicating whether the layer has been marked as needing an update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410958-needsdisplay?language=objc
func (l_ Layer) NeedsDisplay() bool {
	rv := objc.Call[bool](l_, objc.Sel("needsDisplay"))
	return rv
}

// Converts the point from the receiver’s coordinate system to the specified layer’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410881-convertpoint?language=objc
func (l_ Layer) ConvertPointToLayer(p coregraphics.Point, l ILayer) coregraphics.Point {
	rv := objc.Call[coregraphics.Point](l_, objc.Sel("convertPoint:toLayer:"), p, objc.Ptr(l))
	return rv
}

// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410960-geometryflipped?language=objc
func (l_ Layer) IsGeometryFlipped() bool {
	rv := objc.Call[bool](l_, objc.Sel("isGeometryFlipped"))
	return rv
}

// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410960-geometryflipped?language=objc
func (l_ Layer) SetGeometryFlipped(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setGeometryFlipped:"), value)
}

// The opacity of the receiver. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410933-opacity?language=objc
func (l_ Layer) Opacity() float64 {
	rv := objc.Call[float64](l_, objc.Sel("opacity"))
	return rv
}

// The opacity of the receiver. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410933-opacity?language=objc
func (l_ Layer) SetOpacity(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setOpacity:"), value)
}

// The visible region of the layer in its own coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1521892-visiblerect?language=objc
func (l_ Layer) VisibleRect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("visibleRect"))
	return rv
}

// The shape of the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410771-shadowpath?language=objc
func (l_ Layer) ShadowPath() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](l_, objc.Sel("shadowPath"))
	return rv
}

// The shape of the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410771-shadowpath?language=objc
func (l_ Layer) SetShadowPath(value unsafe.Pointer) {
	objc.Call[objc.Void](l_, objc.Sel("setShadowPath:"), value)
}

// A Boolean indicating whether the layer is displayed. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410838-hidden?language=objc
func (l_ Layer) IsHidden() bool {
	rv := objc.Call[bool](l_, objc.Sel("isHidden"))
	return rv
}

// A Boolean indicating whether the layer is displayed. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410838-hidden?language=objc
func (l_ Layer) SetHidden(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setHidden:"), value)
}

// Defines the anchor point of the layer's bounds rectangle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410817-anchorpoint?language=objc
func (l_ Layer) AnchorPoint() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](l_, objc.Sel("anchorPoint"))
	return rv
}

// Defines the anchor point of the layer's bounds rectangle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410817-anchorpoint?language=objc
func (l_ Layer) SetAnchorPoint(value coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("setAnchorPoint:"), value)
}

// The name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410879-name?language=objc
func (l_ Layer) Name() string {
	rv := objc.Call[string](l_, objc.Sel("name"))
	return rv
}

// The name of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410879-name?language=objc
func (l_ Layer) SetName(value string) {
	objc.Call[objc.Void](l_, objc.Sel("setName:"), value)
}

// A constant that specifies how the layer's contents are positioned or scaled within its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410872-contentsgravity?language=objc
func (l_ Layer) ContentsGravity() LayerContentsGravity {
	rv := objc.Call[LayerContentsGravity](l_, objc.Sel("contentsGravity"))
	return rv
}

// A constant that specifies how the layer's contents are positioned or scaled within its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410872-contentsgravity?language=objc
func (l_ Layer) SetContentsGravity(value LayerContentsGravity) {
	objc.Call[objc.Void](l_, objc.Sel("setContentsGravity:"), value)
}

// The filter used when reducing the size of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410898-minificationfilter?language=objc
func (l_ Layer) MinificationFilter() LayerContentsFilter {
	rv := objc.Call[LayerContentsFilter](l_, objc.Sel("minificationFilter"))
	return rv
}

// The filter used when reducing the size of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410898-minificationfilter?language=objc
func (l_ Layer) SetMinificationFilter(value LayerContentsFilter) {
	objc.Call[objc.Void](l_, objc.Sel("setMinificationFilter:"), value)
}

// The scale factor applied to the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410746-contentsscale?language=objc
func (l_ Layer) ContentsScale() float64 {
	rv := objc.Call[float64](l_, objc.Sel("contentsScale"))
	return rv
}

// The scale factor applied to the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410746-contentsscale?language=objc
func (l_ Layer) SetContentsScale(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setContentsScale:"), value)
}

// The offset (in points) of the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410970-shadowoffset?language=objc
func (l_ Layer) ShadowOffset() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](l_, objc.Sel("shadowOffset"))
	return rv
}

// The offset (in points) of the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410970-shadowoffset?language=objc
func (l_ Layer) SetShadowOffset(value coregraphics.Size) {
	objc.Call[objc.Void](l_, objc.Sel("setShadowOffset:"), value)
}

// A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410905-shouldrasterize?language=objc
func (l_ Layer) ShouldRasterize() bool {
	rv := objc.Call[bool](l_, objc.Sel("shouldRasterize"))
	return rv
}

// A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410905-shouldrasterize?language=objc
func (l_ Layer) SetShouldRasterize(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setShouldRasterize:"), value)
}

// The layer’s position on the z axis. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410884-zposition?language=objc
func (l_ Layer) ZPosition() float64 {
	rv := objc.Call[float64](l_, objc.Sel("zPosition"))
	return rv
}

// The layer’s position on the z axis. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410884-zposition?language=objc
func (l_ Layer) SetZPosition(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setZPosition:"), value)
}

// The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410801-rasterizationscale?language=objc
func (l_ Layer) RasterizationScale() float64 {
	rv := objc.Call[float64](l_, objc.Sel("rasterizationScale"))
	return rv
}

// The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410801-rasterizationscale?language=objc
func (l_ Layer) SetRasterizationScale(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setRasterizationScale:"), value)
}

// A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1621277-allowsgroupopacity?language=objc
func (l_ Layer) AllowsGroupOpacity() bool {
	rv := objc.Call[bool](l_, objc.Sel("allowsGroupOpacity"))
	return rv
}

// A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1621277-allowsgroupopacity?language=objc
func (l_ Layer) SetAllowsGroupOpacity(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setAllowsGroupOpacity:"), value)
}

// The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410866-contentsrect?language=objc
func (l_ Layer) ContentsRect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("contentsRect"))
	return rv
}

// The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410866-contentsrect?language=objc
func (l_ Layer) SetContentsRect(value coregraphics.Rect) {
	objc.Call[objc.Void](l_, objc.Sel("setContentsRect:"), value)
}

// The anchor point for the layer’s position along the z axis. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410796-anchorpointz?language=objc
func (l_ Layer) AnchorPointZ() float64 {
	rv := objc.Call[float64](l_, objc.Sel("anchorPointZ"))
	return rv
}

// The anchor point for the layer’s position along the z axis. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410796-anchorpointz?language=objc
func (l_ Layer) SetAnchorPointZ(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setAnchorPointZ:"), value)
}

// An optional dictionary used to store property values that aren't explicitly defined by the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410875-style?language=objc
func (l_ Layer) Style() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](l_, objc.Sel("style"))
	return rv
}

// An optional dictionary used to store property values that aren't explicitly defined by the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410875-style?language=objc
func (l_ Layer) SetStyle(value foundation.Dictionary) {
	objc.Call[objc.Void](l_, objc.Sel("setStyle:"), value)
}

// An array containing the layer’s sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410802-sublayers?language=objc
func (l_ Layer) Sublayers() []Layer {
	rv := objc.Call[[]Layer](l_, objc.Sel("sublayers"))
	return rv
}

// An array containing the layer’s sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410802-sublayers?language=objc
func (l_ Layer) SetSublayers(value []ILayer) {
	objc.Call[objc.Void](l_, objc.Sel("setSublayers:"), value)
}

// The radius to use when drawing rounded corners for the layer’s background. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410818-cornerradius?language=objc
func (l_ Layer) CornerRadius() float64 {
	rv := objc.Call[float64](l_, objc.Sel("cornerRadius"))
	return rv
}

// The radius to use when drawing rounded corners for the layer’s background. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410818-cornerradius?language=objc
func (l_ Layer) SetCornerRadius(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setCornerRadius:"), value)
}

// The layer’s bounds rectangle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410915-bounds?language=objc
func (l_ Layer) Bounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("bounds"))
	return rv
}

// The layer’s bounds rectangle. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410915-bounds?language=objc
func (l_ Layer) SetBounds(value coregraphics.Rect) {
	objc.Call[objc.Void](l_, objc.Sel("setBounds:"), value)
}

// Specifies the transform to apply to sublayers when rendering. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410888-sublayertransform?language=objc
func (l_ Layer) SublayerTransform() Transform3D {
	rv := objc.Call[Transform3D](l_, objc.Sel("sublayerTransform"))
	return rv
}

// Specifies the transform to apply to sublayers when rendering. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410888-sublayertransform?language=objc
func (l_ Layer) SetSublayerTransform(value Transform3D) {
	objc.Call[objc.Void](l_, objc.Sel("setSublayerTransform:"), value)
}

// The layer’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410984-delegate?language=objc
func (l_ Layer) Delegate() LayerDelegateWrapper {
	rv := objc.Call[LayerDelegateWrapper](l_, objc.Sel("delegate"))
	return rv
}

// The layer’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410984-delegate?language=objc
func (l_ Layer) SetDelegate(value PLayerDelegate) {
	po0 := objc.WrapAsProtocol("CALayerDelegate", value)
	objc.SetAssociatedObject(l_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](l_, objc.Sel("setDelegate:"), po0)
}

// The layer’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410984-delegate?language=objc
func (l_ Layer) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The blur radius (in points) used to render the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410819-shadowradius?language=objc
func (l_ Layer) ShadowRadius() float64 {
	rv := objc.Call[float64](l_, objc.Sel("shadowRadius"))
	return rv
}

// The blur radius (in points) used to render the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410819-shadowradius?language=objc
func (l_ Layer) SetShadowRadius(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setShadowRadius:"), value)
}

// The layer’s position in its superlayer’s coordinate space. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410791-position?language=objc
func (l_ Layer) Position() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](l_, objc.Sel("position"))
	return rv
}

// The layer’s position in its superlayer’s coordinate space. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410791-position?language=objc
func (l_ Layer) SetPosition(value coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("setPosition:"), value)
}

// The background color of the receiver. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410966-backgroundcolor?language=objc
func (l_ Layer) BackgroundColor() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](l_, objc.Sel("backgroundColor"))
	return rv
}

// The background color of the receiver. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410966-backgroundcolor?language=objc
func (l_ Layer) SetBackgroundColor(value coregraphics.ColorRef) {
	objc.Call[objc.Void](l_, objc.Sel("setBackgroundColor:"), value)
}

// The superlayer of the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410761-superlayer?language=objc
func (l_ Layer) Superlayer() Layer {
	rv := objc.Call[Layer](l_, objc.Sel("superlayer"))
	return rv
}

// A Boolean value indicating whether the layer contains completely opaque content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410763-opaque?language=objc
func (l_ Layer) IsOpaque() bool {
	rv := objc.Call[bool](l_, objc.Sel("isOpaque"))
	return rv
}

// A Boolean value indicating whether the layer contains completely opaque content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410763-opaque?language=objc
func (l_ Layer) SetOpaque(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setOpaque:"), value)
}

// The object responsible for laying out the layer’s sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410749-layoutmanager?language=objc
func (l_ Layer) LayoutManager() LayoutManagerWrapper {
	rv := objc.Call[LayoutManagerWrapper](l_, objc.Sel("layoutManager"))
	return rv
}

// The object responsible for laying out the layer’s sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410749-layoutmanager?language=objc
func (l_ Layer) SetLayoutManager(value PLayoutManager) {
	po0 := objc.WrapAsProtocol("CALayoutManager", value)
	objc.Call[objc.Void](l_, objc.Sel("setLayoutManager:"), po0)
}

// The object responsible for laying out the layer’s sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410749-layoutmanager?language=objc
func (l_ Layer) SetLayoutManagerObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setLayoutManager:"), objc.Ptr(valueObject))
}

// A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410974-drawsasynchronously?language=objc
func (l_ Layer) DrawsAsynchronously() bool {
	rv := objc.Call[bool](l_, objc.Sel("drawsAsynchronously"))
	return rv
}

// A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410974-drawsasynchronously?language=objc
func (l_ Layer) SetDrawsAsynchronously(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setDrawsAsynchronously:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/2877488-maskedcorners?language=objc
func (l_ Layer) MaskedCorners() CornerMask {
	rv := objc.Call[CornerMask](l_, objc.Sel("maskedCorners"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/2877488-maskedcorners?language=objc
func (l_ Layer) SetMaskedCorners(value CornerMask) {
	objc.Call[objc.Void](l_, objc.Sel("setMaskedCorners:"), value)
}

// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410924-doublesided?language=objc
func (l_ Layer) IsDoubleSided() bool {
	rv := objc.Call[bool](l_, objc.Sel("isDoubleSided"))
	return rv
}

// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410924-doublesided?language=objc
func (l_ Layer) SetDoubleSided(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setDoubleSided:"), value)
}

// An optional layer whose alpha channel is used to mask the layer’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410861-mask?language=objc
func (l_ Layer) Mask() Layer {
	rv := objc.Call[Layer](l_, objc.Sel("mask"))
	return rv
}

// An optional layer whose alpha channel is used to mask the layer’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410861-mask?language=objc
func (l_ Layer) SetMask(value ILayer) {
	objc.Call[objc.Void](l_, objc.Sel("setMask:"), objc.Ptr(value))
}

// A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410896-maskstobounds?language=objc
func (l_ Layer) MasksToBounds() bool {
	rv := objc.Call[bool](l_, objc.Sel("masksToBounds"))
	return rv
}

// A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410896-maskstobounds?language=objc
func (l_ Layer) SetMasksToBounds(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setMasksToBounds:"), value)
}

// The opacity of the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410751-shadowopacity?language=objc
func (l_ Layer) ShadowOpacity() float64 {
	rv := objc.Call[float64](l_, objc.Sel("shadowOpacity"))
	return rv
}

// The opacity of the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410751-shadowopacity?language=objc
func (l_ Layer) SetShadowOpacity(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setShadowOpacity:"), value)
}

// An object that provides the contents of the layer. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410773-contents?language=objc
func (l_ Layer) Contents() objc.Object {
	rv := objc.Call[objc.Object](l_, objc.Sel("contents"))
	return rv
}

// An object that provides the contents of the layer. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410773-contents?language=objc
func (l_ Layer) SetContents(value objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setContents:"), value)
}

// An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410901-filters?language=objc
func (l_ Layer) Filters() []objc.Object {
	rv := objc.Call[[]objc.Object](l_, objc.Sel("filters"))
	return rv
}

// An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410901-filters?language=objc
func (l_ Layer) SetFilters(value []objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setFilters:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/3152596-cornercurve?language=objc
func (l_ Layer) CornerCurve() LayerCornerCurve {
	rv := objc.Call[LayerCornerCurve](l_, objc.Sel("cornerCurve"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/3152596-cornercurve?language=objc
func (l_ Layer) SetCornerCurve(value LayerCornerCurve) {
	objc.Call[objc.Void](l_, objc.Sel("setCornerCurve:"), value)
}

// The layer’s frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410779-frame?language=objc
func (l_ Layer) Frame() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("frame"))
	return rv
}

// The layer’s frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410779-frame?language=objc
func (l_ Layer) SetFrame(value coregraphics.Rect) {
	objc.Call[objc.Void](l_, objc.Sel("setFrame:"), value)
}

// The color of the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410829-shadowcolor?language=objc
func (l_ Layer) ShadowColor() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](l_, objc.Sel("shadowColor"))
	return rv
}

// The color of the layer’s shadow. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410829-shadowcolor?language=objc
func (l_ Layer) SetShadowColor(value coregraphics.ColorRef) {
	objc.Call[objc.Void](l_, objc.Sel("setShadowColor:"), value)
}

// A hint for the desired storage format of the layer contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1792104-contentsformat?language=objc
func (l_ Layer) ContentsFormat() LayerContentsFormat {
	rv := objc.Call[LayerContentsFormat](l_, objc.Sel("contentsFormat"))
	return rv
}

// A hint for the desired storage format of the layer contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1792104-contentsformat?language=objc
func (l_ Layer) SetContentsFormat(value LayerContentsFormat) {
	objc.Call[objc.Void](l_, objc.Sel("setContentsFormat:"), value)
}

// The filter used when increasing the size of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410907-magnificationfilter?language=objc
func (l_ Layer) MagnificationFilter() LayerContentsFilter {
	rv := objc.Call[LayerContentsFilter](l_, objc.Sel("magnificationFilter"))
	return rv
}

// The filter used when increasing the size of the content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410907-magnificationfilter?language=objc
func (l_ Layer) SetMagnificationFilter(value LayerContentsFilter) {
	objc.Call[objc.Void](l_, objc.Sel("setMagnificationFilter:"), value)
}

// A CoreImage filter used to composite the layer and the content behind it. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410748-compositingfilter?language=objc
func (l_ Layer) CompositingFilter() objc.Object {
	rv := objc.Call[objc.Object](l_, objc.Sel("compositingFilter"))
	return rv
}

// A CoreImage filter used to composite the layer and the content behind it. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410748-compositingfilter?language=objc
func (l_ Layer) SetCompositingFilter(value objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setCompositingFilter:"), value)
}

// The transform applied to the layer’s contents. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410836-transform?language=objc
func (l_ Layer) Transform() Transform3D {
	rv := objc.Call[Transform3D](l_, objc.Sel("transform"))
	return rv
}

// The transform applied to the layer’s contents. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410836-transform?language=objc
func (l_ Layer) SetTransform(value Transform3D) {
	objc.Call[objc.Void](l_, objc.Sel("setTransform:"), value)
}

// A Boolean indicating whether the layer is allowed to perform edge antialiasing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1621285-allowsedgeantialiasing?language=objc
func (l_ Layer) AllowsEdgeAntialiasing() bool {
	rv := objc.Call[bool](l_, objc.Sel("allowsEdgeAntialiasing"))
	return rv
}

// A Boolean indicating whether the layer is allowed to perform edge antialiasing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1621285-allowsedgeantialiasing?language=objc
func (l_ Layer) SetAllowsEdgeAntialiasing(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setAllowsEdgeAntialiasing:"), value)
}

// The color of the layer’s border. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410903-bordercolor?language=objc
func (l_ Layer) BorderColor() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](l_, objc.Sel("borderColor"))
	return rv
}

// The color of the layer’s border. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410903-bordercolor?language=objc
func (l_ Layer) SetBorderColor(value coregraphics.ColorRef) {
	objc.Call[objc.Void](l_, objc.Sel("setBorderColor:"), value)
}

// A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410923-needsdisplayonboundschange?language=objc
func (l_ Layer) NeedsDisplayOnBoundsChange() bool {
	rv := objc.Call[bool](l_, objc.Sel("needsDisplayOnBoundsChange"))
	return rv
}

// A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410923-needsdisplayonboundschange?language=objc
func (l_ Layer) SetNeedsDisplayOnBoundsChange(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setNeedsDisplayOnBoundsChange:"), value)
}

// The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410740-contentscenter?language=objc
func (l_ Layer) ContentsCenter() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("contentsCenter"))
	return rv
}

// The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410740-contentscenter?language=objc
func (l_ Layer) SetContentsCenter(value coregraphics.Rect) {
	objc.Call[objc.Void](l_, objc.Sel("setContentsCenter:"), value)
}

// A bitmask defining how the edges of the receiver are rasterized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410892-edgeantialiasingmask?language=objc
func (l_ Layer) EdgeAntialiasingMask() EdgeAntialiasingMask {
	rv := objc.Call[EdgeAntialiasingMask](l_, objc.Sel("edgeAntialiasingMask"))
	return rv
}

// A bitmask defining how the edges of the receiver are rasterized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410892-edgeantialiasingmask?language=objc
func (l_ Layer) SetEdgeAntialiasingMask(value EdgeAntialiasingMask) {
	objc.Call[objc.Void](l_, objc.Sel("setEdgeAntialiasingMask:"), value)
}

// A bitmask defining how the layer is resized when the bounds of its superlayer changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410877-autoresizingmask?language=objc
func (l_ Layer) AutoresizingMask() AutoresizingMask {
	rv := objc.Call[AutoresizingMask](l_, objc.Sel("autoresizingMask"))
	return rv
}

// A bitmask defining how the layer is resized when the bounds of its superlayer changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410877-autoresizingmask?language=objc
func (l_ Layer) SetAutoresizingMask(value AutoresizingMask) {
	objc.Call[objc.Void](l_, objc.Sel("setAutoresizingMask:"), value)
}

// An array of Core Image filters to apply to the content immediately behind the layer. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410827-backgroundfilters?language=objc
func (l_ Layer) BackgroundFilters() []objc.Object {
	rv := objc.Call[[]objc.Object](l_, objc.Sel("backgroundFilters"))
	return rv
}

// An array of Core Image filters to apply to the content immediately behind the layer. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410827-backgroundfilters?language=objc
func (l_ Layer) SetBackgroundFilters(value []objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setBackgroundFilters:"), value)
}

// The width of the layer’s border. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410917-borderwidth?language=objc
func (l_ Layer) BorderWidth() float64 {
	rv := objc.Call[float64](l_, objc.Sel("borderWidth"))
	return rv
}

// The width of the layer’s border. Animatable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410917-borderwidth?language=objc
func (l_ Layer) SetBorderWidth(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setBorderWidth:"), value)
}

// The constraints used to position current layer’s sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1521906-constraints?language=objc
func (l_ Layer) Constraints() []Constraint {
	rv := objc.Call[[]Constraint](l_, objc.Sel("constraints"))
	return rv
}

// The constraints used to position current layer’s sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1521906-constraints?language=objc
func (l_ Layer) SetConstraints(value []IConstraint) {
	objc.Call[objc.Void](l_, objc.Sel("setConstraints:"), value)
}

// A dictionary containing layer actions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410789-actions?language=objc
func (l_ Layer) Actions() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](l_, objc.Sel("actions"))
	return rv
}

// A dictionary containing layer actions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410789-actions?language=objc
func (l_ Layer) SetActions(value foundation.Dictionary) {
	objc.Call[objc.Void](l_, objc.Sel("setActions:"), value)
}

// The bias factor used by the minification filter to determine the levels of detail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410775-minificationfilterbias?language=objc
func (l_ Layer) MinificationFilterBias() float64 {
	rv := objc.Call[float64](l_, objc.Sel("minificationFilterBias"))
	return rv
}

// The bias factor used by the minification filter to determine the levels of detail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410775-minificationfilterbias?language=objc
func (l_ Layer) SetMinificationFilterBias(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setMinificationFilterBias:"), value)
}