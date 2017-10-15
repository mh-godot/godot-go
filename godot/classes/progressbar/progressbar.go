//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package progressbar

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

	"github.com/shadowapex/godot-go/godot/classes/ranges"
)

/*
   General purpose progress bar. Shows fill percentage from right to left.
*/
type ProgressBar struct {
	ranges.Range
}

func (o *ProgressBar) baseClass() string {
	return "ProgressBar"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *ProgressBar) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *ProgressBar) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *ProgressBar) IsPercentVisible() bool {
	log.Println("Calling ProgressBar.IsPercentVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_percent_visible", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ProgressBar) SetPercentVisible(visible bool) {
	log.Println("Calling ProgressBar.SetPercentVisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(visible)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_percent_visible", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ProgressBarImplementer is an interface for ProgressBar objects.
*/
type ProgressBarImplementer interface {
	class.Class
}