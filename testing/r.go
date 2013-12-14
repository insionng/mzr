package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(nameOf(foo()), nameOf((*A).Method))
}

func nameOf(f interface{}) string {
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			return rf.Name()
		}
	}
	return v.String()
}

func foo() func() {
	return func() {}
}

type A struct{ x, y int }

func (*A) Method() {}
