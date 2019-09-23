package config

import (
	"github.com/dollarkillerx/easyutils/clog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type myconf struct {
	App struct {
		Host       string `yaml:"host"`
		Debug      bool   `yaml:"debug"`
		MaxRequest int    `yaml:"max_request"`
	}
	Mysql struct {
		Dsn   string `yaml:"dsn"`
		Cache bool   `yaml:"cache"`
	}
	Redis struct {
		Maxidle     int           `yaml:"maxidle"`
		MaxActive   int           `yaml:"max_active"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
		Port        string        `yaml:"port"`
	}
}

var (
	MyConfig *myconf
)

func init() {
	MyConfig = &myconf{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		clog.PrintWa("请填写配置文件")
		create_config()
		os.Exit(0)
	}

	e = yaml.Unmarshal(bytes, MyConfig)
	if e != nil {
		panic(e.Error())
	}

	if MyConfig.App.Debug {
		log.Println(MyConfig)
	}
}

func create_config() {
	err := ioutil.WriteFile("config.yml", []byte(conf), 00666)
	if err != nil {
		panic(err)
	}
}

var conf = `
# 翻译配置文件

app:
  host: ":8083"
  debug: true
  max_request: 1000
`
