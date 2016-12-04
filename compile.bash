#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABME=0 \
    go build -a -installsuffix cgo -o bin/gin-app
