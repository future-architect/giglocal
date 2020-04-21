#!/usr/bin/env bash

NAME="firestore"
if [[ ${SERVICES} == "" ]] || [[ ${SERVICES} =~  ${NAME} ]];then
  gcloud beta emulators firestore start \
    --host-port=${FIRESTORE_LISTEN_ADDRESS} \
    --project=${FIRESTORE_PROJECT_ID}\
    --quiet
else
  supervisorctl stop $NAME
fi 