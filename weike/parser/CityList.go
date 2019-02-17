package parser

import (
	"regexp"
	"crawler/engine"
	"log"
	"crawler/common"
)

var InfoUrlRule = regexp.MustCompile(`<a href="(/show.asp\?id=[0-9]+)">\W+</a></li>`)
func CityList(context []byte)engine.ParserResult  {
	var result engine.ParserResult
	urlArray :=InfoUrlRule.FindAllSubmatch(context, -1)
	urlArrayLen :=len(urlArray)
	log.Println(urlArrayLen)
	if urlArrayLen>2 {
		for _,m := range urlArray {
			result.Requests = append(result.Requests,engine.Request{
				Url:common.WeikeUrl + string(m[1]),
				ParserFunc:ParseProfile,
				})
			result.Items = append(result.Items, string(m[0]))
		}
		return result
	}
	return result
}
