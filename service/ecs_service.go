package service

import (
	"aliyun-magic/aliyun_struct_wrapper"
	"aliyun-magic/constant"
	"aliyun-magic/dto"
	"aliyun-magic/sdk_wrapper"
	"aliyun-magic/util"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"math"
	"sort"
	"strings"
)

func GetECSCostDTOArray(instances []ecs.Instance, regionId string, orderType string, month string) []dto.ECSCostDTO {
	//定义slice
	var ecsCostDTOArray []dto.ECSCostDTO
	for _, instance := range instances {
		ecsCostDTO := new(dto.ECSCostDTO)
		ecsMarkInfo := getECSMarkInfo(instance, regionId)
		ecsCostDTO.ResourceECSMarkInfo = ecsMarkInfo

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

func GetECSSubResourceUsagePXByYesterday(regionId string, instances []ecs.Instance, ecsSubResourceTypes []constant.ECSSubResourceType, pxes []constant.PX) []dto.ECSSubResourceUsagePXDTO {
	yesterday := util.GetYesterdayZeroTime()
	yesterdayTimePeriod := util.GetTimePeriod(yesterday, 10)

	//定义Slice
	var result []dto.ECSSubResourceUsagePXDTO

	var instanceMonitorData []ecs.InstanceMonitorData
	for _, instance := range instances {
		for _, tp := range yesterdayTimePeriod {
			tpa := strings.Split(tp, ",")
			data := sdk_wrapper.GetInstanceMonitorData(regionId, instance.InstanceId, tpa[0], tpa[1])
			if data != nil {
				for _, td := range data.MonitorData.InstanceMonitorData {
					instanceMonitorData = append(instanceMonitorData, td)
				}
			}
		}

		for _, subType := range ecsSubResourceTypes {
			//目前只实现了CPU
			if subType == constant.ECSCPU {
				cpuUsagePXDTO := getECSResourcePXMonitorData(pxes, instanceMonitorData, instance, regionId)
				result = append(result, *cpuUsagePXDTO)
			}
		}
	}
	return result
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

func getECSMarkInfo(instance ecs.Instance, regionId string) *dto.ResourceECSMarkDTO {
	ecsMarkInfo := new(dto.ResourceECSMarkDTO)
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
	return ecsMarkInfo
}

func getECSResourcePXMonitorData(pxes []constant.PX, instanceMonitorData []ecs.InstanceMonitorData, instance ecs.Instance, regionId string) *dto.ECSSubResourceUsagePXDTO {
	size := len(instanceMonitorData)

	sort.Sort(aliyun_struct_wrapper.ECSInstanceMonitorDataWrapper{instanceMonitorData, func(p, q *ecs.InstanceMonitorData) bool {
		return p.CPU < q.CPU
	}})

	cpuUsagePXDTO := new(dto.ECSSubResourceUsagePXDTO)
	pxDTO := new(dto.PXDTO)
	for _, px := range pxes {
		if px == constant.PXMax {
			pxDTO.MAX = float64(instanceMonitorData[size-1].CPU) * float64(0.01)
		} else if px == constant.PXMin {
			pxDTO.MIN = float64(instanceMonitorData[0].CPU) * float64(0.01)
		} else if px == constant.PXP50 {
			index := int(math.Ceil(float64(size) * float64(0.5)))
			pxDTO.P50 = float64(instanceMonitorData[index].CPU) * float64(0.01)
		} else if px == constant.PXP90 {
			index := int(math.Ceil(float64(size) * float64(0.9)))
			pxDTO.P90 = float64(instanceMonitorData[index].CPU) * float64(0.01)
		} else if px == constant.PXP95 {
			index := int(math.Ceil(float64(size) * float64(0.95)))
			pxDTO.P95 = float64(instanceMonitorData[index].CPU) * float64(0.01)
		} else if px == constant.PXP99 {
			index := int(math.Ceil(float64(size) * float64(0.99)))
			pxDTO.P99 = float64(instanceMonitorData[index].CPU) * float64(0.01)
		}
	}
	cpuUsagePXDTO.ResourceECSMarkInfo = getECSMarkInfo(instance, regionId)
	cpuUsagePXDTO.ResourcePX = pxDTO
	cpuUsagePXDTO.SubResourceType = constant.ECSCPU

	return cpuUsagePXDTO
}
