FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.21 AS builder

WORKDIR /go/src/app
RUN yum install -y make && yum clean all

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build-only

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/tgbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./tgbot", "start"]