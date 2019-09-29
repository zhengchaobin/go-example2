package main

import (
	"errors"
	"fmt"
	"github.com/zhengchaobin/go-example1"
)

func main() {

	// 异常测试
	number, e := errorTest(0)
	if e != nil{
		fmt.Println(number, e.Error())
	}else {
		fmt.Println(number)
	}


	// 异常自定义测试
	num, e := customErrorTest(1)
	if e != nil{
		fmt.Println(num, e.Error())
	}else {
		fmt.Println(num)
	}

	// 异常自定义测试
	num2, e := customErrorTest(0)
	if e != nil{
		fmt.Println(num2, e.Error())
	}else {
		fmt.Println(num2)
	}

	// 异常自定义测试
	num3, e := customErrorTest2(1)
	if e != nil{
		fmt.Println(num3, e.Error())
	}else {
		fmt.Println(num3)
	}


}

func errorTest(num int) (int, error) {
	if num < 1{
		return num, errors.New("异常")
	}
	// 正常情况下error返回nil
	return num, nil
}

func customErrorTest(num int) (int, error) {

	var e zcb.CustomError
	if num < 1{
		e = zcb.New("自定义异常1")
		return num, e
	}
	// 正常情况下error返回nil,注意: 不能返回e,返回e的话会非nil
	return num, nil
}

func customErrorTest2(num int) (int, *zcb.CustomError) {

	//指针类型
	var e *zcb.CustomError
	if num < 1{
		e = zcb.New2("自定义异常2")
		return num, e
	}
	// 正常情况下error返回nil,注意: 函数定义返回类型为指针类型,在外层判断时会非nil
	return num, nil
}


