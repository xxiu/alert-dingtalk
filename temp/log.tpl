{
     "msgtype": "markdown",
     "markdown": {
         "title":"{{ .msg }}",
         "text": "#### {{ .msg }}  \n
**Level**:{{index . "@level" }} \n
**Host**: {{ .beat.hostname }} \n
**Pid**: {{index  . "@pid" }}  \n
**Err**: {{ .err }}  \n
**Codefile**:{{index  . "@file" }}  \n
**Logfile**:{{ .filename }} \n
**Time**: {{index . "@timestamp" }}\n"

     },
      "at": {
          "atMobiles": [
              ""
          ],
          "isAtAll": false
      }
}
