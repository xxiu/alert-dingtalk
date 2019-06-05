## Alertmanager Dingtalk Webhook

Webhook service support send Prometheus 2.0 alert message to Dingtalk.

## How To Use

```
cd cmd/webhook
go build
webhook -defaultRobot=https://oapi.dingtalk.com/robot/send?access_token=xxxx
```

```
go run webhook.go -defaultRobot=https://oapi.dingtalk.com/robot/send?access_token=xxxx
```

* -defaultRobot: default dingtalk webhook url, all notifaction from alertmanager will direct to this webhook address.

Or you can overwrite by add annotations to Prometheus alertrule to special the dingtalk webhook for each alert rule.

```
groups:
- name: hostStatsAlert
  rules:
  - alert: hostCpuUsageAlert
    expr: sum(avg without (cpu)(irate(node_cpu{mode!='idle'}[5m]))) by (instance) > 0.85
    for: 1m
    labels:
      severity: page
    annotations:
      summary: "Instance {{ $labels.instance }} CPU usgae high"
      description: "{{ $labels.instance }} CPU usage above 85% (current value: {{ $value }})"
      dingtalkRobot: "https://oapi.dingtalk.com/robot/send?access_token=xxxx"
```

# build
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o alert-dingtalk
```

# DOCKER RUN 
```
docker-compose build
docker-compose up -d 
```

# webhook_config

https://prometheus.io/docs/alerting/configuration/#webhook_config 



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