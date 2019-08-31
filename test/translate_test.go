package test

import (
	"Translate-Api-Go/backend/service"
	"Translate-Api-Go/defs"
	bytes2 "bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestTranslate(t *testing.T) {
	input := defs.Translate{
		Tl:   "en",
		Sl:   "auto",
		Text: "我的天啊",
	}
	result, e := service.Translate(1, &input)
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

func TestAb(t *testing.T) {
	tagUrl := "http://localhost:8083/translate_cn?sl=&tl=zh-Ch&text=helo"

	ints := make(chan int, 100)
	errNum := 0

	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}

	for i := 0; i <= 10000; i++ {
		ints <- 1
		wg.Add(1)
		go func(i int) {
			defer func() {
				<-ints
				wg.Done()
			}()

			get := httplib.Get(tagUrl)
			response, e := get.Response()
			defer response.Body.Close()
			if e != nil {
				panic(e.Error())
			}

			if response.StatusCode != 200 {
				clog.Println("http status !200")

				rw.Lock()
				errNum += 1
				rw.Unlock()

			} else {
				log.Println("http ok !   ", i)
			}
		}(i)
	}

	clog.Println("err: " + strconv.Itoa(errNum))

	wg.Wait()
}
