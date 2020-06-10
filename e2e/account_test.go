// e2e, black-box testing
package e2e

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro/go-micro/v2/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	// "github.com/xmlking/micro-starter-kit/shared/micro/client/selector/static"
	profilePB "github.com/xmlking/micro-starter-kit/service/account/proto/profile"
	userPB "github.com/xmlking/micro-starter-kit/service/account/proto/user"
)

/**
* set envelopment variables for CI e2e tests with `memory` registry.
* - export MICRO_REGISTRY=memory
* - export MICRO_SELECTOR=static
* (Or) Set envelopment variables for CI e2e tests via gRPC Proxy
* - MICRO_PROXY_ADDRESS="localhost:8081"
* You can also run this test againest your local running service with mDNS. i.e., `make run-account`
**/

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AccountTestSuite struct {
	suite.Suite
	user    userPB.UserService
	profile profilePB.ProfileService
}

// SetupSuite implements suite.SetupAllSuite
func (suite *AccountTestSuite) SetupSuite() {
	suite.T().Log("in SetupSuite")
	suite.user = userPB.NewUserService("mkit.service.account", client.NewClient())
	suite.profile = profilePB.NewProfileService("mkit.service.account", client.NewClient())
}

// TearDownSuite implements suite.TearDownAllSuite
func (suite *AccountTestSuite) TearDownSuite() {
	suite.T().Log("in TearDownSuite")
}

// before each test
func (suite *AccountTestSuite) SetupTest() {
	t := suite.T()
	t.Log("in SetupTest - creating user")
	_, err := suite.user.Create(context.TODO(), &userPB.CreateRequest{
		Username:  &wrappers.StringValue{Value: "sumo"},
		FirstName: &wrappers.StringValue{Value: "sumo"},
		LastName:  &wrappers.StringValue{Value: "demo"},
		Email:     &wrappers.StringValue{Value: "sumo@demo.com"},
	})
	require.Nil(t, err)
}

// after each test
func (suite *AccountTestSuite) TearDownTest() {
	suite.T().Log("in TearDownTest")
}

// All methods that begin with "Test" are run as tests within a suite.
func (suite *AccountTestSuite) TestUserHandler_Exist_E2E() {
	t := suite.T()
	t.Log("in TestUserHandler_Exist_E2E, checking if user Exist")

	rsp, err := suite.user.Exist(context.TODO(), &userPB.ExistRequest{
		Username: &wrappers.StringValue{Value: "sumo"},
	})
	require.Nil(t, err)
	assert.Equal(suite.T(), rsp.GetResult(), true)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAccountTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}
	suite.Run(t, new(AccountTestSuite))
}
