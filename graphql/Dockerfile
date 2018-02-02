FROM node:8.9.1-alpine
WORKDIR /app
COPY ./ .
RUN npm install
RUN ./node_modules/.bin/tsc && rm -rf ./node_modules
RUN npm install --only=production

FROM node:8.9.1-alpine
WORKDIR /app
COPY --from=0 /app .
CMD ["node", "/app/dist/index.js"]

EXPOSE 8000

