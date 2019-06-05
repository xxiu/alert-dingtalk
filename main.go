package main

import (
	"flag"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	model "github.com/xxiu/alert-webhook/model"
	"github.com/xxiu/alert-webhook/notifier"

)

var (
	h            bool
	webHookUrl 		 string
	tempFile  	 string  
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&webHookUrl, "url", "", "webhook url ")
	flag.StringVar(&tempFile,"tpl","temp/default.tpl"," template file  ")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(tempFile)

		err = notifier.Send(notification, webHookUrl,tempFile)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to dingtalk successful!"})

	})
	router.Run(":8080")
}
