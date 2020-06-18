# alert webhook 
1. 根据配置接收一个 json 
2. 指定路由
3. 绑定一个temp 
4. 发送到 webhook




## 配置文件  
配置文件使用 yaml 格式 
```
host: "0.0.0.0"  
prot: 8080

#默认 webhook 
webhook: "https://oapi.dingtalk.com/robot/send?access_token=xxx" 

# 绑定的路由
route:
  - url: "/webhook"    #路由
    tempfile: "temp/default.tpl" # 路由对应的发送模版 
    webhook: "https://oapi.dingtalk.com/robot/send?access_token=xxx" # webhook 不存在使用默认
```


# temp 

prom.tpl 是以前写的 promtheus 的告警 alert，为了对 log 中的某些日志进行告警，做了一些改造。
目的是能够方便的将 log 告警和 prometheus 的告警使用一套推送到钉钉 webhook 。
如果需要改变告警，只需要修改配置和告警模版。 




