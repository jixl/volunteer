package models

import "reflect"

func isNeedReset(field interface{}) bool {
	return reflect.TypeOf(field).String() != "string"
}
