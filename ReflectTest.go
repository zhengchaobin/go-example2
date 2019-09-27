package main

import (
	"fmt"
	"reflect"
)

type user struct {
	
	Name string
	
	Age int
}

//结构体方法，需要在方法名前面增加结构体声明 (u user)
func(u user)Hello()  {
	fmt.Println(u.Name, u.Age)
}

// 反射获取结构体对应的类型、属性、值、方法
func getStructType(inter interface{}) {
	// 获取结构体类型
	t := reflect.TypeOf(inter)
	name := t.Name()
	fmt.Println("interface name:" + name)

	values := reflect.ValueOf(inter)
	for i := 0; i < t.NumField(); i++ {
		// 获取属性
		fie := t.Field(i)
		// 获取属性值
		val := values.Field(i).Interface()
		fmt.Println("field ", fie.Name, fie.Type, val)
	}

	// 获取结构体方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("method ", m.Name, m.Type)
	}
}


func main() {
	u := user{"zcb", 18}
	getStructType(u)
	u.Hello()
}
