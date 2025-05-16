FROM golang:1.24.3

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

# Install sqlc
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Install migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy start script and make it executable
COPY start.sh .
RUN chmod +x start.sh

# Copy the rest of the application code and build the application
COPY . .
RUN go build -v -o main .

# Run migrations before starting the application
CMD ["./start.sh"]
