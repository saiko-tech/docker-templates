FROM golang:1.23-alpine AS builder

WORKDIR /go-build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd/app

FROM alpine:3.20

ARG USER=app
ARG UID=1000
ARG GID=1000
ENV HOME=/home/$USER

# don't run app as root
# hadolint ignore=DL3059
RUN addgroup -g $GID $USER && adduser -D -u $UID -G $USER $USER
USER $USER
WORKDIR $HOME

COPY --from=builder /go-build/app .

ENTRYPOINT ["./app"]
