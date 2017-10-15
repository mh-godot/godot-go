//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package position2d

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

	"github.com/shadowapex/godot-go/godot/classes/node2d"
)

/*
   Generic 2D Position hint for editing. It's just like a plain [Node2D] but displays as a cross in the 2D-Editor at all times.
*/
type Position2D struct {
	node2d.Node2D
}

func (o *Position2D) baseClass() string {
	return "Position2D"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Position2D) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *Position2D) getOwner() *C.godot_object {
	return o.owner
}

/*
   Position2DImplementer is an interface for Position2D objects.
*/
type Position2DImplementer interface {
	class.Class
}