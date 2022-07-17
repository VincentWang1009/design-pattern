package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func initInstance() {
	fmt.Println("init ")
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	once.Do(initInstance)

	return singleton
}
