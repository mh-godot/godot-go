//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptselect

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
   Chooses between two input values based on a Boolean condition. [b]Input Ports:[/b] - Data (boolean): [code]cond[/code] - Data (variant): [code]a[/code] - Data (variant): [code]b[/code] [b]Output Ports:[/b] - Data (variant): [code]out[/code]
*/
type VisualScriptSelect struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptSelect) baseClass() string {
	return "VisualScriptSelect"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptSelect) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptSelect) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *VisualScriptSelect) GetTyped() int64 {
	log.Println("Calling VisualScriptSelect.GetTyped()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_typed", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptSelect) SetTyped(aType int64) {
	log.Println("Calling VisualScriptSelect.SetTyped()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(aType)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_typed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptSelectImplementer is an interface for VisualScriptSelect objects.
*/
type VisualScriptSelectImplementer interface {
	class.Class
}