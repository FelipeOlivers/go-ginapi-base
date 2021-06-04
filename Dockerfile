FROM golang:1.16.4 as builder

# Set the working directory
WORKDIR /go/src/api

COPY . .

RUN go version && go mod tidy
RUN go build -o ./app main.go

EXPOSE 8000

CMD ["./app"]