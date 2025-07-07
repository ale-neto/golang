FROM golang:1.24.2 AS builder

WORKDIR /app
COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY initDependencies.go initDependencies.go
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
 GOOS=linux go build -o meuprimeirocrudgo .

FROM golang:1.24.2 AS runner

RUN useradd -m -s /bin/bash golang

COPY --from=builder /app/meuprimeirocrudgo /app/meuprimeirocrudgo

RUN chown -R golang:golang /app
RUN chmod +x /app/meuprimeirocrudgo

EXPOSE 8080

USER golang

WORKDIR /app

CMD ["./meuprimeirocrudgo"]