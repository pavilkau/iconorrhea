# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /src

COPY . .

RUN go mod download
RUN go build

CMD [ "./iconorrhea" ]
