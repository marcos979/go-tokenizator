FROM golang:1.12-alpine

RUN apk add --no-cache git

WORKDIR /app/go-tokenizator

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./out/go-tokenizator .

EXPOSE 3000

CMD ["./out/go-tokenizator"]