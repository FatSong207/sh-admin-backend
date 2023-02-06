FROM golang:alpine

RUN mkdir -p /app
WORKDIR /app

COPY . .
RUN  go env     && go mod tidy     && go build -o app .
ENTRYPOINT ["./app"]