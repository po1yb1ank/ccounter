FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY . .
COPY config.yaml /app/config.yaml

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

RUN go mod download all
RUN go build ./cmd/main.go

EXPOSE 8080

CMD [ "./main" ]