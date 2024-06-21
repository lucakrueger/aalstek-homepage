FROM golang:1.22

WORKDIR /usr/src/app

RUN apt-get update
RUN apt-get install -y git

RUN git clone https://github.com/lucakrueger/aalstek-homepage.git

RUN cd aalstek-homepage

RUN go get .
RUN go build main.go

RUN ./main