FROM golang:1.12-alpine3.9 AS build
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add --update git make
RUN git clone https://github.com/zacharychang/leetcode /go/src/github.com/zacharychang/leetcode --depth 1
WORKDIR /go/src/github.com/zacharychang/leetcode/
RUN GOPROXY=https://goproxy.io make build

FROM python:3.7-alpine3.10

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && \
    apk add --update \
        bash \
        git \
        openssh \
        ca-certificates \
        gcc \
        g++

RUN adduser -h /leetcode -D leetcode

COPY --from=build /go/src/github.com/zacharychang/leetcode/client /usr/bin/
COPY --from=build /go/src/github.com/zacharychang/leetcode/requirements.txt /leetcode/

RUN pip install -i https://mirrors.aliyun.com/pypi/simple --no-warn-script-location -r /leetcode/requirements.txt

COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh

USER leetcode

ENTRYPOINT ["/entrypoint.sh"]
