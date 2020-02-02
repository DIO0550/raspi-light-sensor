# golang最新インストール
FROM golang:latest

RUN mkdir /web

WORKDIR /web

ADD web/ /web

# mysql のドライバ
RUN go get github.com/go-sql-driver/mysql;

EXPOSE 8080