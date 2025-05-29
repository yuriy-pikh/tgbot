FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.21 AS builder

WORKDIR /go/src/app
RUN apk add --no-cache make

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build-only

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/tgbot .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./tgbot", "start"]