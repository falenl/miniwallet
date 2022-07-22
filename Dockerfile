FROM golang:alpine AS builder
WORKDIR /app
COPY . .

RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 

RUN go build -o miniwallet


FROM alpine
WORKDIR /app
RUN apk add tzdata
ENV TZ=Asia/Jakarta

COPY --from=builder /app/miniwallet /app/miniwallet
CMD ["./miniwallet"]