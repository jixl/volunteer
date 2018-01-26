package models

import (
	"reflect"
)

func isNeedReset(field interface{}) bool {
	return reflect.TypeOf(field).String() != "string"
}

func getSkip(page int, size int) int {
	if page > 1 && size > 0 {
		return (page - 1) * size
	}
	return 0
}

func getPageSize(size int) int {
	if size > 0 {
		return size
	}
	return 20
}
