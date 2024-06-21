FROM golang:1.22

RUN apt-get update
RUN apt-get install -y git

WORKDIR /etc/aalstek-website
COPY . .

RUN go install .
RUN go build main.go

CMD ["./main"]