package reflecttools

import (
	"reflect"
)

// Init wrapper to call InitValue
func Init(i interface{}) {
	InitValue(reflect.ValueOf(i).Elem())
}

// InitValue initialized non-zeros types (Maps, Slices, ...)
func InitValue(v reflect.Value) {
	switch v.Kind() {
	// Func, UnsafePointer, and Interface also have no zero value
	// but we it isn't have any sance to init them
	case reflect.Chan,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice:
		if !v.IsNil() {
			return
		}
	}
	switch v.Type().Kind() {
	case reflect.Struct:
		for idx := 0; idx < v.NumField(); idx++ {
			InitValue(v.Field(idx))
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
	case reflect.Slice:
		v.Set(reflect.MakeSlice(v.Type(), 0, 0))
	case reflect.Chan:
		v.Set(reflect.MakeChan(v.Type(), 0))
	case reflect.Ptr:
		fv := reflect.New(v.Type())
		InitValue(fv.Elem())
		v.Set(fv)
	}
}
