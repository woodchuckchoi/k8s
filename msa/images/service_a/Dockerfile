FROM node:12

WORKDIR /usr/src/app

COPY package.json ./

RUN npm install

COPY server.js ./

EXPOSE 7777

CMD ["node", "server.js"]