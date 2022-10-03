FROM golang:latest AS buildContainer
WORKDIR /go/src/app

# ENV HTTP_PROXY=http://proxy.cht.com.tw:8080
# ENV HTTPS_PROXY=http://proxy.cht.com.tw:8080

COPY . /app
WORKDIR /app
RUN pwd
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=buildContainer /app/bin/main .
COPY --from=buildContainer /app/configYml ./configYml

ENV GIN_MODE release
ENV HOST 0.0.0.0
ENV PORT 9003
EXPOSE 9003

ENV HTTP_PROXY=
ENV HTTPS_PROXY=

CMD ["./main"]
