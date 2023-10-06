# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's working directory
COPY . /app

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Define the command to run your application
CMD ["./main"]