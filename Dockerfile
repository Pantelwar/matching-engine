FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /dist

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

FROM gcr.io/distroless/static-debian10

COPY --from=builder /dist/main .

EXPOSE 9099
# Command to run when starting the container
CMD ["./main"]