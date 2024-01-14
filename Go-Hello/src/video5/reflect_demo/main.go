package main

import (
	"fmt"
	"reflect"
)

//反射

func reflectType(x interface{}) {

	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v", v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Println(ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Println(ret)
	}
}
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind() //根据指针取值
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(110)
	case reflect.Float32:
		v.Elem().SetFloat(3.234)
	}
}

type Cat struct {
}
type Dog struct {
}

func main() {
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 10
	reflectType(b)
	var c Cat
	reflectType(c)
	var d Dog
	reflectType(d)

	var aa int32 = 100
	reflectValue(aa)

	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Println(aaa)
}
