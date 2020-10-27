#!/bin/bash

echo 'Running vemproifood installation script!'

echo 'Running npm install'
npm install

echo 'Docker setup'
docker-compose up

echo 'Project running on: https://localhost:3333'
