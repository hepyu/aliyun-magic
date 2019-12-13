package constant

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetPushGatewayAddress() string {
	return os.Getenv("PushGatewayAddress")
}

func GetAccessKeyID() string {
	return os.Getenv("AccessKeyID")
}

func GetAccessSecret() string {
	return os.Getenv("AccessKeySecret")
}

func GetRegionId() []string {
	tmp := os.Getenv("RegionID")
	return strings.Split(tmp, ",")
}

func GetECSCollectorConcurrent() int {
	concurrent, err := strconv.Atoi(os.Getenv("ECSCollectorConcurrent"))
	if err != nil {
		fmt.Println(err)
	}
	return concurrent
}

func GetECSCollectorPageSize() int {
	pageSize, err := strconv.Atoi(os.Getenv("ECSCollectorPageSize"))
	if err != nil {
		fmt.Println(err)
	}
	return pageSize
}
