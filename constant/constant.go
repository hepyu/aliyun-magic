package constant

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ECSSubResourceType string

const (
	ECSCPU    ECSSubResourceType = "cpu"
	ECSMemory ECSSubResourceType = "memory"
)

type PX string

const (
	PXMax PX = "max"
	PXMin PX = "min"
	PXP99 PX = "99"
	PXP95 PX = "95"
	PXP90 PX = "90"
	PXP50 PX = "50"
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
