FROM golang:alpine

RUN mkdir -p /app
WORKDIR /app

COPY . .
RUN go env -w GO111MODULE=on     && go env     && go mod tidy     && go build -o app .
ENTRYPOINT ["./app"]