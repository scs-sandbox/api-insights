# pull official base image
FROM node:16.17.1-alpine
# set working directory
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

# add app
COPY . ./
RUN sed -i 's/0.0.0.0:8081/backend:8081/g' package.json
RUN npm install --ignore-scripts \
    && npm rebuild node-sass

ENV REACT_APP_API_ENDPOINT_URL $REACT_APP_API_ENDPOINT_URL
ENV PORT 3000
EXPOSE 3000
# start app
CMD [ "npm", "run", "dockerstart" ]