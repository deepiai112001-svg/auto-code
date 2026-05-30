# Stage 1: build binary
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
# CGO_ENABLED=0 -> binary tinh, chay duoc tren image khong co glibc
RUN CGO_ENABLED=0 go build -o server .

# Stage 2: runtime image sieu nho (~10MB)
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
