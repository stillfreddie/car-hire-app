FROM golang:1.24.4-alpine
WORKDIR /app

COPY . .

RUN go build -o car-hire-app

EXPOSE 8080

CMD ["./car-hire-app"]
