# BUILDER STAGE

FROM golang:1.25.6-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/api

# RUNNER STAGE

FROM alpine:3.23 AS runner

RUN apk add --no-cache curl

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080


CMD ["./server"]