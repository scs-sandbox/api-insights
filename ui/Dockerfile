# pull official base image
FROM node:14.18.1-alpine
# set working directory
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

# add app
COPY . ./

RUN apk add --no-cache --virtual .gyp \
    py3-pip \
    make \
    g++ \
    && npm install \
    && apk del .gyp

ENV REACT_APP_API_ENDPOINT_URL $REACT_APP_API_ENDPOINT_URL
ENV PORT 3000
EXPOSE 3000
# start app
CMD [ "npm", "run", "dockerstart" ]