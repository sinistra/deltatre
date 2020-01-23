# Start from golang v1.13 base image
FROM golang:1.13 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the src directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

#Test the Go app
RUN go test -c ./services

# Build the Go app
RUN CGO_ENABLED=0 go build -a -o /out/frontend .


######## Start a new stage from scratch #######
FROM scratch

# copy app into build
WORKDIR /
COPY --from=builder /out/frontend .
# COPY ./.env .
# COPY run.sh .

EXPOSE 8000
ENV GIN_MODE=release
ENV BACKEND_HOST=http://backend:8001/api/v1

# Command to run the executable
CMD ["./frontend"]