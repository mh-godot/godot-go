//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package editorselection

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

/*
   This object manages the SceneTree selection in the editor.
*/
type EditorSelection struct {
	object.Object
}

func (o *EditorSelection) baseClass() string {
	return "EditorSelection"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *EditorSelection) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *EditorSelection) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *EditorSelection) X_NodeRemoved(arg0 *Object) {
	log.Println("Calling EditorSelection.X_NodeRemoved()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_node_removed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Add a node to the selection.
*/
func (o *EditorSelection) AddNode(node *Object) {
	log.Println("Calling EditorSelection.AddNode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(node)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_node", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Clear the selection.
*/
func (o *EditorSelection) Clear() {
	log.Println("Calling EditorSelection.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Get the list of selected nodes.
*/
func (o *EditorSelection) GetSelectedNodes() *Array {
	log.Println("Calling EditorSelection.GetSelectedNodes()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_selected_nodes", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Get the list of selected nodes, optimized for transform operations (ie, moving them, rotating, etc). This list avoids situations where a node is selected and also chid/grandchild.
*/
func (o *EditorSelection) GetTransformableSelectedNodes() *Array {
	log.Println("Calling EditorSelection.GetTransformableSelectedNodes()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_transformable_selected_nodes", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Remove a node from the selection.
*/
func (o *EditorSelection) RemoveNode(node *Object) {
	log.Println("Calling EditorSelection.RemoveNode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(node)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "remove_node", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   EditorSelectionImplementer is an interface for EditorSelection objects.
*/
type EditorSelectionImplementer interface {
	class.Class
}