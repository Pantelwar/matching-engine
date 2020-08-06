FROM golang:latest

ENV GO111MODULE=on

WORKDIR /matching-engine

COPY . .

RUN go build main.go

CMD [ "./main" ]

EXPOSE 9096