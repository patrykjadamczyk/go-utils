package base

import (
	"fmt"
	"reflect"
)

const (
	ENSURE_OBJECT_OK_STATE = "OK"
)

type EnsuredObject interface {
	// Function for checking if Object is according to implementation
	Dev() Result[string]
}

// Ensure that specified object implements provided interface
// OT - Object Type
// IT - Interface Type
// self - Object
// Returns Result type
func EnsureObjectImplementsInterface[OT any, IT any](self OT) Result[string] {
	if any(self) == nil {
		return Err[string](NewError("Object is nil"))
	}
	interfaceValue := reflect.TypeOf((*IT)(nil))
	interfaceType := interfaceValue.Elem()
	selfType := reflect.TypeOf(self)
	if selfType.Implements(interfaceType) {
		return Ok(ENSURE_OBJECT_OK_STATE)
	}
	return Err[string](NewError(fmt.Sprintf("Object is not %s", interfaceValue)))
}
