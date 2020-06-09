package config_test

import (
    "testing"
    "time"

    "github.com/xmlking/micro-starter-kit/shared/config"
)

// CONFIGOR_DEBUG_MODE=true go test -v ./shared/config/... -count=1

func TestNestedConfig(t *testing.T) {
    t.Logf("Environment: %s", config.Configor.GetEnvironment())
    t.Log(config.GetServiceConfig().Database)
    connMaxLifetime := config.GetServiceConfig().Database.ConnMaxLifetime
    if *connMaxLifetime != time.Duration(time.Hour*2) {
        t.Fatalf("Expected %s got %s", "2h0m0s", connMaxLifetime)
    }
}

func TestDefaultValues(t *testing.T) {
    t.Logf("Environment: %s", config.Configor.GetEnvironment())
    t.Log(config.GetServiceConfig().Database)
    connMaxLifetime := config.GetServiceConfig().Database.ConnMaxLifetime
    if *connMaxLifetime != time.Duration(time.Hour*2) {
        t.Fatalf("Expected %s got %s", "2h0m0s", connMaxLifetime)
    }
}
