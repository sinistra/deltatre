# Start from golang v1.13 base image
FROM golang:1.13 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the src directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

#Test the Go app
#RUN go test .
RUN go test -c ./services

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app .
RUN go build .


######## Start a new stage from scratch #######
FROM scratch

# copy app into build
WORKDIR /
COPY --from=builder /app .
# COPY ./.env .
# COPY run.sh .

EXPOSE 8001
ENV GIN_MODE=release

# Command to run the executable
CMD ["./backend"]