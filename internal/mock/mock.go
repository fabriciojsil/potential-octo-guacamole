package mock

import "time"

type FakePresenter struct {
	Result interface{}
}

func (f *FakePresenter) Present(data interface{}) {
	f.Result = data
}

type CounterFake struct {
	acc int
}

func (c *CounterFake) Increment() int {
	c.acc = c.acc + 1
	return c.acc
}

func (c *CounterFake) StopExpire() {

}
func (c *CounterFake) Values() (time []time.Time) {
	return
}

func (c CounterFake) SetState(t []time.Time) {

}
