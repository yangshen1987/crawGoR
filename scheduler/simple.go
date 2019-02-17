package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	WorkerChanSimple chan engine.Request
}

func (s *SimpleScheduler)WorkerChan()chan engine.Request {
	return  s.WorkerChanSimple
}
func (s *SimpleScheduler)Submit(r engine.Request)  {
	go func() {s.WorkerChanSimple<-r}()
}
func (s *SimpleScheduler)Run()  {
	s.WorkerChanSimple = make(chan engine.Request)
}
func (s *SimpleScheduler)ReadyWorker(r chan engine.Request) {
}