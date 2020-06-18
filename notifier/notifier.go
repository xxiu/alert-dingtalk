package notifier

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"alert-webhook/model"
	"alert-webhook/transformer"
)

// 数据模版格式化为字符串
func TempMust(data map[string]interface{}, tempfile string) (str string, err error) {
	tpl := template.Must(template.ParseFiles(tempfile))

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// send data to string
func SendData(text string, webhook string) (err error) {
	if text == "" {
		return errors.New("text is enpty")
	}
	if webhook == "" {
		return errors.New("webhook url is enpty")
	}
	req, err := http.NewRequest(
		"POST",
		webhook,
		bytes.NewBufferString(text))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Send send markdown message to webhook
func Send(notification model.Notification, defaultwebhook string, tempfile string) (err error) {

	// markdown, robotURL, err := transformer.TransformToMarkdown(notification)
	text, webhook, err := transformer.TransformTemplete(notification, tempfile)

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
		bytes.NewBufferString(text))
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
