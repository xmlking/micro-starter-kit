package main

import (
	_ "github.com/micro/go-plugins/client/selector/static"
    _ "github.com/micro/go-plugins/registry/kubernetes"
    // Flag usage of plugins `micro --client=grpc --server=grpc`
    // _ "github.com/micro/go-plugins/client/grpc"
    // _ "github.com/micro/go-plugins/server/grpc"
)

func init() {
	// set values for registry/selector
	// os.Setenv("MICRO_REGISTRY", "kubernetes")
	// os.Setenv("MICRO_SELECTOR", "static")
}
