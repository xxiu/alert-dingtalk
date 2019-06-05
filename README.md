# Alertmanager  Webhook

Prometheus 中的 webhook 触发后的数据结构是固定的。如果要对接到其他提醒的平台,需要一个触发转换的过程。`alert-webhook` 利用json 将 Prometheus 的消息转换为对象后，绑定到指定的模板。然后将模板信息 POST 到指定的地址。 

这样做的好处是只需要一个模板文件就可以匹配到合适的 Webhook 上,方便自定义消息。 

#运行

```
go run main.go -url "https://oapi.dingtalk.com/robot/send?access_token=xxx" -tpl "temp/default.tpl"
```


# docker run 

1. build
   编译 linux 环境使用的文件
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o alert-webhook-amd64 
```

2. docker-compose 编译 image
```
docker-compose build
docker-compose up -d 
``` 

# webhook_config

https://prometheus.io/docs/alerting/configuration/#webhook_config 

prometheus webhook 对应到 go 对象。对象匹配模板后，提交指定的 hook 
模板参考 `temp/default.tpl`

```
{
  "version": "4",
  "groupKey": <string>,    // key identifying the group of alerts (e.g. to deduplicate)
  "status": "<resolved|firing>",
  "receiver": <string>,
  "groupLabels": <object>,
  "commonLabels": <object>,
  "commonAnnotations": <object>,
  "externalURL": <string>,  // backlink to the Alertmanager.
  "alerts": [
    {
      "status": "<resolved|firing>",
      "labels": <object>,
      "annotations": <object>,
      "startsAt": "<rfc3339>",
      "endsAt": "<rfc3339>",
      "generatorURL": <string> // identifies the entity that caused the alert
    },
    ...
  ]
}
```
对应变量 
```
type Alert struct {
	Status      string            `json:"status"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:annotations`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
	GeneratorURL string 		  `json:"generatorURL"`
}

type Notification struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:receiver`
	GroupLabels       map[string]string `json:groupLabels`
	CommonLabels      map[string]string `json:commonLabels`
	CommonAnnotations map[string]string `json:commonAnnotations`
	ExternalURL       string            `json:externalURL`
	Alerts            []Alert           `json:alerts`
}
```
