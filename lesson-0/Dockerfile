# Base build image
FROM golang:1.13 as builder
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN CGO_ENABLED=0 go build -o service .

FROM gcr.io/distroless/base
COPY --from=builder /app/service .
ENTRYPOINT ["./service"]