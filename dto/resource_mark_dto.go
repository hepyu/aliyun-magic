package dto

type ResourceECSMarkDTO struct {
	Env string //所属环境，取值范围：dev, test, staging, product

	BusinessLine string //所归属的业务线，取值范围有两类，第一类：业务线名称；第二类是技术中台或基础服务或业务中台等多共用部分，如技术中台，基础架构，数据中台等，用户中台，支付中台等，业务线是最大单元
	Project      string //项目，取值范围：实际的项目名称。

	RegionId string
	Status   string

	ServerType string //服务类型，取值范围有两类，第一类：elasticsearch, kafka, rocketmq等；第二类：除第一类统称做“业务服务”。
	ServerName string //服务名称，取值范围有两类，第一类：elasticsearch，kafka, rocketmq等；第二类：具体的业务服务名称，如account, pay, crm等。

	Applicant string //申请人，来自资源申请邮件
	Owner     string //目前所有者，owner是ECS实际所有者

	InstanceId   string
	InstanceType string
}
