package common_test

import (
	"Glib/common"
	"fmt"
	"testing"
)

func TestLimiter(t *testing.T) {
	limiter  := common.NewConnLimiter(3)
	stop     := make(chan interface{})
	success  := []int{}
	//执行1000次任务,当达到任务数量阈值时,保存成功的任务编号，并退出任务
	for i := 1; i <= 1000; i++ {
		go func(v int) {
			if err := limiter.GetConn(); err != nil {
				stop <- true
				close(stop)
			} else {
				success = append(success, v)
				stop <- false
			}
		}(i)
		if true == <-stop {
			break
		}
	}
	fmt.Printf("Task is over, success task index is %v", success)
}