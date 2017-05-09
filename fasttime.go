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
	go func() {
		for {
			time.Sleep(time.Second * 1)
			curTime = time.Now()
		}
	}()
}

func main() {
	for {
		fmt.Printf("%v\n", Now())
		time.Sleep(time.Millisecond * 50)
	}

}
