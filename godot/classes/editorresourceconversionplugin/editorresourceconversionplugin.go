//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package editorresourceconversionplugin

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

	"github.com/shadowapex/godot-go/godot/classes/reference"
)

/*

 */
type EditorResourceConversionPlugin struct {
	reference.Reference
}

func (o *EditorResourceConversionPlugin) baseClass() string {
	return "EditorResourceConversionPlugin"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *EditorResourceConversionPlugin) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *EditorResourceConversionPlugin) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *EditorResourceConversionPlugin) X_Convert(resource *Resource) *Resource {
	log.Println("Calling EditorResourceConversionPlugin.X_Convert()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(resource)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_convert", goArguments, "*Resource")

	returnValue := goRet.Interface().(*Resource)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *EditorResourceConversionPlugin) X_ConvertsTo() bool {
	log.Println("Calling EditorResourceConversionPlugin.X_ConvertsTo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_converts_to", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   EditorResourceConversionPluginImplementer is an interface for EditorResourceConversionPlugin objects.
*/
type EditorResourceConversionPluginImplementer interface {
	class.Class
}