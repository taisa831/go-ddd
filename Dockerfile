# syntax = docker/dockerfile:experimental
FROM golang:1.16.4-alpine

RUN apk add --no-cache \
	git

RUN GO111MODULE=off go get -u -v \
	# ホットリロードライブラリ
	github.com/oxequa/realize

# コンテナの起動を待つライブラリ
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
	&& tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
	&& rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/go-ddd

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . .

# ビルド
RUN go build -o /opt/app main.go