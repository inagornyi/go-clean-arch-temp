# Start from golang base image
FROM golang:latest as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build ./cmd

#
FROM alpine:latest

#
WORKDIR /app

#
COPY --from=builder /app /app

# Command to run the executable
CMD ["/app"]