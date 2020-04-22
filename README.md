# giglocal - A partly functional local GCP

*giglocal* provides an easy-to-use and integrated GCP testing environment.

* [日本語](README_JP.md)

## Overview

Several emulators are provided for GCP applications(by [officially](https://cloud.google.com/sdk/gcloud/reference/beta/emulators) or 3rd party). But they are independent.

Therefore, giglocal provides these environments at once.

We currently support the following emulators of the application.

* **Datastore**
* **Firestore**
* **Pub/Sub**

... In the future, more...

## Requirements
* `Docker`
* [Cloud Client Libraries](https://cloud.google.com/apis/docs/client-libraries-explained)(for language you want to use)


## QuickStart
You can use the `docker-compose.yml` file from the repository and use this command:

```bash
docker-compose up
```

After the container starts, you can use emulators through the ports just below.

|emulator  |default port  |service name  |SERVICE|
|----------|------|------|------|
|Datastore | http://localhost:5051 | datastore | DATASTORE |
|Firestore | http://localhost:5052 | firestore | FIRESTORE |
|Pub/Sub   | http://localhost:5053 | pubsub | PUBSUB |


And set environment variables `<SERVICE>_EMULATOR_HOST`, You can connect emulators.

```bash
export DATASTORE_EMULATOR_HOST=localhost:5051
export FIRESTORE_EMULATOR_HOST=localhost:5052
export PUBSUB_EMULATOR_HOST=localhost:5053
```

more information
- datastore : https://cloud.google.com/sdk/gcloud/reference/beta/emulators/datastore
- firestore : https://cloud.google.com/sdk/gcloud/reference/beta/emulators/firestore
- pubsub : https://cloud.google.com/sdk/gcloud/reference/beta/emulators/pubsub 


## Configurations

You can pass the following environment variables to giglocal

* `SERVICES`: Comma-separated list of service names you want to invoke. If you do not set this variable, all emulators will be started. see above for service name.\
  > Example value: `datastore,pubsub` to start datastore and pubsub.
* `<SERVICE>_PORT`: Port number to bind a specific service (defaults to service ports above).
* `<SERVICE>_DIR`:  The host directory to be mounted on the emulator data directory (default ./src/-ServiceName-/.data). This setting is only available for Datastore and Pub/Sub.


## License
This version of giglocal is released under the Apache License, Version 2.0 (see [LICENSE](https://github.com/future-architect/giglocal/blob/master/LICENSE)).
