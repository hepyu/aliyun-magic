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

	ecsCpuUsageData = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "ecs",
		Name:      "cpu_usage",
		Help:      "ecs cpu usage",
	}, []string{"status", "regionId", "instanceId", "instanceType", "applicant", "env", "serverType", "serverName", "owner", "businessLine", "project"})
)

func CollectECS() {
	var collect func()
	var t *time.Timer

	collect = func() {
		registry := prometheus.NewRegistry()
		registry.MustRegister(ecs_cost_by_neworder_per1month_data, ecsCpuUsageData)
		pusher := push.New(constant.GetPushGatewayAddress(), "aliyun-ecs-stat").Gatherer(registry)

		//runtime.GOMAXPROCS(constant.GetECSCollectorConcurrent())
		regionIdArray := constant.GetRegionId()
		pageSize := constant.GetECSCollectorPageSize()
		for _, regionId := range regionIdArray {
			instances := service.GetECSCostDTOArray(regionId, pageSize)
			for _, tobj := range instances {
				ecsMarkInfo := tobj.ResourceECSMarkInfo
				ecs_cost_by_neworder_per1month_data.WithLabelValues(ecsMarkInfo.Status, ecsMarkInfo.RegionId, ecsMarkInfo.InstanceId, ecsMarkInfo.InstanceType, ecsMarkInfo.Applicant, ecsMarkInfo.Env, ecsMarkInfo.ServerType, ecsMarkInfo.ServerName, ecsMarkInfo.Owner, ecsMarkInfo.BusinessLine, ecsMarkInfo.Project).Set(tobj.Price)
			}
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
