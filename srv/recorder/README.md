# Recorder Service

This is the Recorder service
Ability to save transaction request and response to bigtable.
Transactions are published by other remote services i.e., `keylink`, `inquiry`, `search`, `score`, `resolver`.

## Start

> start the BigTable emulator and PubSub emulator first

### Start PubSub emulator

```bash
# start pubsub emulator
gcloud beta emulators pubsub start --project=df-key-kls-npe-8d1f --host-port=localhost:8085

# create topic (dev env only)
export PUBSUB_EMULATOR_HOST=localhost:8085
go run ./cmd/psutil -p df-key-kls-npe-8d1f  -t recordersrv
```

### Start Bigtable emulator

```bash
# start bigtable emulator
gcloud beta emulators bigtable start --host-port=localhost:8086

# create schema for Recorder
export BIGTABLE_EMULATOR_HOST=localhost:8086
cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe createtable translogtable
cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe createfamily translogtable scoresrv
cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe createfamily translogtable searchsrv
cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe createfamily translogtable resolversrv
cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe createfamily translogtable inquirysrv
cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe createfamily translogtable keylinksrv
cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe read translogtable
```

### Start Recorder service

> follow same steps below for other microservices that need `googlepubsub`

#### set dev env

```bash
export PUBSUB_EMULATOR_HOST=localhost:8085
export BIGTABLE_EMULATOR_HOST=localhost:8086

export GOOGLEPUBSUB_PROJECT_ID=df-key-kls-npe-8d1f
```

#### set prod env

```bash
export GOOGLEPUBSUB_PROJECT_ID=df-key-kls-npe-8d1f
export GOOGLE_APPLICATION_CREDENTIALS=config/cbt-dev-key.json
# create topic for prod if needed
gcloud pubsub topics list
gcloud pubsub topics create recordersrv
# gcloud pubsub subscriptions create --topic recordersrv recordersrv-sub
# gcloud pubsub topics publish recordersrv --message "hello"
# gcloud pubsub subscriptions pull --auto-ack recordersrv-sub
```

#### start recorder microservice with googlepubsub broker

```bash
export MICRO_BROKER=googlepubsub
# make run-recorder
make run-recorder ARGS="--server_address=localhost:55012 --broker_address=localhost:55022"
```

> start recorder microservice with default broker <br/>
> Topic may not be created until the first message is published. it that case, restart all services.

```bash
# start recorder microservice with default broker
go run ./srv/recorder/main.go --server_address=localhost:55011 --broker_address=localhost:55021
```

## Usage

```bash
make run-recorder
make build-recorder
make lint-recorder
make test-recorder
make inte-recorder
make docker-recorder
```

## Test

```bash
#  go test -v ./e2e/recorder_test.go
# micro publish recordersrv '{ "Header" : { "a": "aa", "b": "bb" }, "Body" : {"c": "cc"} }'
# micro publish recordersrv '{"id" : "f2c74532-ca52-4fdc-a037-dedc2697e443", "timestamp": 1571673881, "message": "mymessage"}'
# micro publish recordersrv '{ "to" : "sumo@demo.com", "from": "demo@sumo.com", "subject": "sub", "body": "mybody"  }'

# test in e2e k8s env
# call any service, to see if transactions are logged.
k exec -it bigtable-5c7687fdbf-czzlr  -- sh
# for e2e profile
export BIGTABLE_EMULATOR_HOST=bigtable:8086
# for prod profile
export BIGTABLE_EMULATOR_HOST=prod-bigtable-v1:8086
# for local
export BIGTABLE_EMULATOR_HOST=localhost:8086

cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe ls
cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe count translogtable
cbt -project df-key-kls-npe-8d1f -instance kl-bt-key-dev-npe read translogtable
```

## Reference

- <https://github.com/wyp2013/go-every-day/blob/master/servicemesh/microlearn/example/brokertest/example01/server/main.go>
- <https://github.com/ne0z/go-micro-googlepubsub-demo/blob/fd15354c055db07ac6736a99b9c4ce8315c5ee5f/Subscriber/main.go>
