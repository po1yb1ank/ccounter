FROM golang:1.19-stretch

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV CFG_PATH="./"

WORKDIR /go/ccounter
COPY . .
RUN go mod download all
RUN go build ./cmd/main.go

CMD [ "./main" ]