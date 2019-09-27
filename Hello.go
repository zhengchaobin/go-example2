package main

import 	(
	"fmt"
	"github.com/zhengchaobin/go-example1"
)

// :=声明只能在函数中出现
//c,d := 1,2

func main() {


	fmt.Println("Hello World!")

	var a, b = 1, 2
	fmt.Println(a)
	fmt.Println(b)

	var a1 int

	// 使用:=不需要声明变量，但是:=左侧变量均声明了的情况下使用:=会报错，即:=左侧必须存在一个未声明变量
	a1, a2 := 1,2
	fmt.Println(a1)
	fmt.Println(a2)

	var a2Adress = &a2
	fmt.Println(a2Adress)
	var a3 = a2
	fmt.Println(a3)
	// 使用&可获取内存地址
	var a3Adress  = &a3
	fmt.Println(a3Adress)
	//方法调用获取返回值，只需要其中一部分返回值时，另外参数可以用_代替
	_,d,e := return3Value()
	fmt.Println(d)
	fmt.Println(e)
	couponName := getCoupon("zcb", 2)
	// if...else
	if couponName == "zcb"{
		fmt.Println(couponName, "if...")
	}else {
		fmt.Println("else...")
	}

	// for循环
	for i := 1; i < 5; i++ {
		fmt.Println("for:", i)
	}

	//无限循环
	//for true {
	//	fmt.Println("无限循环")
	//}
	// 一维数组
	var array = [5]int{1,2,3,4,5}
	fmt.Println(array)
	// range遍历数组
	for index, value := range array{
		fmt.Println("range for array:",index, value)
	}
	// 将array中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	slice1 := array[2:4]
	fmt.Println("slice1=", slice1)

	// 默认 endIndex 时将表示一直到arr的最后一个元素
	slice2 := array[2:]
	fmt.Println("slice2=", slice2)

	// 默认 startIndex 时将表示从arr的第一个元素开始
	slice3 := array[:4]
	fmt.Println("slice3=", slice3)


	var array2 [3][4]int
	// 二维数组
	for i := 0; i< 3; i++  {
		for j := 0; j< 2; j++ {
			array2[i][j] = i * j
		}
	}

	fmt.Println(array2)
	// 指针 &取址符号  *指向指针地址的值
	var ptr *int
	fmt.Println(ptr)
	var ptrValue = 3
	ptr = &ptrValue
	fmt.Println(ptr)
	fmt.Println(*ptr)

	// 接口和结构体，跨包调用
	var animal zcb.Animal
	//name := "小狗"
	animal = zcb.Dog{Name:"小狗"}
	animal.Sleep()
	animal.Move()

	animal = zcb.Cat{Name:"小猫"}
	animal.Sleep()
	animal.Move()

	//
	animal = zcb.Bird{zcb.Cat{Name: "小鸟"}, "小鸟x"}
	animal.Move()
	animal.Sleep()

	bird := zcb.Bird{zcb.Cat{"小袅袅"}, "小袅袅x"}
	birdName := bird.Name
	birdName1 := bird.Cat.Name
	fmt.Println("直接访问匿名属性:" + birdName, "通过匿名字段访问匿名属性:" + birdName1)

	// 切片: 对数据的抽象，相当于动态数组，长度是不固定的
	// 切片声明'[]'符号不需要带长度，数组的声明需要声明长度
	// 切片声明
	slice := make([]int, 7)
	fmt.Println(cap(slice),len(slice))

	//var mapTest map[int]string
	mapTest := make(map[int]string)
	mapTest[1] = "zcb"
	fmt.Println(mapTest)

}

func return3Value()(string ,int, bool)  {
	a,b,c := "zcb", 1, true
	return a,b,c
}

func getCoupon(myName string, int2 int) (string) {
	var couponName = myName
	return couponName
}
