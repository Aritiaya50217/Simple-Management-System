FROM golang:1.24-alpine


# working directory
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# copy source code ทั้งหมด
COPY . .

# build binary จาก main.go
RUN go build -o app ./cmd/service/main.go

EXPOSE 8080

# run binary
CMD ["./app"]
