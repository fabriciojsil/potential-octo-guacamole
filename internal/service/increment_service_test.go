package service

import (
	"reflect"
	"testing"

	"github.com/fabriciojsil/potential-octo-guacamole/internal/mock"
)

func TestIncrementService(t *testing.T) {

	t.Run("IncrementService Running counter and present result", func(t *testing.T) {
		fakePresenter := &mock.FakePresenter{}
		cs := &IncrementService{
			Counter:   &mock.CounterFake{},
			Presenter: fakePresenter,
		}
		cs.Run()

		count := struct{ Count int }{Count: 1}
		if !reflect.DeepEqual(fakePresenter.Result, count) {
			t.Errorf("The Result is diferent, expected %v Actual %v", fakePresenter.Result, count)
		}
	})
}
