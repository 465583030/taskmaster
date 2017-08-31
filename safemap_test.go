package taskmaster

import (
	"fmt"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	m := sync.Map{}
	m.Store(1, "1")
	m.Store("2", []string{"om"})

	m.Range(func(k, v interface{}) bool {
		fmt.Println(m.Load(k))
		fmt.Println("k=", k, "v=", v)
		return true
	})
}

//output:
//1 true
//1 1
//[om] true
//2 [om]
