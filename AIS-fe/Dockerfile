FROM node:18.10-alpine AS build
WORKDIR /app
COPY / ./
COPY package*.json ./

RUN npm install && \
    npm run build
COPY . .

FROM nginx:1.17.1-alpine
WORKDIR /app
COPY --from=build /app/dist/ais-fe /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
