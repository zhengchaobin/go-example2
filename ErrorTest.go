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


	// 异常测试
	num, e := customErrorTest(0)
	if e != nil{
		fmt.Println(num, e.Error())
	}else {
		fmt.Println(num)
	}


}

func errorTest(num int) (int, error) {
	if num < 1{
		return num, errors.New("异常")
	}
	// 正常情况下error返回nil
	return num, nil
}

func customErrorTest(num int) (int, *zcb.CustomError) {

	var e *zcb.CustomError
	if num < 1{
		e := zcb.New("自定义异常")
		return num, e
	}
	// 正常情况下error返回nil
	return num, e
}


