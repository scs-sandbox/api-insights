# pull official base image
FROM node:18.10.0-alpine as builder
# set working directory
WORKDIR /app

ENV REACT_APP_API_ENDPOINT_URL $REACT_APP_API_ENDPOINT_URL
ENV BACKEND_API_ENDPOINT_URL $BACKEND_API_ENDPOINT_URL
# install package dependencies
COPY package.json ./
RUN apk add --no-cache --virtual .gyp \
    py3-pip \
    make \
    g++ 

RUN npm install --ignore-scripts
RUN npm rebuild node-sass

# load app and build
COPY . ./
RUN npm run build
RUN ls -alh

FROM arm64v8/nginx

COPY ./nginx.conf /etc/nginx/nginx.conf
RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /app/dist /usr/share/nginx/html

ENV PORT 3000
EXPOSE 3000

# start app
ENTRYPOINT ["nginx", "-g", "daemon off;"]
#CMD ["sh", "-c", "sh start.sh ${BACKEND_API_ENDPOINT_URL}"]
