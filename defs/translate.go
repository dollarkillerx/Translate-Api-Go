package defs

// 定义用户传入结构体
type Translate struct {
	Engine string `form:"engine"` // 翻译引擎
	Sl     string `form:"sl"`     // 原语言
	Tl     string `form:"tl"`     // 目标语言
	Text   string `form:"text"`   // 翻译内容
}

// 返回解析内容
type TranslateResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
