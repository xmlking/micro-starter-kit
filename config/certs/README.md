# certificates

identities, certificates and keys

There are 3 identities:

- **CA**: Certificate Authority for **Client** and **Proxy**. It has the
  self-signed certificate *ca-cert.pem*. *ca-key.pem* is its private key.
- **Client**: It has the certificate *client-cert.pem*, which is signed by the
  **CA** using the config *client-cert.cfg*. *client-key.pem* is its private key.
- **Proxy**: It has the certificate *proxy-cert.pem*, which is signed by the
  **CA** using the config *proxy-cert.cfg*. *proxy-key.pem* is its private key.
- **Upstream CA**: Certificate Authority for **Upstream**. It has the self-signed
  certificate *upstream-ca-cert.pem*. *upstream-ca-key.pem* is its private key.
- **Upstream**: It has the certificate *upstream-cert.pem*, which is signed by
  the **Upstream CA** using the config *upstream-cert.cfg*. *upstream-key.pem* is
  its private key.
- **Upstream localhost**: It has the certificate *upstream-localhost-cert.pem*, which is signed by
  the **Upstream CA** using the config *upstream-localhost-cert.cfg*. *upstream-localhost-key.pem* is
  its private key. The different between this certificate and **Upstream** is that this certificate
  has a SAN for "localhost".

## How to update certificates

**certs.sh** has the commands to generate all files. Running `certs.sh` directly
will cause all files to be regenerated. So if you want to regenerate a
particular file, please copy the corresponding commands from `certs.sh` and
execute them in command line.

```bash
# at project root, run:
./config/certs/certs.sh
```

## Reference

- [RSA vs DSA vs ECDSA](https://www.misterpki.com/rsa-dsa-ecdsa/)
