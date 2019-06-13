package config

import (
	"fmt"
	"testing"

	"github.com/micro/go-micro/config"
)

func TestConfig(t *testing.T) {
	fmt.Println(config.Get("database", "dialect").String("postgres"))
}
