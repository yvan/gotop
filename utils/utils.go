package utils

import (
	"fmt"
	"reflect"
)

// little function to print the methods of an object
func PrintMethods (object interface{}) {
	oType := reflect.TypeOf(object)
	for i := 0; i < oType.NumMethod(); i++ {
		method := oType.Method(i)
		fmt.Println(method.Name)
	}
}
