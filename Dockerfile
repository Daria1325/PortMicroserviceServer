FROM golang:1.16

WORKDIR $GOPATH/src/github.com/daria/PortMicroservoce/cmd
COPY . .

RUN go mod tidy
RUN go build ./cmd/server

EXPOSE 9080

ENTRYPOINT ["./server.exe"]

