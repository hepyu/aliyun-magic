package main

import (
	"aliyun-magic/collector"
	"github.com/robfig/cron"
)

//cron语法
//{秒数} {分钟} {小时} {日期} {月份} {星期} {年份(可为空)}

//统一定义所有采集任务，注意时间错位
const (
	cron_collect_ecs_info = "0 0 */1 * * ?" //每天整点执行
)

func main() {
	//创建一个定时器管理器
	crontab := cron.New()
	//添加一个定时任务，第一个参数是cron时间表达式，第二个参数是要触发执行的函数
	//crontab.AddFunc("* * * * * ?", collector.CollectECS)

	//------所有Cron任务------
	crontab.AddFunc(cron_collect_ecs_info, collector.CollectECS)

	//新启一个协程，运行定时任务
	go crontab.Start()
	//是等待停止信号结束任务,注释掉，否则程序终止
	//defer collect_ecs_info.Stop()

	//永久阻塞
	select {}
}
