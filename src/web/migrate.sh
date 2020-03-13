#!/bin/sh
# 初回のみDBマイグレーションを実行します。

cd /go/src/github.com/web
if [ -f /go/src/github.com/web/MIGRATED ]; then
  # 2回目以降は実行しない
  echo MIGRATE SKIPPED.
  exit
fi

# マイグレーション実行
cd db
sql-migrate up

cd ../
touch MIGRATED
echo MIGRATE DONE.
