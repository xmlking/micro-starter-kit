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
micro call recordersrv Transaction.Read  '{"key": "0edc8bb2-33e0-4766-bc13-e62a346465be#accountsrv"}'
# micro publish recordersrv '{ "Header" : { "a": "aa", "b": "bb" }, "Body" : {"c": "cc"} }'
```
