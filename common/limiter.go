package common

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

//ChannelLimiter 定义一个了基于缓冲 channel 的限流器,当请求连接时判断channel里面长度是不是大于设定的容量值,
//如果没有超过容量就存入一个值进入channel,如果超过容量则 channel 自动阻塞。
//最后当请求结束的时候,剔除 channel里面的值，释放容量。
type ChannelLimiter struct {
	concurrentConn int
	bucket         chan int
}

//NewChannelLimiter ...
func NewChannelLimiter(capacity int) *ChannelLimiter {
	return &ChannelLimiter{
		concurrentConn: capacity,
		bucket:         make(chan int, capacity),
	}
}

//GetConn 获取通道里面的值
func (cl *ChannelLimiter) GetConn() error {
	if len(cl.bucket) >= cl.concurrentConn {
		return fmt.Errorf("Reached the rate limitation.")
	}
	cl.bucket <- 1
	return nil
}

//ReleaseConn 释放通道里面的值
func (cl *ChannelLimiter) ReleaseConn() int {
	c := <-cl.bucket
	return c
}


//TokenBucketLimiter 定义一个基于令牌桶的限流器
type TokenBucketLimiter struct {
	limiter *rate.Limiter
	rate    float64
	capacity int
}

//NewTokenBucketLimiter 实例化一个令牌限流器
func NewTokenBucketLimiter(r float64, capacity int) *TokenBucketLimiter  {
	return &TokenBucketLimiter{
		limiter: rate.NewLimiter(rate.Limit(r), capacity),
	}
}

//VisitCallBackWithLimiter 接收一个回调函数，在内部通过限流器来控制该函数是否可以执行
func (tb *TokenBucketLimiter) VisitCallBackWithLimiter(fn func(param ...interface{}), params ...interface{}) error {
	if ! tb.limiter.Allow() {
		return fmt.Errorf("rate limit exceeded, try agin later")
	}
	fn(params)
	return nil
}

//AckToken 返回获取令牌的结果，如果没有任何错误，则认为令牌申请成功,反之则失败
func (tb *TokenBucketLimiter) AckToken(second time.Duration) error {
	if second <= 0 {
		if ! tb.limiter.Allow(){
			return fmt.Errorf("rate limit exceeded, try agin later")
		}
		return nil
	}

	c, _ := context.WithTimeout(context.Background(), second)
	if err := tb.limiter.Wait(c); err != nil {
		return fmt.Errorf("get token timeout")
	}

	return nil
}