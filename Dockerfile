# Use a Go builder image
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /src

# Copy the program source code to the container
COPY . .

# Build the Go program
RUN go build -o /app/main

# Use a lightweight base image
FROM ubuntu:jammy

# Set the working directory
WORKDIR /app

# Copy the compiled program binary to the container
COPY --from=builder /app/main .

# Set the command to run when the container starts
CMD ["/app/main"]