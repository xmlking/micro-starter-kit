# Recorder Service

Recorder service saves `TransactionEvents` to configured go-micro `store`.
Transactions are published by other micro services i.e., `account`, `emailer`, `greeter`

## Start

> (optional) set broker to googlepubsub

```bash
export MICRO_BROKER=googlepubsub
```

```bash
make run-recorder
```

## Test

```bash
# micro publish recordersrv '{ "Header" : { "a": "aa", "b": "bb" }, "Body" : {"c": "cc"} }'
```
