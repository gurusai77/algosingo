package main

import (
	"fmt"
	"reflect"
)

func genericMethod(r any) {
	fmt.Printf("type of r is %v\n", reflect.TypeOf(r))
}

func genericMethodI(r interface{}) {
	fmt.Printf("type of r is %v\n", reflect.TypeOf(r))
}

func genericMethodT[T any](r T) {
	fmt.Printf("type of r is %v\n", reflect.TypeOf(r))
}

func main() {
	genericMethod(4)
	genericMethod(32.3)
	genericMethod("test")
	genericMethod('c')

	genericMethodT(4)
	genericMethodT(4.4)

	genericMethodI(4)

}
