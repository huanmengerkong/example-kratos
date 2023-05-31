package helper

import (
	"fmt"
	"reflect"
)

func StructToGromStruct(a interface{}, b interface{}) interface{} {
	t := reflect.TypeOf(a)
	f := reflect.TypeOf(b)
	if t.Kind() != reflect.Struct && f.Kind() != reflect.Struct {
		return nil
	}
	// valueb := reflect.ValueOf(b)
	// var (
	// 	field []string
	// )
	// 动态创建结构体
	// 在使用 reflect.New(reflect.TypeOf(b)) 函数创建指向 b 类型的指针时，应该使用reflect.ValueOf(b).Type()方法确保获取到的是指针类型。
	p := reflect.New(reflect.TypeOf(b)).Elem() // 后去结构体类型
	valueB := reflect.ValueOf(b)
	valueA := reflect.ValueOf(a)
	for i := 0; i < t.NumField(); i++ {
		// field = append(field, t.Field(i).Name)
		name := t.Field(i).Name
		if !valueB.FieldByName(name).IsValid() {
			continue
		}

		fmt.Println(valueB.FieldByName(name).Type().String(), "==", t.Field(i).Type.String(), "==", valueB.FieldByName(name).Type().String() == t.Field(i).Type.String())
		if valueB.FieldByName(name).Type().String() == t.Field(i).Type.String() {
			p.FieldByName(t.Field(i).Name).Set(valueA.Field(i))
		}
		// fmt.Println(f.Field(i).Type.String()) //string int
	}
	return p.Interface()
}
