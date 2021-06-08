package common

import (
	"fmt"
)

//ConnLimiter 定义一个了限流器,基于缓冲channel实现,当请求连接时判断channel里面长度是不是大于设定的容量值,
//如果没有超过容量就存入一个值进入channel,如果超过容量则 channel 自动阻塞。
//最后当请求结束的时候,剔除 channel里面的值，释放容量。
type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

//NewConnLimiter ...
func NewConnLimiter(capacity int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: capacity,
		bucket:         make(chan int, capacity),
	}
}

//GetConn 获取通道里面的值
func (cl *ConnLimiter) GetConn() error {
	if len(cl.bucket) >= cl.concurrentConn {
		return fmt.Errorf("Reached the rate limitation.")
	}
	cl.bucket <- 1
	return nil
}

//ReleaseConn 释放通道里面的值
func (cl *ConnLimiter) ReleaseConn() int {
	c := <-cl.bucket
	return c
}
