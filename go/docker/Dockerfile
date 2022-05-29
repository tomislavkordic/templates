FROM golang:1.18-bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./

RUN go build -o main.o


FROM debian:bullseye-slim

ENV GIN_MODE=release

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates

WORKDIR /app

COPY --from=build /app .

EXPOSE 8080

ENTRYPOINT ["./main.o"]
