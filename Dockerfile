FROM golang:1.22
WORKDIR /usr/src/app
COPY ./cmd/main/serviceAccountKey.json go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o ./main ./cmd/main/main.go
CMD ["./main"]