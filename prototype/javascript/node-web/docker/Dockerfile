FROM node:16-alpine

ENV TZ="Asia/Shanghai"

WORKDIR /app
COPY server.js package.json package-lock.json ./

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache update && apk --no-cache upgrade
# RUN npm install --registry=https://registry.npm.taobao.org
RUN npm config set registry https://registry.npm.taobao.org
RUN npm install

EXPOSE 8080
CMD ["node", "server.js"]
