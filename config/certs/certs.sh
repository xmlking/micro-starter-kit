#!/bin/bash

pushd config/certs

set -e

# $1=<CA name>
generate_ca() {
  openssl genrsa -out $1-key.pem 2048
  openssl req -new -key $1-key.pem -out $1-cert.csr -config $1-cert.cfg -batch -sha256
  openssl x509 -req -days 730 -in $1-cert.csr -signkey $1-key.pem -out $1-cert.pem \
    -extensions v3_ca -extfile $1-cert.cfg
}

# $1=<certificate name>
generate_rsa_key() {
  openssl genrsa -out $1-key.pem 2048
}

# $1=<certificate name>
generate_ecdsa_key() {
  openssl ecparam -name secp256r1 -genkey -out $1-key.pem
}

# $1=<certificate name> $2=<CA name>
generate_x509_cert() {
  openssl req -new -key $1-key.pem -out $1-cert.csr -config $1-cert.cfg -batch -sha256
  openssl x509 -req -days 730 -in $1-cert.csr -sha256 -CA $2-cert.pem -CAkey \
    $2-key.pem -CAcreateserial -out $1-cert.pem -extensions v3_ca -extfile $1-cert.cfg
  # echo -e "// NOLINT(namespace-envoy)\nconstexpr char TEST_$(echo $1 | tr a-z A-Z)_CERT_HASH[] = \"$(openssl x509 -in $1-cert.pem -noout -fingerprint -sha256 | cut -d"=" -f2)\";" > $1-cert_hash.h
}

# Generate cert for the CA.
generate_ca ca

# Generate cert for the proxy.
generate_rsa_key proxy
generate_x509_cert proxy ca
# Generate ECDSA cert for the server.
cp -f proxy-cert.cfg proxy-ecdsa-cert.cfg
generate_ecdsa_key proxy-ecdsa
generate_x509_cert proxy-ecdsa ca
rm -f proxy-ecdsa-cert.cfg

# Generate cert for the client.
generate_rsa_key client
generate_x509_cert client ca
# Generate ECDSA cert for the client.
cp -f client-cert.cfg client-ecdsa-cert.cfg
generate_ecdsa_key client-ecdsa
generate_x509_cert client-ecdsa ca
rm -f client-ecdsa-cert.cfg

# Generate cert for the upstream CA.
generate_ca upstream-ca
# Generate cert for the upstream node.
generate_rsa_key upstream
generate_x509_cert upstream upstream-ca
generate_rsa_key upstream-localhost
generate_x509_cert upstream-localhost upstream-ca

rm *.csr
rm *.srl

popd

