FROM golang:latest

# Make folder to put our files in
RUN mkdir -p /usr/app-ais
RUN mkdir -p /usr/app-ais/backend

# Set working directory so that all subsequent command runs in this folder
WORKDIR /usr/app-ais/backend

# Enable relevant variables for the build
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPATH /app-ais

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

# RUN go build -o /go-docker-demo
RUN go build -o main && chmod 755 main

EXPOSE 8080

CMD [ "./main" ]