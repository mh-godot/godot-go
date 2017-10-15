//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package inputeventmousemotion

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

	"github.com/shadowapex/godot-go/godot/classes/inputeventmouse"
)

/*

 */
type InputEventMouseMotion struct {
	inputeventmouse.InputEventMouse
}

func (o *InputEventMouseMotion) baseClass() string {
	return "InputEventMouseMotion"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *InputEventMouseMotion) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *InputEventMouseMotion) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *InputEventMouseMotion) GetRelative() *Vector2 {
	log.Println("Calling InputEventMouseMotion.GetRelative()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_relative", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *InputEventMouseMotion) GetSpeed() *Vector2 {
	log.Println("Calling InputEventMouseMotion.GetSpeed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_speed", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *InputEventMouseMotion) SetRelative(relative *Vector2) {
	log.Println("Calling InputEventMouseMotion.SetRelative()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(relative)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_relative", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *InputEventMouseMotion) SetSpeed(speed *Vector2) {
	log.Println("Calling InputEventMouseMotion.SetSpeed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(speed)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_speed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   InputEventMouseMotionImplementer is an interface for InputEventMouseMotion objects.
*/
type InputEventMouseMotionImplementer interface {
	class.Class
}