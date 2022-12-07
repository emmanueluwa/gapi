FROM golang:1.19.0

WORKDIR /usr/src/app

#hot reload for browser
RUN go install github.com/cosmtrek/air@latest

#copy all files from host to container
COPY . .

#ensure all req packages are installed
RUN go mod tidy