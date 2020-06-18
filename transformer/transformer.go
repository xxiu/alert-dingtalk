package transformer

import (
	"alert-webhook/model"
	"bytes"
	"fmt"
	"text/template"
)

//
func TransformTemplete(notification model.Notification, tempfile string) (text string, webhook string, err error) {

	annotations := notification.CommonAnnotations
	webhook = annotations["webhook"]
	tempfile1 := annotations["tempfile"]
	if tempfile == "" {
		tempfile = tempfile1
	}

	fmt.Println(tempfile)

	tpl := template.Must(template.ParseFiles(tempfile))

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, notification)

	fmt.Println(err)

	text = buf.String()

	return
}

// TransformToMarkdown transform alertmanager notification to dingtalk markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.DingTalkMarkdown, robotURL string, err error) {

	groupKey := notification.GroupKey
	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["dingtalkRobot"]

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 通知组%s(当前状态:%s) \n", groupKey, status))

	buffer.WriteString(fmt.Sprintf("#### 告警项:\n"))

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("##### %s\n > %s\n", annotations["summary"], annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Format("15:04:05")))
	}

	markdown = &model.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Title: fmt.Sprintf("通知组：%s(当前状态:%s)", groupKey, status),
			Text:  buffer.String(),
		},
		At: &model.At{
			IsAtAll: false,
		},
	}

	return
}
