FROM golang:latest as builder

WORKDIR /app

COPY math.go go.mod /app/

RUN go build -o math

FROM scratch

COPY --from=builder /app/math /app/math

CMD ["/app/math"]