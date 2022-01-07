FROM node:16.13.1-alpine3.15 as build

# Install project dependencies
WORKDIR /build
COPY package.json .
RUN yarn install

# Copy project files and create production build
COPY . .
RUN yarn build

# Build final image for deployment
FROM nginx:mainline-alpine

# Author labels
LABEL maintainer="Ben Sykes"
LABEL authors="Ben Sykes"

# Copy build result
COPY --from=build /build/dist/ /usr/share/nginx/html