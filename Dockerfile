FROM alpine

LABEL version="1.4.0"

# Copy the distribution files
COPY ./dist /app

# Make /app the work directory
WORKDIR /app

# Expose the listening port
EXPOSE 8080

# Run the program
ENTRYPOINT ["./freepad"]