FROM node:8.9.1-alpine

MAINTAINER BlackBox Vision

RUN rm -rf /usr/app/node_modules

ADD ./src /usr/app/
ADD ./logs /usr/app/
ADD ./data /usr/app/

ADD ./package.json /usr/app/
ADD ./package-lock.json /usr/app/
ADD ./tsconfig.json /usr/app/

WORKDIR /usr/app/

EXPOSE 8000

RUN npm prune
RUN npm install