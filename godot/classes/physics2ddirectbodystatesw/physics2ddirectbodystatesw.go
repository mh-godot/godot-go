//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package physics2ddirectbodystatesw

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

	"github.com/shadowapex/godot-go/godot/classes/physics2ddirectbodystate"
)

/*
   Software implementation of [Physics2DDirectBodyState]. This object exposes no new methods or properties and should not be used, as [Physics2DDirectBodyState] selects the best implementation available.
*/
type Physics2DDirectBodyStateSW struct {
	physics2ddirectbodystate.Physics2DDirectBodyState
}

func (o *Physics2DDirectBodyStateSW) baseClass() string {
	return "Physics2DDirectBodyStateSW"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Physics2DDirectBodyStateSW) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *Physics2DDirectBodyStateSW) getOwner() *C.godot_object {
	return o.owner
}

/*
   Physics2DDirectBodyStateSWImplementer is an interface for Physics2DDirectBodyStateSW objects.
*/
type Physics2DDirectBodyStateSWImplementer interface {
	class.Class
}