################################################
#          STEP 1: BUILD                       #
################################################
FROM golang:1.26-alpine AS builder

RUN apk update && \
  apk add --no-cache git ca-certificates tzdata && \
  update-ca-certificates

ENV GOTOOLCHAIN=auto
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o app .

################################################
#          STEP 2: RUNTIME (NO APK)            #
################################################
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 3032

USER nonroot:nonroot

CMD ["/app/app"]