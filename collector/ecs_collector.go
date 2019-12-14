package collector

import (
	"aliyun-magic/constant"
	//"aliyun-magic/dto"
	"aliyun-magic/service"
	"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/push"
	//"runtime"
	//"sync"
	"fmt"
	"time"
)

var (
	ecsCost = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "ecs",
		Name:      "ecs_cost",
		Help:      "ecs cost",
	}, []string{"status", "regionId", "instanceId", "cpu", "memory", "dsNameEn", "instanceTypeFamily", "instanceName", "instanceType", "ipAddr", "applicant"})

	ecsCpuUsage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "ecs",
		Name:      "ecs_cpu_usage",
		Help:      "ecs cpu usage",
	}, []string{"status", "regionId", "instanceId", "cpu", "memory", "dsNameEn", "instanceTypeFamily", "instanceName", "instanceType", "ipAddr", "applicant"})
)

func CollectECS() {
	var collect func()
	var t *time.Timer

	collect = func() {
		registry := prometheus.NewRegistry()
		registry.MustRegister()
		//pusher := push.New(constant.GetPushGatewayAddress(), "aliyun-ecs").Gatherer(registry)
		//runtime.GOMAXPROCS(constant.GetECSCollectorConcurrent())
		regionIdArray := constant.GetRegionId()
		pageSize := constant.GetECSCollectorPageSize()
		for _, regionId := range regionIdArray {
			instances := service.GetECSCostDTOArray(regionId, pageSize)
			for _, tobj := range instances {
				fmt.Println(tobj.ResourceECSMarkInfo)
			}
		}
		//waitGroup := sync.WaitGroup{}
		//waitGroup.Add()

		//waitGroup.Wait()
		t = time.AfterFunc(time.Duration(1)*time.Second, collect)
	}

	t = time.AfterFunc(time.Duration(1)*time.Second, collect)

	defer t.Stop()
	time.Sleep(time.Minute)
}
