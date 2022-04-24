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

func strOrder(a, b string) (string, string) {
    if len(a) > len(b) {
        return a, b
    } else {
	return b, a
    }
}

func intStrOrder(a, b string) (string, string) {
    for i := 0; i < len(a); i++ {
        inta := int(a[i])
        intb := int(b[i])
        if inta > intb {
            return a, b
        } else if intb > inta {
            return b, a
        }
    }
    return a, b
}

func arrayMax(array []int32) int32{
    var tempMax int32 = 0
    for _, v := range array {
        tempMax = max(tempMax, v)
    }
    return tempMax
}

func arrMin(arr []int) int {
    tempMin := 0
    for i, v := range arr {
        if i == 0 {
            tempMin = v
        }
        tempMin = min(tempMin, v)
    }
    return tempMin
}


