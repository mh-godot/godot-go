//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package navigationpolygon

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
type NavigationPolygon struct {
	resource.Resource
}

func (o *NavigationPolygon) baseClass() string {
	return "NavigationPolygon"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *NavigationPolygon) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *NavigationPolygon) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *NavigationPolygon) X_GetOutlines() *Array {
	log.Println("Calling NavigationPolygon.X_GetOutlines()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_outlines", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *NavigationPolygon) X_GetPolygons() *Array {
	log.Println("Calling NavigationPolygon.X_GetPolygons()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_polygons", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *NavigationPolygon) X_SetOutlines(outlines *Array) {
	log.Println("Calling NavigationPolygon.X_SetOutlines()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(outlines)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_set_outlines", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *NavigationPolygon) X_SetPolygons(polygons *Array) {
	log.Println("Calling NavigationPolygon.X_SetPolygons()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(polygons)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_set_polygons", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NavigationPolygon) AddOutline(outline *PoolVector2Array) {
	log.Println("Calling NavigationPolygon.AddOutline()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(outline)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_outline", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NavigationPolygon) AddOutlineAtIndex(outline *PoolVector2Array, index int64) {
	log.Println("Calling NavigationPolygon.AddOutlineAtIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(outline)
	goArguments[1] = reflect.ValueOf(index)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_outline_at_index", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NavigationPolygon) AddPolygon(polygon *PoolIntArray) {
	log.Println("Calling NavigationPolygon.AddPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(polygon)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_polygon", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NavigationPolygon) ClearOutlines() {
	log.Println("Calling NavigationPolygon.ClearOutlines()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear_outlines", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NavigationPolygon) ClearPolygons() {
	log.Println("Calling NavigationPolygon.ClearPolygons()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear_polygons", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NavigationPolygon) GetOutline(idx int64) *PoolVector2Array {
	log.Println("Calling NavigationPolygon.GetOutline()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_outline", goArguments, "*PoolVector2Array")

	returnValue := goRet.Interface().(*PoolVector2Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NavigationPolygon) GetOutlineCount() int64 {
	log.Println("Calling NavigationPolygon.GetOutlineCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_outline_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NavigationPolygon) GetPolygon(idx int64) *PoolIntArray {
	log.Println("Calling NavigationPolygon.GetPolygon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_polygon", goArguments, "*PoolIntArray")

	returnValue := goRet.Interface().(*PoolIntArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NavigationPolygon) GetPolygonCount() int64 {
	log.Println("Calling NavigationPolygon.GetPolygonCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_polygon_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NavigationPolygon) GetVertices() *PoolVector2Array {
	log.Println("Calling NavigationPolygon.GetVertices()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_vertices", goArguments, "*PoolVector2Array")

	returnValue := goRet.Interface().(*PoolVector2Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *NavigationPolygon) MakePolygonsFromOutlines() {
	log.Println("Calling NavigationPolygon.MakePolygonsFromOutlines()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "make_polygons_from_outlines", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NavigationPolygon) RemoveOutline(idx int64) {
	log.Println("Calling NavigationPolygon.RemoveOutline()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "remove_outline", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NavigationPolygon) SetOutline(idx int64, outline *PoolVector2Array) {
	log.Println("Calling NavigationPolygon.SetOutline()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(outline)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_outline", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *NavigationPolygon) SetVertices(vertices *PoolVector2Array) {
	log.Println("Calling NavigationPolygon.SetVertices()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(vertices)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_vertices", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   NavigationPolygonImplementer is an interface for NavigationPolygon objects.
*/
type NavigationPolygonImplementer interface {
	class.Class
}