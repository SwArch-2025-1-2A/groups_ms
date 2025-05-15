FROM golang:1.24.3

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

# Install sqlc
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . .
RUN go build -v -o main .

CMD ["/app/main"]
