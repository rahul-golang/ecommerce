package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/rahul-golang/ecommerce/users/pkg/models"
)

type MockMySQLClientConn struct {
	clientConn *gorm.DB
}

//MySQLClientConnInterface is MySQLClientConn interfcae
type MockMySQLClientConnInterface interface {
	NewClientConnection() *gorm.DB
}

//NewMySQLClientConn inject dependancies for
func NewMokeMySQLClientConn() MockMySQLClientConnInterface {

	return &MockMySQLClientConn{}
}

//NewClientConnection  new mysql client connection
func (mySQLClientConn MockMySQLClientConn) NewClientConnection() *gorm.DB {

	client, err := gorm.Open("mysql", "rahul:password@tcp(127.0.0.1:3306)/ecommerce?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Error in Create client connection", err)
		panic("Error In Create Client Connection")
	}
	return client

}

// func TestNewUserRepository(t *testing.T) {

// 	type args struct {
// 		mysqlInterface MockMySQLClientConnInterface
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want UserRepositoryInterface
// 	}{
// 		{
// 			"Create Record",
// 			args{mysqlInterface: NewClientConnection()},
// 			UserRepositoryInterface{},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewUserRepository(tt.args.mysqlInterface); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestUserRepository_Create(t *testing.T) {

	connection := NewUserRepository(NewMokeMySQLClientConn())

	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name           string
		userRepository UserRepositoryInterface
		args           args
		//want           *models.User
		wantErr bool
	}{
		{
			"Create User ' Test :1 '",
			connection,
			args{
				ctx: context.Background(),
				user: models.User{
					FirstName:     "Rahul",
					LastName:      "Shinde",
					AccountNumber: "1232434354546",
				},
			},
			false,
		},
		{
			//this test throw an error because in database Mobile
			// number lent should be less than twelve i.e 12
			// Error Data too long for column 'user_mobile' at row 1
			"Create User ' Test :2 '",
			connection,
			args{
				ctx: context.Background(),
				user: models.User{
					FirstName:     "Rahul",
					LastName:      "Shinde",
					AccountNumber: "12324343",
					UserMobile:    "5454624323545657897654675465464646464646456456456546546545",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.userRepository.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("UserRepository.Create() = %v, want %v", got, tt.want)
			// }
			fmt.Println(got)
		})
	}
}

func TestUserRepository_Get(t *testing.T) {

	connection := NewUserRepository(NewMokeMySQLClientConn())
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name           string
		userRepository UserRepositoryInterface
		args           args
		//	want           *models.User
		wantErr bool
	}{
		{
			"' Test 1 ': Get user details for user Id :   ",
			connection,
			args{
				context.Background(),
				"10200",
			},
			false,
		},
		{
			// This test Will fail Because this id is not present in databse
			"' Test 2 ': Get user details for user Id :   ",
			connection,
			args{
				context.Background(),
				"100278",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.userRepository.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("UserRepository.Get() = %v, want %v", got, tt.want)
			// }
			fmt.Println(got)
		})
	}
}

func TestUserRepository_Delete(t *testing.T) {

	connection := NewUserRepository(NewMokeMySQLClientConn())

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name           string
		userRepository UserRepositoryInterface
		args           args
		//want           *models.DeleteUserResp
		wantErr bool
	}{
		{
			"'Test 1 ': Delete user details for user Id :10278 ",
			connection,
			args{
				context.Background(),
				"10276",
			},
			false,
		},
		{
			// This test Will fail Because this id is not present in databse
			"' Test 2 ': Delete user details for user Id : 10278",
			connection,
			args{
				context.Background(),
				"10277",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.userRepository.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("UserRepository.Delete() = %v, want %v", got, tt.want)
			// }
			fmt.Println("Result :", got)
		})
	}
}

func TestUserRepository_Update(t *testing.T) {

	connection := NewUserRepository(NewMokeMySQLClientConn())
	user := models.User{
		FirstName:     "Rahul",
		LastName:      "Shinde",
		AccountNumber: "1232434354546",
	}
	user.ID = 1053
	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name           string
		userRepository UserRepositoryInterface
		args           args
		//want           *models.User
		wantErr bool
	}{
		{
			"Create User ' Test :1 '",
			connection,
			args{
				ctx:  context.Background(),
				user: user,
			},
			false,
		},
		{
			//this test throw an error because in database Mobile
			// number lent should be less than twelve i.e 12
			// Error Data too long for column 'user_mobile' at row 1
			"Create User ' Test :2 '",
			connection,
			args{
				ctx:  context.Background(),
				user: user,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.userRepository.Update(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("UserRepository.Update() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestUserRepository_All(t *testing.T) {
	connection := NewUserRepository(NewMokeMySQLClientConn())
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name           string
		userRepository UserRepositoryInterface
		args           args
		//wantGetAllResp []*models.User
		wantErr bool
	}{
		{
			"Test Get All User ",
			connection,
			args{
				context.Background(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.userRepository.All(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotGetAllResp, tt.wantGetAllResp) {
			// 	t.Errorf("UserRepository.All() = %v, want %v", gotGetAllResp, tt.wantGetAllResp)
			// }
			//fmt.Println(gotGetAllResp)
		})
	}
}
