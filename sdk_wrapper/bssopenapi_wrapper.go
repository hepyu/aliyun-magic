package sdk_wrapper

import (
	"aliyun-magic/constant"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func GetSubscriptionPrice(regionId string, instanceId string, productCode string, orderType string, servicePeriodUnit string, servicePeriodQuantity int, moduleList *[]bssopenapi.GetSubscriptionPriceModuleList) *bssopenapi.GetSubscriptionPriceResponse {

	client, err := bssopenapi.NewClientWithAccessKey("cn-hangzhou", constant.GetAccessKeyID(), constant.GetAccessSecret())

	if err != nil {
		fmt.Println(err)
	}

	request := bssopenapi.CreateGetSubscriptionPriceRequest()
	request.Scheme = "https"

	request.SubscriptionType = "Subscription"
	request.ProductCode = productCode
	request.OrderType = orderType
	//request.ModuleList = &[]bssopenapi.GetSubscriptionPriceModuleList{
	//	{
	//		ModuleCode: "InstanceType",
	//		Config:     "InstanceType:ecs.g5.xlarge",
	//	},
	//}
	request.ModuleList = moduleList
	request.ServicePeriodUnit = servicePeriodUnit
	request.ServicePeriodQuantity = requests.NewInteger(servicePeriodQuantity)
	request.Region = regionId
	request.InstanceId = instanceId

	response, err := client.GetSubscriptionPrice(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	return response
}
