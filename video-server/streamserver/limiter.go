package main

import (
	"log"
)

//rate limit (token bucket)

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(concurrentConn int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: concurrentConn,
		bucket:         make(chan int, concurrentConn),
	}
}

func (connLimiter *ConnLimiter) GetConn() bool {
	if len(connLimiter.bucket) >= connLimiter.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	connLimiter.bucket <- 1
	return true
}

func (connLimiter *ConnLimiter) ReleaseConn() {
	currentBucket := <-connLimiter.bucket
	log.Printf("New Connection coming : %d\n", currentBucket)
}
