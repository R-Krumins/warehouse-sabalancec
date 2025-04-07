FROM golang:1.24

# Install dependencies
# sqlite3
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev && rm -rf /var/lib/apt/lists/*

# goose migration tool
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Create data directory for SQLite database with proper permissions
RUN mkdir -p /data && chmod 755 /data

WORKDIR /usr/src/warehouse

# Enviroment variables
ENV PORT=3000
ENV DB_PATH=/data/warehouse.db
ENV GOOSE_DBSTRING=/data/warehouse.db

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/warehouse ./cmd

EXPOSE 3000

CMD ["sh", "-c", "goose up && warehouse"]

