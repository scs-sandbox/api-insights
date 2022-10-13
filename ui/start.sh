#!/bin/bash
if [ "$#" -eq  "0" ]
then
    echo "No arguments supplied"
    npm run dockerstart
else
    echo "API Endpoint arg: $1"
    endpoint=$1
    sed -i "s#http://0.0.0.0:8081#$1#g" package.json
    npm run dockerstart
fi
