# pull official base image
FROM node:14.18.1-alpine
# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH
ENV endpoint=https://devnet-testing.cisco.com
# install app dependencies
COPY .env ./
COPY ./package.json ./
#COPY ./server.sh ./

# add app
COPY . ./
#RUN chmod +x ./server.sh

RUN apk add --no-cache --virtual .gyp \
    py3-pip \
    make \
    g++ \
    && npm install \
    && apk del .gyp
# RUN npm run build:docker


ENV REACT_APP_API_ENDPOINT_URL $REACT_APP_API_ENDPOINT_URL
ENV PORT 3000
EXPOSE 3000
# start app
CMD [ "npm", "start" ]
#CMD ["sh", "-c", "sh server.sh ${endpoint}"]