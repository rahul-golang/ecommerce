package httphandler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/rahul-golang/ecommerce/users/pkg/models"
	"github.com/rahul-golang/ecommerce/users/pkg/service"
)

type MokeBasicUsersService struct {
}

func NewMokeBasicUsersService() service.UsersService {
	return &MokeBasicUsersService{}
}

func (mokeBasicUsersService MokeBasicUsersService) CreateUser(ctx context.Context, createReq models.CreateUserReq) (createResp *models.CreateUserResp, err error) {
	return
}
func (mokeBasicUsersService MokeBasicUsersService) GetAllUser(ctx context.Context) (allRecordResp []*models.User, err error) {
	return
}
func (mokeBasicUsersService MokeBasicUsersService) UpdateUser(ctx context.Context, upadteReq models.User) (updateResp *models.User, err error) {
	return
}
func (mokeBasicUsersService MokeBasicUsersService) DeleteUser(ctx context.Context, id string) (deleteResp *models.DeleteUserResp, err error) {
	return
}
func (mokeBasicUsersService MokeBasicUsersService) GetUser(ctx context.Context, id string) (createResp *models.User, err error) {

	return
}
func TestNewUserRestHandler(t *testing.T) {
	type args struct {
		userService service.UsersService
	}
	tests := []struct {
		name string
		args args
		want *UserHttpHandlers
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRestHandler(tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRestHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHttpHandlers_CreateUser(t *testing.T) {
	basicService := NewMokeBasicUsersService()
	userHttpHandlers := NewUserRestHandler(basicService)
	// createReq := models.CreateUserReq{
	// 	models.User{
	// 		FirstName:     "Rahul",
	// 		LastName:      "Shinde",
	// 		AccountNumber: "12324343",
	// 		UserMobile:    "5454624323545657897654675465464646464646456456456546546545",
	// 	},
	// }
	type args struct {
		Method string
		URL    string
		Body   io.Reader
	}
	tests := []struct {
		name             string
		userHttpHandlers *UserHttpHandlers
		args             args
		want             error
	}{
		{
			"Test 1 ",
			userHttpHandlers,
			args{
				"GET",
				"users/1",
				nil,
			},
			nil,
		},
		{
			"Test 2 ",
			userHttpHandlers,
			args{
				"GET",
				"users/",
				nil,
			},
			nil,
		},
	}
	for _, tt := range tests {

		req, err := http.NewRequest(tt.args.Method, tt.args.URL, tt.args.Body)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tt.userHttpHandlers.GetUser)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		fmt.Println("" + rr.Body.String())

		// Check the response body is what we expect.
		expected := `{"message":"Sucess","user":null}`
		if rr.Body.String() == expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

}

func TestUserHttpHandlers_GetUser(t *testing.T) {

	basicService := NewMokeBasicUsersService()
	userHttpHandlers := NewUserRestHandler(basicService)

	type args struct {
		Method string
		URL    string
		Body   io.Reader
	}
	tests := []struct {
		name             string
		userHttpHandlers *UserHttpHandlers
		args             args
		want             error
	}{
		{
			"Test 1 ",
			userHttpHandlers,
			args{
				"GET",
				"users/1",
				nil,
			},
			nil,
		},
		{
			"Test 2 ",
			userHttpHandlers,
			args{
				"GET",
				"users/",
				nil,
			},
			nil,
		},
	}
	for _, tt := range tests {

		req, err := http.NewRequest(tt.args.Method, tt.args.URL, tt.args.Body)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tt.userHttpHandlers.GetUser)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		fmt.Println("" + rr.Body.String())

		// Check the response body is what we expect.
		expected := `{"message":"Sucess","user":null}`
		if rr.Body.String() == expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

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

	basicService := NewMokeBasicUsersService()
	userHttpHandlers := NewUserRestHandler(basicService)

	type args struct {
		Method string
		URL    string
		Body   io.Reader
	}
	tests := []struct {
		name             string
		userHttpHandlers *UserHttpHandlers
		args             args
		want             error
	}{
		{
			"Test 1 ",
			userHttpHandlers,
			args{
				"GET",
				"users/1",
				nil,
			},
			nil,
		},
		{
			"Test 2 ",
			userHttpHandlers,
			args{
				"GET",
				"users/",
				nil,
			},
			nil,
		},
	}
	for _, tt := range tests {

		req, err := http.NewRequest(tt.args.Method, tt.args.URL, tt.args.Body)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tt.userHttpHandlers.DeleteUser)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		fmt.Println("" + rr.Body.String())

		// Check the response body is what we expect.
		expected := `{"message":"Sucess","user":null}`
		if rr.Body.String() == expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter
}

func TestUserHttpHandlers_GetAllUser(t *testing.T) {
	basicService := NewMokeBasicUsersService()
	userHttpHandlers := NewUserRestHandler(basicService)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHttpHandlers.GetAllUser)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	fmt.Println("" + rr.Body.String())

	// Check the response body is what we expect.
	expected := `{"message":"Sucess","user":null}`
	if rr.Body.String() == expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
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
