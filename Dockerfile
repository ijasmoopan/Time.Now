FROM golang:alpine

LABEL maintainer="Ijas Moopan"

WORKDIR /go/src/github.com/postgres-go

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
COPY ./.gitignore/.env .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]