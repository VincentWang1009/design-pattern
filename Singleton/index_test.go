package singleton

import (
	"sync"
	"testing"
)

// verify whether the instances generated are same
func TestInit(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

// veryify whether the func is pa
func TestParallelSingleon(t *testing.T) {
	wg := sync.WaitGroup{}
	const count = 100
	wg.Add(count)
	instances := [count]*Singleton{}
	for i := 0; i < count; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < count-1; i++ {
		if instances[i] != instances[i+1] {
			t.Fatal("instance is not equal")
		}
	}
}
