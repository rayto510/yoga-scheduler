FROM golang:1.24.3-alpine

# Install git and other deps needed for air (and building)
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum and download dependencies first (cache)
COPY apps/api/go.mod apps/api/go.sum ./
RUN go mod download

# Install air for live reload
RUN go install github.com/air-verse/air@latest

# Add GOPATH/bin to PATH so air is found
ENV PATH="/go/bin:${PATH}"

# Copy the rest of your app source
COPY apps/api .

# Build your app binary (optional, air will build on its own)
RUN go build -o server .

# Run air (it will watch files and build/run your app)
CMD ["air"]
