FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .

RUN go build -o main .

CMD ["./main"]