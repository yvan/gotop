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

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int{
    if a < b {
        return a
    }
    return b
}

func maxOrder(a, b int) (int, int){
    if a > b {
        return a, b
    }
    return b, a
}

func minOrder(a, b int) (int, int){
    if a < b {
        return a, b
    }
    return b, a
}

func arrayMax(array []int32) int32{
    var tempMax int32 = 0
    for _, v := range array {
        tempMax = max(tempMax, v)
    }
    return tempMax
}


