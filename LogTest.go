package main

import (
	"log"
	"github.com/sirupsen/logrus"
)

// log日志
func main() {

	//var mapTest map[int]string
	mapTest := make(map[int]string)
	mapTest[1] = "zcb"

	// Fatal会输出日志并调用os.exit退出，导致后面部分代码逻辑不执行，开发时需要慎重
	//log.Fatal(cap(slice), len(slice))

	log.Println("map", mapTest)

	// logrus日志库
	log1 := logrus.New()
	log1.Debug("map", mapTest)
	log1.Info("map", mapTest)
}
