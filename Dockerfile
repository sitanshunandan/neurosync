# 1. Start with a lightweight "Base Image" (Think of this as the empty lunchbox)
# We use Alpine Linux because it's tiny and secure.
FROM golang:alpine

# 2. Install necessary system tools
# We need 'gcc' and 'musl-dev' if we ever switch back to CGO-based SQLite.
# Even for pure Go, it's good practice to have the basics.
RUN apk add --no-cache git build-base

# 3. Set the "Working Directory" inside the container
# This is like creating a folder /app inside the lunchbox.
WORKDIR /app

# 4. Copy the Dependency files first
# We do this separately to use Docker's "Layer Caching" (makes builds faster).
COPY go.mod go.sum ./

# 5. Download the Go modules
RUN go mod download

# 6. Copy the rest of your source code into the container
COPY . .

# 7. Build the application
# We name the binary "neurosync-engine"
RUN go build -o neurosync-engine ./cmd/server/main.go

# 8. Expose the port
# Tell the world that this container listens on port 8080.
EXPOSE 8080

# 9. The Command to Run
# What happens when you open the lunchbox? The engine starts.
CMD ["./neurosync-engine"]