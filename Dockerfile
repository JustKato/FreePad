FROM golang:1.18 as builder

# Go into the app dir
WORKDIR /app

# Copy all of the source to /app
COPY . ./

# Download dependencies
RUN go mod download

# Build
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o freepad .

FROM alpine

LABEL version="1.4.0"

# Copy the files from the builder to the new image
COPY --from=builder /app/freepad /app/freepad
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static

# Make /app the work directory
WORKDIR /app

# Expose the listening port
EXPOSE 8080

# Run the program
ENTRYPOINT ["./freepad"]