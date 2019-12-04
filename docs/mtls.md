# Certs

## CA Certs

```bash
# Generate CA key
# pass-phrase: sumodemo
openssl genrsa -des3 -out ca.key 2048
# Generate CA crt from key
openssl req -new -x509 -key ca.key -out ca.crt
```

## Micro Certs

```bash
# Create microservice key
openssl genrsa -out micro.key 2048
# Generate CSR
openssl  req -new -key micro.key -out micro.csr
# Sign CSR with CA
openssl x509 -req -in micro.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out micro.crt
```

## CSR Prompts

```
Country Name (2 letter code) []:US
State or Province Name (full name) []:CA
Locality Name (eg, city) []:Riverside
Organization Name (eg, company) []:Sumo
Organizational Unit Name (eg, section) []:Demo
Common Name (eg, fully qualified host name) []:*
Email Address []:demo@sumo.com

A challenge password []:
```
