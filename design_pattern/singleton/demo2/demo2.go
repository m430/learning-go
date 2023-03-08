package main

import (
	"fmt"
	"sync"
)

// 我们可以利用sync.Once仅执行一次的特性实现单例模式
var once sync.Once

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
