package dto

import (
	"aliyun-magic/constant"
)

//----(1).ECS费用DTO----//
type ECSCostDTO struct {
	ResourceECSMarkInfo *ResourceECSMarkDTO
	Price               float64
}

//----(2).ECS各分类资源使用率情况
type ECSSubResourceUsagePXDTO struct {
	SubResourceType     constant.ECSSubResourceType
	ResourceECSMarkInfo *ResourceECSMarkDTO
	ResourcePX          *PXDTO
}
