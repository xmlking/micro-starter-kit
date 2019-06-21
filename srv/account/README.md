# Account Service

This is the Account service

showcase

1. Implements basic CRUD API
2. Multiple handlers, repositories, subscribers  
3. Publishing events
4. GORM data access
5. Config Managment
6. Custom Logging

Generated with

```
micro new --namespace=go.micro --type=srv --gopath=false --alias=account srv/account
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.account
- Type: srv
- Alias: account

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./account-srv
```

Build a docker image
```
make docker
```

## Test

> start web tool to test:  `micro web --namespace=go.micro.srv`

### User Test (Postman)

```
$http://localhost:8080/rpc
```

#### Create

```json
{
    "service": "go.micro.srv.account",
    "method": "DcMgr.Create",
    "request": {
        "username": "sumo",
        "firstName": "sumo",
        "lastName": "demo",
        "email": "sumo@demo.com"
    }
}
```

### User Test (Curl)

#### Create

```bash
curl -d 'service=go.micro.srv.account' \
     -d 'method=DcMgr.Create' \
     -d 'request={"username": "sumo", "firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}' \
     http://localhost:8080/rpc
     ```
```

### User Test (Micro Web)

#### Create

```json
{
"username": "sumo",
"firstName": "sumoto",
"lastName": "demo",
"email": "sumo@demo.com"
}
```

#### List

```json
{
"limit": 10,
"page": 1
}
```

```json
{
"limit": 10,
"page": 1,
"sort" : "username",
"lastName": "demo"
}
```

```json
{
"limit": 10,
"username": "sumo",
"firstName": "sumo",
"lastName": "demo",
"email": "sumo@demo.com"
}
```
#### Get

```json
{
    "id": 1
}
```

```json
{
"id": 1,
"username": "sumo",
"firstName": "sumo",
"lastName": "demo",
"email": "sumo@demo.com"
}
```

#### Exist

>with any combination

```json
{
"username": "sumo",
"firstName": "sumoto",
"lastName": "demo",
"email": "sumo@demo.com"
}
```

#### Update

```json
{
"id": 1,
"username": "sumo222",
"firstName": "sumoto222",
"lastName": "demo222",
"email": "sumo222@demo.com"
}
```

#### Delete

```json
{
    "id": 1
}
```

### Profile Test

#### Create

```json
{
"user_id": 2,
"tx" : "PST",
"avatar": "sumo2.jpg",
"gender": "F"
}
```

#### List

```json
{
"limit": 10,
"page": 1,
"sort" : "gender",
"user_id": 1,
"gender": "M"
}
```
```json
{
"limit": 10,
"page": 1
}
```
```json
{
"limit": 10,
"page": 1,
"gender": "F"
}
```

#### Get

```json
{
"id": 2
}
```
