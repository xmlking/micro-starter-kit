# Certs

## CA Certs

```bash
mkdir -p config/certs
# Generate CA key
# Use pass-phrase: sumodemo
openssl genrsa -des3 -out config/certs/ca.key 4096
# Generate Root CA crt from key
openssl req -x509 -new -days 1825 -key config/certs/ca.key -out config/certs/ca.crt
```

> Prompts

```
Country Name (2 letter code) []:US
State or Province Name (full name) []:CA
Locality Name (eg, city) []:Riverside
Organization Name (eg, company) []:Sumo
Organizational Unit Name (eg, section) []:Demo
Common Name (eg, fully qualified host name) []:*
Email Address []:ca@sumo.com
```

## Micro Certs

```bash
# Create microservice key
openssl genrsa -out config/certs/micro.key 2048
# Generate CSR
openssl  req -new -key config/certs/micro.key -out config/certs/micro.csr
```

> Prompts, Provide `empty` challenge password

```
Country Name (2 letter code) []:US
State or Province Name (full name) []:CA
Locality Name (eg, city) []:Riverside
Organization Name (eg, company) []:Sumo
Organizational Unit Name (eg, section) []:Demo
Common Name (eg, fully qualified host name) []:*
Email Address []:micro@sumo.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
```

### Sign CSR with CA

```bash
openssl x509 -req -days 365 -in config/certs/micro.csr -CA config/certs/ca.crt -CAkey config/certs/ca.key -CAcreateserial -out config/certs/micro.crt
```
