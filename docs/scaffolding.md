# Scaffolding

> setup new golang project workspace from scratch using `go-micro` cli

```bash
mkdir -p /Developer/Work/go/micro-starter-kit
cd /Developer/Work/go/micro-starter-kit
go mod init github.com/xmlking/micro-starter-kit
mkdir srv api fnc

# scaffold modules
micro new --fqdn="account-srv" --type="srv" --gopath=false \
--alias="account" --plugin=registry=kubernetes srv/account

micro new --fqdn="emailer-srv" --type="srv" --gopath=false \
--alias="emailer"  --plugin=registry=kubernetes:broker=nats srv/emailer

micro new --fqdn="greeter-srv" --type="srv" --gopath=false \
--alias="greeter"  --plugin=registry=kubernetes srv/greeter

micro new --fqdn="account-api" --type="api" --gopath=false \
--alias="account" --plugin=registry=kubernetes api/account
```
