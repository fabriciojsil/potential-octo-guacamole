package counter

import (
	"sync"
	"testing"
	"time"
)

func TestCounterRequest(t *testing.T) {
	t.Run("Increment into counter", func(t *testing.T) {
		expected := 1
		ticker := time.NewTicker(time.Millisecond)
		counter := NewCounterRequest(time.Minute, ticker)
		acc := counter.Increment()

		if acc != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, acc)
		}
	})

	t.Run("Increment into counter concurrently", func(t *testing.T) {
		expected := 10
		ticker := time.NewTicker(time.Millisecond)
		counter := NewCounterRequest(time.Minute, ticker)

		waitIncrement := executFuncWithGoRoutines(counter.Increment, expected)
		waitIncrement()

		length := len(counter.Values())
		if length != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, length)
		}
	})

	t.Run("Decrement after expiration time", func(t *testing.T) {
		expected := 1
		future := time.Now().Add(time.Minute)
		past := time.Date(2017, 06, 1, 1, 1, 06, 66666, time.UTC)

		ticker := time.NewTicker(time.Nanosecond)
		counter := NewCounterRequest(time.Millisecond, ticker)
		counter.append(past)
		counter.append(future)

		time.Sleep(time.Millisecond)

		length := len(counter.Values())
		if length != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, length)
		}
	})

	t.Run("Restore state", func(t *testing.T) {
		expected := 2
		ticker := time.NewTicker(time.Millisecond)
		counter := NewCounterRequest(time.Millisecond, ticker)
		past := time.Date(2017, 06, 1, 1, 1, 06, 66666, time.UTC)
		future := time.Now().Add(time.Second)
		storaged := []time.Time{past, future}
		counter.SetState(storaged)

		length := len(counter.Values())
		if length != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, length)
		}
	})

	t.Run("Stop Decrement after call Stop func", func(t *testing.T) {
		expected := 2
		ticker := time.NewTicker(time.Millisecond)

		counter := NewCounterRequest(time.Millisecond, ticker)
		past := time.Date(2017, 06, 1, 1, 1, 06, 66666, time.UTC)
		future := time.Now().Add(time.Second)
		counter.append(past, future)

		counter.StopExpire()
		time.Sleep(time.Millisecond)

		length := len(counter.Values())
		if length != expected {
			t.Errorf("The length is diferent, expected %v Actual %v", expected, length)
		}
	})
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
