// +build integration

package service

import (
	"testing"
)

func TestCloudEmailVersionIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
}
