FROM golang:alpine AS builder
LABEL stage=gobuilder

EXPOSE 80

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /app/main ./cmd/api/main.go


FROM scratch

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /build/.env /app/.env

CMD ["./main"]
