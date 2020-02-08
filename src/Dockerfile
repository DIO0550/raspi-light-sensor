# golang最新インストール
FROM golang:latest

WORKDIR /go/src/github.com/
RUN mkdir /web
WORKDIR /go/src/github.com/web
ADD web/ /go/src/github.com/web

# mysql のドライバ
RUN go get github.com/go-sql-driver/mysql;
# realizeダウンロードしてホットリロードできるようにする
RUN go get github.com/oxequa/realize
# migrationツール
RUN go get github.com/rubenv/sql-migrate/...

EXPOSE 8080 5002 3306