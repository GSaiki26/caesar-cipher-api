# Basics
FROM golang:1.22.4-alpine
WORKDIR /app

# Env
ENV TZ="America/Sao_Paulo"

# Update the container
RUN apk upgrade --no-cache --update
RUN apk add --no-cache tzdata
RUN date

# Configure the user
RUN adduser -D user
RUN chown user /app
USER user

# Install the dependencies
COPY --chown=user ./go.mod .
RUN go mod download

# Copy the project
COPY --chown=user ./cmd ./cmd
COPY --chown=user ./internal ./internal

# Run the project
ENTRYPOINT go run ./cmd/api/main.go
