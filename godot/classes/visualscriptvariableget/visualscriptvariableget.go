//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package visualscriptvariableget

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
   Returns a variable's value. "Var Name" must be supplied, with an optional type. [b]Input Ports:[/b] none [b]Output Ports:[/b] - Data (variant): [code]value[/code]
*/
type VisualScriptVariableGet struct {
	visualscriptnode.VisualScriptNode
}

func (o *VisualScriptVariableGet) baseClass() string {
	return "VisualScriptVariableGet"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptVariableGet) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VisualScriptVariableGet) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *VisualScriptVariableGet) GetVariable() string {
	log.Println("Calling VisualScriptVariableGet.GetVariable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_variable", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VisualScriptVariableGet) SetVariable(name string) {
	log.Println("Calling VisualScriptVariableGet.SetVariable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_variable", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptVariableGetImplementer is an interface for VisualScriptVariableGet objects.
*/
type VisualScriptVariableGetImplementer interface {
	class.Class
}