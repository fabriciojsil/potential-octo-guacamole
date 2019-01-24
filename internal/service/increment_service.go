package service

import (
	"github.com/fabriciojsil/potential-octo-guacamole/internal/gateway/counter"
	"github.com/fabriciojsil/potential-octo-guacamole/internal/presenter"
)

//IncrementService It is a Service to Increment and present the result
type IncrementService struct {
	Counter   counter.Counter
	Presenter presenter.Presenter
}

//Run execute the process to add and present
func (c *IncrementService) Run() {
	total := c.Counter.Increment()
	data := struct{ Count int }{Count: total}
	c.Presenter.Present(data)
}
