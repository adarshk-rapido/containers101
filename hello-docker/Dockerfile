FROM golang:1.18.5-alpine3.16 AS build

COPY . /app
WORKDIR /app
RUN go build -o hello-world main.go

FROM alpine:3.16
COPY --from=build /app/hello-world /app/hello-world
CMD [ "/app/hello-world" ]
