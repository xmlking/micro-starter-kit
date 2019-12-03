package util

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"

	maddr "github.com/micro/go-micro/util/addr"
	mls "github.com/micro/go-micro/util/tls"
)

func GetSelfSignedTLSConfig(address string) (*tls.Config, error) {
	hosts := []string{address}

	// check if its a valid host:port
	if host, _, err := net.SplitHostPort(address); err == nil {
		if len(host) == 0 {
			hosts = maddr.IPs()
		} else {
			hosts = []string{host}
		}
	}

	// generate a certificate
	cert, err := mls.Certificate(hosts...)
	if err != nil {
		return nil, err
	}

	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}

func GetTLSConfig(certFile string, keyFile string, caFile string, address string) (tlsConfig *tls.Config, err error) {
	var cert tls.Certificate
	cert, err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}
	var caCert []byte
	caCert, err = ioutil.ReadFile(caFile)
	if err != nil {
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   address,
		RootCAs:      caCertPool,
	}
	return
}
