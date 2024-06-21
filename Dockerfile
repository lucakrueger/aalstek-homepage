FROM golang:1.22

RUN apt-get update
RUN apt-get install -y git

RUN ls
RUN echo "$PWD"

RUN go install .
RUN go build main.go

RUN ./main