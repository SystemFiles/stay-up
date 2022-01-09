FROM golang:1.17.5-alpine3.15 as build

# Install project dependencies
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy project files and create production build
COPY api ./api
COPY main.go .
RUN go build -o /build/docker-stayup

# Final image build
FROM bash:alpine3.14
LABEL maintainer="Ben Sykes"
LABEL authors="Ben Sykes"

# Configure environment (Note: Tini allows us to avoid several Docker edge cases, see https://github.com/krallin/tini.)
RUN apk add --no-cache tini bind-tools
RUN addgroup -g 10001 -S stayup && adduser -u 10000 -S -G stayup -h /home/stayup stayup

# Exposes API port
EXPOSE 5555

# Copy build artifacts
COPY --from=build /build/docker-stayup /docker-stayup

# Use nonroot user stayup
USER stayup
ENTRYPOINT ["/sbin/tini", "--", "/docker-stayup"]