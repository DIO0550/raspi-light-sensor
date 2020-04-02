## 開発環境作成
1. DockerfileLambdaを使用して、イメージを作成。
2. 作成したイメージでコンテナを起動する。
```bash
$ docker run -v hoge/raspberry-light-sensor/src/Lambda/:/go/src/github.com/Lambda -itd タグ名 /bin/sh 
```

## fetch
### ビルド
上記Dockerのコンテナに入って以下のコマンドを実行
```
$ cd fetch
$ GOOS=linux GOARCH=amd64 go build -o fetch
```
上記で「fetch」というファイルができる。

## update
### ビルド
上記Dockerのコンテナに入って以下のコマンドを実行
```
$ cd update
$ GOOS=linux GOARCH=amd64 go build -o update
```
上記で「update」というファイルができる。