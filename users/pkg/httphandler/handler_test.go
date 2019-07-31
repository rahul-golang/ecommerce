package httphandler

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/rahul-golang/ecommerce/users/pkg/service"
)

func TestNewUserRestHandler(t *testing.T) {
	type args struct {
		userService service.UsersService
	}
	tests := []struct {
		name string
		args args
		want *UserHttpHandlers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRestHandler(tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRestHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHttpHandlers_CreateUser(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name             string
		userHttpHandlers UserHttpHandlers
		args             args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.userHttpHandlers.CreateUser(tt.args.w, tt.args.req)
		})
	}
}

func TestUserHttpHandlers_GetUser(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name             string
		userHttpHandlers UserHttpHandlers
		args             args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.userHttpHandlers.GetUser(tt.args.w, tt.args.req)
		})
	}
}

func TestUserHttpHandlers_UpdateUser(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name             string
		userHttpHandlers UserHttpHandlers
		args             args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.userHttpHandlers.UpdateUser(tt.args.w, tt.args.req)
		})
	}
}

func TestUserHttpHandlers_DeleteUser(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name             string
		userHttpHandlers UserHttpHandlers
		args             args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.userHttpHandlers.DeleteUser(tt.args.w, tt.args.req)
		})
	}
}

func TestUserHttpHandlers_GetAllUser(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name             string
		userHttpHandlers UserHttpHandlers
		args             args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.userHttpHandlers.GetAllUser(tt.args.w, tt.args.req)
		})
	}
}

func Test_writeResponse(t *testing.T) {
	type args struct {
		w         http.ResponseWriter
		errorCode int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeResponse(tt.args.w, tt.args.errorCode)
		})
	}
}
