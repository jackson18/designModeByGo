package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance *Singleton
)

type Singleton struct {
}

// 通过atomic + mutex方式实现
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		fmt.Println("instance...")
	})
	return instance
}

func main() {
	s := GetInstance()
	s2 := GetInstance()
	fmt.Println("s: ", s, " s2: ", s2)
}
