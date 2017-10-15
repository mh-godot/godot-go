//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptbuiltinfunc

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

	"github.com/shadowapex/godot-go/godot/classes/visualscriptnode"
)

/*
   A built-in function used inside a [VisualScript]. It is usually a math function or an utility function. See also [@GDScript], for the same functions in the GDScript language.
*/
type VisualScriptBuiltinFunc struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptBuiltinFunc) baseClass() string {
	return "VisualScriptBuiltinFunc"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptBuiltinFunc) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptBuiltinFunc) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *VisualScriptBuiltinFunc) GetFunc() int64 {
	log.Println("Calling VisualScriptBuiltinFunc.GetFunc()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_func", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptBuiltinFunc) SetFunc(which int64) {
	log.Println("Calling VisualScriptBuiltinFunc.SetFunc()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(which)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_func", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptBuiltinFuncImplementer is an interface for VisualScriptBuiltinFunc objects.
*/
type VisualScriptBuiltinFuncImplementer interface {
	class.Class
}