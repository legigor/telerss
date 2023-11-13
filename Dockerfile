FROM golang:1.21.4 as builder
WORKDIR /src
COPY go.mod . 
COPY go.sum .
RUN go mod download
COPY . .
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

FROM scratch
COPY --from=builder /app /app
ENTRYPOINT ["/app"]

