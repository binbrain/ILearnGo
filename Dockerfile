FROM golang:latest

RUN go get -d -v -u github.com/go-delve/delve/cmd/dlv
RUN go get -d -v -u github.com/golangci/golangci-lint/cmd/golangci-lint
