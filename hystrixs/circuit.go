package hystrix

import (
	"errors"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

const (
	CircuitNewerTask = "newer_task"
)

var HystrixMap = map[string]hystrix.CommandConfig{
	CircuitNewerTask: hystrix.CommandConfig{
		Timeout:                500,  // 超时时间
		MaxConcurrentRequests:  10,   // 并发请求数
		SleepWindow:            5000, // 熔断开启后，尝试打开的时间，ms
		RequestVolumeThreshold: 10,   // 统计的请求跨度，最少达到这个值，才会计算
		ErrorPercentThreshold:  50,   // 错误熔断比例
	},
}

func InitHystrix() {
	hystrix.Configure(HystrixMap)
}

func taskApi() error {
	return nil
}

func taskApiError() error {
	return errors.New("error")
}

func GetCircuit(name string) *hystrix.CircuitBreaker {
	cir, _, _ := hystrix.GetCircuit(name)
	return cir
}

func test1() {
	var success bool
	cir := GetCircuit(CircuitNewerTask)
	for i := 0; i < 50; i++ {
		start1 := time.Now()
		hystrix.Do("newer_task", func() error {
			err := taskApi()
			if err != nil {
				fmt.Println("live.NewerTaskList.Error", err)
				return err
			}
			success = true
			return nil
		}, func(err error) error {
			fmt.Println("live.NewerTaskList.failed", err)
			return nil
		})

		fmt.Println("hystrix", "请求次数:", i+1, ";用时:", time.Now().Sub(start1), ";请求状态 :", success, ";熔断器开启状态:", cir.IsOpen(), "请求是否允许：", cir.AllowRequest())
	}
}