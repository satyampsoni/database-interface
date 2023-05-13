# specifying the base Image
FROM golang:1.16-alpine

# working directory
WORKDIR /dbinterface


# Copy the current directory contents into the container at /dbinterface
COPY . .

# Build the Go application
RUN go build -o main .

# Exposing port 

EXPOSE 8080

# Port 8080 is often used as a default port for running web servers or API endpoints.

#run the executable
CMD ["./main"]






