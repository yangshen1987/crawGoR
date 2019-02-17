package main


//ProviceList
import (
	"crawler/engine"
	"crawler/weike/parser"
	"crawler/scheduler"
)

func main()  {
	e :=engine.ConCurrent{
		Scheduler:&scheduler.QueueScheduler{},
		WorkerCount:5,
	}
	e.Run(engine.Request{
		Url:"http://weike46.com/",
		ParserFunc: parser.ProviceList,
	})
}
