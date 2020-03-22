# Scaffolding

> setup new golang project workspace from scratch using `go-micro` cli

```bash
mkdir -p ~/Developer/Work/go/micro-starter-kit
cd ~/Developer/Work/go/micro-starter-kit
go mod init github.com/xmlking/micro-starter-kit
mkdir srv api fnc

# scaffold modules
micro new --fqdn="accountsrv" --type="service" --alias="account" srv/account
micro new --fqdn="greetersrv" --type="service" --alias="greeter" srv/greeter
micro new --fqdn="emailersrv" --type="service" --alias="emailer" srv/emailer
# micro new --fqdn="emailersrv" --type="service"  \
# --alias="emailer"  --plugin=client/selector=static:broker=nats srv/emailer
```

## Setup project

### GitFlow setup

```bash
git flow init -D
```

### CHANGELOG generator setup

```bash
git-chglog --init
```
