# Define our base image to use
FROM golang:1.22-alpine

# Set the working directory which is the destination for copying files
WORKDIR /app

# Copy go files from the host machine to the docker image
COPY go.mod ./
COPY go.sum ./

# Download Go modules
#RUN go mod download

# Copy the source code to the docker image
COPY app/ ./

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/digiapp

# The port that the application is going to listen on
EXPOSE 8080

# Run the app
CMD ["/app/digiapp"]