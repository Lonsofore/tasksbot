FROM golang:alpine AS builder
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a cmd/bot/main.go

FROM alpine AS final
WORKDIR /
COPY --from=builder /workspace/main .
CMD [ "./main" ]
