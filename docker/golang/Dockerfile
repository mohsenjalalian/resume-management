FROM golang:1.13.6 as builder

# Set the Current Working Directory inside the container
WORKDIR /src/resume-management

# Copy go mod and sum files
COPY src/go.mod .
COPY src/go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY src .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk add bash

RUN apk add poppler-utils

RUN apk --no-cache add ca-certificates

WORKDIR /src/resume-management

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /src/resume-management/main .

# Copy the env file for working with binary file
COPY ./src/.env .

COPY ./src/statics statics

EXPOSE 5600

# Command to run the executable
CMD ["./main"]