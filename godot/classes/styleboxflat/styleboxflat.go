//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package styleboxflat

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

	"github.com/shadowapex/godot-go/godot/classes/stylebox"
)

/*
   This stylebox can be used to achieve all kinds of looks without the need of a texture. Those properties are customizable: - Color - Border width (individual width for each border) - Rounded corners (individual radius for each corner) - Shadow About corner radius: Setting corner radius to high values is allowed. As soon as corners would overlap the stylebox will switch to a relative system. Example: [codeblock] height = 30 corner_radius_top_left = 50 corner_radius_bottom_left = 100 [/codeblock] The relative system now would take the 1:2 ratio of the two left corners to calculate the actual corner width. Both corners added will [b]never[/b] be more than the height. Result: [codeblock] corner_radius_top_left: 10 corner_radius_bottom_left: 20 [/codeblock]
*/
type StyleBoxFlat struct {
	stylebox.StyleBox
}

func (o *StyleBoxFlat) baseClass() string {
	return "StyleBoxFlat"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *StyleBoxFlat) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *StyleBoxFlat) getOwner() *C.godot_object {
	return o.owner
}

/*

 */
func (o *StyleBoxFlat) GetAaSize() int64 {
	log.Println("Calling StyleBoxFlat.GetAaSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_aa_size", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetBgColor() *Color {
	log.Println("Calling StyleBoxFlat.GetBgColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_bg_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetBorderBlend() bool {
	log.Println("Calling StyleBoxFlat.GetBorderBlend()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_border_blend", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetBorderColor() *Color {
	log.Println("Calling StyleBoxFlat.GetBorderColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_border_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetBorderWidth(margin int64) int64 {
	log.Println("Calling StyleBoxFlat.GetBorderWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(margin)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_border_width", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetBorderWidthMin() int64 {
	log.Println("Calling StyleBoxFlat.GetBorderWidthMin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_border_width_min", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetCornerDetail() int64 {
	log.Println("Calling StyleBoxFlat.GetCornerDetail()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_corner_detail", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetCornerRadius(corner int64) int64 {
	log.Println("Calling StyleBoxFlat.GetCornerRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(corner)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_corner_radius", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetExpandMargin(margin int64) float64 {
	log.Println("Calling StyleBoxFlat.GetExpandMargin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(margin)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_expand_margin", goArguments, "float64")

	returnValue := goRet.Interface().(float64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetShadowColor() *Color {
	log.Println("Calling StyleBoxFlat.GetShadowColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_shadow_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetShadowSize() int64 {
	log.Println("Calling StyleBoxFlat.GetShadowSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_shadow_size", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) IsAntiAliased() bool {
	log.Println("Calling StyleBoxFlat.IsAntiAliased()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_anti_aliased", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) IsDrawCenterEnabled() bool {
	log.Println("Calling StyleBoxFlat.IsDrawCenterEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_draw_center_enabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) SetAaSize(size int64) {
	log.Println("Calling StyleBoxFlat.SetAaSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_aa_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetAntiAliased(antiAliased bool) {
	log.Println("Calling StyleBoxFlat.SetAntiAliased()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(antiAliased)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_anti_aliased", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetBgColor(color *Color) {
	log.Println("Calling StyleBoxFlat.SetBgColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_bg_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetBorderBlend(blend bool) {
	log.Println("Calling StyleBoxFlat.SetBorderBlend()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(blend)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_border_blend", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetBorderColor(color *Color) {
	log.Println("Calling StyleBoxFlat.SetBorderColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_border_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetBorderWidth(margin int64, width int64) {
	log.Println("Calling StyleBoxFlat.SetBorderWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(margin)
	goArguments[1] = reflect.ValueOf(width)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_border_width", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetBorderWidthAll(width int64) {
	log.Println("Calling StyleBoxFlat.SetBorderWidthAll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(width)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_border_width_all", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetCornerDetail(detail int64) {
	log.Println("Calling StyleBoxFlat.SetCornerDetail()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(detail)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_corner_detail", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetCornerRadius(corner int64, radius int64) {
	log.Println("Calling StyleBoxFlat.SetCornerRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(corner)
	goArguments[1] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_corner_radius", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetCornerRadiusAll(radius int64) {
	log.Println("Calling StyleBoxFlat.SetCornerRadiusAll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_corner_radius_all", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetCornerRadiusIndividual(radiusTopLeft int64, radiusTopRight int64, radiusBottonRight int64, radiusBottomLeft int64) {
	log.Println("Calling StyleBoxFlat.SetCornerRadiusIndividual()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(radiusTopLeft)
	goArguments[1] = reflect.ValueOf(radiusTopRight)
	goArguments[2] = reflect.ValueOf(radiusBottonRight)
	goArguments[3] = reflect.ValueOf(radiusBottomLeft)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_corner_radius_individual", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetDrawCenter(drawCenter bool) {
	log.Println("Calling StyleBoxFlat.SetDrawCenter()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(drawCenter)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_draw_center", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetExpandMargin(margin int64, size float64) {
	log.Println("Calling StyleBoxFlat.SetExpandMargin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(margin)
	goArguments[1] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_expand_margin", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetExpandMarginAll(size float64) {
	log.Println("Calling StyleBoxFlat.SetExpandMarginAll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_expand_margin_all", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetExpandMarginIndividual(sizeLeft float64, sizeTop float64, sizeRight float64, sizeBottom float64) {
	log.Println("Calling StyleBoxFlat.SetExpandMarginIndividual()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(sizeLeft)
	goArguments[1] = reflect.ValueOf(sizeTop)
	goArguments[2] = reflect.ValueOf(sizeRight)
	goArguments[3] = reflect.ValueOf(sizeBottom)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_expand_margin_individual", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetShadowColor(color *Color) {
	log.Println("Calling StyleBoxFlat.SetShadowColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_shadow_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetShadowSize(size int64) {
	log.Println("Calling StyleBoxFlat.SetShadowSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_shadow_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   StyleBoxFlatImplementer is an interface for StyleBoxFlat objects.
*/
type StyleBoxFlatImplementer interface {
	class.Class
}