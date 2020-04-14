FROM golang:alpine

# # Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /github.com/siesgstarena/epicentre/src
WORKDIR /github.com/siesgstarena/epicentre/src

# Copy and download dependency using go mod
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

# Copy the code into the container
COPY src .

# Move to working directory /github.com/siesgstarena/epicentre/src
WORKDIR /

# Build the application
RUN go build -o main github.com/siesgstarena/epicentre/src

# Export necessary port
EXPOSE 8000

# Command to run when starting the container
CMD ["main"]