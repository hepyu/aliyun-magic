package constant

import (
	"os"
)

func GetPushGatewayAddress() string {
	return os.Getenv("PushGatewayAddress")
}

func GetAccessKeyID() string {
	return os.Getenv("AccessKeyID")
}

func GetAccessSecret() string {
	return os.Getenv("AccessKeySecret")
}

func GetRegionId() string {
	return os.Getenv("RegionID")
}
