FROM alpine:latest
MAINTAINER zhangwei<chxxiu@gmail.com>
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.8/main" > /etc/apk/repositories
RUN apk add --update curl bash && rm -rf /var/cache/apk/*
WORKDIR /data
COPY ./alert-webhook /data/
EXPOSE 8088
ENTRYPOINT ["/data/alert-webhook"]


 
 