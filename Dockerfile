# Specify the base image for the go app.
FROM golang

# RUN mkdir /app
# ADD . /app

# # Specify that we now need to execute any commands in this directory.
# WORKDIR /app

WORKDIR /go/src/github.com/postgres-go

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code. Not the slash at the end, as explained in https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

COPY . .

# Importing packages
RUN go get -u github.com/go-chi/chi/v5
RUN go get -u github.com/shopspring/decimal
RUN go get -u golang.org/x/tools
RUN go get -u github.com/dgrijalva/jwt-go
RUN go get -u github.com/google/uuid
RUN go get -u github.com/joho/godotenv
RUN go get -u github.com/lib/pq
RUN go get -u github.com/pressly/goose/v3/cmd/goose
RUN go get -u github.com/pressly/goose/v3

# Build
RUN go build -o main .

# Trying to run goose status
# RUN goose -dir=migrations postgres postgresql://postgres:ijasmoopan@postgres:5432/timenow?sslmode=disable status

# This is for documetnation purposes only. To actually open the port, runtime parameters must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised application can make use of. The value of enviroment variables
# can also be set via parameters supplied to the docker command on the command line.
# ENV HTTP_PORT=8080

# Run
# CMD ["/app/main"]

CMD ["./main"]