package calc

import "log"

type Calc struct {
	log *log.Logger
}

func NewCalc(log *log.Logger) *Calc {
	return &Calc{
		log: log,
	}
}
