package tls

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"

    "github.com/markbates/pkger"
    "github.com/markbates/pkger/pkging"
)

func GetTLSConfig(certFile string, keyFile string, caFile string, address string) (tlsConfig *tls.Config, err error) {
    var cert tls.Certificate
    var certPEMBlock, keyPEMBlock, caPEMBlock []byte
    var certF, keyF, caF pkging.File
    // cert, err = tls.LoadX509KeyPair(certFile, keyFile)
    certF, err = pkger.Open(certFile)
    if err != nil {
        return
    }
    certPEMBlock, err = ioutil.ReadAll(certF)
    if err != nil {
        return
    }
    keyF, err = pkger.Open(keyFile)
    if err != nil {
        return
    }
    keyPEMBlock, err = ioutil.ReadAll(keyF)
    if err != nil {
        return
    }
    cert, err = tls.X509KeyPair(certPEMBlock, keyPEMBlock)
    if err != nil {
        return
    }
    caF, err = pkger.Open(caFile)
    if err != nil {
        return
    }
    caPEMBlock, err = ioutil.ReadAll(caF)
    if err != nil {
        return
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caPEMBlock)

    tlsConfig = &tls.Config{
        Certificates: []tls.Certificate{cert},
        ServerName:   address,
        RootCAs:      caCertPool,
    }
    return
}
