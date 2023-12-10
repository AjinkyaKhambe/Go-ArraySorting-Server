# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Expose the port the app runs on
EXPOSE 8000

# Set the default command to run when the container starts
CMD ["./myapp"]
