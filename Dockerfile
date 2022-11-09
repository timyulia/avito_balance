FROM golang:1.17

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o balance ./cmd/main.go

CMD ["./balance"]
