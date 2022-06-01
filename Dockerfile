# Importing golang 1.18 to use as a builder for our source
FROM golang:1.18 as builder

# Use the /src directory as a workdir
WORKDIR /src

# Copy the src to /src
COPY . ./

# Download dependencies
RUN go mod download

# Build the executable
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o freepad .

# Import alpine linux as a base
FROM scratch

LABEL version="1.4.0"

# Copy the files from the builder to the new image
COPY --from=builder /src/freepad /app/freepad
COPY --from=builder /src/templates /app/templates
COPY --from=builder /src/static /app/static

# Make /app the work directory
WORKDIR /app

# Expose the listening port
EXPOSE 8080

# Run the program
ENTRYPOINT ["./freepad"]