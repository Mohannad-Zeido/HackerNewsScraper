FROM golang:latest

WORKDIR /go/src/hackerNews
COPY . .

RUN go get -d -v ./...
RUN go test -v ./...
RUN go install -v ./...

RUN cp /go/bin/HackerNewsScraper hackernews

RUN ./hackernews --posts 35

CMD []