{
    "msgtype": "text", 
    "text": {
        "content": " Version:{{.Version }} 
            GroupKey: {{ .GroupKey }} 
            Status: {{ .Status }}
            Receiver: {{ .Receiver }}
            GroupLabels:[
            {{range $i,$v := .GroupLabels}}
                {{$i}}:{{$v}}
            {{end}}]

            CommonLabels:[
            {{range $i,$v := .CommonLabels}}
                {{$i}}:{{$v}}
            {{end}}]    

            CommonAnnotations:[
            {{range $i,$v := .CommonAnnotations}}
                {{$i}}:{{$v}}
            {{end}}]             

            ExternalURL : {{ .ExternalURL }}

            Alerts:[
                {{range .Alerts}}
                    Status: {{.Status}}
                    Labels:[
                        {{range $i,$v := .Labels}}
                            {{$i}}:{{$v}}
                        {{end}}] 
                    Annotations:[
                        {{range $i,$v := .Annotations}}
                            {{$i}}:{{$v}}
                        {{end}}] 
                    StartsAt: {{.StartsAt}}
                    EndsAt: {{.EndsAt}}
                    GeneratorURL: {{.GeneratorURL}}
                {{end}}
            ]

        
        "
    }, 
    "at": {
        "atMobiles": [
            "156xxxx8827", 
            "189xxxx8325"
        ], 
        "isAtAll": false
    }
}