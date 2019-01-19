package counter

import (
	"sync"
	"time"
)

//Counter Interface to implements a Counter
type Counter interface {
	Increment() int
	Decrement() int
}

// CounterRequest Struct a concurrency safe for count request
type CounterRequest struct {
	Accumulator []time.Time `json:accumulator`
	sync.Mutex
	expiration time.Duration
}

//Increment Add 1 to Accumulator
func (c *CounterRequest) Increment() int {
	c.Lock()
	defer c.Unlock()
	c.Accumulator = append(c.Accumulator, time.Now().Add(c.expiration))
	return len(c.Accumulator)
}

//Decrement remove first register from Accumulator
func (c *CounterRequest) Decrement() int {
	c.Lock()
	defer c.Unlock()
	c.Accumulator = c.Accumulator[1:len(c.Accumulator)]
	return len(c.Accumulator)
}

//Len retrieve length from Accumulator
func (c *CounterRequest) len() int {
	c.Lock()
	defer c.Unlock()
	return len(c.Accumulator)
}

func (c *CounterRequest) firstIsExpired() bool {
	c.Lock()
	defer c.Unlock()
	if len(c.Accumulator) > 0 && c.Accumulator[0].Before(time.Now()) {
		return true
	}
	return false
}

// WithExpiration add expirationTime to CounterRequest and Decrement when is expired
func (c *CounterRequest) withExpiration(expirationTime time.Duration) {
	c.expiration = expirationTime
	go func() {
		for {
			if c.firstIsExpired() {
				c.Decrement()
			}
		}
	}()
}

//NewCounterRequest Returns a pointer to CounterRequest with a expiration time
func NewCounterRequest(expiration time.Duration) *CounterRequest {
	c := &CounterRequest{}
	c.withExpiration(expiration)
	return c
}
