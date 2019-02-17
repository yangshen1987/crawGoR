package parser

import (
	"crawler/engine"
	"regexp"
	"log"
	"crawler/weike/model"
)



func ParseProfile(contents []byte) engine.ParserResult  {
	msg := model.MessageObject{}
	var err error
	msg.Address = extractString(
		contents,
		msg.GetMatchRule("address"))
	msg.Name =extractString(
		contents,
		msg.GetMatchRule("name"))
	msg.Age = extractString(
		contents,
		msg.GetMatchRule("age"))
	if err != nil {
		log.Printf("age is empty")
	}
	msg.ContactWay =extractString(
		contents,
		msg.GetMatchRule("source"))
	msg.Appearance =extractString(
		contents,
		msg.GetMatchRule("appearance"))
	msg.ServiceItems =extractString(
		contents,
		msg.GetMatchRule("serviceitems"))
	msg.Pricelist =extractString(
		contents,
		msg.GetMatchRule("pricelist"))
	msg.BusinessHours =extractString(
		contents,
		msg.GetMatchRule("businesshours"))
	msg.Environmentalequipment =extractString(
		contents,
		msg.GetMatchRule("environmentalequipment"))
	msg.Comprehensiveevaluation =extractString(
		contents,
		msg.GetMatchRule("comprehensiveevaluation"))
	msg.Detailedintroduction =extractString(
		contents,
		msg.GetMatchRule("detailedintroduction"))
	msg.Imgs =extractStringForSplipt(
		contents,
		msg.GetMatchRule("imgs"))
	msg.Title = extractString(
		contents,
		msg.GetMatchRule("title"))
	result := engine.ParserResult{
		Items:[]interface{}{msg},
	}
	return result
}

func extractString(contents []byte, rule *regexp.Regexp) string {
	msg := rule.FindSubmatch(contents)
	log.Printf("rule is %v", rule)
	log.Printf("len is %d", len(msg))
	if len(msg)>=2 {
		log.Printf("msg one is %s", string(msg[1]))
		return string(msg[1])
	}else {
		return ""
	}
}
func extractStringForSplipt(contents []byte, rule *regexp.Regexp)[]string  {
	msg := rule.FindSubmatch(contents)
	var result []string
	if len(msg)>=2 {
		for i :=1;i<=len(msg) ;i++  {
			result = append(result,string(msg[i]))
		}
	}
	return result
}
