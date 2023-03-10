# Build environment
# -----------------
FROM golang:1.13-alpine as build-env
#Doesn't matter what name you use for the workdir as long as it is consistent throughout the dockerfile
WORKDIR /Bookshelf

RUN apk update && apk add --no-cache gcc musl-dev git 

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash

COPY --from=build-env /Bookshelf/bin/app /Bookshelf/
COPY --from=build-env /Bookshelf/bin/migrate /Bookshelf/
COPY --from=build-env /Bookshelf/migrations /Bookshelf/migrations
EXPOSE 8080

CMD ["/Bookshelf/bin/app"]

# Always run docker-compose down before running docker-compose up to stop running containers 
# and remove containers, networks, volumes, and images created by previous up commands.