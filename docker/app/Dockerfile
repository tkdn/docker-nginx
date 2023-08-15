# syntax=docker/dockerfile:1

##### build
FROM golang:1.21-alpine AS build

WORKDIR /go-app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /app

##### deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app /app

EXPOSE 8080

USER nonroot:nonroot

CMD [ "/app" ]
