# Development stage
FROM golang:1.23.1-alpine AS development

# Install Air for live reloading
RUN go install github.com/air-verse/air@v1.60.0

# Set the working directory
WORKDIR /app

# Copy the .env file explicitly if it exists
COPY .env /app/.env

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Expose port 8081
EXPOSE 8081

# Run Air for live reloading
CMD ["air"]