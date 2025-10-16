# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vint .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/vint .

# Create directories for VintLang files
RUN mkdir -p /usr/local/share/vintlang/examples
RUN mkdir -p /usr/local/share/vintlang/docs

# Copy documentation and examples
COPY --from=builder /app/examples/ /usr/local/share/vintlang/examples/
COPY --from=builder /app/docs/ /usr/local/share/vintlang/docs/
COPY --from=builder /app/README.md /usr/local/share/vintlang/

# Add to PATH
RUN ln -s /root/vint /usr/local/bin/vint

# Create a non-root user
RUN adduser -D -s /bin/sh vint

# Switch to non-root user
USER vint
WORKDIR /home/vint

# Default command
CMD ["vint", "--help"]
