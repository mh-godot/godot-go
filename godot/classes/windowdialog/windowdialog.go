//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package windowdialog

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

	"github.com/shadowapex/godot-go/godot/classes/popup"
)

/*
   Windowdialog is the base class for all window-based dialogs. It's a by-default toplevel [Control] that draws a window decoration and allows motion and resizing.
*/
type WindowDialog struct {
	popup.Popup
}

func (o *WindowDialog) baseClass() string {
	return "WindowDialog"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *WindowDialog) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *WindowDialog) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *WindowDialog) X_Closed() {
	log.Println("Calling WindowDialog.X_Closed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_closed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *WindowDialog) X_GuiInput(arg0 *InputEvent) {
	log.Println("Calling WindowDialog.X_GuiInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_gui_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the close [TextureButton].
*/
func (o *WindowDialog) GetCloseButton() *TextureButton {
	log.Println("Calling WindowDialog.GetCloseButton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_close_button", goArguments, "*TextureButton")

	returnValue := goRet.Interface().(*TextureButton)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *WindowDialog) GetResizable() bool {
	log.Println("Calling WindowDialog.GetResizable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_resizable", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the title of the window.
*/
func (o *WindowDialog) GetTitle() string {
	log.Println("Calling WindowDialog.GetTitle()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_title", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *WindowDialog) SetResizable(resizable bool) {
	log.Println("Calling WindowDialog.SetResizable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(resizable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_resizable", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the title of the window.
*/
func (o *WindowDialog) SetTitle(title string) {
	log.Println("Calling WindowDialog.SetTitle()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(title)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_title", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   WindowDialogImplementer is an interface for WindowDialog objects.
*/
type WindowDialogImplementer interface {
	class.Class
}