# TODO

## Capabilities

- [x] Monorepo - Sharing Code Between Microservices
- [x] gRPC microservices with REST Gateway
- [x] Input Validation with [protoc-gen-validate (PGV)](https://github.com/envoyproxy/protoc-gen-validate)
- [x] Config - Pluggable Sources, Mergeable Config, Safe Recovery
- [x] Customizable Logging
- [x] CRUD Example using [GORM](https://gorm.io/), [benchmarks](https://github.com/kihamo/orm-benchmark), [XORM](https://xorm.io/) next?
- [x] GORM code gen via [protoc-gen-gorm](https://github.com/infobloxopen/protoc-gen-gorm) or use [protoc-go-inject-tag](https://github.com/favadi/protoc-go-inject-tag)?
- [x] Dependency injection [Container](https://github.com/sarulabs/di), Try [wire](https://itnext.io/mastering-wire-f1226717bbac) next?
- [x] multi-stage-multi-target Dockerfile
- [x] One Step _build/publish/deploy_ with [ko](https://github.com/google/ko)
- [x] BuildInfo with [govvv](https://github.com/ahmetb/govvv)
- [x] Linting with [GolangCI](https://github.com/golangci/golangci-lint) linters aggregator
- [x] Linting Protos with [Buf](https://buf.build/docs/introduction)
- [x] CICD Pipelines with [GitHub Actions](https://github.com/features/actions)
- [x] Kubernetes _Matrix Deployment_ with [Kustomize](https://kustomize.io/)
- [ ] Add k8s [healthchecks](https://github.com/heptiolabs/healthcheck) with [cmux](https://medium.com/@drgarcia1986/listen-grpc-and-http-requests-on-the-same-port-263c40cb45ff)
- [x] Feature Flags (enable/disable with zero cost)
- [ ] Observability
- [ ] Service Mesh with [Istio](https://istio.io/)
- [ ] GraphQL Gateway with [gqlgen](https://gqlgen.com/), [rejoiner](https://github.com/google/rejoiner)
- [ ] Graph-Based ORM [ent](https://entgo.io/)
- [ ] Switch to [Bazel Build](https://bazel.build/)
