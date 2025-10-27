FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o minusone-demo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

COPY --from=builder /app/minusone-demo .
RUN chown appuser:appuser minusone-demo

USER appuser
EXPOSE 8080

CMD ["./minusone-demo"]
