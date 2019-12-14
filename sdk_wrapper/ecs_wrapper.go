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

	fmt.Println(constant.GetAccessKeyID())
	fmt.Println(constant.GetAccessSecret())

	client, err := ecs.NewClientWithAccessKey(regionId, constant.GetAccessKeyID(), constant.GetAccessSecret())
	if err != nil {
		fmt.Println(err)
	}

	request := ecs.CreateDescribePriceRequest()

	request.RegionId = regionId
	request.InstanceType = instanceType
	request.PriceUnit = priceUnit
	request.SystemDiskSize = requests.NewInteger(systemDiskSize)

	fmt.Println(regionId)
	fmt.Println(request)

	response, err := client.DescribePrice(request)
	if err != nil {
		fmt.Println("here")
		fmt.Println(err)
	}

	return response
}

func GetInstance(regionId string, pageSize int, pageNum int) *ecs.DescribeInstancesResponse {
	client, err := ecs.NewClientWithAccessKey(regionId, constant.GetAccessKeyID(), constant.GetAccessSecret())
	if err != nil {
		fmt.Println(err)
	}

	request := ecs.CreateDescribeInstancesRequest()
	request.RegionId = regionId
	request.PageSize = requests.NewInteger(pageSize)
	request.PageNumber = requests.NewInteger(pageNum)

	response, err := client.DescribeInstances(request)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

func GetDisks(regionId string, instanceId string) *ecs.DescribeDisksResponse {
	client, err := ecs.NewClientWithAccessKey(regionId, constant.GetAccessKeyID(), constant.GetAccessSecret())
	if err != nil {
		fmt.Println(err)
	}

	request := ecs.CreateDescribeDisksRequest()
	//request.AcceptFormat = "json"
	request.InstanceId = instanceId

	response, err := client.DescribeDisks(request)
	if err != nil {
		fmt.Println(err)
	}

	return response
}
