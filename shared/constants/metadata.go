package constants

// metadata constants

const (
	// There are certain requirements for metadata to be passed in the http header:
	// gRPC recommended Key format: `lowercase alphanumeric characters and hyphen`
	// but go-micro use `camelcase alphanumeric characters and hyphen`
    TraceIDKey = "Micro-Trace-Id"
    TenantIdKey = "Tenant-Id"
)
