// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An encoder that encodes commands that modify resource configurations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder?language=objc
type PResourceStateCommandEncoder interface {
	// optional
	UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions(texture TextureObject, mode SparseTextureMappingMode, regions *Region, mipLevels *uint, slices *uint, numRegions uint)
	HasUpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions() bool

	// optional
	WaitForFence(fence FenceObject)
	HasWaitForFence() bool

	// optional
	UpdateTextureMappingModeIndirectBufferIndirectBufferOffset(texture TextureObject, mode SparseTextureMappingMode, indirectBuffer BufferObject, indirectBufferOffset uint)
	HasUpdateTextureMappingModeIndirectBufferIndirectBufferOffset() bool

	// optional
	UpdateFence(fence FenceObject)
	HasUpdateFence() bool

	// optional
	UpdateTextureMappingModeRegionMipLevelSlice(texture TextureObject, mode SparseTextureMappingMode, region Region, mipLevel uint, slice uint)
	HasUpdateTextureMappingModeRegionMipLevelSlice() bool
}

// ensure impl type implements protocol interface
var _ PResourceStateCommandEncoder = (*ResourceStateCommandEncoderObject)(nil)

// A concrete type for the [PResourceStateCommandEncoder] protocol.
type ResourceStateCommandEncoderObject struct {
	objc.Object
}

func (r_ ResourceStateCommandEncoderObject) HasUpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions() bool {
	return r_.RespondsToSelector(objc.Sel("updateTextureMappings:mode:regions:mipLevels:slices:numRegions:"))
}

// Encodes a command to update memory mappings for multiple regions inside a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder/3043995-updatetexturemappings?language=objc
func (r_ ResourceStateCommandEncoderObject) UpdateTextureMappingsModeRegionsMipLevelsSlicesNumRegions(texture TextureObject, mode SparseTextureMappingMode, regions *Region, mipLevels *uint, slices *uint, numRegions uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](r_, objc.Sel("updateTextureMappings:mode:regions:mipLevels:slices:numRegions:"), po0, mode, regions, mipLevels, slices, numRegions)
}

func (r_ ResourceStateCommandEncoderObject) HasWaitForFence() bool {
	return r_.RespondsToSelector(objc.Sel("waitForFence:"))
}

// Prevents the command encoder from enqueuing further GPU commands until the given fence is reached. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder/3114015-waitforfence?language=objc
func (r_ ResourceStateCommandEncoderObject) WaitForFence(fence FenceObject) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](r_, objc.Sel("waitForFence:"), po0)
}

func (r_ ResourceStateCommandEncoderObject) HasUpdateTextureMappingModeIndirectBufferIndirectBufferOffset() bool {
	return r_.RespondsToSelector(objc.Sel("updateTextureMapping:mode:indirectBuffer:indirectBufferOffset:"))
}

// Encodes a command to update a texture’s memory mappings, specifying the parameters indirectly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder/3043993-updatetexturemapping?language=objc
func (r_ ResourceStateCommandEncoderObject) UpdateTextureMappingModeIndirectBufferIndirectBufferOffset(texture TextureObject, mode SparseTextureMappingMode, indirectBuffer BufferObject, indirectBufferOffset uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	po2 := objc.WrapAsProtocol("MTLBuffer", indirectBuffer)
	objc.Call[objc.Void](r_, objc.Sel("updateTextureMapping:mode:indirectBuffer:indirectBufferOffset:"), po0, mode, po2, indirectBufferOffset)
}

func (r_ ResourceStateCommandEncoderObject) HasUpdateFence() bool {
	return r_.RespondsToSelector(objc.Sel("updateFence:"))
}

// Updates the given fence to capture all GPU work that the command encoder has enqueued up to this  point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder/3114014-updatefence?language=objc
func (r_ ResourceStateCommandEncoderObject) UpdateFence(fence FenceObject) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](r_, objc.Sel("updateFence:"), po0)
}

func (r_ ResourceStateCommandEncoderObject) HasUpdateTextureMappingModeRegionMipLevelSlice() bool {
	return r_.RespondsToSelector(objc.Sel("updateTextureMapping:mode:region:mipLevel:slice:"))
}

// Encodes a command to update the texture mappings for a region in a single texture mipmap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourcestatecommandencoder/3043994-updatetexturemapping?language=objc
func (r_ ResourceStateCommandEncoderObject) UpdateTextureMappingModeRegionMipLevelSlice(texture TextureObject, mode SparseTextureMappingMode, region Region, mipLevel uint, slice uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](r_, objc.Sel("updateTextureMapping:mode:region:mipLevel:slice:"), po0, mode, region, mipLevel, slice)
}