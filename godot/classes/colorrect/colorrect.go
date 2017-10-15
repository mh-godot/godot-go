//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package colorrect

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

	"github.com/shadowapex/godot-go/godot/classes/control"
)

/*
   An object that is represented on the canvas as a rect with color. [Color] is used to set or get color info for the rect.
*/
type ColorRect struct {
	control.Control
}

func (o *ColorRect) baseClass() string {
	return "ColorRect"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *ColorRect) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *ColorRect) getOwner() *C.godot_object {
	return o.owner
}

/*
   Return the color in RGBA format. [codeblock] var cr = get_node("colorrect_node") var c = cr.get_frame_color() # Default color is white [/codeblock]
*/
func (o *ColorRect) GetFrameColor() *Color {
	log.Println("Calling ColorRect.GetFrameColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_frame_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set new color to ColorRect. [codeblock] var cr = get_node("colorrect_node") cr.set_frame_color(Color(1, 0, 0, 1)) # Set color rect node to red [/codeblock]
*/
func (o *ColorRect) SetFrameColor(color *Color) {
	log.Println("Calling ColorRect.SetFrameColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_frame_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ColorRectImplementer is an interface for ColorRect objects.
*/
type ColorRectImplementer interface {
	class.Class
}