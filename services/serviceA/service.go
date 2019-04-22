package serviceA

import "github.com/smoreg/test_cycle/services/serviceB"

func New() Service {
	return Service{}
}

type Service struct{}

func (s Service) MethodA1() {
	/* some logic */
}

func (s Service) MethodA2() {
	/* some logic */

	b := serviceB.New()
	b.MethodB1()

	/* some more logic */
}

/* some tons of code */
