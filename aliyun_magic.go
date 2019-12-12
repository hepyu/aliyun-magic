package main

import (
	"aliyun-magic/service"
	"fmt"
	"os"
)

func main() {
	fmt.Println("111")
	fmt.Println(os.Getenv("AccessKeyID"))
	fmt.Println(os.Getenv("JAVA_HOME"))
	service.GetECSInfo("cn-zhangjiakou", 100)
}
