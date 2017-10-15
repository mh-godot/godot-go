//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package vehiclebody

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

	"github.com/shadowapex/godot-go/godot/classes/physicsbody"
)

/*

 */
type VehicleBody struct {
	physicsbody.PhysicsBody
}

func (o *VehicleBody) baseClass() string {
	return "VehicleBody"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VehicleBody) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *VehicleBody) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *VehicleBody) X_DirectStateChanged(arg0 *Object) {
	log.Println("Calling VehicleBody.X_DirectStateChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_direct_state_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VehicleBody) GetBrake() float64 {
	log.Println("Calling VehicleBody.GetBrake()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_brake", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VehicleBody) GetEngineForce() float64 {
	log.Println("Calling VehicleBody.GetEngineForce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_engine_force", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VehicleBody) GetFriction() float64 {
	log.Println("Calling VehicleBody.GetFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_friction", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the VehicleBody's velocity vector. To get the absolute speed in scalar value, get the length of the return vector in pixels/second. Example: [codeblock] # vehicle is an instance of VehicleBody var speed = vehicle.get_linear_velocity().length() [/codeblock]
*/
func (o *VehicleBody) GetLinearVelocity() *Vector3 {
	log.Println("Calling VehicleBody.GetLinearVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_linear_velocity", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VehicleBody) GetMass() float64 {
	log.Println("Calling VehicleBody.GetMass()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_mass", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the steering angle (in radians).
*/
func (o *VehicleBody) GetSteering() float64 {
	log.Println("Calling VehicleBody.GetSteering()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_steering", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *VehicleBody) SetBrake(brake float64) {
	log.Println("Calling VehicleBody.SetBrake()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(brake)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_brake", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VehicleBody) SetEngineForce(engineForce float64) {
	log.Println("Calling VehicleBody.SetEngineForce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(engineForce)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_engine_force", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VehicleBody) SetFriction(friction float64) {
	log.Println("Calling VehicleBody.SetFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(friction)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_friction", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *VehicleBody) SetMass(mass float64) {
	log.Println("Calling VehicleBody.SetMass()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mass)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_mass", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the steering angle (in radians).
*/
func (o *VehicleBody) SetSteering(steering float64) {
	log.Println("Calling VehicleBody.SetSteering()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(steering)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_steering", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VehicleBodyImplementer is an interface for VehicleBody objects.
*/
type VehicleBodyImplementer interface {
	class.Class
}