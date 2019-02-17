package engine

import (
	"crawler/fetch"
	"log"
)

type SimpleEngine struct {

}
func (e SimpleEngine)Run(seeds ...Request)  {
	var requests []Request
	for _,m := range seeds{
		requests = append(requests, m)
	}
	for len(requests)>0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		parserResult, err := e.Worker(r)
		if err!= nil {
			continue;
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items{
			log.Printf("get it item %s",item)
		}

	}

}
func (e SimpleEngine)Worker(r Request)(ParserResult,error){
	body, err := fetch.Fetch(r.Url)
   if err != nil {
       log.Printf("Fetcher wrong url is %s err is %s", r.Url, err.Error())
      return ParserResult{},err
	}
		parserResult := r.ParserFunc(body)
		return parserResult, nil
}
