package counter

import (
	"sync"
	"time"
)

//Counter Interface to implements a Counter
type Counter interface {
	Increment() int
	StopExpire()
	Values() []time.Time
	RestoreState([]time.Time)
}

//CounterRequest Struct a concurrency safe for count request
type CounterRequest struct {
	Accumulator []time.Time
	expiration  time.Duration
	ticker      *time.Ticker
	stop        chan struct{}
	sync.RWMutex
}

//Increment Add 1 to Accumulator
func (c *CounterRequest) Increment() int {
	c.append(time.Now().Add(c.expiration))
	c.RLock()
	defer c.RUnlock()
	return len(c.Accumulator)
}

//Values retrieve slice from Accumulator
func (c *CounterRequest) Values() (acc []time.Time) {
	c.RLock()
	defer c.RUnlock()
	acc = c.Accumulator
	return acc
}

//SetState Set a previous state to Accumulator
func (c *CounterRequest) SetState(t []time.Time) {
	c.append(t...)
}

//StopExpire Stops expires
func (c *CounterRequest) StopExpire() {
	c.stop <- struct{}{}
}

func (c *CounterRequest) removeExpireds() {
	c.Lock()
	defer c.Unlock()
	y := c.Accumulator[:0]
	for _, n := range c.Accumulator {
		if n.After(time.Now()) {
			y = append(y, n)
		}
	}
	c.Accumulator = y
}

// WithExpiration add expirationTime to CounterRequest and Decrement when is expired
func (c *CounterRequest) withExpiration() {
	c.stop = make(chan struct{})
	go func() {
		for {
			select {
			case <-c.ticker.C:
				c.removeExpireds()
			case <-c.stop:
				c.ticker.Stop()
			}
		}
	}()
}

func (c *CounterRequest) append(t ...time.Time) {
	c.Lock()
	c.Accumulator = append(c.Accumulator, t...)
	c.Unlock()
}

//NewCounterRequest Returns a pointer to CounterRequest with a expiration time
func NewCounterRequest(expiration time.Duration, ticker *time.Ticker) *CounterRequest {
	c := &CounterRequest{}
	c.ticker = ticker
	c.expiration = expiration
	c.withExpiration()
	return c
}
