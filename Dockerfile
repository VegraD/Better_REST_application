FROM golang:1.19-alpine

# Set up execution environment in container's GOPATH
WORKDIR /go/src/app/cmd

# Copy relevant folders into container
COPY ./cmd /go/src/app/cmd
COPY ./constants /go/src/app/constants
COPY ./database /go/src/app/database
COPY ./handlers /go/src/app/handlers
COPY ./json_coder /go/src/app/json_coder
COPY ./res /go/src/app/res
COPY ./responses /go/src/app/responses
COPY ./static /go/src/app/static
COPY ./structs /go/src/app/structs
COPY ./utils /go/src/app/utils
COPY ./webhooks /go/src/app/webhooks
COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server

# Retrieve binary from builder container
COPY ./assignment-2-key.json /go/src/app/assignment-2-key.json

# Instantiate server
CMD ["./server"]

