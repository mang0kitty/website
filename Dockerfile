FROM golang:alpine AS builder
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main
ADD ./src/store/ /app/src/store/
EXPOSE 8001

ENTRYPOINT ["/app/main"]



