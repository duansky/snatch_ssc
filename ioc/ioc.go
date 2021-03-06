package ioc

import "reflect"

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

var objs = make(map[string]interface{})

func RegisterObj(name string, o interface{}) {
	objs[name] = o
}

func CreateObj(name string) (interface{}, bool) {
	o, ok := objs[name]

	if ok {
		return o, true
	} else {
		return nil, false
	}
}
