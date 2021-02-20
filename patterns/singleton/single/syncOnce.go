package single

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleTwo struct {
}

var singleInstanceTwo *singleTwo

func GetInstanceTwo(wg *sync.WaitGroup) *singleTwo {
	defer wg.Done()

	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating singleTwo instance now...")
			singleInstanceTwo = &singleTwo{}
		})
	} else {
		fmt.Println("singleTwo already created")
	}

	return singleInstanceTwo
}
