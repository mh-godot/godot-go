//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package boxcontainer

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

	"github.com/shadowapex/godot-go/godot/classes/container"
)

/*
   Arranges child controls vertically or horizontally, and rearranges the controls automatically when their minimum size changes.
*/
type BoxContainer struct {
	container.Container
}

func (o *BoxContainer) baseClass() string {
	return "BoxContainer"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *BoxContainer) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *BoxContainer) getOwner() *C.godot_object {
	return o.owner
}

/*
   Adds a control to the box as a spacer. If [code]true[/code], [i]begin[/i] will insert the spacer control in front of other children.
*/
func (o *BoxContainer) AddSpacer(begin bool) {
	log.Println("Calling BoxContainer.AddSpacer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(begin)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_spacer", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the alignment of children in the container.
*/
func (o *BoxContainer) GetAlignment() int64 {
	log.Println("Calling BoxContainer.GetAlignment()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_alignment", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the alignment of children in the container(Must be one of ALIGN_BEGIN, ALIGN_CENTER or ALIGN_END).
*/
func (o *BoxContainer) SetAlignment(alignment int64) {
	log.Println("Calling BoxContainer.SetAlignment()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(alignment)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_alignment", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   BoxContainerImplementer is an interface for BoxContainer objects.
*/
type BoxContainerImplementer interface {
	class.Class
}