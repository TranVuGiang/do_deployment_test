FROM golang:1.25-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -i server ./cmd/main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 9190
ENTRYPOINT [ "./server" ]