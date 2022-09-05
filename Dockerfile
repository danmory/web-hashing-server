FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

EXPOSE 8080

ENV GIN_MODE release
ENV USE_DB false

CMD ["sh", "-c", "go run . -d=${USE_DB}"]