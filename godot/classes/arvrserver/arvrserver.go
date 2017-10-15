//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package arvrserver

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

	"github.com/shadowapex/godot-go/godot/classes/object"
)

func newSingletonARVRServer() *arvrServer {
	obj := &arvrServer{}
	ptr := C.godot_global_get_singleton(C.CString("ARVRServer"))
	obj.owner = (*C.godot_object)(ptr)
	return obj
}

/*
   The AR/VR Server is the heart of our AR/VR solution and handles all the processing.
*/
var ARVRServer = newSingletonARVRServer()

/*
   The AR/VR Server is the heart of our AR/VR solution and handles all the processing.
*/
type arvrServer struct {
	object.Object
}

func (o *arvrServer) baseClass() string {
	return "ARVRServer"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *arvrServer) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *arvrServer) getOwner() *C.godot_object {
	return o.owner
}

/*
   Mostly exposed for GDNative based interfaces, this is called to register an available interface with the AR/VR server.
*/
func (o *arvrServer) AddInterface(arg0 *ARVRInterface) {
	log.Println("Calling ARVRServer.AddInterface()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_interface", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Find an interface by its name. Say that you're making a game that uses specific capabilities of an AR/VR platform you can find the interface for that platform by name and initialize it.
*/
func (o *arvrServer) FindInterface(name string) *ARVRInterface {
	log.Println("Calling ARVRServer.FindInterface()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "find_interface", goArguments, "*ARVRInterface")

	returnValue := goRet.Interface().(*ARVRInterface)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the interface registered at a given index in our list of interfaces.
*/
func (o *arvrServer) GetInterface(idx int64) *ARVRInterface {
	log.Println("Calling ARVRServer.GetInterface()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_interface", goArguments, "*ARVRInterface")

	returnValue := goRet.Interface().(*ARVRInterface)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the number of interfaces currently registered with the AR/VR server. If you're game supports multiple AR/VR platforms you can look throught the available interface and either present the user with a selection or simply try an initialize each interface and use the first one that returns true.
*/
func (o *arvrServer) GetInterfaceCount() int64 {
	log.Println("Calling ARVRServer.GetInterfaceCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_interface_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Gets our reference frame transform, mostly used internally and exposed for GDNative build interfaces.
*/
func (o *arvrServer) GetReferenceFrame() *Transform {
	log.Println("Calling ARVRServer.GetReferenceFrame()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_reference_frame", goArguments, "*Transform")

	returnValue := goRet.Interface().(*Transform)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the positional tracker at the given ID.
*/
func (o *arvrServer) GetTracker(idx int64) *ARVRPositionalTracker {
	log.Println("Calling ARVRServer.GetTracker()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_tracker", goArguments, "*ARVRPositionalTracker")

	returnValue := goRet.Interface().(*ARVRPositionalTracker)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the number of trackers currently registered.
*/
func (o *arvrServer) GetTrackerCount() int64 {
	log.Println("Calling ARVRServer.GetTrackerCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_tracker_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns our world scale (see ARVROrigin for more information).
*/
func (o *arvrServer) GetWorldScale() float64 {
	log.Println("Calling ARVRServer.GetWorldScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_world_scale", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Removes a registered interface, again exposed mostly for GDNative based interfaces.
*/
func (o *arvrServer) RemoveInterface(arg0 *ARVRInterface) {
	log.Println("Calling ARVRServer.RemoveInterface()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "remove_interface", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *arvrServer) RequestReferenceFrame(ignoreTilt bool, keepHeight bool) {
	log.Println("Calling ARVRServer.RequestReferenceFrame()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(ignoreTilt)
	goArguments[1] = reflect.ValueOf(keepHeight)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "request_reference_frame", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Changes the primary interface to the specified interface. Again mostly exposed for GDNative interfaces.
*/
func (o *arvrServer) SetPrimaryInterface(arg0 *ARVRInterface) {
	log.Println("Calling ARVRServer.SetPrimaryInterface()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_primary_interface", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Changing the world scale, see the ARVROrigin documentation for more information.
*/
func (o *arvrServer) SetWorldScale(arg0 float64) {
	log.Println("Calling ARVRServer.SetWorldScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_world_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}