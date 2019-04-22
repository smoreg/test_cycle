package serviceB

import "github.com/smoreg/test_cycle/services/serviceA"

func New() Service {
	return Service{}
}

type Service struct{}

func (s Service) MethodB1() {
	/* some logic */
}

func (s Service) MethodB2() {
	/* some logic */

	a := serviceA.New()
	a.MethodA1()

	/* some more logic */
}

/* some tons of code */
