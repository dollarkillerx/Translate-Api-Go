package test

import (
	"Translate-Api-Go/backend/service"
	"Translate-Api-Go/defs"
	bytes2 "bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/httplib"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestTranslate(t *testing.T) {
	input := defs.Translate{
		Tl:   "en",
		Sl:   "auto",
		Text: "我的天啊",
	}
	result, e := service.Translate(1,&input)
	if e != nil {
		t.Log(e.Error())
		return
	}
	t.Log(result)
}

func TestStr(t *testing.T) {
	tagUrl := fmt.Sprintf("https://translate.google.cn/m?hl=%s&sl=%s&q=%s", "pt", "en", "hello dollarkiller")
	url, i := easyutils.UrlEncoding(tagUrl)
	if i != nil {
		panic(i)
	}

	bytes, e := httplib.EuUserGet(url)
	if e != nil {
		panic(e.Error())
	}

	ioutil.WriteFile("one.html", bytes, 00666)
}

func TestAsl(t *testing.T) {
	bytes, e := ioutil.ReadFile("one.html")
	if e != nil {
		panic(e.Error())
	}

	document, e := goquery.NewDocumentFromReader(bytes2.NewBuffer(bytes))
	if e != nil {
		panic(e.Error())
	}

	text := document.Find("div[class='t0']").Text()
	space := strings.TrimSpace(text)

	log.Println(space)

}
