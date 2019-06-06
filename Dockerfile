FROM alpine:latest
MAINTAINER zhangwei<chxxiu@gmail.com>
RUN echo "https://mirrors.aliyun.com/alpine/latest-stable/main" > /etc/apk/repositories
RUN apk add --update curl bash && rm -rf /var/cache/apk/*
WORKDIR /data
COPY ./alert-webhook-amd64 /data/alert-webhook
EXPOSE 8088
ENTRYPOINT ["/data/alert-webhook"]


 
 
