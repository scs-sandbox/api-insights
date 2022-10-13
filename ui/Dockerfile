# pull official base image
FROM node:18.10.0-alpine
# set working directory
WORKDIR /app

ENV REACT_APP_API_ENDPOINT_URL $REACT_APP_API_ENDPOINT_URL
ENV BACKEND_API_ENDPOINT_URL $BACKEND_API_ENDPOINT_URL
# add app
COPY . ./

RUN apk add --no-cache --virtual .gyp \
    py3-pip \
    make \
    g++ 

RUN npm install --ignore-scripts
RUN npm rebuild node-sass

ENV PORT 3000
EXPOSE 3000

# start app
CMD ["sh", "-c", "sh start.sh ${BACKEND_API_ENDPOINT_URL}"]
