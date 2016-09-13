package ioc

import (
	"fmt"
	"reflect"
)

var factory = make(map[string]reflect.Type)

func Register(name string, t reflect.Type) {
	factory[name] = t
}

func Create(name string) (interface{}, bool) {

	t, ok := factory[name]

	if ok {
		v := reflect.New(t).Elem()
		// Maybe fill in fields here if necessary
		return v.Interface(), true
	} else {
		return nil, false
	}
}

func Print() {
	fmt.Println("====ioc_map:", factory)
}
