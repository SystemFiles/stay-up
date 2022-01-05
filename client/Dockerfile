FROM node:16.13.1-alpine3.15

LABEL maintainer="Ben Sykes"
LABEL authors="Ben Sykes"

# Configure environment (Note: Tini allows us to avoid several Docker edge cases, see https://github.com/krallin/tini.)
RUN apk add --no-cache tini bind-tools
RUN addgroup -g 10001 -S stayup && adduser -u 10000 -S -G stayup -h /home/stayup stayup

# Install project dependencies
WORKDIR /app
COPY package.json .
RUN yarn install

# Add production serving
RUN yarn global add serve --prefix /usr/local

# Copy project files and create production build
COPY . .
RUN yarn build

# Change this port to your configured port in `config/index.ts`
EXPOSE 3000

# Use nonroot user stayup
USER stayup
ENTRYPOINT ["/sbin/tini", "--", "serve", "-s", "dist" ]