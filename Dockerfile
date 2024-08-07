FROM golang:alpine as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /golang-app

FROM alpine:latest
WORKDIR /
COPY --from=build /golang-app /golang-app
EXPOSE 8080
ENTRYPOINT [ "/golang-app" ]