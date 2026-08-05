# Builds the CHACONTAINER web UI as static files served by nginx, which also
# reverse-proxies /api and /webhooks to the api container so the whole app
# is reachable from a single origin (no CORS needed, no cloud CDN involved).
FROM node:22-alpine AS build
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:1.27-alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
