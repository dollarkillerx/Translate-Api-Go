package service

import (
	"Translate-Api-Go/defs"
	bytes2 "bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/gcache"
	"github.com/dollarkillerx/easyutils/httplib"
	"strings"
)

// 翻译
func Translate(tag int, data *defs.Translate) (*defs.TranslateResult, error) {

	// 查询缓存中是否存在对应的内容 如果存在 就返回 反之请求 api
	key := easyutils.Sha256Encode(data.Text + data.Sl + data.Tl)
	outData := defs.TranslateResult{}

	get, b := gcache.CacheGet(key)
	if !b {
		var url string
		if tag == 0 {
			url = fmt.Sprintf("https://translate.google.com/m?hl=%s&sl=%s&q=%s", data.Tl, data.Sl, data.Text)
		} else {
			url = fmt.Sprintf("https://translate.google.cn/m?hl=%s&sl=%s&q=%s", data.Tl, data.Sl, data.Text)
		}

		clog.Println(url)

		tagUrl, i := easyutils.UrlEncoding(url)
		if i != nil {
			clog.Println(i.Error())
			return nil, i
		}

		var bytes []byte
		var e error
		for i:=0;i<20;i++ {
			bytes, e = httplib.EuUserGet(tagUrl)
			if e == nil {
				break
			}else if i<3 {
				continue
			}else if i > 3 {
				return nil,errors.New("not data")
			}
		}

		document, i := goquery.NewDocumentFromReader(bytes2.NewBuffer(bytes))
		if i != nil {
			clog.Println(i.Error())
			return nil, i
		}

		text := document.Find("div[class='t0']").Text()

		space := strings.TrimSpace(text)
		gcache.CacheSet(key,space)

		return Translate(tag,data)
	}

	space := get.(string)

	outData.Code = 200
	outData.Msg = space

	return &outData, nil
}
