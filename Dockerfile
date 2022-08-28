# Specify the base image for the go app.
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Ijas Moopan"

# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/github.com/postgres-go

COPY go.mod .
COPY go.sum .

# Add the go mod download to pull in any dependencies
RUN go mod download

# Copy everything from this project into the filesystem of the container.
COPY . .

# Obtain the package needed to run code. Alternatively use GO modules.
# RUN go get -u github.com/lib/pq

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Compile the binary exe for our app.
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Start the application.
CMD ["./main"]
