# goバージョン
FROM golang:1.17.6

# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

COPY ../cmd/go.mod /go/src/app
COPY ../cmd/go.sum /go/src/app

COPY ../cmd/qrCode/conf.xml /go/src/app
COPY ../cmd/qrCode/main.go  /go/src/app
COPY ../cmd/qrCode/font /go/src/app/font

ENTRYPOINT ["go", "run", "main.go"]
CMD ["",""]