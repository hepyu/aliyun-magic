package collector

import (
	"aliyun-magic/constant"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"time"
)

var (
	ecsCost = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: regionId,
		Name:      "ecs_cost",
		Help:      "ecs cost",
	}, []string{"status", "regionId", "instanceId", "cpu", "memory", "dsNameEn", "instanceTypeFamily", "instanceName", "instanceType", "ipAddr", "applicant"})

	ecsCpuUsage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: regionId,
		Name:      "ecs_cpu_usage",
		Help:      "ecs cpu usage",
	}, []string{"status", "regionId", "instanceId", "cpu", "memory", "dsNameEn", "instanceTypeFamily", "instanceName", "instanceType", "ipAddr", "applicant"})
)

func main() {
	var collect func()
	var t *time.Timer

	collect = func() {
		registry := prometheus.NewRegistry()
		registry.MustRegister()
		pusher := push.New(constant.GetPushGatewayAddress, "aliyun-ecs").Gatherer(registry)

	}
}
