package main

import "fmt"

// 管道channel 定义，

//ci := make(chan int)
//cs := make(chan string)
//cf := make(chan interface{})

func main() {

	// 非缓存channel
	c := make(chan int)

	array := []int{1, 2, 3, 7, -9, 0}

	go sum(array[:len(array)/2], c)

	go sum(array[len(array)/2:], c)

	// 从channel接收值
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// 缓冲channel, 把2改为1后会报错
	c1 := make(chan int, 2)
	c1 <- 1
	c1 <- 2
	fmt.Println(<-c1)
	fmt.Println(<-c1)

	// range 和 close
	c2 := make(chan int, 2)
	// 开启线程向管道增加数据
	go add(c2)

	// range 循环获取管道值, go add函数是多线程，所以range循环获取管道数据不需要等待add执行完再获取
	for v := range c2 {
		fmt.Println(v)
	}

	//// 以下代码看不懂
	//// 利用 close 广播
	//c3 := make(chan int, 5)
	//for i := 0; i < 5; i++ {
	//	go do(c3)
	//}
	//close(c3)
	//
	//fmt.Println("finished")

}
func do(c chan int) {
	<-c //利用 close 广播
	fmt.Println("do something")
}

func add(c2 chan int) {
	for i := 90; i <= 95; i++ {
		c2 <- i
		fmt.Println("add ", i)
	}
	close(c2)
}
func sum(array []int, c chan int) {

	total := 0
	for _, v := range array {
		total = total + v
	}
	// channel通过操作符<-来接收和发送数据
	// total发送到channel
	c <- total
}
