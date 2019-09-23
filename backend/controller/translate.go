package controller

import (
	"Translate-Api-Go/backend/service"
	"Translate-Api-Go/defs"
	"github.com/kataras/iris"
)

func Translate(ctx iris.Context) {
	data := defs.Translate{}

	err := ctx.ReadForm(&data)
	if err != nil {
		ctx.StatusCode(200)
		ctx.JSON(data)
		return
	}

	if data.Sl == "" {
		data.Sl = "auto"
	}

	if data.Text == "" || data.Tl == "" {
		ctx.StatusCode(400)
		ctx.JSON(defs.TranslateResult{Code: 400, Msg: ""})
		return
	}

	result, err := service.Translate(0, &data)

	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(defs.TranslateResult{Code: 400, Msg: ""})
		return
	}

	ctx.JSON(result)
}

func TranslateCn(ctx iris.Context) {
	data := defs.Translate{}

	err := ctx.ReadForm(&data)
	if err != nil {
		ctx.StatusCode(200)
		ctx.JSON(data)
	}

	if data.Sl == "" {
		data.Sl = "auto"
	}

	result, err := service.Translate(1, &data)

	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(defs.TranslateResult{Code: 400, Msg: ""})
		return
	}

	ctx.JSON(result)
}

func DocumentApi(ctx iris.Context) {


	ctx.Header("content-type", "text/html")
	ctx.WriteString(string(htm))
}


var htm = `<!DOCTYPE html>
<html lang=zh-CN>
<head>
    <meta charset=utf-8>
    <meta http-equiv=X-UA-Compatible content="IE=edge">
    <meta name=viewport content="width=device-width,initial-scale=1">
    <title>DollarKiller翻译API</title>

    <meta name=keywords content=mvvm,browser,framework>
    <meta name=author content=DollarKiller翻译API>
    <meta name=founder content=DollarKiller>
    <link href=https://cdn.bootcss.com rel=dns-prefetch>
    <link href=https://api.bootcdn.cn rel=dns-prefetch>
    <link href=https://cdn.jsdelivr.net rel=dns-prefetch>
    <link href=https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css rel=stylesheet>
    <link href=https://cdn.jsdelivr.net/npm/font-awesome@4.7.0/css/font-awesome.min.css rel=stylesheet>
    <link href=https://www.bootcdn.cn/assets/css/site.min.css?1566793064021 rel=stylesheet>
    <!--[if lt IE 9]>
    <script src="https://cdn.jsdelivr.net/npm/html5shiv@3.7.3/dist/html5shiv.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/respond.js@1.4.2/dest/respond.min.js"></script>
    <![endif]-->
    <link rel=apple-touch-icon-precomposed sizes=144x144
          href=https://www.bootcdn.cn/assets/ico/apple-touch-icon-144-precomposed.png?1566793064021>
    <link rel="shortcut icon" href=https://www.bootcdn.cn/assets/ico/favicon.ico?1566793064021>

</head>
<body class=package-template>
<header class=site-header>
    <div class="container jumbotron"><h1>DollarKiller Translate Api</h1>
        <p>高性能翻译API</p></div>
    <div class="package-info well well-md text-center">
    </div>
</header>
<div class="container protocal-notice hidden-xs">
    <div class=row>
        <div class=col-xs-12></div>
    </div>
</div>

<main class="container">
    翻译API文档：
    <br>
    [get|post] https://translate.dollarkiller.com/translate_cn
    <br>
    [get|post] https://translate.dollarkiller.com/translate_cn?sl=&tl=&text=
    <br>
    [get|post] https://translate.dollarkiller.com/translate
    <br>
    [get|post] https://translate.dollarkiller.com/translate?sl=&tl=&text=
    <br>


    sl: 源语言  为空是auto
    <br>
    tl: 目标语言
    <br>
    text: 内容

    <table class="table table-hover">
        <thead>
        <tr>
            <th>#</th>
            <th>语言</th>
            <th>简写</th>
        </tr>
        </thead>
        <tbody>
        <tr>
            <th scope="row">1</th>
            <td>英文</td>
            <td>en</td>
        </tr>
        <tr>
            <th scope="row">2</th>
            <td>中文</td>
            <td>zh-CN</td>
        </tr>
        <tr>
            <th scope="row">3</th>
            <td>法语</td>
            <td>fr</td>
        </tr>
        <tr>
            <th scope="row">4</th>
            <td>德语</td>
            <td>de</td>
        </tr>
        <tr>
            <th scope="row">5</th>
            <td>波兰语</td>
            <td>pl</td>
        </tr>
        <tr>
            <th scope="row">6</th>
            <td>葡萄牙语</td>
            <td>pt</td>
        </tr>
        <tr>
            <th scope="row">7</th>
            <td>俄语</td>
            <td>ru</td>
        </tr>
        <tr>
            <th scope="row">8</th>
            <td>西语</td>
            <td>es</td>
        </tr>
        <tr>
            <th scope="row">9</th>
            <td>丹麦语</td>
            <td>da</td>
        </tr>
        <tr>
            <th scope="row">10</th>
            <td>荷兰语</td>
            <td>nl</td>
        </tr>
        <tr>
            <th scope="row">11</th>
            <td>日语</td>
            <td>ja</td>
        </tr>
        <tr>
            <th scope="row">12</th>
            <td>泰语</td>
            <td>th</td>
        </tr>
        <tr>
            <th scope="row">13</th>
            <td>捷克语</td>
            <td>cs</td>
        </tr>
        <tr>
            <th scope="row">14</th>
            <td>马来语</td>
            <td>ms</td>
        </tr>
        <tr>
            <th scope="row">15</th>
            <td>希伯来语</td>
            <td>iw</td>
        </tr>
        </tbody>
    </table>
</main>

<footer id=footer class="footer hidden-print">

    <div class=copy-right><span>&copy; 9102</span> <a
                href=https://www.github.com/dollarkillerx>Github</a> <span>DollarKiller</span></div>
</footer>
<a href=# id=back-to-top><i class="fa fa-angle-up"></i></a>
<script src=https://cdn.jsdelivr.net/npm/jquery@1.12.4/dist/jquery.min.js></script>
<script src=https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js></script>
<script src=https://cdn.jsdelivr.net/npm/@msmiley/geopattern@1.2.3-1/js/geopattern.min.js></script>
<script src=https://cdn.jsdelivr.net/npm/clipboard@1.5.16/dist/clipboard.min.js></script>
<script src=https://cdn.jsdelivr.net/npm/localforage@1.4.2/dist/localforage.min.js></script>
<script src=https://cdn.jsdelivr.net/npm/lodash@4.17.4/lodash.min.js></script>
<script src=https://www.bootcdn.cn/assets/js/site.min.js?1566793064021></script>
</body>
</html>`