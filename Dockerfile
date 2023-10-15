FROM golang:1.19

EXPOSE 8000

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /mini-balancer

CMD ["/mini-balancer"]
