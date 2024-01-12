FROM golang:1.21.0 as builder

WORKDIR /go/src

COPY go.mod go.sum ./

RUN go mod download all

COPY . .

RUN go build -o /go/bin/as.bin -ldflags "-w -s" .

ENTRYPOINT [ "/go/bin/as.bin" ]
