FROM golang:1.13-alpine
WORKDIR /Bookshelf

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

CMD ["/Bookshelf/bin/app"]
EXPOSE 8080