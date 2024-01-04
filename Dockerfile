FROM golang:1.20 AS build
LABEL authors="josephelbacha"

FROM golang:1.20

WORKDIR /go/src/mower-project

# Copy the current directory contents into the container at /go/src/mower-project
COPY . .

# Run the tests in the container
RUN go test -v ./...

# Build the Go app
RUN go get -d -v ./...
RUN go install -v ./...

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the entry point for the container
CMD ["/go/bin/mower-project"]

