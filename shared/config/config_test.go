package config_test

import (
	"testing"

	"github.com/xmlking/micro-starter-kit/shared/config"
)

func TestConfig(t *testing.T) {
	config.DefaultConfig = config.NewConfig(config.WithConfigDir("/config"), config.WithConfigFile("config.prod.yaml"))
	config.IsProduction()
	dialect := config.GetServiceConfig().Database.Dialect
	if dialect != "postgres" {
		t.Fatalf("Expected %s got %s", "postgres", dialect)
	}
}
