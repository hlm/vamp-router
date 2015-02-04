#!/bin/bash
set -x

# build the app for linux/i386 an build the docker container
GOOS=linux GOARCH=386 go build
mv vamp-loadbalancer target/linux_i386/
docker build -t magneticio/vamp-loadbalancer:$1 .