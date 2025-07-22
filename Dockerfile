# --- Build stage ---
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o fi-mcp-server main.go

# --- Runtime stage ---
FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/fi-mcp-server ./
COPY --from=builder /app/static ./static
COPY --from=builder /app/test_data_dir ./test_data_dir
EXPOSE 8080
ENV FI_MCP_PORT=8080
ENTRYPOINT ["/app/fi-mcp-server"] 