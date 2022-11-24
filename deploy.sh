#! /bin/sh

docker build -t ge0r/dice-pouch .
docker run -d -p 80:9001 ge0r/dice-pouch
