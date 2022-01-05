FROM golang:1.17.5-alpine3.15

LABEL maintainer="Ben Sykes"
LABEL authors="Ben Sykes"

# Configure environment (Note: Tini allows us to avoid several Docker edge cases, see https://github.com/krallin/tini.)
RUN apk add --no-cache tini bind-tools
RUN addgroup -g 10001 -S stayup && adduser -u 10000 -S -G stayup -h /home/stayup stayup

# Install project dependencies
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy project files and create production build
COPY . .
RUN go build -o /docker-stayup

# Exposes API port
EXPOSE 5555

# Use nonroot user stayup
USER stayup
ENTRYPOINT ["/sbin/tini", "--", "/docker-stayup"]