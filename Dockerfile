FROM golang:1.18

WORKDIR /usr/src/blog-service

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/blog-service .
# RUN go run . &
CMD ["blog-service"]