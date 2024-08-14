FROM golang:1.23-rc-bookworm

WORKDIR /simple_backend

RUN go mod init simple_backend

RUN go get github.com/gin-gonic/gin

RUN go get github.com/gin-contrib/cors

COPY main.go .

CMD ["go", "run", "main.go"]