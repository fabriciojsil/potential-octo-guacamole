package service

import (
	"github.com/fabriciojsil/potential-octo-guacamole/internal/gateway/counter"
	"github.com/fabriciojsil/potential-octo-guacamole/internal/presenter"
)

//CounterService It is a Service to Increment and present the result
type CounterService struct {
	Counter   counter.Counter
	Presenter presenter.Presenter
}

//Run execute the process to add and present
func (c *CounterService) Run() {
	acc := c.Counter.Increment()
	data := struct{ Count int }{Count: acc}
	c.Presenter.Present(data)
}
