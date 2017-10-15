//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package staticbody

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
   Static body for 3D Physics. A static body is a simple body that is not intended to move. They don't consume any CPU resources in contrast to a [RigidBody3D] so they are great for scenario collision. A static body can also be animated by using simulated motion mode. This is useful for implementing functionalities such as moving platforms. When this mode is active the body can be animated and automatically computes linear and angular velocity to apply in that frame and to influence other bodies. Alternatively, a constant linear or angular velocity can be set for the static body, so even if it doesn't move, it affects other bodies as if it was moving (this is useful for simulating conveyor belts or conveyor wheels).
*/
type StaticBody struct {
	physicsbody.PhysicsBody
}

func (o *StaticBody) baseClass() string {
	return "StaticBody"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *StaticBody) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *StaticBody) getOwner() *C.godot_object {
	return o.owner
}

/*
   Return the body bounciness.
*/
func (o *StaticBody) GetBounce() float64 {
	log.Println("Calling StaticBody.GetBounce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_bounce", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the constant angular velocity for the body.
*/
func (o *StaticBody) GetConstantAngularVelocity() *Vector3 {
	log.Println("Calling StaticBody.GetConstantAngularVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_constant_angular_velocity", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the constant linear velocity for the body.
*/
func (o *StaticBody) GetConstantLinearVelocity() *Vector3 {
	log.Println("Calling StaticBody.GetConstantLinearVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_constant_linear_velocity", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the body friction.
*/
func (o *StaticBody) GetFriction() float64 {
	log.Println("Calling StaticBody.GetFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_friction", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Set the body bounciness, from 0 (not bouncy) to 1 (bouncy).
*/
func (o *StaticBody) SetBounce(bounce float64) {
	log.Println("Calling StaticBody.SetBounce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bounce)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_bounce", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set a constant angular velocity for the body. This does not rotate the body, but affects other bodies touching it, as if it was rotating.
*/
func (o *StaticBody) SetConstantAngularVelocity(vel *Vector3) {
	log.Println("Calling StaticBody.SetConstantAngularVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(vel)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_constant_angular_velocity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set a constant linear velocity for the body. This does not move the body, but affects other bodies touching it, as if it was moving.
*/
func (o *StaticBody) SetConstantLinearVelocity(vel *Vector3) {
	log.Println("Calling StaticBody.SetConstantLinearVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(vel)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_constant_linear_velocity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the body friction, from 0 (frictionless) to 1 (full friction).
*/
func (o *StaticBody) SetFriction(friction float64) {
	log.Println("Calling StaticBody.SetFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(friction)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_friction", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   StaticBodyImplementer is an interface for StaticBody objects.
*/
type StaticBodyImplementer interface {
	class.Class
}