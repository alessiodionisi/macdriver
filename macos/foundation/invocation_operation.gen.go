// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [InvocationOperation] class.
var InvocationOperationClass = _InvocationOperationClass{objc.GetClass("NSInvocationOperation")}

type _InvocationOperationClass struct {
	objc.Class
}

// An interface definition for the [InvocationOperation] class.
type IInvocationOperation interface {
	IOperation
	Result() objc.Object
	Invocation() Invocation
}

// An operation that manages the execution of a single encapsulated task specified as an invocation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocationoperation?language=objc
type InvocationOperation struct {
	Operation
}

func InvocationOperationFrom(ptr unsafe.Pointer) InvocationOperation {
	return InvocationOperation{
		Operation: OperationFrom(ptr),
	}
}

func (i_ InvocationOperation) InitWithInvocation(inv IInvocation) InvocationOperation {
	rv := objc.Call[InvocationOperation](i_, objc.Sel("initWithInvocation:"), objc.Ptr(inv))
	return rv
}

// Returns an NSInvocationOperation object initialized with the specified invocation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocationoperation/1543647-initwithinvocation?language=objc
func InvocationOperation_InitWithInvocation(inv IInvocation) InvocationOperation {
	return InvocationOperationClass.Alloc().InitWithInvocation(inv)
}

func (i_ InvocationOperation) InitWithTargetSelectorObject(target objc.IObject, sel objc.Selector, arg objc.IObject) InvocationOperation {
	rv := objc.Call[InvocationOperation](i_, objc.Sel("initWithTarget:selector:object:"), target, sel, arg)
	return rv
}

// Returns an NSInvocationOperation object initialized with the specified target and selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocationoperation/1543653-initwithtarget?language=objc
func InvocationOperation_InitWithTargetSelectorObject(target objc.IObject, sel objc.Selector, arg objc.IObject) InvocationOperation {
	return InvocationOperationClass.Alloc().InitWithTargetSelectorObject(target, sel, arg)
}

func (ic _InvocationOperationClass) Alloc() InvocationOperation {
	rv := objc.Call[InvocationOperation](ic, objc.Sel("alloc"))
	return rv
}

func InvocationOperation_Alloc() InvocationOperation {
	return InvocationOperationClass.Alloc()
}

func (ic _InvocationOperationClass) New() InvocationOperation {
	rv := objc.Call[InvocationOperation](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInvocationOperation() InvocationOperation {
	return InvocationOperationClass.New()
}

func (i_ InvocationOperation) Init() InvocationOperation {
	rv := objc.Call[InvocationOperation](i_, objc.Sel("init"))
	return rv
}

// The result of the invocation or method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocationoperation/1543615-result?language=objc
func (i_ InvocationOperation) Result() objc.Object {
	rv := objc.Call[objc.Object](i_, objc.Sel("result"))
	return rv
}

// The receiver’s invocation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinvocationoperation/1543655-invocation?language=objc
func (i_ InvocationOperation) Invocation() Invocation {
	rv := objc.Call[Invocation](i_, objc.Sel("invocation"))
	return rv
}