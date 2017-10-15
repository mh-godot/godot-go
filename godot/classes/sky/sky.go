//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package sky

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

	"github.com/shadowapex/godot-go/godot/classes/resource"
)

/*

 */
type Sky struct {
	resource.Resource
}

func (o *Sky) baseClass() string {
	return "Sky"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Sky) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *Sky) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *Sky) GetRadianceSize() int64 {
	log.Println("Calling Sky.GetRadianceSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_radiance_size", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *Sky) SetRadianceSize(size int64) {
	log.Println("Calling Sky.SetRadianceSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_radiance_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   SkyImplementer is an interface for Sky objects.
*/
type SkyImplementer interface {
	class.Class
}