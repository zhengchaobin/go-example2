package main

import (
	"fmt"
	"runtime"
)

func main() {

	go saySth("hello")
	saySth("world")

}

func saySth(str string) {
	for i := 1; i <= 5; i++ {
		runtime.Gosched()
		fmt.Println(str)
	}
}
