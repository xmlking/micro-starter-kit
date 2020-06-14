# Recorder Service

Recorder service saves `TransactionEvents` to configured go-micro `store`.
Transactions are published by other micro services i.e., `account`, `emailer`, `greeter`

## Usage

Run the service

> (optional) set broker to googlepubsub

```bash
export MICRO_BROKER=googlepubsub
export GOOGLEPUBSUB_PROJECT_ID=<my-project-id>
export GOOGLE_APPLICATION_CREDENTIALS=<path_to.json>
```

```bash
make run-recorder
make run-recorder ARGS="--server_address=:8083"
```

Test the service

```bash
micro call mkit.service.recorder Transaction.Read  '{"key": "0edc8bb2-33e0-4766-bc13-e62a346465be#accountsrv"}'
# micro publish mkit.service.recorder '{ "Header" : { "a": "aa", "b": "bb" }, "Body" : {"c": "cc"} }'
```
