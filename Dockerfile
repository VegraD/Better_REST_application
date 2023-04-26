FROM golang:1.19-alpine

# Set up execution environment in container's GOPATH
WORKDIR /go/src/app/cmd

# Copy relevant folders into container
COPY . /go/src/app/.

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server

EXPOSE 8080

WORKDIR /go/src/app

# Instantiate server
CMD ["./cmd/server"]

