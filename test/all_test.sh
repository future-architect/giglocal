#!/bin/bash

echo "[INFO] docker-compose down"
docker-compose down

echo "[INFO] move .data -> .tmpdata"
if [ -e ../src/datastore/.data ]; then
    mv ../src/datastore/.data ../src/datastore/.tmpdata
fi

if [ -e  ../src/pubsub/.data ]; then
    mv ../src/pubsub/.data ../src/pubsub/.tmpdata
fi

echo "[INFO] docker-compose up"
export SERVICES=""
docker-compose up -d

echo "[INFO] run testscript"
go test -v | tee runtest.log

echo "[INFO] remove .data"
if [ -e ../src/datastore/.data ]; then
    rm -R -f ../src/datastore/.data
fi

if [ -e  ../src/pubsub/.data ]; then
    rm -R -f ../src/pubsub/.data
fi

echo "[INFO] move .tmpdata -> .data"
if [ -e ../src/datastore/.tmpdata ]; then
    mv ../src/datastore/.tmpdata ../src/datastore/.data
fi

if [ -e  ../src/pubsub/.tmpdata ]; then
    mv ../src/pubsub/.tmpdata ../src/pubsub/.data
fi

echo "[INFO] test succeeded. check runtest.log"
