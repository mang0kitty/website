FROM golang:alpine AS builder
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/main /app/main
ADD ./store/ /app/store/
EXPOSE 8001

ENTRYPOINT ["/app/main"]



