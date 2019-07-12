package config

import (
	"testing"

	microConfig "github.com/micro/go-micro/config"
)

func TestConfig(t *testing.T) {
	// For Tests, always use relative paths
	InitConfig("../../config", "config.prod.yaml")
	dialect := microConfig.Get("database", "dialect").String("no dialect")
	if dialect != "postgres" {
		t.Fatalf("Expected %s got %s", "postgres", dialect)
	}
}
