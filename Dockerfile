FROM golang:1.20-alpine AS build
WORKDIR /
COPY . .
RUN go build -o /app /main.go

FROM alpine:3.18
WORKDIR /
COPY --from=build /app /app
ENTRYPOINT [ "/app" ]
