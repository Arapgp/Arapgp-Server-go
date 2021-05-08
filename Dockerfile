FROM golang

ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct
WORKDIR /gopath

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 3000
CMD ["./Arapgp-Server-go"]
