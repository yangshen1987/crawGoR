package model

import "regexp"

var Title  = regexp.MustCompile(`<li class="neirong2">信息标题：.*<b>[\W]+</b>.*`)

type MessageObject struct {
	Title string
	WeiKe
}

func (m *MessageObject)GetMatchRule(attribute string)*regexp.Regexp  {
	switch attribute {
	case "title":
		return Title
	default:
		return m.WeiKe.GetMatchRule(attribute)
	}
}
