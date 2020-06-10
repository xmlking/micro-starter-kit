# Transaction Logger

Transaction Logger wrapper publish req/res to [recorder](../../../service/recorder/README.md) service

Transaction Logger wrapper works along with `recorder` Service(broker) that listen for published `TransactionEvents`,<br/>
and store then into configured `go-micro` store for future analysis.
