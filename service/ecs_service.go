package service

import (
	"aliyun-magic/dto"
	"aliyun-magic/sdk_wrapper"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

//func GetECSCostDTOArray(regionId string, pageSize int) []dto.ECSCostDTO {
func GetECSCostDTOArray(instances []ecs.Instance, regionId string, orderType string, month string) []dto.ECSCostDTO {
	//instances := getECSInfoArray(regionId, pageSize)
	//定义slice
	var ecsCostDTOArray []dto.ECSCostDTO
	//var ecsCostDTO dto.ECSCostDTO
	for _, instance := range instances {
		ecsCostDTO := new(dto.ECSCostDTO)
		ecsMarkInfo := new(dto.ResourceECSMarkDTO)
		ecsCostDTO.ResourceECSMarkInfo = ecsMarkInfo

		ecsMarkInfo.Status = instance.Status
		ecsMarkInfo.RegionId = regionId
		ecsMarkInfo.InstanceId = instance.InstanceId
		ecsMarkInfo.InstanceType = instance.InstanceType

		tagArray := instance.Tags.Tag
		for _, tag := range tagArray {
			if tag.TagKey == "applicant" {
				ecsMarkInfo.Applicant = tag.TagValue
			} else if tag.TagKey == "env" {
				ecsMarkInfo.Env = tag.TagValue
			} else if tag.TagKey == "serverType" {
				ecsMarkInfo.ServerType = tag.TagValue
			} else if tag.TagKey == "serverName" {
				ecsMarkInfo.ServerName = tag.TagValue
			} else if tag.TagKey == "owner" {
				ecsMarkInfo.Owner = tag.TagValue
			} else if tag.TagKey == "businessLine" {
				ecsMarkInfo.BusinessLine = tag.TagValue
			} else if tag.TagKey == "project" {
				ecsMarkInfo.Project = tag.TagValue
			}
		}
		fmt.Println(*ecsCostDTO.ResourceECSMarkInfo)
		//ecsCostDTO.Price = sdk_wrapper.GetPrice(regionId, ecsMarkInfo.InstanceType, "Month", 40).PriceInfo.Price.OriginalPrice

		moduleList := &[]bssopenapi.GetSubscriptionPriceModuleList{
			{
				ModuleCode: "InstanceType",
				Config:     "InstanceType:" + ecsMarkInfo.InstanceType,
			},
		}
		ecsCostDTO.Price = sdk_wrapper.GetSubscriptionPrice(regionId, ecsMarkInfo.InstanceId, "ecs", orderType, month, 1, moduleList).Data.OriginalPrice

		ecsCostDTOArray = append(ecsCostDTOArray, *ecsCostDTO)
	}
	return ecsCostDTOArray
}

func GetECSInfoArray(regionId string, pageSize int) []ecs.Instance {
	//定义slice
	var instances []ecs.Instance
	for pageNum := 1; ; pageNum++ {
		response := sdk_wrapper.GetInstance(regionId, pageSize, pageNum)
		if response != nil && response.Instances.Instance != nil {
			tarray := response.Instances.Instance
			for _, t := range tarray {
				instances = append(instances, t)
			}
			if len(response.Instances.Instance) < pageSize {
				break
			}
		} else {
			break
		}
	}
	return instances
}
