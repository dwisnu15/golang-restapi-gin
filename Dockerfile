FROM golang:1.13.8-alpine

#Set env variable for docker image
#GOOS should be your OS
ENV GO111MODULE=on \
    CGO_ENABLED=0
    GOOS=linux \
    GOARCH=amd64

#working dir
WORKDIR /build

#Copy dependencies on go.mod and go.sum
COPY go.mod .
COPY go.sum .

#Copy folder into docker container
COPY . .

#Build app command
RUN go build -o main .

#Command for running .exe file

CMD ["/build/main"]


