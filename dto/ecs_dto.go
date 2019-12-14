package dto

type ECSCostDTO struct {
	ResourceECSMarkInfo *ResourceECSMarkDTO
	Price               float64
}

type ECSCpuUsageDTO struct {
	ResourceECSMarkInfo *ResourceECSMarkDTO
	CPUUsage            float64
}
