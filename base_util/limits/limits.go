package limits

import "github.com/afex/hystrix-go/hystrix"

func InitLimiter() {
	// TODO 初始化限流器
	hystrix.ConfigureCommand("/v1/order", hystrix.CommandConfig{
		Timeout:               10,
		MaxConcurrentRequests: 50,
		ErrorPercentThreshold: 25,
	})
}
