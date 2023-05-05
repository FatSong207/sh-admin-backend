FROM golang:alpine

RUN mkdir -p /application
WORKDIR /application

COPY . .
RUN  go env \
     && go mod tidy \
     && go build -o application .
ENTRYPOINT ["./application"]