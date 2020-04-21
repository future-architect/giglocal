#!/usr/bin/env bash

NAME="pubsub"
if [[ ${SERVICES} == "" ]] || [[ ${SERVICES} =~  ${NAME} ]];then
  gcloud beta emulators pubsub start \
    --data-dir=/gcplocal/pubsub/.data \
    --host-port=${PUBSUB_LISTEN_ADDRESS} \
    --project=${PUBSUB_PROJECT_ID}
else
  supervisorctl stop $NAME
fi 