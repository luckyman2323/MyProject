package reflect

import (
	"fmt"
	"reflect"
)

func ForeachStruct(obj interface{}) {
	t := reflect.TypeOf(obj) // 注意，obj不能为指针类型，否则会报：panic recovered: reflect: NumField of non-struct type
	v := reflect.ValueOf(obj)
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%+v -- %v \n", t.Field(k), v.Field(k).Interface())
	}
}
