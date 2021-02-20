package single

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func GetInstance(wg *sync.WaitGroup) *single {
	defer wg.Done()

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating singleton now...")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}

	return singleInstance
}
