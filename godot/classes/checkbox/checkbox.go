//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package checkbox

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

	"github.com/shadowapex/godot-go/godot/classes/button"
)

/*
   A checkbox allows the user to make a binary choice (choosing only one of two possible options), for example Answer 'yes' or 'no'.
*/
type CheckBox struct {
	button.Button
}

func (o *CheckBox) baseClass() string {
	return "CheckBox"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *CheckBox) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *CheckBox) getOwner() *C.godot_object {
	return o.owner
}

/*
   CheckBoxImplementer is an interface for CheckBox objects.
*/
type CheckBoxImplementer interface {
	class.Class
}