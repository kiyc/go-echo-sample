FROM golang:alpine

WORKDIR /go/src

COPY ./src /go/src

#ENV GOPATH ''

RUN apk update && \
  apk add --no-cache git && \
  go get github.com/labstack/echo/v4 && \
  go get github.com/go-sql-driver/mysql && \
  go get github.com/jinzhu/gorm && \
  go get github.com/joho/godotenv && \
  go get github.com/oxequa/realize

RUN go build -o server main.go

#CMD ["go", "run", "main.go"]
