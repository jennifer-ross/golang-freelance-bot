FROM golang:1.23.4-alpine3.20 AS builder

# Update package manager and install necessary tools
RUN apk update

# Set the working directory in the container
WORKDIR /opt/app

ARG CGO_ENABLED
ARG GO111MODULE

ENV CGO_ENABLED=${CGO_ENABLED}
ENV GO111MODULE=${GO111MODULE}

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into
COPY . .

# Build the Go application
RUN GOOS=linux go build -o ./build/app

FROM golang:1.23.4-alpine3.20 AS final

# Set the working directory in the container
WORKDIR /opt/app/

COPY --from=builder /opt/app/build/app /opt/app/app

ENTRYPOINT ["/opt/app/app"]