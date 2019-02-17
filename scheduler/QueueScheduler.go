package scheduler

import (
	"crawler/engine"
)

type QueueScheduler struct {
	RequestChan  chan engine.Request
	WorkeChane chan chan engine.Request
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.RequestChan<-r
}
func (s *QueueScheduler) ReadyWorker(w chan engine.Request) {
	s.WorkeChane<-w
}

func (s *QueueScheduler)WorkerChan()chan engine.Request {
	return  make(chan engine.Request)
}

func (s *QueueScheduler)Run()  {
	s.WorkeChane = make(chan chan engine.Request)
	s.RequestChan = make(chan engine.Request)
	go func() {
		var requrstQ []engine.Request//REUQEST队列
		var workerQ []chan engine.Request//WORIKER队列
		for   {
			 var activeRequest engine.Request
			 var activeWorker chan engine.Request
			 //当REQUST和WORKER队列都排队的时候将REUQEST托付给WQRKER
			if len(requrstQ)>0 && len(workerQ)>0 {
				activeRequest = requrstQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.RequestChan:
				requrstQ = append(requrstQ, r)
			case w :=<-s.WorkeChane:
				workerQ = append(workerQ, w)
			case activeWorker<-activeRequest:// 将REUQEST托付给WQRKER 动作
				requrstQ = requrstQ[1:]
				workerQ = workerQ[1:]
			}


		}
	}()
}