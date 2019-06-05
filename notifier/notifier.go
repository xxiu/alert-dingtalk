package notifier

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/xxiu/alert-webhook/model"
	"github.com/xxiu/alert-webhook/transformer"
)

 
// Send send markdown message to webhook
func Send(notification model.Notification, defaultwebhook string,tempfile string) (err error) {

	// markdown, robotURL, err := transformer.TransformToMarkdown(notification)
	text,webhook,err := transformer.TransformTemplete(notification,tempfile)
 
	fmt.Println(text)
	
	if err != nil {
		return
	}

	if webhook == "" {
		webhook = defaultwebhook
	}

  
	req, err := http.NewRequest(
		"POST",
		webhook,
		bytes.NewBufferString(text) )
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}
