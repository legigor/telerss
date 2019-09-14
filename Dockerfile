FROM golang:1.13.0-stretch as builder
WORKDIR /src
COPY . .
RUN go test ./...
RUN go build -o /telerss

FROM scratch
COPY --from=builder /telerss /telerss
ENTRYPOINT ["/telerss"]

