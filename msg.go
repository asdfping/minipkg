package minipkg

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func HelloMessage() string {
	return "hello"
}

func EventFire(ch chan<- int) {
	i := 0
	for {
		ch <- i
		i++
	}
}

func RateLimitUsing(ch <-chan int) {
	limit := rate.Every(5 * time.Millisecond)
	fmt.Println(limit)

	limiter := rate.NewLimiter(limit, 1)

	start := time.Now()
	total := 0
	cantExecCnt := 0
	for {
		<-ch
		//fmt.Println("recv:", recv, total)
		if limiter.Allow() {
			//fmt.Println("proc data")
			total++
		}
		if total >= 1000 {
			break
		}
	}

	end := time.Now()
	fmt.Println("use time:", end.Sub(start).Milliseconds(), cantExecCnt)
}

func MockFireAndRecv() {
	ch := make(chan int)

	go EventFire(ch)
	go RateLimitUsing(ch)

	time.Sleep(6 * time.Second)
}
