FROM golang:latest AS builder

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o example .

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/example .
EXPOSE 8080
CMD ["./example"]
