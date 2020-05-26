package config

type FeatureFlags interface {
	IsTLSEnabled() bool
	IsMetricsEnabled() bool
	IsTracingEnabled() bool
	IsValidatorEnabled() bool
	IsReqlogsEnabled() bool
	IsTranslogsEnabled() bool
}

type featureFlags struct {
	features map[string]Feature
}

func (f *featureFlags) IsTLSEnabled() bool {
	return f.features["mtls"].Enabled
}
func (f *featureFlags) IsMetricsEnabled() bool {
	return f.features["metrics"].Enabled
}
func (f *featureFlags) IsTracingEnabled() bool {
	return f.features["tracing"].Enabled
}
func (f *featureFlags) IsValidatorEnabled() bool {
	return f.features["validator"].Enabled
}
func (f *featureFlags) IsReqlogsEnabled() bool {
	return f.features["reqlogs"].Enabled
}
func (f *featureFlags) IsTranslogsEnabled() bool {
	return f.features["translogs"].Enabled
}
