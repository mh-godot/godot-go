//------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "templates/*.go.template" so they can be included in the generated
//   code.
//------------------------------------------------------------------------------

package popupmenu

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
   PopupMenu is the typical Control that displays a list of options. They are popular in toolbars or context menus.
*/
type PopupMenu struct {
	popup.Popup
}

func (o *PopupMenu) baseClass() string {
	return "PopupMenu"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *PopupMenu) setOwner(object *C.godot_object) {
	o.owner = object
}

func (o *PopupMenu) getOwner() *C.godot_object {
	return o.owner
}

/*
   Undocumented
*/
func (o *PopupMenu) X_GetItems() *Array {
	log.Println("Calling PopupMenu.X_GetItems()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "_get_items", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PopupMenu) X_GuiInput(arg0 *InputEvent) {
	log.Println("Calling PopupMenu.X_GuiInput()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_gui_input", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PopupMenu) X_SetItems(arg0 *Array) {
	log.Println("Calling PopupMenu.X_SetItems()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_set_items", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PopupMenu) X_SubmenuTimeout() {
	log.Println("Calling PopupMenu.X_SubmenuTimeout()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "_submenu_timeout", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Add a new checkable item with text "label". An id can optionally be provided, as well as an accelerator. If no id is provided, one will be created from the index. Note that checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
func (o *PopupMenu) AddCheckItem(label string, id int64, accel int64) {
	log.Println("Calling PopupMenu.AddCheckItem()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(label)
	goArguments[1] = reflect.ValueOf(id)
	goArguments[2] = reflect.ValueOf(accel)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_check_item", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *PopupMenu) AddCheckShortcut(shortcut *ShortCut, id int64, global bool) {
	log.Println("Calling PopupMenu.AddCheckShortcut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(shortcut)
	goArguments[1] = reflect.ValueOf(id)
	goArguments[2] = reflect.ValueOf(global)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_check_shortcut", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Add a new checkable item with text "label" and icon "texture". An id can optionally be provided, as well as an accelerator. If no id is provided, one will be created from the index. Note that checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
func (o *PopupMenu) AddIconCheckItem(texture *Texture, label string, id int64, accel int64) {
	log.Println("Calling PopupMenu.AddIconCheckItem()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(texture)
	goArguments[1] = reflect.ValueOf(label)
	goArguments[2] = reflect.ValueOf(id)
	goArguments[3] = reflect.ValueOf(accel)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_icon_check_item", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *PopupMenu) AddIconCheckShortcut(texture *Texture, shortcut *ShortCut, id int64, global bool) {
	log.Println("Calling PopupMenu.AddIconCheckShortcut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(texture)
	goArguments[1] = reflect.ValueOf(shortcut)
	goArguments[2] = reflect.ValueOf(id)
	goArguments[3] = reflect.ValueOf(global)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_icon_check_shortcut", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Add a new item with text "label" and icon "texture". An id can optionally be provided, as well as an accelerator keybinding. If no id is provided, one will be created from the index.
*/
func (o *PopupMenu) AddIconItem(texture *Texture, label string, id int64, accel int64) {
	log.Println("Calling PopupMenu.AddIconItem()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(texture)
	goArguments[1] = reflect.ValueOf(label)
	goArguments[2] = reflect.ValueOf(id)
	goArguments[3] = reflect.ValueOf(accel)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_icon_item", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *PopupMenu) AddIconShortcut(texture *Texture, shortcut *ShortCut, id int64, global bool) {
	log.Println("Calling PopupMenu.AddIconShortcut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(texture)
	goArguments[1] = reflect.ValueOf(shortcut)
	goArguments[2] = reflect.ValueOf(id)
	goArguments[3] = reflect.ValueOf(global)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_icon_shortcut", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Add a new item with text "label". An id can optionally be provided, as well as an accelerator keybinding. If no id is provided, one will be created from the index.
*/
func (o *PopupMenu) AddItem(label string, id int64, accel int64) {
	log.Println("Calling PopupMenu.AddItem()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(label)
	goArguments[1] = reflect.ValueOf(id)
	goArguments[2] = reflect.ValueOf(accel)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_item", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Add a separator between items. Separators also occupy an index.
*/
func (o *PopupMenu) AddSeparator() {
	log.Println("Calling PopupMenu.AddSeparator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_separator", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *PopupMenu) AddShortcut(shortcut *ShortCut, id int64, global bool) {
	log.Println("Calling PopupMenu.AddShortcut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(shortcut)
	goArguments[1] = reflect.ValueOf(id)
	goArguments[2] = reflect.ValueOf(global)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_shortcut", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds an item with a submenu. The submenu is the name of a child PopupMenu node that would be shown when the item is clicked. An id can optionally be provided, but if is isn't provided, one will be created from the index.
*/
func (o *PopupMenu) AddSubmenuItem(label string, submenu string, id int64) {
	log.Println("Calling PopupMenu.AddSubmenuItem()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(label)
	goArguments[1] = reflect.ValueOf(submenu)
	goArguments[2] = reflect.ValueOf(id)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "add_submenu_item", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Clear the popup menu, in effect removing all items.
*/
func (o *PopupMenu) Clear() {
	log.Println("Calling PopupMenu.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Return the accelerator of the item at index "idx". Accelerators are special combinations of keys that activate the item, no matter which control is focused.
*/
func (o *PopupMenu) GetItemAccelerator(idx int64) int64 {
	log.Println("Calling PopupMenu.GetItemAccelerator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_accelerator", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the amount of items.
*/
func (o *PopupMenu) GetItemCount() int64 {
	log.Println("Calling PopupMenu.GetItemCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_count", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the icon of the item at index "idx".
*/
func (o *PopupMenu) GetItemIcon(idx int64) *Texture {
	log.Println("Calling PopupMenu.GetItemIcon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_icon", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the id of the item at index "idx".
*/
func (o *PopupMenu) GetItemId(idx int64) int64 {
	log.Println("Calling PopupMenu.GetItemId()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_id", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Find and return the index of the item containing a given id.
*/
func (o *PopupMenu) GetItemIndex(id int64) int64 {
	log.Println("Calling PopupMenu.GetItemIndex()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(id)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_index", goArguments, "int64")

	returnValue := goRet.Interface().(int64)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the metadata of an item, which might be of any type. You can set it with [method set_item_metadata], which provides a simple way of assigning context data to items.
*/
func (o *PopupMenu) GetItemMetadata(idx int64) *Variant {
	log.Println("Calling PopupMenu.GetItemMetadata()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_metadata", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PopupMenu) GetItemShortcut(idx int64) *ShortCut {
	log.Println("Calling PopupMenu.GetItemShortcut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_shortcut", goArguments, "*ShortCut")

	returnValue := goRet.Interface().(*ShortCut)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the submenu name of the item at index "idx".
*/
func (o *PopupMenu) GetItemSubmenu(idx int64) string {
	log.Println("Calling PopupMenu.GetItemSubmenu()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_submenu", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the text of the item at index "idx".
*/
func (o *PopupMenu) GetItemText(idx int64) string {
	log.Println("Calling PopupMenu.GetItemText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_text", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PopupMenu) GetItemTooltip(idx int64) string {
	log.Println("Calling PopupMenu.GetItemTooltip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "get_item_tooltip", goArguments, "string")

	returnValue := goRet.Interface().(string)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a boolean that indicates whether or not the PopupMenu will hide on checkable item selection.
*/
func (o *PopupMenu) IsHideOnCheckableItemSelection() bool {
	log.Println("Calling PopupMenu.IsHideOnCheckableItemSelection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_hide_on_checkable_item_selection", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a boolean that indicates whether or not the PopupMenu will hide on item selection.
*/
func (o *PopupMenu) IsHideOnItemSelection() bool {
	log.Println("Calling PopupMenu.IsHideOnItemSelection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_hide_on_item_selection", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the item at index "idx" has a checkbox. Note that checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
func (o *PopupMenu) IsItemCheckable(idx int64) bool {
	log.Println("Calling PopupMenu.IsItemCheckable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_item_checkable", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the checkstate status of the item at index "idx".
*/
func (o *PopupMenu) IsItemChecked(idx int64) bool {
	log.Println("Calling PopupMenu.IsItemChecked()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_item_checked", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the item at index "idx" is disabled. When it is disabled it can't be selected, or its action invoked.
*/
func (o *PopupMenu) IsItemDisabled(idx int64) bool {
	log.Println("Calling PopupMenu.IsItemDisabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_item_disabled", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the item is a seperator. If it is, it would be displayed as a line.
*/
func (o *PopupMenu) IsItemSeparator(idx int64) bool {
	log.Println("Calling PopupMenu.IsItemSeparator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := class.CallParentMethod(o, o.baseClass(), "is_item_separator", goArguments, "bool")

	returnValue := goRet.Interface().(bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Removes the item at index "idx" from the menu. Note that the indexes of items after the removed item are going to be shifted by one.
*/
func (o *PopupMenu) RemoveItem(idx int64) {
	log.Println("Calling PopupMenu.RemoveItem()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "remove_item", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets whether or not the PopupMenu will hide on checkable item selection.
*/
func (o *PopupMenu) SetHideOnCheckableItemSelection(enable bool) {
	log.Println("Calling PopupMenu.SetHideOnCheckableItemSelection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_hide_on_checkable_item_selection", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets whether or not the PopupMenu will hide on item selection.
*/
func (o *PopupMenu) SetHideOnItemSelection(enable bool) {
	log.Println("Calling PopupMenu.SetHideOnItemSelection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_hide_on_item_selection", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the accelerator of the item at index "idx". Accelerators are special combinations of keys that activate the item, no matter which control is focused.
*/
func (o *PopupMenu) SetItemAccelerator(idx int64, accel int64) {
	log.Println("Calling PopupMenu.SetItemAccelerator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(accel)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_accelerator", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set whether the item at index "idx" has a checkbox. Note that checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
func (o *PopupMenu) SetItemAsCheckable(idx int64, enable bool) {
	log.Println("Calling PopupMenu.SetItemAsCheckable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_as_checkable", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Mark the item at index "idx" as a seperator, which means that it would be displayed as a mere line.
*/
func (o *PopupMenu) SetItemAsSeparator(idx int64, enable bool) {
	log.Println("Calling PopupMenu.SetItemAsSeparator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_as_separator", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the checkstate status of the item at index "idx".
*/
func (o *PopupMenu) SetItemChecked(idx int64, checked bool) {
	log.Println("Calling PopupMenu.SetItemChecked()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(checked)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_checked", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets whether the item at index "idx" is disabled or not. When it is disabled it can't be selected, or its action invoked.
*/
func (o *PopupMenu) SetItemDisabled(idx int64, disabled bool) {
	log.Println("Calling PopupMenu.SetItemDisabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(disabled)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_disabled", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *PopupMenu) SetItemIcon(idx int64, icon *Texture) {
	log.Println("Calling PopupMenu.SetItemIcon()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(icon)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_icon", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the id of the item at index "idx".
*/
func (o *PopupMenu) SetItemId(idx int64, id int64) {
	log.Println("Calling PopupMenu.SetItemId()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(id)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_id", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the metadata of an item, which might be of any type. You can later get it with [method get_item_metadata], which provides a simple way of assigning context data to items.
*/
func (o *PopupMenu) SetItemMetadata(idx int64, metadata *Variant) {
	log.Println("Calling PopupMenu.SetItemMetadata()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(metadata)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_metadata", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *PopupMenu) SetItemShortcut(idx int64, shortcut *ShortCut, global bool) {
	log.Println("Calling PopupMenu.SetItemShortcut()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(shortcut)
	goArguments[2] = reflect.ValueOf(global)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_shortcut", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the submenu of the item at index "idx". The submenu is the name of a child PopupMenu node that would be shown when the item is clicked.
*/
func (o *PopupMenu) SetItemSubmenu(idx int64, submenu string) {
	log.Println("Calling PopupMenu.SetItemSubmenu()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(submenu)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_submenu", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the text of the item at index "idx".
*/
func (o *PopupMenu) SetItemText(idx int64, text string) {
	log.Println("Calling PopupMenu.SetItemText()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(text)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_text", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *PopupMenu) SetItemTooltip(idx int64, tooltip string) {
	log.Println("Calling PopupMenu.SetItemTooltip()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(idx)
	goArguments[1] = reflect.ValueOf(tooltip)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "set_item_tooltip", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *PopupMenu) ToggleItemChecked(idx int64) {
	log.Println("Calling PopupMenu.ToggleItemChecked()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	o.callParentMethod(o.baseClass(), "toggle_item_checked", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   PopupMenuImplementer is an interface for PopupMenu objects.
*/
type PopupMenuImplementer interface {
	class.Class
}