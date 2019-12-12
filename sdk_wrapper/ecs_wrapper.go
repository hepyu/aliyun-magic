package sdk_wrapper

import (
	"aliyun-magic/constant"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func GetInstanceMonitorData(regionId string, instanceId string, startTime string, endTime string) *ecs.DescribeInstanceMonitorDataResponse {
	client, err := ecs.NewClientWithAccessKey(regionId, constant.GetAccessKeyID(), constant.GetAccessSecret())
	if err != nil {
		fmt.Println(err)
	}

	request := ecs.CreateDescribeInstanceMonitorDataRequest()

	request.RegionId = regionId
	request.InstanceId = instanceId
	request.StartTime = startTime
	request.EndTime = endTime

	response, err := client.DescribeInstanceMonitorData(request)
	if err != nil {
		fmt.Println(err)
	}

	return response
}

func GetPrice(regionId string, instanceType string, priceUnit string, systemDiskSize int) *ecs.DescribePriceResponse {

	client, err := ecs.NewClientWithAccessKey(regionId, constant.GetAccessKeyID(), constant.GetAccessSecret())
	if err != nil {
		fmt.Println(err)
	}

	request := ecs.CreateDescribePriceRequest()

	request.RegionId = regionId
	request.InstanceType = instanceType
	request.PriceUnit = priceUnit
	request.SystemDiskSize = requests.NewInteger(systemDiskSize)

	response, err := client.DescribePrice(request)
	if err != nil {
		fmt.Println(err)
	}

	return response
}

func GetInstance(regionId string, pageSize int) *ecs.DescribeInstancesResponse {
	fmt.Println("GetAccessKeyID:" + constant.GetAccessKeyID())
	fmt.Println("GetAccessSecret:" + constant.GetAccessSecret())
	client, err := ecs.NewClientWithAccessKey(regionId, constant.GetAccessKeyID(), constant.GetAccessSecret())
	if err != nil {
		fmt.Println(err)
	}

	request := ecs.CreateDescribeInstancesRequest()
	request.RegionId = regionId
	request.PageSize = requests.NewInteger(pageSize)

	response, err := client.DescribeInstances(request)
	if err != nil {
		fmt.Println(err)
	}
	return response
}
