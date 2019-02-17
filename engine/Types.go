package engine

import "net/http"

type Request struct {
	Url string
	ParserFunc func([]byte)ParserResult
	Cookie *http.Cookie
}
type ParserResult struct {
	Requests []Request
	Items []interface{}
}

func NilParser([]byte)ParserResult  {
	return ParserResult{}
}


