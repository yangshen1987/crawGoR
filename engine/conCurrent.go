package engine

import (
	"log"
	"crawler/fetch"
)

type ConCurrent struct {
	Scheduler Scheduler
	WorkerCount int
}
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}
type ReadyNotifier interface {
	ReadyWorker(chan Request)
} 
func (c *ConCurrent)Run(seeds ...Request)  {
	our := make(chan  ParserResult)
	c.Scheduler.Run()
	for i:=0;i<c.WorkerCount ;i++  {
		createWorker(c.Scheduler.WorkerChan(), our, c.Scheduler)
	}
	for _, r := range seeds{
		c.Scheduler.Submit(r)
	}
	for  {
		result := <-our
		for _,item := range result.Items {
			log.Printf("get it item %s",item)
		}
		for _,request := range result.Requests{
			c.Scheduler.Submit(request)
		}
	}

}
func createWorker(in chan Request,
	out chan ParserResult,
	r ReadyNotifier)  {
	go func() {
		for  {
			r.ReadyWorker(in)
			request := <-in
			parserResult , err := Worker(request)
			if err != nil{
				continue
			}
			out<-parserResult
		}
	}()
}
func Worker(r Request)(ParserResult,error){
	body, err := fetch.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher wrong url is %s err is %s", r.Url, err.Error())
		return ParserResult{},err
	}
	parserResult := r.ParserFunc(body)
	return parserResult, nil
}
