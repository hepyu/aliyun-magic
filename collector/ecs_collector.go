package collector

import (
	"aliyun-magic/constant"
	//"aliyun-magic/dto"
	"aliyun-magic/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	//"runtime"
	//"sync"
	"fmt"
	"time"
)

var (
	ecs_cost_by_neworder_per1month_data = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "ecs",
		Name:      "cost_by_neworder_per1month",
		Help:      "ecs cost",
	}, []string{"status", "regionId", "instanceId", "instanceType", "applicant", "env", "serverType", "serverName", "owner", "businessLine", "project"})

	ecs_cpu_usage_p50_data = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "ecs_cpu_usage",
		Name:      "p50",
		Help:      "ecs cpu usage",
	}, []string{"status", "regionId", "instanceId", "instanceType", "applicant", "env", "serverType", "serverName", "owner", "businessLine", "project"})

	ecs_cpu_usage_p90_data = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "ecs_cpu_usage",
		Name:      "p90",
		Help:      "ecs cpu usage",
	}, []string{"status", "regionId", "instanceId", "instanceType", "applicant", "env", "serverType", "serverName", "owner", "businessLine", "project"})

	ecs_cpu_usage_p95_data = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "ecs_cpu_usage",
		Name:      "p95",
		Help:      "ecs cpu usage",
	}, []string{"status", "regionId", "instanceId", "instanceType", "applicant", "env", "serverType", "serverName", "owner", "businessLine", "project"})

	ecs_cpu_usage_p99_data = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "ecs_cpu_usage",
		Name:      "p99",
		Help:      "ecs cpu usage",
	}, []string{"status", "regionId", "instanceId", "instanceType", "applicant", "env", "serverType", "serverName", "owner", "businessLine", "project"})
)

func CollectECS() {
	var collect func()
	var t *time.Timer
	collect = func() {
		registry := prometheus.NewRegistry()
		registry.MustRegister(ecs_cost_by_neworder_per1month_data, ecs_cpu_usage_p50_data, ecs_cpu_usage_p90_data, ecs_cpu_usage_p95_data, ecs_cpu_usage_p99_data)
		pusher := push.New(constant.GetPushGatewayAddress(), "aliyun-ecs-stat").Gatherer(registry)

		//runtime.GOMAXPROCS(constant.GetECSCollectorConcurrent())
		regionIdArray := constant.GetRegionId()
		pageSize := constant.GetECSCollectorPageSize()
		for _, regionId := range regionIdArray {
			//获取当前regionId的所有ecs机器
			ecsInstances := service.GetECSInfoArray(regionId, pageSize)

			//根据新购维度计算所有ecs的月成本
			ecsCostDTOArray := service.GetECSCostDTOArray(ecsInstances, regionId, "NewOrder", "Month")
			for _, tobj := range ecsCostDTOArray {
				ecsMarkInfo := tobj.ResourceECSMarkInfo
				ecs_cost_by_neworder_per1month_data.WithLabelValues(ecsMarkInfo.Status, ecsMarkInfo.RegionId, ecsMarkInfo.InstanceId, ecsMarkInfo.InstanceType, ecsMarkInfo.Applicant, ecsMarkInfo.Env, ecsMarkInfo.ServerType, ecsMarkInfo.ServerName, ecsMarkInfo.Owner, ecsMarkInfo.BusinessLine, ecsMarkInfo.Project).Set(tobj.Price)
			}

			//计算昨日cpu使用率的day95,即ecs_cpu_usage_lastday_p95
			//定义Slice
			var ecsSubResourceTypes = []constant.ECSSubResourceType{constant.ECSCPU, constant.ECSMemory}
			var pxes = []constant.PX{constant.PXMax, constant.PXMin, constant.PXP99, constant.PXP95, constant.PXP90, constant.PXP50}

			ecsSubResourceUsageArray := service.GetECSSubResourceUsagePXByYesterday(regionId, ecsInstances, ecsSubResourceTypes, pxes)
			for _, obj := range ecsSubResourceUsageArray {
				if obj.SubResourceType == constant.ECSCPU {
					pxDTO := obj.ResourcePX
					ecsMarkInfo := obj.ResourceECSMarkInfo
					ecs_cpu_usage_p50_data.WithLabelValues(ecsMarkInfo.Status, ecsMarkInfo.RegionId, ecsMarkInfo.InstanceId, ecsMarkInfo.InstanceType, ecsMarkInfo.Applicant, ecsMarkInfo.Env, ecsMarkInfo.ServerType, ecsMarkInfo.ServerName, ecsMarkInfo.Owner, ecsMarkInfo.BusinessLine, ecsMarkInfo.Project).Set(pxDTO.P50)
					ecs_cpu_usage_p90_data.WithLabelValues(ecsMarkInfo.Status, ecsMarkInfo.RegionId, ecsMarkInfo.InstanceId, ecsMarkInfo.InstanceType, ecsMarkInfo.Applicant, ecsMarkInfo.Env, ecsMarkInfo.ServerType, ecsMarkInfo.ServerName, ecsMarkInfo.Owner, ecsMarkInfo.BusinessLine, ecsMarkInfo.Project).Set(pxDTO.P90)
					ecs_cpu_usage_p95_data.WithLabelValues(ecsMarkInfo.Status, ecsMarkInfo.RegionId, ecsMarkInfo.InstanceId, ecsMarkInfo.InstanceType, ecsMarkInfo.Applicant, ecsMarkInfo.Env, ecsMarkInfo.ServerType, ecsMarkInfo.ServerName, ecsMarkInfo.Owner, ecsMarkInfo.BusinessLine, ecsMarkInfo.Project).Set(pxDTO.P95)
					ecs_cpu_usage_p99_data.WithLabelValues(ecsMarkInfo.Status, ecsMarkInfo.RegionId, ecsMarkInfo.InstanceId, ecsMarkInfo.InstanceType, ecsMarkInfo.Applicant, ecsMarkInfo.Env, ecsMarkInfo.ServerType, ecsMarkInfo.ServerName, ecsMarkInfo.Owner, ecsMarkInfo.BusinessLine, ecsMarkInfo.Project).Set(pxDTO.P99)
				}
			}

			//计算昨日内存使用率的day95,即ecs_memory_usage_lastday_p95
		}
		//waitGroup := sync.WaitGroup{}
		//waitGroup.Add()

		//waitGroup.Wait()

		if err := pusher.Add(); err != nil {
			fmt.Println("Could not push to Pushgateway:", err)
		} else {

			fmt.Println("success this time")
		}

		t = time.AfterFunc(time.Duration(1)*time.Second, collect)
	}

	t = time.AfterFunc(time.Duration(1)*time.Second, collect)

	defer t.Stop()
	time.Sleep(time.Minute)
}
