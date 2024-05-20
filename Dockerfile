FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o proxy-api

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/proxy-api .

ENTRYPOINT ["./proxy-api"]

ENV PROXY_API_PORT=80

EXPOSE 80