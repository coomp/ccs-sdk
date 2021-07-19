FROM golang:alpine3.13

RUN apk update \
    && apk add tzdata git \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR $GOPATH/src/github.com/coomp
RUN git clone https://github.com/coomp/ccs-sdk.git && cd ccs-sdk/example && go build .
EXPOSE 8080
ENTRYPOINT ["./example"]
