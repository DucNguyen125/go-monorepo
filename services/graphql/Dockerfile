FROM golang:1.15.6-alpine3.12 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./services/graphql/main.go

# Start a new stage from scratch

FROM alpine:3.12

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main ./

CMD ["./main"] 