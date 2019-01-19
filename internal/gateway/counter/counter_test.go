package counter

import (
	"sync"
	"testing"
	"time"
)

func TestCounterRequest(t *testing.T) {
	t.Run("Increment into counter", func(t *testing.T) {
		expected := 1

		counter := NewCounterRequest(time.Minute)
		acc := counter.Increment()

		if acc != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, acc)
		}
	})

	t.Run("Decrement from counter", func(t *testing.T) {
		expected := 1

		counter := NewCounterRequest(time.Minute)
		counter.Increment()
		counter.Increment()
		counter.Decrement()

		if counter.len() != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, counter.len())
		}
	})

	t.Run("Increment into counter concurrently", func(t *testing.T) {
		expected := 10
		counter := NewCounterRequest(time.Minute)

		waitIncrement := executFuncWithGoRoutines(counter.Increment, expected)
		waitIncrement()

		length := counter.len()
		if length != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, length)
		}
	})

	t.Run("Decrement from counter concurrently", func(t *testing.T) {
		expected := 0
		counter := NewCounterRequest(time.Minute)

		createIncrementDecrementWG(counter, 10)
		length := counter.len()
		if length != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, length)
		}
	})

	t.Run("Decrement after expiration time", func(t *testing.T) {
		expected := 0
		counter := NewCounterRequest(time.Millisecond)
		counter.Increment()
		counter.Increment()
		time.Sleep(time.Millisecond * 3)
		length := counter.len()
		if length != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, length)
		}
	})

}

func createIncrementDecrementWG(counter *CounterRequest, times int) {
	waitIncrement := executFuncWithGoRoutines(counter.Increment, times)
	defer waitIncrement()
	waitDecrement := executFuncWithGoRoutines(counter.Decrement, times)
	defer waitDecrement()
}

func executFuncWithGoRoutines(function func() int, times int) func() {
	wg := sync.WaitGroup{}
	for i := 0; i < times; i++ {
		wg.Add(1)
		func(function func() int) {
			defer wg.Done()
			function()
		}(function)
	}
	return wg.Wait
}
