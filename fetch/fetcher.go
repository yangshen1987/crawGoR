package fetch

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"bufio"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
)

var rateLimitWait = time.Tick(10 * time.Millisecond)
func Fetch(url string)([]byte, error)  {
	log.Printf("url is %s", url)
	if url == "http://weike46.com/show.asp?id=111263" {
		//return nil, nil
	}
	<-rateLimitWait
	requset, err := http.NewRequest(http.MethodGet, url , nil)
	if err != nil {
		log.Printf("httpRequest client error error is %s", err)
		return nil, err
	}
	cookie1 := &http.Cookie{Name: "urlrefer",Value: "", HttpOnly: true}
	cookie2 := &http.Cookie{Name: "ASPSESSIONIDACDSDBRR",Value: "NCEKPGDDNMBIDPJHHPCAAHPI", HttpOnly: true}
	cookie3 := &http.Cookie{Name: "usercookies%5F1095996",Value: "dayarticlenum=0&daysoftnum=0&userip=183%2E184%2E70%2E55", HttpOnly: true}
	cookie4 := &http.Cookie{Name: "NewAspUsers",Value: "LastTime=2019%2F2%2F4+11%3A53%3A32&LastTimeDate=2019%2F2%2F3+13%3A30%3A49&LastTimeIP=183%2E184%2E68%2E18&userid=1095996&username=ysaibobo&UserClass=0&password=b9f7ae2acaddc151&nickname=ysaibobo&UserGrade=1&UserLogin=54&usermail=my%40email%2Ecom&UserGroup=%C6%D5%CD%A8%BB%E1%D4%B1&userlastip=183%2E184%2E70%2E55&UserToday=0%2C0%2C0%2C0%2C0%2C0&RegDateTime=2018%2D01%2D22+20%3A33%3A58", HttpOnly: true}
	cookie5 := &http.Cookie{Name: "ASPSESSIONIDQQSDCCTS",Value: "OICCFBLDIPMIFEJNMCKDHNGK", HttpOnly: true}
	cookie6 := &http.Cookie{Name: "ASPSESSIONIDQSRDBDTS",Value: "ASPSESSIONIDQSRDBDTS=IHCJCEHDOOCDDLHGLFDGLJGJ", HttpOnly: true}
	requset.AddCookie(cookie1)
	requset.AddCookie(cookie2)
	requset.AddCookie(cookie3)
	requset.AddCookie(cookie4)
	requset.AddCookie(cookie5)
	requset.AddCookie(cookie6)
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			return nil
		},
	}
	resp, err := client.Do(requset)
	if err != nil {
		log.Printf("httpRequest get error error is %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	bodyBufioRead := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyBufioRead)
	utf8Reader :=transform.NewReader(resp.Body, e.NewDecoder())
	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			return nil,err
		}
		return all, nil
	}else {
		log.Printf("httpRequest get error error is %s", resp.StatusCode)
		return nil, fmt.Errorf("获取网站错误 HTTPCODE IS %d", resp.StatusCode)
	}
	return nil, fmt.Errorf("未知错误")
}
func determineEncoding(r *bufio.Reader)encoding.Encoding  {
	contect, err := r.Peek(1024)
	if err != nil {
		log.Printf("determineEncoding is %s", err)
		return unicode.UTF8
	}
	e, _,_ :=charset.DetermineEncoding(contect, "")
	return e
}
