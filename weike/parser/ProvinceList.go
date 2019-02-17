package parser

import (
	"crawler/engine"
	"regexp"
	"crawler/common"
)

const  regexpRule  = `<li><a href="(/class.asp\?typeid=&areaid=[0-9]+)">([^<]+)</a></li>`

func ProviceList(contents []byte) engine.ParserResult {
	ret, err := regexp.Compile(regexpRule)
	if err != nil {
		panic(err)
	}
	mathes := ret.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _,m := range mathes {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: common.WeikeUrl + string(m[1]),
			ParserFunc:CityList,
			})
	}
	return result
}