# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

ARG aero_host
ARG aero_port
ARG aero_namespace
ARG aero_set

ENV HOST=${aero_host}
ENV PORT=${aero_port}
ENV NAMESPACE=${aero_namespace}
ENV SET=${aero_set}

# Set the Current Working Directory inside the container
WORKDIR /app


RUN mkdir /server
RUN mkdir /utility

COPY server server/
COPY utility utility/
COPY .env .
# Copy everything from the current directory to the Working Directory inside the container
COPY *.go .
COPY go.* ./
COPY *.html .

# Build the Go app
RUN go build -o /main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/main"]
