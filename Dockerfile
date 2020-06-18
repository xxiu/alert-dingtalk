FROM golang:1.14 AS builder
WORKDIR /code
COPY ./ /code
RUN export GOPROXY=https://goproxy.io \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Package
# Use scratch image
FROM scratch
WORKDIR /root/
COPY --from=builder /code/app /root/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY conf.yaml /root/conf.yaml
EXPOSE 8080
ENTRYPOINT ["/root/app"]

 
 
