package main

import (
	"aliyun-magic/collector"
	//"aliyun-magic/service"
)

func main() {
	//service.GetECSInfo("cn-zhangjiakou", 100)
	collector.CollectECS()
}
