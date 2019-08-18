package main

import (
	// Flags usage of k8s plugins `micro --registry=kubernetes --selector=static`
	_ "github.com/micro/go-plugins/client/selector/static"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	// Flags usage of cors plugin `micro --cors-allowed-headers=X-Custom-Header --cors-allowed-origins=someotherdomain.com  --cors-allowed-methods=POST`
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"
	// Flags usage of grpc plugin `micro --client=grpc --server=grpc`
	// _ "github.com/micro/go-plugins/client/grpc"
	// _ "github.com/micro/go-plugins/server/grpc"
)

func init() {
	// lets not hard-code kubernetes, so that we can use same `micro` binary with consul
	// set values for registry/selector
	// os.Setenv("MICRO_REGISTRY", "kubernetes")
	// os.Setenv("MICRO_SELECTOR", "static")

	// setup cors plugin
	plugin.Register(cors.NewPlugin())
}
