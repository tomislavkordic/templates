FROM node:16.15-slim

ENV NODE_ENV=production

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . ./

EXPOSE 3000

ENTRYPOINT ["node", "index.js"]
