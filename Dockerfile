FROM golang:alpine AS builder
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/main /app/main
ADD ./src/store/ /app/src/store/
EXPOSE 8000

ENTRYPOINT ["/app/main"]



