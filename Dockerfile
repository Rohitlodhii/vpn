FROM golang:1.25-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY cmd ./cmd
COPY protocol ./protocol

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/server ./cmd/server

FROM alpine:3.22

WORKDIR /app

COPY --from=build /out/server /app/server

EXPOSE 9000

ENTRYPOINT ["/app/server"]
