package main

import (
	"aliyun-magic/util"
	"fmt"
)

func main() {
	result := util.GetYesterdayTimePeriod()
	fmt.Println(result)
}
