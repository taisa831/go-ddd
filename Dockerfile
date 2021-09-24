FROM golang:1.14

# コンテナログイン時のディレクトリ指定
WORKDIR /opt/go-ddd

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . .

# ビルド
RUN go build -o app main.go

# 起動
CMD ["/opt/go-ddd/app"]