#build stage (Go)
FROM golang:alpine AS backBuilder
WORKDIR /go/src/app
COPY go.* ./
COPY *.go ./
COPY pkg pkg
COPY cmd cmd
RUN apk add --no-cache git
RUN apk add build-base
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /go/bin/app
RUN ls -lah
RUN pwd

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=backBuilder /go/bin/app /app
ENV GIN_MODE=release
ENTRYPOINT ./app
LABEL Name=barker-worker
