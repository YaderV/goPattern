package singleton_test

import (
	"sync"
	"testing"

	"github.com/YaderV/goPattern/singleton"
)

func TestSingleton(t *testing.T) {
	var f1 *singleton.File
	var f2 *singleton.File

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		f1, _ = singleton.ReadData()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f2, _ = singleton.ReadData()
	}()

	wg.Wait()

	if f1 == nil {
		t.Fatal("f1 object should not be nil")
	}

	if f2 == nil {
		t.Fatal("f2 object should not be nil")
	}

	if f1.Open != f2.Open {
		t.Fatal("the singleton pattern is wrong")
	}

}
