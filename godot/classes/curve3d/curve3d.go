//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package curve3d

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
   This class describes a Bezier curve in 3D space. It is mainly used to give a shape to a [Path], but can be manually sampled for other purposes. It keeps a cache of precalculated points along the curve, to speed further calculations up.
*/
type Curve3D struct {
	resource.Resource
}

func (o *Curve3D) baseClass() string {
	return "Curve3D"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Curve3D) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *Curve3D) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *Curve3D) X_GetData() *Dictionary {
	log.Println("Calling Curve3D.X_GetData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_data", goArguments, "*Dictionary")

	returnValue := goRet.Interface().(*Dictionary)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Curve3D) X_SetData(arg0 *Dictionary) {
	log.Println("Calling Curve3D.X_SetData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_set_data", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a point to a curve, at "position", with control points "in" and "out". If "at_position" is given, the point is inserted before the point number "at_position", moving that point (and every point after) after the inserted point. If "at_position" is not given, or is an illegal value (at_position <0 or at_position >= [method get_point_count]), the point will be appended at the end of the point list.
*/
func (o *Curve3D) AddPoint(position *Vector3, in *Vector3, out *Vector3, atPosition int64) {
	log.Println("Calling Curve3D.AddPoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(position)
	goArguments[1] = reflect.ValueOf(in)
	goArguments[2] = reflect.ValueOf(out)
	goArguments[3] = reflect.ValueOf(atPosition)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_point", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *Curve3D) ClearPoints() {
	log.Println("Calling Curve3D.ClearPoints()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear_points", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns the distance between two adjacent cached points.
*/
func (o *Curve3D) GetBakeInterval() float64 {
	log.Println("Calling Curve3D.GetBakeInterval()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_bake_interval", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the total length of the curve, based on the cached points. Given enough density (see [method set_bake_interval]), it should be approximate enough.
*/
func (o *Curve3D) GetBakedLength() float64 {
	log.Println("Calling Curve3D.GetBakedLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_baked_length", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the cache of points as a [PoolVector3Array].
*/
func (o *Curve3D) GetBakedPoints() *PoolVector3Array {
	log.Println("Calling Curve3D.GetBakedPoints()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_baked_points", goArguments, "*PoolVector3Array")

	returnValue := goRet.Interface().(*PoolVector3Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the cache of tilts as a [RealArray].
*/
func (o *Curve3D) GetBakedTilts() *PoolRealArray {
	log.Println("Calling Curve3D.GetBakedTilts()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_baked_tilts", goArguments, "*PoolRealArray")

	returnValue := goRet.Interface().(*PoolRealArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the number of points describing the curve.
*/
func (o *Curve3D) GetPointCount() int64 {
	log.Println("Calling Curve3D.GetPointCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_point_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the position of the control point leading to the vertex "idx". If the index is out of bounds, the function sends an error to the console, and returns (0, 0, 0).
*/
func (o *Curve3D) GetPointIn(idx int64) *Vector3 {
	log.Println("Calling Curve3D.GetPointIn()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_point_in", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the position of the control point leading out of the vertex "idx". If the index is out of bounds, the function sends an error to the console, and returns (0, 0, 0).
*/
func (o *Curve3D) GetPointOut(idx int64) *Vector3 {
	log.Println("Calling Curve3D.GetPointOut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_point_out", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the position of the vertex "idx". If the index is out of bounds, the function sends an error to the console, and returns (0, 0, 0).
*/
func (o *Curve3D) GetPointPosition(idx int64) *Vector3 {
	log.Println("Calling Curve3D.GetPointPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_point_position", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the tilt angle in radians for the point "idx". If the index is out of bounds, the function sends an error to the console, and returns 0.
*/
func (o *Curve3D) GetPointTilt(idx int64) float64 {
	log.Println("Calling Curve3D.GetPointTilt()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_point_tilt", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the position between the vertex "idx" and the vertex "idx"+1, where "t" controls if the point is the first vertex (t = 0.0), the last vertex (t = 1.0), or in between. Values of "t" outside the range (0.0 >= t <=1) give strange, but predictable results. If "idx" is out of bounds it is truncated to the first or last vertex, and "t" is ignored. If the curve has no points, the function sends an error to the console, and returns (0, 0, 0).
*/
func (o *Curve3D) Interpolate(idx int64, t float64) *Vector3 {
	log.Println("Calling Curve3D.Interpolate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(t)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "interpolate", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a point within the curve at position "offset", where "offset" is measured as a distance in 3D units along the curve. To do that, it finds the two cached points where the "offset" lies between, then interpolates the values. This interpolation is cubic if "cubic" is set to true, or linear if set to false. Cubic interpolation tends to follow the curves better, but linear is faster (and often, precise enough).
*/
func (o *Curve3D) InterpolateBaked(offset float64, cubic bool) *Vector3 {
	log.Println("Calling Curve3D.InterpolateBaked()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(offset)
	goArguments[1] = reflect.ValueOf(cubic)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "interpolate_baked", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the position at the vertex "fofs". It calls [method interpolate] using the integer part of fofs as "idx", and its fractional part as "t".
*/
func (o *Curve3D) Interpolatef(fofs float64) *Vector3 {
	log.Println("Calling Curve3D.Interpolatef()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(fofs)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "interpolatef", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Deletes the point "idx" from the curve. Sends an error to the console if "idx" is out of bounds.
*/
func (o *Curve3D) RemovePoint(idx int64) {
	log.Println("Calling Curve3D.RemovePoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "remove_point", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the distance in 3D units between two adjacent cached points. Changing it forces the cache to be recomputed the next time a xxx_baked_xxx function is called. The less distance, the more points the cache will have, and the more memory it will consume, so use with care.
*/
func (o *Curve3D) SetBakeInterval(distance float64) {
	log.Println("Calling Curve3D.SetBakeInterval()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(distance)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_bake_interval", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the position of the control point leading to the vertex "idx". If the index is out of bounds, the function sends an error to the console.
*/
func (o *Curve3D) SetPointIn(idx int64, position *Vector3) {
	log.Println("Calling Curve3D.SetPointIn()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(position)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_point_in", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the position of the control point leading out of the vertex "idx". If the index is out of bounds, the function sends an error to the console.
*/
func (o *Curve3D) SetPointOut(idx int64, position *Vector3) {
	log.Println("Calling Curve3D.SetPointOut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(position)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_point_out", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the position for the vertex "idx". If the index is out of bounds, the function sends an error to the console.
*/
func (o *Curve3D) SetPointPosition(idx int64, position *Vector3) {
	log.Println("Calling Curve3D.SetPointPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(position)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_point_position", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the tilt angle in radians for the point "idx". If the index is out of bounds, the function sends an error to the console. The tilt controls the rotation along the look-at axis an object traveling the path would have. In the case of a curve controlling a [PathFollow], this tilt is an offset over the natural tilt the PathFollow calculates.
*/
func (o *Curve3D) SetPointTilt(idx int64, tilt float64) {
	log.Println("Calling Curve3D.SetPointTilt()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(tilt)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_point_tilt", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns a list of points along the curve, with a curvature controlled point density. That is, the curvier parts will have more points than the straighter parts. This approximation makes straight segments between each point, then subdivides those segments until the resulting shape is similar enough. "max_stages" controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care! "tolerance_degrees" controls how many degrees the midpoint of a segment may deviate from the real curve, before the segment has to be subdivided.
*/
func (o *Curve3D) Tessellate(maxStages int64, toleranceDegrees float64) *PoolVector3Array {
	log.Println("Calling Curve3D.Tessellate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(maxStages)
	goArguments[1] = reflect.ValueOf(toleranceDegrees)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "tessellate", goArguments, "*PoolVector3Array")

	returnValue := goRet.Interface().(*PoolVector3Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Curve3DImplementer is an interface for Curve3D objects.
*/
type Curve3DImplementer interface {
	class.Class
}