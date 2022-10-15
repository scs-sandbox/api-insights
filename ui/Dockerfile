# Node layer
FROM node:18.10.0 as builder
# set working directory
WORKDIR /app

ENV REACT_APP_API_ENDPOINT_URL $REACT_APP_API_ENDPOINT_URL
ENV BACKEND_API_ENDPOINT_URL $BACKEND_API_ENDPOINT_URL

# install package dependencies
COPY package.json ./
COPY docker-package-lock.json ./package-lock.json
RUN npm install --ignore-scripts

# load app and build
COPY . ./
RUN npm run build

# Nginx layer 
#For Apple M1 machines use `FROM arm64v8/nginx`
#FROM arm64v8/nginx
FROM nginx:1.23.1-alpine

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./start.sh /start.sh

RUN rm -rf /usr/share/nginx/html/*
COPY --from=builder /app/dist /usr/share/nginx/html
COPY --from=builder /app/package-lock.json ./
ENV PORT 3000
EXPOSE 3000

# start app
CMD ["sh", "-c", "sh start.sh ${BACKEND_API_ENDPOINT_URL}"]

