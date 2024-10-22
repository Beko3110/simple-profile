#Containerize the go application that we have created

# Start with a base image
FROM golang:1.22.5 as base

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files to the contanier
COPY go.mod ./

# Download all the dependencies for the application
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the go application
RUN go build -o main .

#######################
# Reduce the image size using multi-stage builds
# We will use a distroless image to run the application
FROM gcr.io/distroless/base

# Copy the built binary from the base image to the distroless image
COPY --from=base /app/main .

# Copy the static files from the base image to the distroless image
COPY --from=base /app/static ./static

# Expose the port that the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]