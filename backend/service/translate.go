package service

import (
	"Translate-Api-Go/defs"
	bytes2 "bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"strings"
)

// 翻译
func Translate(tag int,data *defs.Translate) (*defs.TranslateResult, error) {

	outData := defs.TranslateResult{}

	var url string
	if tag == 0 {
		url = fmt.Sprintf("https://translate.google.com/m?hl=%s&sl=%s&q=%s", data.Tl, data.Sl, data.Text)
	}else {
		url = fmt.Sprintf("https://translate.google.cn/m?hl=%s&sl=%s&q=%s", data.Tl, data.Sl, data.Text)
	}


	tagUrl, i := easyutils.UrlEncoding(url)
	if i != nil {
		clog.Println(i.Error())
		return nil, i
	}

	bytes, e := httplib.EuUserGet(tagUrl)
	if e != nil {
		clog.Println(e.Error())
		return nil, e
	}

	document, i := goquery.NewDocumentFromReader(bytes2.NewBuffer(bytes))
	if i != nil {
		clog.Println(i.Error())
		return nil, i
	}

	text := document.Find("div[class='t0']").Text()

	space := strings.TrimSpace(text)

	outData.Code = 200
	outData.Msg = space

	return &outData,nil
}
