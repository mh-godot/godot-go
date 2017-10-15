//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package physicsshapequeryresult

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

import (
	"github.com/shadowapex/godot-go/godot/classes/class"
	"log"
	"reflect"

	"github.com/shadowapex/godot-go/godot/classes/reference"
)

/*

 */
type PhysicsShapeQueryResult struct {
	reference.Reference
}

func (o *PhysicsShapeQueryResult) baseClass() string {
	return "PhysicsShapeQueryResult"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *PhysicsShapeQueryResult) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *PhysicsShapeQueryResult) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultCount() int64 {
	log.Println("Calling PhysicsShapeQueryResult.GetResultCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_result_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultObject(idx int64) *Object {
	log.Println("Calling PhysicsShapeQueryResult.GetResultObject()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_result_object", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultObjectId(idx int64) int64 {
	log.Println("Calling PhysicsShapeQueryResult.GetResultObjectId()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_result_object_id", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultObjectShape(idx int64) int64 {
	log.Println("Calling PhysicsShapeQueryResult.GetResultObjectShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_result_object_shape", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultRid(idx int64) *RID {
	log.Println("Calling PhysicsShapeQueryResult.GetResultRid()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_result_rid", goArguments, "*RID")

	returnValue := goRet.Interface().(*RID)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   PhysicsShapeQueryResultImplementer is an interface for PhysicsShapeQueryResult objects.
*/
type PhysicsShapeQueryResultImplementer interface {
	class.Class
}