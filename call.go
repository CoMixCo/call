package call

import (
	"errors"
	"reflect"
)

type Handle func() any

type CallMap struct {
	m map[string]Handle
}

func NewCall() CallMap {
	return CallMap{
		m: map[string]Handle{},
	}
}

func (c *CallMap) Register(register_name string, func_new Handle) {
	c.m[register_name] = func_new
}

func (c *CallMap) Invok(register_name string, method_name string, params ...interface{}) ([]reflect.Value, error) {
	bind_handle, ok := c.m[register_name]
	if !ok {
		return []reflect.Value{}, errors.New("handle is not registed")
	}
	bind_struct := bind_handle()

	r_struct := reflect.ValueOf(bind_struct)

	//必须是结构体指针
	if r_struct.Kind() != reflect.Ptr || r_struct.Elem().Kind() != reflect.Struct {
		return []reflect.Value{}, errors.New("handle must return struct ptr")
	}
	r_method, ok := r_struct.Type().MethodByName(method_name)
	if !ok {
		return []reflect.Value{}, errors.New("struct method is not exits")
	}

	if len(params)+1 != r_method.Type.NumIn() {
		return []reflect.Value{}, errors.New("struct method number of params is not adapted")
	}
	in := make([]reflect.Value, len(params)+1)
	in[0] = r_struct
	for k, param := range params {
		in[k+1] = reflect.ValueOf(param)
	}
	result := r_method.Func.Call(in)
	return result, nil
}
