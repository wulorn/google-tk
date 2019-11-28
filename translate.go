package tk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var URL string = "http://translate.google.cn/translate_a/single?client=t" +
	"&sl=en&tl=zh-CN&hl=zh-CN&dt=at&dt=bd&dt=ex&dt=ld&dt=md&dt=qca" +
	"&dt=rw&dt=rm&dt=ss&dt=t&ie=UTF-8&oe=UTF-8&clearbtn=1&otf=1&pc=1" +
	"&srcrom=0&ssel=0&tsel=0&kc=2&tk=%s&q=%s"

func Translate(s string) (buf []byte, err error) {
	qStr := url.PathEscape(s)
	tkk, err := GetTKK()
	if err != nil {
		fmt.Println("get tkk err", err)
		return
	}
	t := GetTK(s, tkk)

	urls := fmt.Sprintf(URL, t, qStr)

	client := &http.Client{}
	request, err := http.NewRequest("GET", urls, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	if err != nil {
		panic(err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	buf, err = ioutil.ReadAll(response.Body)
	//fmt.Println(string(buf))
	return
}
