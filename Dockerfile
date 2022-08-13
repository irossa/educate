# Build Stage
FROM golang:1.19.0-alpine3.16 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN go build -o /main-ed

# Run Stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /main-ed /main-ed
COPY app.env .
EXPOSE 5000
CMD ["/main-ed"]
