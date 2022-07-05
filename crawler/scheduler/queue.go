package scheduler

import "crawler.com/oys/learngo/engine"

type QueueScheduler struct {
	RequestChan chan engine.Request
	WorkerChan chan chan engine.Request
}

func (q *QueueScheduler) Submit(request engine.Request) {
	q.RequestChan <- request
}

func (q *QueueScheduler) WorkReady(w chan engine.Request) {
	q.WorkerChan <- w
}

func (q *QueueScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {

}

func (s *QueueScheduler) Run() {
	s.WorkerChan = make(chan chan engine.Request)
	s.RequestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-s.WorkerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}




