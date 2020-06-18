package main

import (
	"alert-webhook/notifier"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host    string        `yaml:"host"`
	Prot    string        `yaml:"prot"`
	WebHook string        `yaml:"webhook"`
	Route   []ConfigRoute `yaml:"route"`
}

type ConfigRoute struct {
	Url      string `yaml:"url"`
	TempFile string `yaml:"tempfile"`
	WebHook  string `yaml:"webhook"`
}

var (
	conf      string
	routeDict map[string]ConfigRoute
)

func init() {
	flag.StringVar(&conf, "c", "conf.yaml", "config path ,defaule: conf.yaml")
	flag.Parse()

	routeDict = make(map[string]ConfigRoute)

}

func getConf() *Config {

	yamlFile, err := ioutil.ReadFile(conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	var c Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &c
}

func Handle(c *gin.Context) {

	path := c.Request.URL.Path
	if r, ok := routeDict[path]; ok {
		// data, _ := ioutil.ReadAll(c.Request.Body)
		var jsonData map[string]interface{}
		// err := json.Unmarshal(data, &jsonData)

		err := c.BindJSON(&jsonData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("%v ,%v,%s", jsonData, r, path)
		txt, _ := notifier.TempMust(jsonData, r.TempFile)
		fmt.Printf(txt)
		err = notifier.SendData(txt, r.WebHook)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "500", "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "200", "msg": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": "404"})
	}
}

func main() {

	config := getConf()
	fmt.Println(config.Host)
	fmt.Println(config.Prot)
	fmt.Println(config.WebHook)

	gin.SetMode(gin.DebugMode)
	ginroute := gin.Default()
	for _, route := range config.Route {
		if route.WebHook == "" {
			route.WebHook = config.WebHook
		}

		routeDict[route.Url] = route
		ginroute.POST(route.Url, Handle)
	}
	ginroute.Run(fmt.Sprintf("%s:%s", config.Host, config.Prot))
}
