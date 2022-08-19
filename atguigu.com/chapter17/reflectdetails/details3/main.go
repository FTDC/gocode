package main

import (
	"fmt"
	"reflect"
)

func test(age interface{}) {

	rType := reflect.TypeOf(age)
	fmt.Printf("kind %v, type %v \n", rType.Kind(), rType)

	rval := reflect.ValueOf(age)
	fmt.Println("rval: ", rval)

	itype := rval.Interface()
	// fmt.Println("itype:", itype.Kind())

	ivalue := itype.(float64)
	fmt.Println("ivalue:", ivalue)
}

func main() {
	age := 16.5
	test(age)
}
