package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//MySQLClientConn defines mysql connection
type MySQLClientConn struct {
	clientConn *gorm.DB
}

//MySQLClientConnInterface is MySQLClientConn interfcae
type MySQLClientConnInterface interface {
	NewClientConnection() *gorm.DB
}

var conn func() *gorm.DB

//NewMySQLClientConn inject dependancies for
func NewMySQLClientConn() MySQLClientConnInterface {

	conn = NewDBConnection()
	return &MySQLClientConn{}
}

//NewClientConnection  new mysql client connection
func (mySQLClientConn MySQLClientConn) NewClientConnection() *gorm.DB {

	// client, err := gorm.Open("mysql", "rahul:password@tcp(127.0.0.1:3306)/ecommerce?charset=utf8&parseTime=True&loc=Local")

	// if err != nil {
	// 	fmt.Println("Error in Create client connection", err)
	// 	panic("Error In Create Client Connection")
	// }
	// return client

	return conn()

}

//NewDBConnection return same connection or return new Connection
func NewDBConnection() func() *gorm.DB {
	var conn *gorm.DB
	var err error
	conn, err = gorm.Open("mysql", "rahul:password@tcp(127.0.0.1:3306)/ecommerce?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("connot establish connections")
	}
	//fmt.Println(conn)
	return func() *gorm.DB {
		err1 := conn.DB().Ping()
		if err1 != nil {
			//fmt.Println("Error", err1)
			conn, err = gorm.Open("mysql", "rahul:password@tcp(127.0.0.1:3306)/ecommerce?charset=utf8&parseTime=True&loc=Local")
			if err != nil {
				panic("connot establish connections")
			}

		} else {
			println("GOT IN PING")
			//fmt.Println("GOT IN PING")
		}
		return conn
	}
}
