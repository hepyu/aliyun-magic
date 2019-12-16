package aliyun_struct_wrapper

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

//----(1).根据InstanceMonitorData中的字段排序----
type ECSInstanceMonitorDataWrapper struct {
	Data  []ecs.InstanceMonitorData
	ASCBy func(p, q *ecs.InstanceMonitorData) bool
}

func (o ECSInstanceMonitorDataWrapper) Len() int { //重写Len()方法
	return len(o.Data)
}

func (o ECSInstanceMonitorDataWrapper) Swap(i, j int) { //重写Swap()方法
	o.Data[i], o.Data[j] = o.Data[j], o.Data[i]
}

func (o ECSInstanceMonitorDataWrapper) Less(i, j int) bool { //重写Less()方法
	return o.ASCBy(&o.Data[i], &o.Data[j])
}
