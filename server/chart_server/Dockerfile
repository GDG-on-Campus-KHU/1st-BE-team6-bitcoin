# syntax=docker/dockerfile:1
FROM golang:1.20 as builder

WORKDIR /chart_server

# 프로젝트 루트 디렉토리에서 go.mod와 go.sum을 복사
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM gcr.io/distroless/base-debian10
COPY --from=builder /chart_server/main /
CMD ["/main"]
