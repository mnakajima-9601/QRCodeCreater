FROM golang:1.7 as builder

WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./
RUN go mod install

COPY QRCode.go  ./

ARG CGO_ENABLED=0
ARG GOARCH=amd64
RUN /bin/sh -c /bin/sh -c go build

CMD ["go", "run", "QRCode.go"]