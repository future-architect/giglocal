#!/usr/bin/env bash

NAME="datastore"
if [[ ${SERVICES} == "" ]] || [[ ${SERVICES} =~  ${NAME} ]];then
  gcloud beta emulators datastore start \
    --data-dir=/gcplocal/datastore/.data \
    --host-port=${DATASTORE_LISTEN_ADDRESS} \
    --project=${DATASTORE_PROJECT_ID}
else
  supervisorctl stop $NAME
fi 