package scheduler

import (
	"crawler.com/oys/learngo/engine"
)

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.WorkChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request)  {
	//send request done to work chan
	go func() {
		s.WorkChan <- request
	}()
}
