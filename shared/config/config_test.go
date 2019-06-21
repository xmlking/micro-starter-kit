package config

import (
	"testing"

	"github.com/micro/go-micro/config"
)

func TestConfig(t *testing.T) {
	dialect := config.Get("database", "dialect").String("postgres")
	if dialect != "postgres" {
		t.Fatalf("Expected %s got %s", "postgres", dialect)
	}
}
