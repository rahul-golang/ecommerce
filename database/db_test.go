package database

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Mockconn func() *gorm.DB
var clientConnection = NewDBConnection()

func TestNewDBConnection(t *testing.T) {
	tests := []struct {
		name string
		want func() *gorm.DB
	}{
		{
			"T1",
			Mockconn,
		},

		{
			"T2",
			Mockconn,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mockconn = NewDBConnection()
			if reflect.DeepEqual(Mockconn, tt.want) {
				//t.Errorf("NewDBConnection() = %v, want %v", got, tt.want)
				t.Error("errpor")
			}
			//sfmt.Println("", Mockconn)
		})
	}
}

func TestNewMySQLClientConn(t *testing.T) {
	tests := []struct {
		name string
		want MySQLClientConnInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMySQLClientConn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMySQLClientConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySQLClientConn_NewClientConnection(t *testing.T) {
	tests := []struct {
		name            string
		mySQLClientConn MySQLClientConn
		want            *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mySQLClientConn.NewClientConnection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQLClientConn.NewClientConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
