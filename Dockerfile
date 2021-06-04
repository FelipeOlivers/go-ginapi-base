FROM golang:1.16.4 as builder

# Set the working directory
WORKDIR /go/src/api

# Copy all required source, blacklist files that are not required through `.dockerignore`
COPY . .

# Get the vendor library
RUN go version && go mod tidy


CMD ["go", "run", "main.go"]