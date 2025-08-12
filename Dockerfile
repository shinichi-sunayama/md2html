# ===== builder =====
FROM golang:1.22.4-bullseye AS builder
WORKDIR /app

# 依存取得（キャッシュ）
COPY go.mod go.sum ./
RUN go mod download

# ソースコピー & ビルド（静的リンク）
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o md2html .

# ===== runner =====
FROM gcr.io/distroless/static:nonroot
WORKDIR /work
COPY --from=builder /app/md2html /usr/local/bin/md2html
USER nonroot
ENTRYPOINT ["/usr/local/bin/md2html"]
