FROM golang:1.24
WORKDIR /app
COPY ./app .
RUN go mod download; go build -o server
EXPOSE 8080
CMD ["./server"]