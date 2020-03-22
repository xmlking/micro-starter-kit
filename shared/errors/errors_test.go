package errors

// TODO: https://gochronicles.dev/posts/datastructures/list/singlylinkedlist/part-ii/

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	myErrors "github.com/micro/go-micro/v2/errors"
)

func TestErrors(t *testing.T) {
	testCases := []struct {
		name   string
		values string
		want   error
	}{
		{"EC1", "not good", AppError(EC1)},
	}

	testData := []*myErrors.Error{
		{
			Id:     "mkit.service.account",
			Code:   422,
			Detail: "proto validation: sumo-val-error",
			Status: http.StatusText(422),
		},
	}

	appErrTestData := []*myErrors.Error{
		{
			Id:     "EC1",
			Code:   500,
			Detail: "not good",
			Status: http.StatusText(500),
		},
	}

	// test AppError
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := AppError(EC1)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}

	// test ValidationError
	for _, e := range testData {
		ne := ValidationError("mkit.service.account", "proto validation: sumo-val-error")
		if e.Error() != ne.Error() {
			t.Fatalf("Expected %s got %s", e.Error(), ne.Error())
		}

		ne2 := ValidationError("mkit.service.account", "proto validation: %v", errors.New("sumo-val-error"))

		if e.Error() != ne2.Error() {
			t.Fatalf("Expected %s got %s", e.Error(), ne2.Error())
		}

		pe := myErrors.Parse(ne.Error())

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
