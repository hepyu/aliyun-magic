package service

import (
	"aliyun-magic/sdk_wrapper"
	"fmt"
	//"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func GetECSInfo(regionId string, pageSize int) {
	//ecs.DescribeInstancesResponse
	response := sdk_wrapper.GetInstance(regionId, pageSize)
	fmt.Println(response)
}
