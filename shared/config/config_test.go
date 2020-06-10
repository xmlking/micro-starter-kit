package config_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xmlking/micro-starter-kit/shared/config"
)

// CONFIGOR_DEBUG_MODE=true go test -v ./shared/config/... -count=1

func TestNestedConfig(t *testing.T) {
	t.Logf("Environment: %s", config.Configor.GetEnvironment())
	t.Log(config.GetConfig().Database)
	connMaxLifetime := config.GetConfig().Database.ConnMaxLifetime
	if *connMaxLifetime != time.Duration(time.Hour*2) {
		t.Fatalf("Expected %s got %s", "2h0m0s", connMaxLifetime)
	}
}

func TestDefaultValues(t *testing.T) {
	t.Logf("Environment: %s", config.Configor.GetEnvironment())
	t.Log(config.GetConfig().Database)
	connMaxLifetime := config.GetConfig().Database.ConnMaxLifetime
	if *connMaxLifetime != time.Duration(time.Hour*2) {
		t.Fatalf("Expected %s got %s", "2h0m0s", connMaxLifetime)
	}
}

func ExampleGetConfig() {
	fmt.Println(config.GetConfig().Email)
	// fmt.Println(config.GetConfig().Services["account"].Deadline)

	// Output:
	// username:"yourGmailUsername" password:"yourGmailAppPassword" email_server:"smtp.gmail.com" port:587 from:"xmlking-test@gmail.com"
}

func ExampleGetConfig_check_defaults() {
	fmt.Println(config.GetConfig().Services.Account.Endpoint)
	fmt.Println(config.GetConfig().Services.Account.Version)
	fmt.Println(config.GetConfig().Services.Account.Deadline)

	// Output:
	// mkit.service.account:8080
	// v0.1.0
	// 8888
}
