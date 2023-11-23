FROM golang:latest as build
WORKDIR /app
COPY . .
## Para nao gerar dependencia do C
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api/main.go

## MultiStage build
FROM scratch
COPY --from=build /app/api /
CMD ["./api"]
