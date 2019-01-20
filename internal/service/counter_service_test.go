package service

import (
	"reflect"
	"testing"
)

func TestCounterService(t *testing.T) {

	t.Run("CounterService Running counter and present result", func(t *testing.T) {
		fakePresenter := &FakePresenter{}
		cs := &CounterService{
			Counter:   &CounterFake{},
			Presenter: fakePresenter,
		}
		cs.Run()

		if !reflect.DeepEqual(fakePresenter.result, struct{ Count int }{Count: 1}) {
			t.Errorf("The Result is diferent, expected %v Actual %v", fakePresenter.result, fakePresenter.result)
		}

	})
}

type FakePresenter struct {
	result interface{}
}

func (f *FakePresenter) Present(data interface{}) {
	f.result = data
}

type CounterFake struct {
	acc int
}

func (c *CounterFake) Increment() int {
	c.acc = c.acc + 1
	return c.acc
}

func (c *CounterFake) Decrement() int {
	c.acc = c.acc - 1
	return c.acc
}
