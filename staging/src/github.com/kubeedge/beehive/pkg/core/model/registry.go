package model

import (
	"fmt"
	"reflect"
)

var defaultRegistry = newRegistry()

func Register(name string, val interface{}) error {
	return defaultRegistry.Register(name, val)
}

type registry struct {
	nameToType map[string]reflect.Type
	typeToName map[reflect.Type]string
}

func newRegistry() *registry {
	return &registry{
		nameToType: map[string]reflect.Type{},
		typeToName: map[reflect.Type]string{},
	}
}

func (r *registry) Register(name string, val interface{}) error {
	typ, err := getTypeFromVal(val)
	if err != nil {
		return err
	}
	if name == "" {
		name = typ.PkgPath() + "." + typ.Name()
	}
	return r.register(name, typ)
}

func (r *registry) register(name string, typ reflect.Type) error {
	r.typeToName[typ] = name
	r.nameToType[name] = typ
	return nil
}

func (r *registry) New(name string) interface{} {
	typ, ok := r.nameToType[name]
	if !ok {
		return nil
	}
	return reflect.New(typ).Interface()
}

func (r *registry) GetName(val interface{}) (string, bool) {
	typ, err := getTypeFromVal(val)
	if err != nil {
		return "", false
	}
	name, ok := r.typeToName[typ]
	return name, ok
}

func getTypeFromVal(val interface{}) (reflect.Type, error) {
	typ := reflect.TypeOf(val)
	if typ == nil {
		return nil, fmt.Errorf("cannot be nil")
	}
	return getType(typ)
}

func getType(typ reflect.Type) (reflect.Type, error) {
	switch typ.Kind() {
	case reflect.Interface:
		return nil, fmt.Errorf("does not support registering an interface %q", typ.Name())
	case reflect.Ptr:
		return getType(typ.Elem())
	}
	return typ, nil
}
