package endpoints

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/rahul-golang/ecommerce/users/pkg/httphandler"
)

func TestNewUserEndpoints(t *testing.T) {
	type args struct {
		router      *mux.Router
		restHandler *httphandler.UserHttpHandlers
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewUserEndpoints(tt.args.router, tt.args.restHandler)
		})
	}
}
