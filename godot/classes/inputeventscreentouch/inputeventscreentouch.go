//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package inputeventscreentouch

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

	"github.com/shadowapex/godot-go/godot/classes/inputevent"
)

/*

 */
type InputEventScreenTouch struct {
	inputevent.InputEvent
}

func (o *InputEventScreenTouch) baseClass() string {
	return "InputEventScreenTouch"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *InputEventScreenTouch) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *InputEventScreenTouch) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *InputEventScreenTouch) GetIndex() int64 {
	log.Println("Calling InputEventScreenTouch.GetIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_index", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *InputEventScreenTouch) GetPosition() *Vector2 {
	log.Println("Calling InputEventScreenTouch.GetPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_position", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *InputEventScreenTouch) SetIndex(index int64) {
	log.Println("Calling InputEventScreenTouch.SetIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(index)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_index", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *InputEventScreenTouch) SetPosition(position *Vector2) {
	log.Println("Calling InputEventScreenTouch.SetPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(position)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_position", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *InputEventScreenTouch) SetPressed(pressed bool) {
	log.Println("Calling InputEventScreenTouch.SetPressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(pressed)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_pressed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   InputEventScreenTouchImplementer is an interface for InputEventScreenTouch objects.
*/
type InputEventScreenTouchImplementer interface {
	class.Class
}