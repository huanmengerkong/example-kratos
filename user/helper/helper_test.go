package helper

import (
	"fmt"
	"testing"
)

func TestStructToGromStruct(T *testing.T) {
	type User struct {
		Name string `json:"name"`
		Sex  string `json:"sex"`
		Age  int    `json:"age"`
	}
	var a = User{
		Name: "xiaoming",
		Sex:  "男",
		Age:  10,
	}
	type B struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user B
	fmt.Println(StructToGromStruct(a, user))
}
