# Google PubSub

```bash
export GOOGLEPUBSUB_PROJECT_ID=my-project-id
gcloud beta emulators pubsub start --project=$GOOGLEPUBSUB_PROJECT_ID --host-port=localhost:8085
# Create topic `emailersrv` (optional)
# Note: Second time when you run below service, it will automatically create topic
gcloud pubsub topics create emailersrv
```

Run service

```bash
export GOOGLEPUBSUB_PROJECT_ID=my-project-id
# PUBSUB_EMULATOR_HOST for Dev
$(gcloud beta emulators pubsub env-init)
export PUBSUB_EMULATOR_HOST=localhost:8085

# set GOOGLE_APPLICATION_CREDENTIALS for Prod
# export GOOGLE_APPLICATION_CREDENTIALS=~/path/your_project_credentials.json
export MICRO_BROKER=googlepubsub
make run-emailer ARGS="--server_address=localhost:8080"
```
