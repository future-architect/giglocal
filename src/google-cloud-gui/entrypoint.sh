#!/usr/bin/env bash

NAME="google-cloud-gui"
if [[ ${SERVICES} == "" ]] || ( [[ ${SERVICES} =~  "datastore" ]] && [[ ${SERVICES} =~  $NAME ]] );then
  google-cloud-gui --port ${GOOGLE_CLOUD_GUI_PORT} --skip-browser
else
  supervisorctl stop $NAME
fi 