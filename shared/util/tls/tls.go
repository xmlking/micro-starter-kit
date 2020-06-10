package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"

	"github.com/markbates/pkger"
	"github.com/markbates/pkger/pkging"
	"github.com/micro/go-micro/v2"
	gc "github.com/micro/go-micro/v2/client/grpc"
	gs "github.com/micro/go-micro/v2/server/grpc"
	maddr "github.com/micro/go-micro/v2/util/addr"
	mls "github.com/micro/go-micro/v2/util/tls"
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

func WithTLS(t *tls.Config) micro.Option {
	return func(o *micro.Options) {
		o.Client.Init(
			gc.AuthTLS(t),
		)
		o.Server.Init(
			gs.AuthTLS(t),
		)
	}
}
