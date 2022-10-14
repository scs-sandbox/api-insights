#!/bin/bash
if [ "$#" -eq  "0" ]
then
    echo "No arguments supplied"
    nginx -g 'daemon off;'
else
    echo "API Endpoint arg: $1"
    sed -i "s#http://backend:8081#$1#g" /etc/nginx/nginx.conf
    nginx -g 'daemon off;'
fi
