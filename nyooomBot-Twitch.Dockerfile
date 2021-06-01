FROM golang:1.16-alpine

WORKDIR /go/src/app

COPY ./nyooomBot-Twitch .
COPY ./modules ./modules

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["nyooomBot-Twitch"]
