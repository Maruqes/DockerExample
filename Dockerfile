FROM golang:1.21-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
ENV PORT=8080

ENTRYPOINT ["./server"]
