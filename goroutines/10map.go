package main

import (
	"fmt"
	"sync"
	"time"
)

type S struct {
	sid string
	num int
}

var res = []S{
	{sid: "a", num: 1},
	{sid: "b", num: 2},
	{sid: "c", num: 3},
	{sid: "a", num: 4},
	{sid: "a", num: 5},
	{sid: "b", num: 6},
	{sid: "c", num: 7},
	{sid: "d", num: 8},
	{sid: "e", num: 9},
}

const (
	routineCount = 2
)

func main() {
	var group sync.WaitGroup
	defer group.Wait()
	chanLimit := make(chan struct{}, routineCount+1)
	for i := 0; i < len(res); i++ {
		group.Add(1)
		chanLimit <- struct{}{}
		go routine(i, chanLimit, &group)
	}
}

func routine(i int, chanLimit chan struct{}, group *sync.WaitGroup) {
	defer func(chanLimit chan struct{}) {
		<-chanLimit
	}(chanLimit)
	defer group.Done()

	fmt.Println(i)
	time.Sleep(3 * time.Second)
}
