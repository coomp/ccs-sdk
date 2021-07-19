FROM golang:alpine3.13

RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR $GOPATH/src/github.com/coomp/ccs-sdk/example
COPY . $GOPATH/src/github.com/coomp/ccs-sdk/example
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./example"]
