FROM google/cloud-sdk:289.0.0-alpine

RUN apk update \
    && apk add --update --no-cache openjdk8-jre \
    && apk add --no-cache supervisor openssh nginx nodejs yarn \
    && gcloud components update beta --quiet \
    && gcloud components install cloud-datastore-emulator cloud-firestore-emulator pubsub-emulator --quiet \ 
    && yarn global add google-cloud-gui && apk del yarn

ADD ./resource/supervisord.conf /etc/
ADD ./src /gcplocal/src/

RUN chmod u+x /gcplocal/src/datastore/entrypoint.sh /gcplocal/src/pubsub/entrypoint.sh /gcplocal/src/firestore/entrypoint.sh /gcplocal/src/google-cloud-gui/entrypoint.sh

# create log directory
RUN mkdir /gcplocal/log

WORKDIR /gcplocal/

# run supervisord
ENTRYPOINT ["/usr/bin/supervisord","-c","/etc/supervisord.conf"]
