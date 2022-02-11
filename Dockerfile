# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

ARG environment
ARG aero_host
ARG aero_port
ARG aero_namespace
ARG aero_set

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD "./main" ${environment} ${aero_host} ${aero_port} ${aero_namespace} ${aero_set}
