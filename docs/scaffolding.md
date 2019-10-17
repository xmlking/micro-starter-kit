# Scaffolding

> setup new golang project workspace from scratch using `go-micro` cli

```bash
mkdir -p /Developer/Work/go/micro-starter-kit
cd /Developer/Work/go/micro-starter-kit
go mod init github.com/xmlking/micro-starter-kit
mkdir srv api fnc

# scaffold modules
micro new --fqdn="accountsrv" --type="srv" --gopath=false \
--alias="account" --plugin=client/selector=static srv/account1

micro new --fqdn="emailersrv" --type="srv" --gopath=false \
--alias="emailer"  --plugin=client/selector=static:broker=nats srv/emailer

micro new --fqdn="greetersrv" --type="srv" --gopath=false \
--alias="greeter"  ---plugin=client/selector=static srv/greeter
```
