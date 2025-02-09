### Dockerfile
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o voicebot

FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/voicebot .
CMD ["./voicebot"]