# golang最新インストール
FROM golang:latest

RUN apt-get update
RUN apt-get install -y python3 python3-pip
RUN pip3 install wait-for-it

# mysql のドライバ
RUN go get github.com/go-sql-driver/mysql;
# realizeダウンロードしてホットリロードできるようにする
RUN go get github.com/oxequa/realize
# migrationツール
RUN go get github.com/rubenv/sql-migrate/...

# アプリケーション
WORKDIR /go/src/github.com/
RUN mkdir /web
WORKDIR /go/src/github.com/web
ADD web/ /go/src/github.com/web
RUN chmod 770 /go/src/github.com/web/migrate.sh

EXPOSE 8080 3306
