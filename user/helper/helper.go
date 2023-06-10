package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func NewHelper() *Helper {
	return &Helper{}
}

type Helper struct {
}

func (h *Helper) StructToStruct(a interface{}, b interface{}) interface{} {
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

func (h *Helper) Md5(str, salt string) string {
	t := md5.Sum([]byte(str + salt))
	return hex.EncodeToString(t[:])
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func (h *Helper) RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
