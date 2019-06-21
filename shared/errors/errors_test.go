package errors

import (
	errs "errors"
	"net/http"
	"testing"

	"github.com/micro/go-micro/errors"
)

func TestErrors(t *testing.T) {
	testData := []*errors.Error{
		&errors.Error{
			Id:     "go.micro.srv.account",
			Code:   422,
			Detail: "proto validation: sumo-val-error",
			Status: http.StatusText(422),
		},
	}

	appErrTestData := []*errors.Error{
		&errors.Error{
			Id:     "EC1",
			Code:   500,
			Detail: "not good",
			Status: http.StatusText(500),
		},
	}

	// test ValidationError
	for _, e := range testData {
		ne := ValidationError("go.micro.srv.account", "proto validation: sumo-val-error")
		if e.Error() != ne.Error() {
			t.Fatalf("Expected %s got %s", e.Error(), ne.Error())
		}

		ne2 := ValidationError("go.micro.srv.account", "proto validation: %v", errs.New("sumo-val-error"))

		if e.Error() != ne2.Error() {
			t.Fatalf("Expected %s got %s", e.Error(), ne2.Error())
		}

		pe := errors.Parse(ne.Error())

		if pe == nil {
			t.Fatalf("Expected error got nil %v", pe)
		}

		if pe.Id != e.Id {
			t.Fatalf("Expected %s got %s", e.Id, pe.Id)
		}

		if pe.Detail != e.Detail {
			t.Fatalf("Expected %s got %s", e.Detail, pe.Detail)
		}

		if pe.Code != e.Code {
			t.Fatalf("Expected %d got %d", e.Code, pe.Code)
		}

		if pe.Status != e.Status {
			t.Fatalf("Expected %s got %s", e.Status, pe.Status)
		}
	}

	// test AppError
	ae := AppError(EC1)

	if appErrTestData[0].Error() != ae.Error() {
		t.Fatalf("Expected %s got %s", appErrTestData[0].Error(), ae.Error())
	}

}
