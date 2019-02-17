package model

import (
	"crawler/modelCommon"
	"regexp"
)


var Name =  regexp.MustCompile(`<li class="neirong2">信息标题.*<b>([^<]+)<`)
var Source  = regexp.MustCompile(`<li class="neirong2">信息来源[^>]+">([^<]+)<`)
var Address  =regexp.MustCompile(`<li class="neirong2">详细地址[^>]+">([^<]+)<`)
var Num  =regexp.MustCompile(`<li class="neirong2">详细地址[^>]+">([^<]+)<`)
var Age  =regexp.MustCompile(`<li class="neirong2">小姐年龄[^>]+">([^<]+)<`)
var Quality  =regexp.MustCompile(`<li class="neirong2">小姐素质[^>]+">([^<]+)<`)
var Appearance  =regexp.MustCompile(`<li class="neirong2">小姐外形[^>]+">([^<]+)<`)
var ServiceItems  =regexp.MustCompile(`<li class="neirong2">服务项目[^>]+">([^<]+)<`)
var Pricelist  =regexp.MustCompile(`<li class="neirong2">价格一览[^>]+">([^<]+)<`)
var BusinessHours  =regexp.MustCompile(`<li class="neirong2">营业时间[^>]+">([^<]+)<`)
var Environmentalequipment  =regexp.MustCompile(`<li class="neirong2">环境设备[^>]+">([^<]+)<`)
var Safetyassessment  =regexp.MustCompile(`<li class="neirong2">安全评估[^>]+">([^<]+)<`)
var Comprehensiveevaluation  =regexp.MustCompile(`<li class="neirong2">综合评价[^>]+">([^<]+)<`)
var Detailedintroduction  =regexp.MustCompile(`><p>(.*)</p`)
//var Detailedintroduction  =regexp.MustCompile(`><p>[^<]+<`)
var Imgs  = regexp.MustCompile(`<img[^>]*style=[^>]*src[=\"\'\s]+（[^\.]*\/([^\.]+)\.[^\"\']+）[\"\']?`)
//var Imgs  = regexp.MustCompile(`<img[^>]*style=[^>]*src[=\"\'\s]+（[^\.]*\/([^\.]+)\.[^\"\']+）[\"\']?`)


type WeiKe struct {
	modelCommon.People
	num int
	PeopleQuality string
	Appearance string
	ServiceItems string
	BusinessHours string
	Environmentalequipment string
	Safetyassessment string
	Comprehensiveevaluation string
	Detailedintroduction string
	Imgs []string
}

func (m *WeiKe)GetMatchRule(attribute string )*regexp.Regexp  {
	switch attribute {
	case "name":
		return Name
	case "address":
		return Address
	case "num":
		return Num
	case "quality":
		return Quality
	case "age":
		return Age
	case "source":
		return Source
	case "appearance":
		return Appearance
	case "serviceitems":
		return ServiceItems
	case "pricelist":
		return Pricelist
	case "businesshours":
		return BusinessHours
	case "environmentalequipment":
		return Environmentalequipment
	case "safetyassessment":
		return Safetyassessment
	case "comprehensiveevaluation":
		return Comprehensiveevaluation
	case "detailedintroduction":
		return Detailedintroduction
	case "imgs":
		return Imgs
	}
	return nil
}
