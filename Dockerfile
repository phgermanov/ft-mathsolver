FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /mathsolver ./cmd/mathsolver

EXPOSE 8080

CMD ["/mathsolver"]