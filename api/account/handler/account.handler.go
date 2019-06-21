package handler

import (
	"context"
	"encoding/json"
	"strconv"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"

	accountPB "github.com/xmlking/micro-starter-kit/api/account/proto/account"
	userPB "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

// AccountHandler struct
type accountHandler struct {
	userSrvClient userPB.UserService
	profSrvClient userPB.ProfileService
}

// NewAccountHandler returns an instance of `AccountServiceHandler`.
func NewAccountHandler(userSrv userPB.UserService, profSrv userPB.ProfileService) accountPB.AccountServiceHandler {
	return &accountHandler{
		userSrvClient: userSrv,
		profSrvClient: profSrv,
	}
}

// List is a method which will be served by http request /account/list
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /account/list goes to go.micro.api.account AccountHandler.List
func (h *accountHandler) List(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// parse values from the get request
	limitStr, ok := req.Get["limit"]

	if !ok || len(limitStr.Values) == 0 {
		return errors.BadRequest("go.micro.api.account", "no content")
	}

	limit, _ := strconv.Atoi(limitStr.Values[0])
	// make request
	response, err := h.userSrvClient.List(ctx, &userPB.UserListQuery{
		// Limit: &wrappers.UInt32Value{Value: uint32(limit)},
		// Page:  &wrappers.UInt32Value{Value: 1},
		Limit: uint32(limit),
		Page:  1,
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.account.call", err.Error())
	}
	log.Log(response)

	// set response status
	rsp.StatusCode = 200

	// respond with some json
	b, _ := json.Marshal(response)

	// set json body
	rsp.Body = string(b)

	return nil
}
