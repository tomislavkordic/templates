FROM node:16.15-slim AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . ./
RUN npm run build


FROM nginx:1.22-alpine

COPY docker/nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/out /usr/share/nginx/html

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]
