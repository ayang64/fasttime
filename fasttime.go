package main

import (
	"fmt"
	"time"
)

var curTime time.Time

func Now() time.Time {
	return curTime
}

func init() {
	curTime = time.Now()
	go updateTime()
}

func updateTime() {
	for {
		curTime = time.Now()
		time.Sleep(time.Second * 1)
	}
}

func main() {
	time.Sleep(1 * time.Second)
	fmt.Printf("%v\n", curTime)
	time.Sleep(5 * time.Second)
	fmt.Printf("%v\n", curTime)
}
