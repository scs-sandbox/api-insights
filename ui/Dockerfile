# pull official base image
FROM node:18.10.0-alpine
# set working directory
WORKDIR /app

ARG REACT_APP_API_ENDPOINT_URL $REACT_APP_API_ENDPOINT_URL

# add app
COPY . ./

RUN sed -i 's/0.0.0.0:8081/backend:8081/g' package.json
RUN npm install --ignore-scripts

ENV PORT 3000
EXPOSE 3000

# start app
#CMD [ "npm", "run", "dockerstart" ]

CMD ["sh", "-c", "sh start.sh ${REACT_APP_API_ENDPOINT_URL}"]
