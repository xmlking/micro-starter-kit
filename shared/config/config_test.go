package config_test

import (
	"fmt"
	"testing"
	"time"

    "google.golang.org/grpc/resolver"

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

func TestParseTargetString(t *testing.T) {
    for _, test := range []struct {
        targetStr string
        want      resolver.Target
    }{
        {targetStr: "", want: resolver.Target{Scheme: "", Authority: "", Endpoint: ""}},
        {targetStr: ":///", want: resolver.Target{Scheme: "", Authority: "", Endpoint: ""}},
        {targetStr: "a:///", want: resolver.Target{Scheme: "a", Authority: "", Endpoint: ""}},
        {targetStr: "://a/", want: resolver.Target{Scheme: "", Authority: "a", Endpoint: ""}},
        {targetStr: ":///a", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "a"}},
        {targetStr: "a://b/", want: resolver.Target{Scheme: "a", Authority: "b", Endpoint: ""}},
        {targetStr: "a:///b", want: resolver.Target{Scheme: "a", Authority: "", Endpoint: "b"}},
        {targetStr: "://a/b", want: resolver.Target{Scheme: "", Authority: "a", Endpoint: "b"}},
        {targetStr: "a://b/c", want: resolver.Target{Scheme: "a", Authority: "b", Endpoint: "c"}},
        {targetStr: "dns:///google.com", want: resolver.Target{Scheme: "dns", Authority: "", Endpoint: "google.com"}},
        {targetStr: "dns:///google.com:8080", want: resolver.Target{Scheme: "dns", Authority: "", Endpoint: "google.com:8080"}},
        {targetStr: "dns://a.server.com/google.com", want: resolver.Target{Scheme: "dns", Authority: "a.server.com", Endpoint: "google.com"}},
        {targetStr: "dns://a.server.com/google.com/?a=b", want: resolver.Target{Scheme: "dns", Authority: "a.server.com", Endpoint: "google.com/?a=b"}},

        {targetStr: "/", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "/"}},
        {targetStr: "google.com", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "google.com"}},
        {targetStr: "google.com/?a=b", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "google.com/?a=b"}},
        {targetStr: "/unix/socket/address", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "/unix/socket/address"}},
        {targetStr: "unix:///tmp/mysrv.sock", want: resolver.Target{Scheme: "unix", Authority: "", Endpoint: "tmp/mysrv.sock"}},

        // If we can only parse part of the target.
        {targetStr: "://", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "://"}},
        {targetStr: "unix://domain", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "unix://domain"}},
        {targetStr: "a:b", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "a:b"}},
        {targetStr: "a/b", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "a/b"}},
        {targetStr: "a:/b", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "a:/b"}},
        {targetStr: "a//b", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "a//b"}},
        {targetStr: "a://b", want: resolver.Target{Scheme: "", Authority: "", Endpoint: "a://b"}},
    } {
        got := config.ParseTarget(test.targetStr)
        if got != test.want {
            t.Errorf("ParseTarget(%q) = %+v, want %+v", test.targetStr, got, test.want)
        }
    }
}
