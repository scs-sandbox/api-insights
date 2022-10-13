# pull official base image
FROM node:16.17.1-alpine
# set working directory
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH
ENV REACT_APP_API_ENDPOINT_URL $REACT_APP_API_ENDPOINT_URL

# add app
COPY . ./
RUN rm -rf ./package-lock.json

RUN sed -i 's/http:\/\/0.0.0.0:8081/${REACT_APP_API_ENDPOINT_URL}/g' package.json
RUN npm install --ignore-scripts \
    && npm rebuild node-sass

ENV PORT 3000
EXPOSE 3000
# start app
CMD [ "npm", "run", "dockerstart" ]